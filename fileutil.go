package main

import (
	"io/fs"
	"os"

	"gopkg.in/yaml.v3"
)

func backupFile(systemUsername *string) {
	backupconfig := make(map[string]any)

	openedExistedConfig, error := os.ReadFile("/home/" + *systemUsername + "/.kube/config")
	checkError(&error)

	if openedExistedConfig != nil {
		yaml.Unmarshal(openedExistedConfig, backupconfig)

		marshalledBackupYamlValue := marshalYaml(backupconfig)

		error = os.WriteFile("/home/"+*systemUsername+"/.kube/backupconfig", *marshalledBackupYamlValue, 0644)
		checkError(&error)
	}

}

func fetchFiles(path *string) *[]fs.DirEntry {
	foundFiles, error := os.ReadDir(*path)
	checkError(&error)

	return &foundFiles
}
