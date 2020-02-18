package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Env struct {
	Name string
	Path string
}

func ShowEnv() []Env {
	var envs []Env
	var name string
	var path string

	for _, e := range os.Environ() {
		l := strings.Split(e, "=")
		name = l[0]
		path = l[1]
		envs = append(envs, Env{name, path})
	}
	return envs
}

func main() {
	envs := ShowEnv()
	for _, e := range envs {
		fmt.Printf("Name ： %v\nPath: %v \n", e.Name, e.Path)
	}

	// ピンポイントで取ることも可能
	var env = os.Getenv("USER")
	log.Println(env)
}
