// wrapper application to wrap the `tofu` binary and capture its outputs
// to set STDOUT, STDERR and EXITCODE

package main

import (
	"log"
	"os"
)

func main() {
	if err := setGithubOutput("foobar"); err != nil {
		log.Fatalf("cannot set envvar ")
	}
}

func setGithubOutput(s string) error {
	const key = "GITHUB_OUTPUT"
	o := os.Getenv(key)
	return os.Setenv(key, o+"\n"+s)
}
