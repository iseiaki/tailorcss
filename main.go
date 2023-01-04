package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/atotto/clipboard"
)

type Config struct {
	// importing Json Data
	Name      string `json:"Name"`
	Title     string `json:"title"`
	Authors   string `json:"Authors"`
	Version   string `json:"Version"`
	Html      string `json:"htmlfile"`

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

	Prebuilt struct {
        Navbar struct {
         Links string `json:"links"`
		 Navclass string `json:"navbarclass"`
		 ElementClass string `json:"elementclass"`
		 Content string `json:"content"`
		 Display string `json:"display"`
         Justifycontent string `json:"justify-content"`
         Position string `json:"position"`
		 Backgroundcolor string `json:"background-color"`
		}`json:"navbar"`
	} `json:"Prebuilt-elements"`

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
    
	// creating a variable so user could copy after a command
	navHTML := fmt.Sprintf(`
		<nav>
			<ul>
				<li><a href="#">%s</a></li>
				<li><a href="#">%s</a></li>
				<li><a href="#">%s</a></li>
			</ul>
		</nav>
	`, config.Prebuilt.Navbar.Content, config.Prebuilt.Navbar.Content, config.Prebuilt.Navbar.Content)


	fmt.Print("Enter Command: ") // taking command
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	command := scanner.Text()

	if command == "navbar -c" {
		err := clipboard.WriteAll(navHTML)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Copied to clipboard!")
	}

	if command == "navbar -w" {
		var output string
		n , err  := strconv.Atoi(config.Prebuilt.Navbar.Links)
		fmt.Println(n)
		if err != nil {
			fmt.Println(err)
			return
		}
		for i := 0; i < n; i++ {
			output += fmt.Sprintf("\n<li class='%v'><a href=\"#\">%v</a></li>", config.Prebuilt.Navbar.ElementClass, config.Prebuilt.Navbar.Content)
		}
		html := fmt.Sprintf(`
<nav>
<ul>
		%s
</ul>
</nav>
	`, output)

	errd := clipboard.WriteAll(html)
	if errd != nil {
		fmt.Println(errd)
		return
	}
	fmt.Println("Copied to clipboard!")
	}

	// Open the file for appending
	if command == "setup -html" {
    htmlsetup := fmt.Sprintf(`	
	<html>
    <head>
        <link rel="stylesheet" href="./tailor.css">
        <title>%v</title>
    </head>
    <body>

    </body>
    </html>`, config.Title)

	// Write the navbar HTML to the file
	err = ioutil.WriteFile(config.Html, []byte(htmlsetup), 0644,)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("setup is complete! Please check %s\n", config.Html)
}
}