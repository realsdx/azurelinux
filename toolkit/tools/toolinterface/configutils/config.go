// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package configutils

import (
	"fmt"
	"os"
	"strings"

	"github.com/microsoft/azurelinux/toolkit/tools/internal/file"

)

var config map[string]string
const configFile = "configutils/config.txt"

func PopulateConfigFromFile() (err error) {
	config = make(map[string]string)
	// TODO: get base_dir from pwd
	wd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get working directory:\n%w", err)
	}
	base_dir := strings.Split(wd, "toolkit")[0]
	fmt.Println("base is ", base_dir)
	fmt.Println("wd is ", wd)
	SetConfig("PROJECT_ROOT", base_dir)

	lines, err := file.ReadLines(wd+"/"+configFile)
	fmt.Println("opened file: ", len(lines))
	if err != nil {
		return fmt.Errorf("failed to open file:\n%w", err)
	}
	for _, line := range lines {
		fmt.Println("line is", line)
		entry := strings.Split(line,":")
		if len(entry) != 2 {
			fmt.Println("not a config entry", entry[0])
			continue
		}
		entry[1] = strings.Replace(entry[1], "<PROJECT_ROOT>/", base_dir, -1)
		SetConfig(entry[0], entry[1])
		fmt.Println("entry is is", entry[0], ":",entry[1] )
		i,_ := GetConfig(entry[0])
		fmt.Println("returnied ",i)
	}
	return
}

func SetConfig(key, val string) {
	config[key] = val
	return
}

func GetConfig(key string) (val string, err error) {
	val, exists := config[key]
	if !exists {
		err = fmt.Errorf("key does not exist")
	}
	return
}
