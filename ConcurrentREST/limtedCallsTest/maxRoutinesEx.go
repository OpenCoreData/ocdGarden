package main

import (
	"flag"
	"fmt"
	"time"
)

// Fake a long and difficult work.
func DoWork() {
	time.Sleep(500 * time.Millisecond)
}

func main() {
	maxNbConcurrentGoroutines := flag.Int("maxNbConcurrentGoroutines", 5, "the number of goroutines that are allowed to run concurrently")
	nbJobs := flag.Int("nbJobs", 100, "the number of jobs that we need to do")
	flag.Parse()

	// Dummy channel to coordinate the number of concurrent goroutines.
	// This channel should be buffered otherwise we will be immediately blocked
	// when trying to fill it.
	concurrentGoroutines := make(chan struct{}, *maxNbConcurrentGoroutines)
	// Fill the dummy channel with maxNbConcurrentGoroutines empty struct.
	for i := 0; i < *maxNbConcurrentGoroutines; i++ {
		concurrentGoroutines <- struct{}{}
	}

	// The done channel indicates when a single goroutine has
	// finished its job.
	done := make(chan bool)
	// The waitForAllJobs channel allows the main program
	// to wait until we have indeed done all the jobs.
	waitForAllJobs := make(chan bool)

	// Collect all the jobs, and since the job is finished, we can
	// release another spot for a goroutine.
	go func() {
		for i := 0; i < *nbJobs; i++ {
			<-done
			// Say that another goroutine can now start.
			concurrentGoroutines <- struct{}{}
		}
		// We have collected all the jobs, the program
		// can now terminate
		waitForAllJobs <- true
	}()

	// Try to start nbJobs jobs
	for i := 1; i <= *nbJobs; i++ {
		fmt.Printf("ID: %v: waiting to launch!\n", i)
		// Try to receive from the concurrentGoroutines channel. When we have something,
		// it means we can start a new goroutine because another one finished.
		// Otherwise, it will block the execution until an execution
		// spot is available.
		<-concurrentGoroutines
		fmt.Printf("ID: %v: it's my turn!\n", i)
		go func(id int) {
			DoWork()
			fmt.Printf("ID: %v: all done!\n", id)
			done <- true
		}(i)
	}

	// Wait for all jobs to finish
	<-waitForAllJobs
}
