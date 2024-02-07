// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package main

import (
	"fmt"
	"os"

	"github.com/microsoft/CBL-Mariner/toolkit/tools/toolinterface/interfaceutils"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app          = kingpin.New("toolinterface", "A command-line interface for Mariner toolkit")
//	package      = app.Command("package", "Build packages(s)")
//	spec         = package.Flag("spec", "space separated name(s) of spec(s) to build").Default("all").String()
  	image        = app.Command("image", "Build image(s)")
  	config       = image.Flag("config", "image config to build").String()
    tools        = app.Command("tools", "Build tool(s)")
    toolchain    = app.Command("toolchain", "Build toolchain")
	deleteCommand     = app.Command("delete", "Delete an object.")
	deleteUserCommand = deleteCommand.Command("user", "Delete a user.")
	deletePostCommand = deleteCommand.Command("post", "Delete a post.")
  )

func main() {
	fmt.Println("Hello world!")
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
		// Register user
		case image.FullCommand():
		  fmt.Println("in image ")
		  fmt.Println("foncgi file is ", *config)
		  interfaceutils.BuildPackage(*config)
		case tools.FullCommand():
			fmt.Println("in tools ")
		case "delete user":
			fmt.Println("in delete user ")
	}
}

