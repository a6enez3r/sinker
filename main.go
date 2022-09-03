package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"strings"
	// "path"
	"github.com/tidwall/gjson"
	"github.com/akamensky/argparse"
)

func sync(multi []string, byteValue []byte) {
	fmt.Println(multi)
	for _, single := range multi {
		results := gjson.GetBytes(byteValue, single).Array()

		for k, v := range results {
			fmt.Println(k, v)
			// fmt.Println(v.Map()["filename"])
		}
	}
}

func reader(path string) []byte {
	jsonFile, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	return byteValue
}


func main() {
	parser := argparse.NewParser("sinker", "sync files between local directories")

	config := parser.String("c", "config", &argparse.Options{Required: true, Help: "sync JSON configuration", Default: "./sinker.json"})
	filter := parser.String("f", "filter", &argparse.Options{Required: false, Help: "filter to specify which items to sync", Default: "all"})
	
	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
	}

	byteValue := reader(*config)

	if *filter == "all" || strings.Contains(*filter, "|") {
		if *filter == "all" {
			var multi []string
			for _, item := range gjson.GetBytes(byteValue, *filter).Array() {
				multi = append(multi, item.String())
			}
			sync(multi, byteValue)
		} else {
			multi := strings.Split(*filter, "|")
			sync(multi, byteValue)
		}
	} else {
		multi := []string{*filter}
		sync(multi, byteValue)
	}
}
