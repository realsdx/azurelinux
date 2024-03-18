// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package main

import (
	"fmt"
	"os"

	"github.com/microsoft/azurelinux/toolkit/tools/toolinterface/interfaceutils"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app          = kingpin.New("toolinterface", "A command-line interface for azurelinux toolkit")
	build        = app.Command("build", "Build azurelinux")
	setup        = app.Command("setup", "Setup machine")
	ready        = app.Command("ready", "Ready changes to contribute to opensource")
	
	buildPackage = build.Command("package", "Build package(s)")
	spec         = buildPackage.Flag("spec", "space separated \"\" enclosed name(s) of spec(s) to build").Default("").String()
//	specDir      = buildPackage.Flag("specDir", "directory containing spec files").String()
	rebuild      = buildPackage.Flag("rebuild", "rebuild forcefully even if already built").Default(false).Boolean()
	
	buildImage   = build.Command("image", "Build image(s)")
	config       = buildImage.Flag("config", "image config to build").String()
	configDir    = buildImage.Flag("configDir", "directory containing image config").String()
	
	buildTool     = build.Command("tools", "Build tool(s)")

    buildToolchain = build.Command("toolchain", "Build toolchain")
		
  )

func main() {
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
		case buildPackage.FullCommand():
			fmt.Println("in build_package ")
			fmt.Println("spec list is ", *spec)
			fmt.Println("specDir is ", *specDir)
			err = interfaceutils.BuildPackage(*spec)
			if err != nil {
				fmt.Println("error building package %v", err)
			}
		case buildImage.FullCommand():
		  fmt.Println("in image ")
		  fmt.Println("config file is ", *config)
		  err = interfaceutils.BuildImage(*config)
		  if err != nil {
			fmt.Println("error building image %v", err)
		}
		case buildTool.FullCommand():
			fmt.Println("in tools ")
		default:
			fmt.Println("Invalid call")
	}
}

