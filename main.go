package main

import (
	"gomeetup/levelthree"
	"log"
)

func main() {
	client, err := createS3Client()
	if err != nil {
		log.Fatal(err)
	}

	levelthree.GetFile(client, "some file")
}

// placeholder
func createS3Client() (levelthree.FileClient, error) {
	return nil, nil
}
