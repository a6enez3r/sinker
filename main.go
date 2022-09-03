package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

type Config struct {
	Tools           []Location `json:"tools"`
	Learning        []Location `json:"learning"`
	Packages        []Location `json:"packages"`
	WebappsPersonal []Location `json:"webapps-personal"`
	WebappsProjects []Location `json:"webapps-projects"`
}

type Location struct {
	Name     string `json:"name"`
	Path     string `json:"path"`
	Filename string `json:"filename"`
	Source   string `json:"source"`
}

func main() {
	var config string
	flag.StringVar(&config, "config", "config.json", "sync config file")

	// Open our jsonFile
	jsonFile, err := os.Open(config)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened config.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	var configJson Config

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &configJson)

	// we iterate through every user within our users array and
	// print out the user Type, their name, and their facebook url
	// as just an example
	for group, items := range configJson {
		fmt.Println("Group: " + group)
		for i := 0; i < len(items); i++ {
			fmt.Println("Destination: " + path.Join(items[i].Path, items[i].Filename))
			fmt.Println("Source: " + items[i].Source)
		}
    }
}
