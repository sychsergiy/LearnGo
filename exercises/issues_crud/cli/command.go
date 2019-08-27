package cli

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

type Command interface {
	Execute()
}

func RetrieveArgs(argsLen int) []string {
	if len(os.Args) < argsLen+2 {
		log.Fatal("Not enough arguments")
	}
	args := make([]string, 0, argsLen)

	for index, value := range os.Args {
		if index > 1 {
			args = append(args, value)
		}
	}
	return args
}

func RetrieveArgFromEditor() string {
	fpath := "./temp.txt"
	f, err := os.Create(fpath)
	if err != nil {
		log.Printf("1")
		log.Fatal(err)
	}
	_ = f.Close()

	cmd := exec.Command("nano", fpath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Start()
	if err != nil {
		log.Printf("2")
		log.Fatal(err)
	}
	err = cmd.Wait()
	if err != nil {
		log.Printf("Error while editing. Error: %v\n", err)
	} else {
		log.Printf("Successfully edited.")
	}

	f, err = os.Open(fpath)
	bytes, _ := ioutil.ReadAll(f)
	_ = os.Remove(fpath)
	return string(bytes)
}
