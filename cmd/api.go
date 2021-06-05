package main

import (
	"fmt"

	"github.com/mrityunjaygr8/go-pass/api"
)

func main() {
	fmt.Println("yo")

	app := api.App{}
	app.Initialize("mgr8", "dr0w.Ssap", "pass")

	app.Run(":8000")
}
