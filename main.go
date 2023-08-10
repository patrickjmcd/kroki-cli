package main

import (
	"bytes"
	"compress/zlib"
	"encoding/base64"
	"flag"
	"fmt"
	"github.com/patrickjmcd/kroki-cli/internal"
	"github.com/pkg/errors"
	"os"
)

var krokiURL = "https://kroki.io"

// Encode takes a string and returns an encoded string in deflate + base64 format
func Encode(input string) (string, error) {
	var buffer bytes.Buffer
	writer, err := zlib.NewWriterLevel(&buffer, 9)
	if err != nil {
		return "", errors.Wrap(err, "fail to create the writer")
	}
	_, err = writer.Write([]byte(input))
	writer.Close()
	if err != nil {
		return "", errors.Wrap(err, "fail to create the payload")
	}
	result := base64.URLEncoding.EncodeToString(buffer.Bytes())
	return result, nil
}

func GetKrokiURL(inputType, outputType, encoded string) string {
	return krokiURL + "/" + inputType + "/" + outputType + "/" + encoded
}

func main() {
	// Parse command line flags
	fileName := flag.String("file", "", "The name of the file to be processed.")
	inputType := flag.String("input", "mermaid", "The input type for Kroki (default: mermaid).")
	outputType := flag.String("output", "svg", "The output type for Kroki (default: svg).")
	extract := flag.Bool("extract", false, "Extract the content from the file in a code block. eg: ```mermaid ... ``` (default: false)")

	if v, ok := os.LookupEnv("KROKI_URL"); ok {
		krokiURL = v
	}

	flag.Parse()

	// Validate file name
	if *fileName == "" {
		fmt.Println("Please provide a valid file name using the -file flag.")
		return
	}

	// Read file content
	content, err := os.ReadFile(*fileName)
	if err != nil {
		fmt.Printf("Failed to read the file: %v\n", err)
		return
	}

	var contentString string

	if *extract {
		content, err := internal.ExtractContent(string(content), *inputType)
		if err != nil {
			fmt.Printf("Failed to extract the content: %v\n", err)
			return
		} else {
			contentString = content
		}
	} else {
		contentString = string(content)
	}

	// Encode the content
	encoded, err := Encode(contentString)
	if err != nil {
		fmt.Printf("Failed to encode the content: %v\n", err)
		return
	}

	// Generate URL
	url := GetKrokiURL(*inputType, *outputType, encoded)
	fmt.Println("Generated Kroki URL:", url)
}
