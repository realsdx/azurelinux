// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package interfaceutils

import (
	"fmt"
)

func BuildToolchain(rebuild bool) (err error) {
	fmt.Println("Building toolchain ...")
	// check if we need to build toolchain
	var shouldBuild bool
	shouldBuild = buildStatusToolchain(rebuild)
	if shouldBuild {
		err = buildToolchain()
		if err != nil {
			fmt.Println("failed to build toolchain %v", err)
			return err
		}
	}
}

