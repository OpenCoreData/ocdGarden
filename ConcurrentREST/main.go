// concurrent.go
// http://blog.narenarya.in/concurrent-http-in-go.html
// see also  http://stackoverflow.com/questions/33104192/how-to-run-10000-goroutines-in-parallel-where-each-routine-calls-an-api

package main

import (
     "fmt"
     "os"
     "net/http"
     "time"
     "io/ioutil"
)

func MakeRequest(url string, ch chan<-string) {
  start := time.Now()
  resp, _ := http.Get(url)

  secs := time.Since(start).Seconds()
  body, _ := ioutil.ReadAll(resp.Body)
  ch <- fmt.Sprintf("%.2f elapsed with response length: %d %s", secs, len(body), url)
}

func main() {
  start := time.Now()
  ch := make(chan string)
  for _,url := range os.Args[1:]{
      go MakeRequest(url, ch)
  }

  for range os.Args[1:]{
    fmt.Println(<-ch)
  }
  fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}