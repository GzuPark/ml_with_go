package main

import (
	"archive/zip"
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	tf "github.com/tensorflow/tensorflow/tensorflow/go"
	"github.com/tensorflow/tensorflow/tensorflow/go/op"
)

func main() {
	modelDir := "model"

	imageFile := flag.String("image", "", "Path of a JPEG-image to extract labels for")
	flag.Parse()

	if *imageFile == "" {
		flag.Usage()
		return
	}

	modelFile, labelsFile, err := modelFiles(modelDir)
	if err != nil {
		log.Fatal(err)
	}

	model, err := ioutil.ReadFile(modelFile)
	if err != nil {
		log.Fatal(err)
	}

	graph := tf.NewGraph()
	if err := graph.Import(model, ""); err != nil {
		log.Fatal(err)
	}

	// session
	// tf go는 raw level로 작성하는 것 같음
	session, err := tf.NewSession(graph, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	tensor, err := makeTensorFromImage(*imageFile)
	if err != nil {
		log.Fatal(err)
	}

	// inference
	output, err := session.Run(
		map[tf.Output]*tf.Tensor{
			graph.Operation("input").Output(0): tensor,
		},
		[]tf.Output{
			graph.Operation("output").Output(0),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}

	probabilities := output[0].Value().([][]float32)[0]
	printBestLabel(probabilities, labelsFile)
}

func printBestLabel(probabilities []float32, labelsFile string) {
	bestIdx := 0

	for i, p := range probabilities {
		if p > probabilities[bestIdx] {
			bestIdx = i
		}
	}

	f, err := os.Open(labelsFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var labels []string

	for scanner.Scan() {
		labels = append(labels, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Printf("ERROR: failed to read %s: %v", labelsFile, err)
	}

	fmt.Printf("Best Match: (%2.0f%% likely) %s\n", probabilities[bestIdx]*100.0, labels[bestIdx])
}

func makeTensorFromImage(fileName string) (*tf.Tensor, error) {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	tensor, err := tf.NewTensor(string(bytes))
	if err != nil {
		return nil, err
	}

	graph, input, output, err := constructGraphToNormalizeImage()
	if err != nil {
		return nil, err
	}

	session, err := tf.NewSession(graph, nil)
	if err != nil {
		return nil, err
	}
	defer session.Close()

	normalized, err := session.Run(
		map[tf.Output]*tf.Tensor{input: tensor},
		[]tf.Output{output},
		nil)
	if err != nil {
		return nil, err
	}

	return normalized[0], nil
}

func constructGraphToNormalizeImage() (graph *tf.Graph, input, output tf.Output, err error) {
	const (
		H, W  = 224, 224
		Mean  = float32(117)
		Scale = float32(1)
	)

	s := op.NewScope()
	input = op.Placeholder(s, tf.String)
	output = op.Div(s,
		op.Sub(s,
			op.ResizeBilinear(s, 
				op.ExpandDims(s,
					op.Cast(s,
						op.DecodeJpeg(s, input, op.DecodeJpegChannels(3)),
						tf.Float),
					op.Const(s.SubScope("make_batch"), int32(0))),
				op.Const(s.SubScope("size"), []int32{H, W})),
			op.Const(s.SubScope("mean"), Mean)),
		op.Const(s.SubScope("scale"), Scale))

	graph, err = s.Finalize()

	return graph, input, output, err
}

func modelFiles(dir string) (mnodelFile, labelsFile string, err error) {
	const URL = "https://storage.googleapis.com/download.tensorflow.org/models/inception5h.zip"

	var (
		model   = filepath.Join(dir, "tensorflow_inception_graph.pb")
		labels  = filepath.Join(dir, "imagenet_comp_graph_label_strings.txt")
		zipFile = filepath.Join(dir, "inception5h.zip")
	)

	if filesExist(model, labels) == nil {
		return model, labels, nil
	}
	log.Println("Did not find model in", dir, "downloading from", URL)
	
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", "", err
	}

	if err := download(URL, zipFile); err != nil {
		return "", "", fmt.Errorf("failed to download %v - %v", URL, err)
	}

	if err := unzip(dir, zipFile); err != nil {
		return "", "", fmt.Errorf("failed to extract contents from model archive: %v", err)
	}

	os.Remove(zipFile)

	return model, labels, filesExist(model, labels)
}

func filesExist(files ...string) error {
	for _, f := range files {
		if _, err := os.Stat(f); err != nil {
			return fmt.Errorf("unable to stat %s: %v", f, err)
		}
	}

	return nil
}

func download(URL, filename string) error {
	resp, err := http.Get(URL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, resp.Body)

	return err
}

func unzip(dir, zipFile string) error {
	r, err := zip.OpenReader(zipFile)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		src, err := f.Open()
		if err != nil {
			return err
		}

		log.Println("Extracting", f.Name)

		dst, err := os.OpenFile(filepath.Join(dir, f.Name), os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			return err
		}
		if _, err := io.Copy(dst, src); err != nil {
			return err
		}
		dst.Close()
	}

	return nil
}
