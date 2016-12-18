package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

type config struct {
	Os         string `json:"os"`
	Arch       string `json:"arch"`
	OutputFile string `json:"out"`
}

func main() {
	flag.Parse()

	for _, val := range flag.Args() {
		data, err := ioutil.ReadFile(val)

		if err != nil {
			log.Fatal(err.Error())
		}

		s := string(data)

		// If not an array, pretend it is
		if s[0] != '[' {
			s = "[" + s + "]"
		}

		d := json.NewDecoder(strings.NewReader(s))

		var configs []config
		err = d.Decode(&configs)

		if err != nil {
			log.Fatal(err.Error())
		}

		for _, config := range configs {
			parse(&config)
			run(&config)
		}
	}
}

func parse(c *config) {
	if c.Arch == "x86" {
		c.Arch = "386"
	}
}

func run(c *config) error {
	fmt.Println("Building for", c.Os, c.Arch)

	env := os.Environ()

	// Remove possible GOOS and GOARCH variables
	for k, v := range env {
		if v[:4] == "GOOS" || v[:6] == "GOARCH" {
			env = append(env[:k], env[k+1:]...)
		}
	}

	env = append(env, "GOOS="+c.Os)
	env = append(env, "GOARCH="+c.Arch)

	p := exec.Command("go", "build", "-o", c.OutputFile)
	p.Env = env

	p.Run()

	return nil
}
