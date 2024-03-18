// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package interfaceutils

import (
	"fmt"
	"strings"

	"github.com/microsoft/azurelinux/toolkit/tools/internal/packagelist"
	"github.com/microsoft/azurelinux/toolkit/tools/pkg/specreaderutils"
)

const (
	azlSpecsDirs := []string {"SPECS", "SPECS-EXTENDED", "SPECS-SIGNED"}
)

func BuildPackage(spec string) (err error) {

	var specListFile string

	fmt.Println("Building packages specs are...", spec)

	// check specs exist
	// first, write space delimited string of specs to file
	// TODO: we can probably get rid of writing srpm_pack_list to a file?
	specList := strings.Fields(spec)
	specListFile, err := getSpecListFile(spec)
	if err != nil {
		fmt.Println("error in getting spec file name %v", err)
		return err
	}

	err = validateSpecExistance(specList, specListFile)
	if err != nil {
		fmt.Println("error in validating specs %v", err)
		return err
	}

	// any other checks

	// build toolchain if required

	// put toolchain rpms into toolchain_archive and use it

	// build tools if required

	// set extra configs

	// show dependency graph

	// build package

	// show output
}

// validateSpecExistance checks if each spec in specList exists
// and assigns it the correct specsDir in which it exists
func validateSpecExistance(specList []string, specListFile string) (specsDir string, err error) {
	fmt.Println("Checking if spec exists for (%s)", specList)
	specListSet, err := packagelist.ParsePackageListFile(*specListFile)
	if err != nil {
		fmt.Println("error in parsing package list file %v", err)
		return nil, err
	}
	
	// TODO: currently, we have a limitation that all specs to be built must be present in the same specsDir
	for _, specsDir := range azlSpecsDirs {
		myString, err := specreaderutils.FindSpecFiles(specsDir, specListSet)
		if err != nil {
			fmt.Println("error in FindSpecFiles %v", err)
			return nil, err
		} else {
			fmt.Println("ERROR freeeeee ", myString[0])
			return specsDir, nil
		}
	}	
	fmt.Println("done with specreader ")
}

// write space delimited spec to file and send back filename
func getSpecListFile(spec string) (packageListFileName string, err error) {

	return packageListFileName, nil
}

