package main // import "github.com/weebagency/goload/cmd/goload"

import (
	"log"
	"os"
	"os/exec"

	"github.com/weebagency/goload/pkg/looper"
)

func main() {

	if len(os.Args) > 1 {
		if os.Args[1] == "loop" {
			looper.Loop()
		}
		return
	}

	for {
		// Start a process:
		loopCmd := exec.Command("goload", "loop")
		buildCmd := exec.Command("vgo", "install", "./cmd/goload")

		loopCmd.Stdout = os.Stdout
		loopCmd.Stderr = os.Stderr

		buildCmd.Stdout = os.Stdout
		buildCmd.Stderr = os.Stderr

		err := loopCmd.Start()
		if err != nil {
			log.Fatal(err)
		}

		err = loopCmd.Wait()
		if err != nil {
			log.Fatal(err)
		}

		log.Println("reload...")

		err = buildCmd.Start()
		if err != nil {
			log.Fatal(err)
		}

		err = buildCmd.Wait()
		if err != nil {
			log.Fatal(err)
		}
	}
}
