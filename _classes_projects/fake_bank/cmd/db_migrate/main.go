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

	args := []string{
		"migrate",
		"--migrations",
		"./internal/store/pgstore/migrations",
		"--config",
		"./internal/store/pgstore/migrations/tern.conf",
	}
	args = append(args, os.Args...)

	cmd := exec.Command("tern", args...)

	output, err := cmd.CombinedOutput()

	if (err) != nil {
		fmt.Println("Command exec failed: ", err)
		fmt.Println("Output:\n", string(output))
		return
	}

	if len(output) == 0 {
		fmt.Println("Migration is already up to date!")
		return
	}

	fmt.Println("Migration success!", string(output))
}
