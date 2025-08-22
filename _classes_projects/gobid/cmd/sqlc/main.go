package main

import (
	"fmt"
	"os/exec"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	cmd := exec.Command("sqlc", "generate", "-f", "internal", "store", "pgstore", "sqlc.yaml")

	fmt.Println("Running SQLC generate:", cmd.String())

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("SQLC generate failed:", err)
		fmt.Println("Output", string(output))
		panic(err)
	}

	fmt.Println("SQLC generate successfully")
}
