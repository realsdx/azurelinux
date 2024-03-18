// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package interfaceutils

import (
	"fmt"
	"strings"

)

// buildStatus checks if spec needs to be rebuilt
// reasons for rebuild:
//					- spec has never been built
//					- there is a change in spec from last build
//					- there is a change in toolchain manifest
//					- user wants to rebuild spec
func buildStatus() {

}

// buildStatusToolchain checks if toolchain should be (re)built
// returns true if
//					- toolchain rpms have never been built
//					- there is a change in toolchain spec(s)
//					- there is a change in toolchain manifest
//					- user wants to rebuild toolchain
func buildStatusToolchain() {
	buildStatus()
	// added if change in manifest files

}