package main

import (
	"fmt"
	"os"
)

type Environment uint32

// Define a list of all environments
const (
	InvalidEnvironment Environment = iota + 0
	ProductionEnvironment
	StagingEnvironment
	TestingEnvironment
	DevelopmentEnvironment
)

func main() {
	if len(os.Args) <= 1 {
		os.Exit(1)
	}

	env := environment(os.Args[1])
	fmt.Println(env)
}

// Convert from environment type to string
func (e Environment) String() string {
	return map[Environment]string{
		InvalidEnvironment:     "invalid",
		ProductionEnvironment:  "production",
		StagingEnvironment:     "staging",
		TestingEnvironment:     "test",
		DevelopmentEnvironment: "development",
	}[e]
}

// Get all environments
func environments() []Environment {
	return []Environment{
		ProductionEnvironment,
		StagingEnvironment,
		TestingEnvironment,
		DevelopmentEnvironment,
	}
}

// Convert from string to environment type
func environment(env string) Environment {
	for _, v := range environments() {
		if fmt.Sprint(v) == env {
			return v
		}
	}
	return InvalidEnvironment
}
