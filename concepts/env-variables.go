package concepts

import (
	"fmt"
	"os"
)

// os has env variable compatibility

func EnvVars() {

	// Setting an env variable
	os.Setenv("USER", "1")

	// Getting an env variable
	fmt.Println("USER env variable: ", os.Getenv("USER"))

	// Getting all the env variables
	for _, e := range os.Environ() {
		fmt.Println(e)
	}
}
