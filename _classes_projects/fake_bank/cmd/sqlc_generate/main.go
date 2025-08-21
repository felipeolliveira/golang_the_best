package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command(
		"sqlc",
		"generate",
		"-f",
		"internal/store/pgstore/sqlc.yaml",
	)

	output, err := cmd.CombinedOutput()

	if (err) != nil {
		fmt.Println("Command exec failed: ", err)
		fmt.Println("Output:\n", string(output))
		return
	}

	fmt.Println("sqlc generated files has been success!", string(output))
}
