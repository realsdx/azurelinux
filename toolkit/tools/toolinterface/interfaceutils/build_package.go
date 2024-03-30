// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package interfaceutils

import (
	"fmt"

	packagelist "github.com/microsoft/azurelinux/toolkit/tools/internal/packlist"
	"github.com/microsoft/azurelinux/toolkit/tools/pkg/specreaderutils"
	"github.com/microsoft/azurelinux/toolkit/tools/toolinterface/configutils"
)

var (
	azlSpecsDirs = [...] string {"SPECS", "SPECS-EXTENDED", "SPECS-SIGNED"}
	// get relevant configs
	toolkit_dir string
	project_dir string
)

func BuildPackage(spec string) (err error) {
	// build global config map
	configutils.PopulateConfigFromFile()
	toolkit_dir, _ = configutils.GetConfig("toolkit_root")
	project_dir, _ = configutils.GetConfig("PROJECT_ROOT")

	fmt.Println("[debug] Building packages: specs are (%s)", spec)

	// check specs exist
	specsDir, err := validateSpecExistance(spec)
	if err != nil {
		err = fmt.Errorf("failed to validate specs:\n%w", err)
		return err
	}

	// TODO: set sepcs dir in config

	// any other checks

	// build toolchain if required

	// put toolchain rpms into toolchain_archive and use it

	// build tools if required

	// set extra configs

	// show dependency graph - use graphanalytics tool

	// build package
	err = buildSpecs(spec, specsDir)
	if err != nil {
		err = fmt.Errorf("failed to build specs:\n%w", err)
		return err
	}

	// show output

	return
}

// validateSpecExistance checks if each spec in specList exists
// If the spec exists, it assigns it the correct specsDir
func validateSpecExistance(specList string) (specsDir string, err error) {
	fmt.Println("[debug] Checking if spec exists for (%s)", specList)
	specMap, err := packagelist.ParsePackageList(specList)
	if err != nil {
		err = fmt.Errorf("failed to parse package list:\n%w", err)
		return
	}

	// TODO: currently, we have a limitation that all specs to be built must be present in the same specsDir
	// TODO: return error only if spec is not found in any specsDir
	for _, specsDir := range azlSpecsDirs {
		specFiles, err := specreaderutils.FindSpecFiles(project_dir+specsDir, specMap)
		if err != nil {
			err = fmt.Errorf("failed to FindSpecFiles:\n%w", err)
			return "", err
		} else {
			fmt.Println("[debug] done with specreader, returned specFiles (%s)", specFiles)
			return specsDir, nil
		}
	}
	fmt.Println("[debug] done with specreader")
	return
}

func buildSpecs (specs, specsDir string) (err error) {
	// TODO: use a command builder
	// TODO: some of these arguments can be removed if/when tools start reading directly from config
	srpm_pack_list := "SRPM_PACK_LIST=\""
	srpm_pack_list +=specs
	srpm_pack_list +="\""

	err = execCommands("make",
		toolkit_dir,
		"build-packages",
		"REBUILD_TOOLS=y",
		srpm_pack_list)

	if err != nil {
		fmt.Println(err)
	}
	return
}
