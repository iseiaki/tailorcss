package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
	Name      string `json:"Name"`
	Authors   string `json:"Authors"`
	Version   string `json:"Version"`
	Html   string `json:"htmlfile"`
	MainCSS   struct {
		Primary   string `json:"Primary"`
		Icons     string `json:"icons"`
		Style     string `json:"style"`
		Footer    string `json:"footer"`
		Navbar    string `json:"navbar"`
		Animations string `json:"animations"`
	} `json:"maincss"`
	Display struct {
		Grid   string `json:"Grid"`
		Flex   string `json:"Flex"`
		Block  string `json:"Block"`
		InlineB string `json:"InlineB"`
	} `json:"Display"`
	JustifyContent struct {
		Center       string `json:"center"`
		SpaceAround  string `json:"space-around"`
		SpaceBetween string `json:"space-between"`
		SpaceEvenly  string `json:"space-evenly"`
		FlexEnd      string `json:"flex-end"`
		FlexStart    string `json:"flex-start"`
	} `json:"Justify-Content"`
}

func main() {
	// Read the JSON file
	data, err := ioutil.ReadFile("costumize.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Unmarshal the JSON data into a Config struct
	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		fmt.Println(err)
		return
	}


	fmt.Println(config.Name)
	fmt.Println(config.Html)
	fmt.Println(config.Authors)
	fmt.Println(config.Version)
	fmt.Println(config.MainCSS.Primary)
	fmt.Println(config.MainCSS.Icons)
	fmt.Println(config.MainCSS.Style)
}