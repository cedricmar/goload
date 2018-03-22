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
			looper.Loop()
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
		loopCmd := exec.Command("goload", "loop")
		buildCmd := exec.Command("vgo", "install", config)

		loopCmd.Stdout = os.Stdout
		loopCmd.Stderr = os.Stderr

		buildCmd.Stdout = os.Stdout
		buildCmd.Stderr = os.Stderr

		err = loopCmd.Start()
		if err != nil {
			log.Fatal(err)
		}

		err = loopCmd.Wait()
		if err != nil {
			log.Fatal(err)
		}

		log.Println("reload...")

		_ = buildCmd.Start()

		_ = buildCmd.Wait()
	}
}
