package sooslib

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func prefixAndJoin(list []string, prefix string) string {

	alteredList := []string{}
	for _, value := range list {
		alteredList = append(alteredList, fmt.Sprintf("%v %v", prefix, value))
	}
	joinedList := strings.Join(alteredList, " ")

	return joinedList

}

func execCmd(cmdString string) bool {

	cmd := exec.Command(cmdString)
	var out bytes.Buffer
	cmd.Stdout = &out
	runErr := cmd.Run()

	check(runErr)

	if out.String() == "" {
		return false
	}

	return true

}

// BuildRunCmd generate docker run command
func BuildRunCmd(config BuildConfig) string {

	portsBind := prefixAndJoin(config.Ports, "-p")
	volumesBind := prefixAndJoin(config.Volumes, "-v")
	hashTag := Tokenizer(config.Hashfiles)

	cmd := fmt.Sprintf("docker run -it %v %v %v:%v", volumesBind, portsBind, config.Image, hashTag)

	return cmd
}

// Run build container
func Run(configFile string) {

	config := Config(configFile)

	cmd := BuildRunCmd(config)

	execCmd(cmd)

}
