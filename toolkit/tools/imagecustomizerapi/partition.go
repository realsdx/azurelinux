// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package imagecustomizerapi

import (
	"fmt"
	"unicode"

	"github.com/microsoft/azurelinux/toolkit/tools/internal/sliceutils"
)

type Partition struct {
	// ID is used to correlate `Partition` objects with `PartitionSetting` objects.
	Id string `yaml:"id"`
	// Name is the label to assign to the partition.
	Label string `yaml:"label"`
	// Start is the offset where the partition begins (inclusive), in MiBs.
	Start uint64 `yaml:"start"`
	// End is the offset where the partition ends (exclusive), in MiBs.
	End *uint64 `yaml:"end"`
	// Size is the size of the partition in MiBs.
	Size *uint64 `yaml:"size"`
	// Flags assigns features to the partition.
	Flags []PartitionFlag `yaml:"flags"`
}

func (p *Partition) IsValid() error {
	err := isGPTNameValid(p.Label)
	if err != nil {
		return err
	}

	if p.End != nil && p.Size != nil {
		return fmt.Errorf("cannot specify both end and size on partition (%s)", p.Id)
	}

	if (p.End != nil && p.Start >= *p.End) || (p.Size != nil && *p.Size <= 0) {
		return fmt.Errorf("partition's (%s) size can't be 0 or negative", p.Id)
	}

	for _, f := range p.Flags {
		err := f.IsValid()
		if err != nil {
			return err
		}
	}

	if p.IsBiosBoot() {
		if p.Start != 1 {
			return fmt.Errorf("BIOS boot partition must start at block 1")
		}
	}

	return nil
}

func (p *Partition) GetEnd() (uint64, bool) {
	if p.End != nil {
		return *p.End, true
	}

	if p.Size != nil {
		return p.Start + *p.Size, true
	}

	return 0, false
}

func (p *Partition) IsESP() bool {
	return sliceutils.ContainsValue(p.Flags, PartitionFlagESP)
}

func (p *Partition) IsBiosBoot() bool {
	return sliceutils.ContainsValue(p.Flags, PartitionFlagBiosGrub)
}

// isGPTNameValid checks if a GPT partition name is valid.
func isGPTNameValid(name string) error {
	// The max partition name length is 36 UTF-16 code units, including a null terminator.
	// Since we are also restricting the name to ASCII, this means 35 ASCII characters.
	const maxLength = 35

	// Restrict the name to only ASCII characters as some tools (e.g. parted) work better
	// with only ASCII characters.
	for _, char := range name {
		if char > unicode.MaxASCII {
			return fmt.Errorf("partition name (%s) contains a non-ASCII character (%c)", name, char)
		}
	}

	if len(name) > maxLength {
		return fmt.Errorf("partition name (%s) is too long", name)
	}

	return nil
}
