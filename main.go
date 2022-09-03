package main

import (
	"fmt"
	"os"
	"io/ioutil"
	// "path"
	"github.com/tidwall/gjson"
	"github.com/akamensky/argparse"
)


func main() {
	parser := argparse.NewParser("sinker", "sync files between local directories")

	config := parser.String("c", "config", &argparse.Options{Required: true, Help: "sync JSON configuration", Default: "./sinker.json"})
	filter := parser.String("f", "filter", &argparse.Options{Required: false, Help: "filter to specify which items to sync", Default: "tools.#(name=='diagrams')"})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
	}

	jsonFile, err := os.Open(*config)

	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	results := gjson.GetBytes(byteValue, *filter).Array()
	for k, v := range results {
        fmt.Println(k, v)
    }

	fmt.Println(*config)
	fmt.Println(*filter)	
}
