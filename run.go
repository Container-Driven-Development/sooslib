package sooslib

import (
	"bytes"
	"fmt"
	"os/exec"
	"os/user"
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

var execCommand = exec.Command

// ExecCmd to execute command
func ExecCmd(cmdString string) bool {

	cmdStringSplitted := strings.Split(cmdString, " ")
	command, args := cmdStringSplitted[0], cmdStringSplitted[1:]

	cmd := execCommand(command, args...)
	var out bytes.Buffer
	cmd.Stdout = &out
	runErr := cmd.Run()

	check(runErr)

	if out.String() == "" {
		return false
	}

	fmt.Print(out.String())

	return true

}

var userCurrent = user.Current

func getUserBind() string {

	usr, userCurrentErr := userCurrent()

	check(userCurrentErr)

	return ("-u " + usr.Uid + ":" + usr.Gid)
}

// BuildRunCmd generate docker run command
func BuildRunCmd(config BuildConfig) string {

	userBind := getUserBind()
	portsBind := prefixAndJoin(config.Ports, "-p")

	volumes := append(config.Volumes, cwd()+":/build/app")

	volumesBind := prefixAndJoin(volumes, "-v")
	hashTag := Tokenizer(config.Hashfiles)

	cmd := fmt.Sprintf("docker run -it %v %v %v %v:%v", userBind, volumesBind, portsBind, config.Image, hashTag)

	return cmd
}
