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
	build        = app.Command("build", "Build in Mariner")
	buildPackage = build.Command("package", "Build package(s)")
	spec         = buildPackage.Flag("spec", "space separated \"\" enclosed name(s) of spec(s) to build").Default("all").String()
	
	buildImage   = build.Command("image", "Build image(s)")
	config       = buildImage.Flag("config", "image config to build").String()
	
	buildTool     = build.Command("tools", "Build tool(s)")

    buildToolchain = build.Command("toolchain", "Build toolchain")
		
  )

func main() {
	fmt.Println("Hello world!")
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
		case buildPackage.FullCommand():
			fmt.Println("in build_package ")
			fmt.Println("spec list is ", *spec)
			interfaceutils.BuildPackage(*spec)
		case buildImage.FullCommand():
		  fmt.Println("in image ")
		  fmt.Println("config file is ", *config)
		  interfaceutils.BuildImage(*config)
		case buildTool.FullCommand():
			fmt.Println("in tools ")		
	}
}

