package main

import (
	"fmt"
	"flag"
	"log"

	"github.com/pachyderm/pachyderm/src/client"
	"github.com/pachyderm/pachyderm/src/client/pps"
)

func main() {
	user := flag.String("user", "", "User name of hub.docker.com")
	flag.Parse()

	if *user == "" {
		flag.Usage()
		return
	}

	c, err := client.NewFromAddress("0.0.0.0:30650")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	// https://godoc.org/github.com/pachyderm/pachyderm/src/client#APIClient.CreatePipeline
	// if err := c.CreatePipeline(
	// 	"prediction",                   // name of pipeline
	// 	*user + "/goregpredict:latest", // docker image
	// 	[]string{
	// 		"/goregpredict",
	// 		"-inModelDir=/pfs/model",
	// 		"-inVarDir=/pfs/attributes",
	// 		"-outDir=/pfs/out",
	// 	},                              // command run
	// 	[]string{},                     // stdin
	// 	&pps.ParallelismSpec{
	// 		Constant: 1,
	// 	},                              // parallelism
	// 	client.NewCrossInput(
	// 		client.NewPFSInput("attributes", "/*"),
	// 		client.NewPFSInput("model", "/"),
	// 	),
	// 	"master",
	// 	true,
	// ); err != nil {
	// 	log.Fatal(err)
	// }

	// c.PpsAPIClient.CreatePipeline()
	// https://github.com/pachyderm/pachyderm/blob/dfe62153c8cfb39331b0034c2bf916f2da62f1b7/src/client/pps.go#L643
	// type CreatePipelineRequest struct {}
	// https://github.com/pachyderm/pachyderm/blob/dfe62153c8cfb39331b0034c2bf916f2da62f1b7/src/client/pps/pps.pb.go#L4531
	if _, err := c.PpsAPIClient.CreatePipeline(
		c.Ctx(),
		&pps.CreatePipelineRequest{
			Pipeline:        client.NewPipeline("prediction"),
			Transform:       &pps.Transform{
				Image: *user + "/goregpredict:latest",
				Cmd:   []string{"/goregpredict", "-inModelDir=/pfs/model", "-inVarDir=/pfs/attributes", "-outDir=/pfs/out"},
				Stdin: []string{},
			},
			ParallelismSpec: &pps.ParallelismSpec{Constant: 1},
			Input:           client.NewCrossInput(
				client.NewPFSInput("attributes", "/"),
				client.NewPFSInput("model", "/"),
			),
			OutputBranch:    "master",
			Update:          true,
			Reprocess:       true,
		},
	); err != nil {
		log.Fatal(err)
	}

	pipelines, err := c.ListPipeline()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(pipelines)
}
