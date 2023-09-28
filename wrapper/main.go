// wrapper application to wrap the `tofu` binary and capture its outputs
// to set STDOUT, STDERR and EXITCODE

package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	if err := setGithubOutput("foo", "bar"); err != nil {
		log.Fatalf("cannot set envvar ")
	}
}

func setGithubOutput(key, val string) error {
	const tag = "GITHUB_OUTPUT"
	outputPath := os.Getenv(tag)
	if outputPath == "" {
		cmd := exec.Command("echo", `"`+key+"="+val+`" >> $`+tag)
		return cmd.Run()
	}

	f, err := os.OpenFile(outputPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer func() { _ = f.Close() }()

	_, err = f.WriteString("\n" + key + "=" + val)
	return err
}
