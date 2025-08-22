package main

import (
	"fmt"
	"os/exec"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	cmd := exec.Command(
		"tern",
		"migrate",
		"--migrations",
		"./internal/store/pgstore/migrations",
		"--config",
		"./internal/store/pgstore/migrations/tern.conf",
	)

	fmt.Println("Running migrations with command:", cmd.String())

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Migrations failed to execute:", err)
		fmt.Println("Output", string(output))
		panic(err)
	}

	fmt.Println("Migrations executed successfully:", string(output))
}
