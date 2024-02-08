// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package interfaceutils

import (
	"fmt"
	"strings"

	"github.com/microsoft/CBL-Mariner/toolkit/tools/specreader/specreaderutils"
)

func BuildPackage(spec string) {
	fmt.Println("Building packages specs are...", spec)

	// convert space delimited string to list
	specList := strings.Fields(spec)

	// check specs exist
	checkSpecExist(specList)

	// any other checks

	// build toolchain if missing

	// build tools if missing

	// set extra configs
}


func checkSpecExist(specList []string) {
	fmt.Println("Checking if spec exists for ", specList)
	elementMap := make(map[string]bool)
	spec_dir := "/home/neha/repos/test/CBL-Mariner/SPECS"
	for _, spec := range specList {
		fmt.Println("spec is ", spec)
		elementMap[spec] = true
		myString, err := specreaderutils.FindSpecFiles(spec_dir, elementMap)
		if err != nil {
			fmt.Println("found error is ", err)
		} else {
			fmt.Println("ERROR freeeeee ", myString[0])
		}
		delete(elementMap, spec)
	}
	
	
	fmt.Println("done with specreader ")
}