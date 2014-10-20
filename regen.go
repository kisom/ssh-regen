// regenerate an SSH private key.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/gokyle/readpass"
	"github.com/gokyle/sshkey"
)

func init() {
	sshkey.PasswordPrompt = readpass.PasswordPrompt
}

// ReadLine reads a line of input from the user.
func ReadLine(prompt string) (line string, err error) {
	fmt.Printf(prompt)
	rd := bufio.NewReader(os.Stdin)
	line, err = rd.ReadString('\n')
	if err != nil {
		return
	}
	line = strings.TrimSpace(line)
	return
}

func errorf(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format, args...)
}

func fatalf(format string, args ...interface{}) {
	errorf(format, args...)
	os.Exit(1)
}

func errfatal(err error) {
	if err != nil {
		fatalf("%v\n", err)
	}
}

func storePublic(filename string) {
	priv, _, err := sshkey.LoadPrivateKeyFile(filename)
	if err != nil {
		errorf("%v\n", err)
		return
	}

	publicName := filename + ".pub"
	comment, err := ReadLine("Comment: ")
	if err != nil {
		errorf("%v\n", err)
		return
	}

	pub := sshkey.NewPublic(priv, comment)
	pubBytes := sshkey.MarshalPublic(pub)
	err = ioutil.WriteFile(publicName, pubBytes, 0644)
	if err != nil {
		errorf("%v\n", err)
		return
	}
}

func main() {
	flag.Parse()

	for _, filename := range flag.Args() {
		storePublic(filename)
	}
}
