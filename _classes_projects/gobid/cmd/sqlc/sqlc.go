package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	sqlcConfig := os.Getenv("SQLC_CONFIG")

	cmd := exec.Command("sqlc", "generate", "-f", sqlcConfig)

	fmt.Println("Running SQLC generate:", cmd.String())

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("SQLC generate failed:", err)
		fmt.Println("Output", string(output))
		panic(err)
	}

	fmt.Println("SQLC generate successfully")
}
