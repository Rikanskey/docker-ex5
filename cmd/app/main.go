package main

import "docker_ex5/internal/runner"

const configDir = "./config/"

func main() {
	runner.Start(configDir)
}
