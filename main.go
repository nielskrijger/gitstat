package main

import (
	"fmt"
	"github.com/nielskrijger/gitstat/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		fmt.Printf("error: %s", err);
	}
}
