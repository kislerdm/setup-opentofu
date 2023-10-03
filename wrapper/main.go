// wrapper application to wrap the `tofu` binary and capture its outputs
// to set STDOUT, STDERR and EXITCODE

package main

import (
	"log"
	"os/exec"
)

func main() {
	if err := setGithubOutput("foo", "bar"); err != nil {
		log.Fatalf("cannot set envvar %#v", err)
	}
}

func setGithubOutput(key, val string) error {
	const tag = "GITHUB_OUTPUT"

	assignment := key + "=" + val
	cmd := assignment + " >> $" + tag
	command := exec.Command("echo", cmd)
	return command.Run()
}
