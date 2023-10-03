// wrapper application to wrap the `tofu` binary and capture its outputs
// to set STDOUT, STDERR and EXITCODE

package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if err := setGithubOutput("foo", "bar"); err != nil {
		log.Fatalf("cannot set envvar %#v", err)
	}
}

func setGithubOutput(key, val string) error {
	const tag = "GITHUB_OUTPUT"

	assignment := key + "=" + val
	out := os.Getenv(tag)
	if out == "" {
		cmd := "echo " + assignment + " >> " + tag
		_, err := fmt.Fprintln(os.Stdout, cmd)
		return err
	}

	f, err := os.OpenFile(out, os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer func() { _ = f.Close() }()

	_, err = fmt.Fprintln(f, assignment)
	return err
}
