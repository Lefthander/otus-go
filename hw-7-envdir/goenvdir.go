package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

// ReadDir reads a specified folder and returns map of  env variables. Variables represented as files where filename - name of variable, file content -  value.
func ReadDir(dir string) (map[string]string, error) {

	files := []string{}

	env := make(map[string]string)

	rootpath, err := filepath.Abs(dir)
	if err != nil {
		log.Println("Error processing path ", err)
		return nil, err
	}

	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		files = append(files, info.Name())
		return nil
	})

	if err != nil {
		log.Println("Error reading directory", err)
		return nil, err
	}

	fmt.Println(files)

	for _, file := range files {

		data, err := ioutil.ReadFile(filepath.Join(rootpath, file))
		if err != nil {
			return nil, err
		}
		env[file] = string(data)

	}

	return env, nil

}

// RunCmd runs a command + arguments (cmd) with environment variables - env
func RunCmd(cmd []string, env map[string]string) error {

	envl := []string{}

	for k, v := range env {

		envl = append(envl, k+"="+v)
	}

	var command *exec.Cmd

	if len(cmd) > 1 {
		command = exec.Command(cmd[0], cmd[1:]...)
	} else {
		command = exec.Command(cmd[0])
	}
	command.Env = envl
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()

	if err != nil {
		log.Println("Error executing command", err)
		return err
	}

	return nil
}

func main() {

	if len(os.Args) < 3 {
		log.Println("Error Should be two arguments")
		return
	}
	result, err := ReadDir(os.Args[1])
	if err != nil {
		log.Println("Error reading directory")
		return
	}
	err = RunCmd(os.Args[2:], result)
	if err != nil {
		log.Println("An error ocured executing command", err)

	}
}
