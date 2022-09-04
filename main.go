package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	// "path"
	"github.com/akamensky/argparse"
	"github.com/tidwall/gjson"
)

func copy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

func sync(multi []string, byteValue []byte, inverse bool) {

	for _, single := range multi {
		results := gjson.GetBytes(byteValue, single).Array()

		for _, v := range results {
			source, destination := v.Map()["source"].String(), v.Map()["destination"].String()
			if inverse {
				copy(destination, source)
			} else {
				copy(source, destination)
			}
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
	inverse := parser.Flag("i", "inverse", &argparse.Options{Required: false, Help: "sync from destination to source", Default: false})

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
			sync(multi, byteValue, *inverse)
		} else {
			multi := strings.Split(*filter, "|")
			sync(multi, byteValue, *inverse)
		}
	} else {
		multi := []string{*filter}
		sync(multi, byteValue, *inverse)
	}
}
