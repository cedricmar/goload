package main // import "github.com/cedricmar/goload/cmd/goload"

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/cedricmar/goload/pkg/config"
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
		c := config.LoadConfig()
		dir := "./" + c.Get("main_dir")

		fmt.Println(dir)

		// Start a process:

		buildCmd := exec.Command("vgo", "build", "-o", "./tmp/prg", dir)

		buildCmd.Stdout = os.Stdout
		buildCmd.Stderr = os.Stderr

		log.Println("Running", "CC=gcc", "&&", "vgo", "build", "-x", "-o", "./tmp/prg", dir)

		err := buildCmd.Start()
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
