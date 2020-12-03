package main

import (
	"log"

	"github.com/pachyderm/pachyderm/src/client"
	"github.com/pachyderm/pachyderm/src/client/pfs"
)

func main() {
	c, err := client.NewFromAddress("0.0.0.0:30650")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()
	
	if _, err := c.PfsAPIClient.CreateRepo(
		c.Ctx(),
		&pfs.CreateRepoRequest{
			Repo:        client.NewRepo("training"),
			Description: "Diabetes training data (.csv)",
			Update:      true,
		},
	); err != nil {
		log.Fatal(err)
	}

	if _, err := c.PfsAPIClient.CreateRepo(
		c.Ctx(),
		&pfs.CreateRepoRequest{
			Repo:        client.NewRepo("attributes"),
			Description: "Attributes data (.json)",
			Update:      true,
		},
	); err != nil {
		log.Fatal(err)
	}

	repos, err := c.ListRepo()
	if err != nil {
		log.Fatal(err)
	}

	if len(repos) != 2 {
		log.Fatal("Unexpected number of data repositories")
	}

	if repos[0].Repo.Name != "attributes" || repos[1].Repo.Name != "training" {
		log.Fatal("Unexpected data repository name")
	}
}
