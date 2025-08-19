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

	cmd := exec.Command(
		"sqlc",
		"generate",
		"-f",
		"./internal/store/pgstore/sqlc.yaml",
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Command exec filed:", err)
		fmt.Println("Output:\n", string(output))
		return
	}

	fmt.Println("Executed with success\n", string(output))
}
