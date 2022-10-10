package main

import "docker-ex5/internal/runner"

const configDir = "./config/"

func main() {
	runner.Start(configDir)
}
