package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/pachyderm/pachyderm/src/client"
)

func main() {
	c, err := client.NewFromAddress("0.0.0.0:30650")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()
	
	// go to parent of parent path which can access to storage directory in local
	exPath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i <= 2; i++ {
		exPath = filepath.Dir(exPath)
	}

	repoName := "attributes"
	path := filepath.Join(exPath, "storage", "attributes")
	attrFiles := []string{"1.json", "2.json", "3.json"}
	CustomPutFile(attrFiles, path, repoName, "master", c)

	repoName = "training"
	path = filepath.Join(exPath, "storage", "data")
	attrFiles = []string{"diabetes_training.csv"}
	CustomPutFile(attrFiles, path, repoName, "master", c)
}

func CustomPutFile(files []string, path string, repoName string, commitID string, c *client.APIClient) {
	for _, attrFile := range files {
		attrPath := filepath.Join(path, attrFile)

		f, err := os.Open(attrPath)
		if err != nil {
			log.Fatal(err)
		}

		// https://godoc.org/github.com/pachyderm/pachyderm/src/client#APIClient.PutFileOverwrite
		// PutFileOverwrite를 사용하지 않으면(=PutFile) 같은 이름의 파일은 append됨
		// PutFileOverwrite 마지막 옵션을 0 으로 설정할 경우 파일을 새로 쓰는 것 (1은 append)
		if _, err := c.PutFileOverwrite(repoName, commitID, attrFile, f, 0); err != nil {
			log.Fatal(err)
		}
	}
}
