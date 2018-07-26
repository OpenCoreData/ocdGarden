package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	minio "github.com/minio/minio-go"
)

// Config holds minio info
type Config struct {
	Minio struct {
		Endpoint        string `json:"endpoint"`
		AccessKeyID     string `json:"accessKeyID"`
		SecretAccessKey string `json:"secretAccessKey"`
	} `json:"minio"`
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		// Read the keyboad input.
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		// Remove the newline character.
		input = strings.TrimSuffix(input, "\n")

		// Skip an empty input.
		if input == "" {
			continue
		}

		// Handle the execution of the input.
		err = execInput(input)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func setMort() {

}

// ErrNoPath is returned when 'cd' was called without a second argument.
var ErrNoPath = errors.New("path required")

func execInput(input string) error {
	// Split the input separate the command and the arguments.
	args := strings.Split(input, " ")

	// Set up minio
	cs := loadConfiguration("./config.json") // needs to be a file pointer
	mc := miniConnection(cs)

	// Check for built-in commands.
	switch args[0] {
	case "cd":
		// 'cd' to home with empty path not yet supported.
		if len(args) < 2 {
			return ErrNoPath
		}
		err := os.Chdir(args[1])
		if err != nil {
			return err
		}
		// Stop further processing.
		return nil
	case "exit":
		os.Exit(0)
	case "stat":
		objstat, err := mc.StatObject(args[1], args[2], minio.StatObjectOptions{})
		if err != nil {
			log.Println(err)
		}
		fmt.Println(objstat)
		return nil
	case "morton":
		bucket := args[1]
		object := args[2]

		// first copy object
		src := minio.NewSourceInfo(bucket, object, nil)

		usermeta := make(map[string]string) // what do I want to know?
		usermeta["moratoriumstatus"] = args[3]
		usermeta["moratoriumdate"] = "a date string here DDMMYYYY"

		// Upload the file with FPutObject
		//n, err := minioClient.PutObject(bucketName, objectName, b, int64(b.Len()), minio.PutObjectOptions{ContentType: contentType, UserMetadata: usermeta})

		// set dest
		dst, err := minio.NewDestinationInfo(bucket, object, nil, usermeta)
		if err != nil {
			log.Fatalln(err)
		}

		// Initiate copy object.
		err = mc.CopyObject(dst, src)
		if err != nil {
			log.Fatalln(err)
		}
		log.Println("Copied source object to destination Successfully.")

		return nil
	}

	// Prepare the command to execute.
	cmd := exec.Command(args[0], args[1:]...)

	// Set the correct output device.
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	// Execute the command and save it's output.
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

// utils below  (move to another file later)

// load config file
func loadConfiguration(file string) Config {
	var config Config
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}

// Set up minio and initialize client
func miniConnection(cs Config) *minio.Client {
	endpoint := cs.Minio.Endpoint
	accessKeyID := cs.Minio.AccessKeyID
	secretAccessKey := cs.Minio.SecretAccessKey
	useSSL := false
	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		log.Fatalln(err)
	}
	return minioClient
}
