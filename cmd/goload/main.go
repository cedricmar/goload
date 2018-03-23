package main // import "github.com/weebagency/goload/cmd/goload"

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/weebagency/goload/pkg/looper"
)

func main() {

	if len(os.Args) > 1 {
		if os.Args[1] == "loop" {

			log.Println("looping...")

			prgCmd := exec.Command("./tmp/prg")

			prgCmd.Stdout = os.Stdout
			prgCmd.Stderr = os.Stderr

			err := prgCmd.Start()
			if err != nil {
				log.Fatal(err)
			}

			/*
				err = prgCmd.Wait()
				if err != nil {
					log.Fatal(err)
				}
			*/

			looper.Loop() // Blocking

			if err := prgCmd.Process.Kill(); err != nil {
				log.Println("failed to kill process: ", err)
			}

		}
		return
	}

	for {
		// Get the config
		file, err := ioutil.ReadFile("config.json")
		if err != nil {
			log.Fatal(err)
		}

		var rawConfig map[string]*json.RawMessage
		if err = json.Unmarshal(file, &rawConfig); err != nil {
			log.Fatal(err)
		}

		config := fmt.Sprintf("./%s", strings.Trim(string(*rawConfig["main_dir"]), "\""))

		// Start a process:

		buildCmd := exec.Command("vgo", "build", "-o", "./tmp/prg", config)

		buildCmd.Stdout = os.Stdout
		buildCmd.Stderr = os.Stderr

		log.Println("Running", "CC=gcc", "&&", "vgo", "build", "-x", "-o", "./tmp/prg", config)

		err = buildCmd.Start()
		if err != nil {
			log.Println(err)
		}

		err = buildCmd.Wait()
		if err != nil {
			log.Println(err)
		}

		log.Println("Running loop")
		loopCmd := exec.Command("goload", "loop")

		loopCmd.Stdout = os.Stdout
		loopCmd.Stderr = os.Stderr

		err = loopCmd.Start()
		if err != nil {
			log.Fatal(err)
		}

		err = loopCmd.Wait()
		if err != nil {
			log.Fatal(err)
		}

		log.Println("reload...")

	}
}
