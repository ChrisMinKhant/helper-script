package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/fs"
	"os"

	"gopkg.in/yaml.v3"
)

func main() {
	var path string
	var systemUserName string

	asciiArt := `	·············································································
	:  ____ _____  ____  ____    _       ____ _____ _____ ____ _   _    _ ____  :
	: / _  |____ |/ _  |/ _  |  / \     / _  |____ |_   _|___ | | | |  | |___ \ :
	:| (_| | |_  | | | | | | | / _ \   | (_| | |_  | | | / ___| | | |  | |   | |:
	: > _  |___| | |_| | |_| |/ ___ \   > _  |___| | | || (___| |_| ___| |___| |:
	:/_/ |_|_____|\____|\____/_/   \_\ /_/ |_|_____| |_| \____|\___|_____|____/ :
	·············································································`

	fmt.Println(asciiArt)

	fmt.Print("Please enter system username ::: ")
	fmt.Scan(&systemUserName)
	fmt.Print("Please enter downloaded rancher yaml file path ::: ")
	fmt.Scan(&path)

	// Take a backup of the existed config file
	backupFile(&systemUserName)

	fetchedFiles := fetchFiles(&path)

	// Add cluster config to the existed config file
	for _, singleFile := range *fetchedFiles {
		singleFileAbsolutePath := path + "/" + singleFile.Name()
		addCluster(&singleFileAbsolutePath, &systemUserName)
	}

	fmt.Println("I got go. You got shits to be done, HE HE HE.")
}

func addCluster(path *string, systemUsername *string) {
	downloadedConfig := make(map[string]any)
	existedConfig := make(map[string]any)

	openedDownloadedConfig, error := os.ReadFile(*path)
	checkError(&error)

	openedExistedConfig, error := os.ReadFile("/home/" + *systemUsername + "/.kube/config")
	checkError(&error)

	yaml.Unmarshal(openedDownloadedConfig, downloadedConfig)
	yaml.Unmarshal(openedExistedConfig, existedConfig)

	foundFactor := 0

	for _, cluster := range existedConfig["clusters"].([]any) {
		if checkClusterExistance(&cluster, &downloadedConfig["clusters"].([]any)[0]) {
			foundFactor++
		}
	}

	if foundFactor == 0 {
		existedConfig["clusters"] = append(existedConfig["clusters"].([]any), downloadedConfig["clusters"].([]any)[0])
		existedConfig["users"] = append(existedConfig["users"].([]any), downloadedConfig["users"].([]any)[0])
		existedConfig["contexts"] = append(existedConfig["contexts"].([]any), downloadedConfig["contexts"].([]any)[0])

		marshalledYamlValue := marshalYaml(existedConfig)

		error = os.WriteFile("/home/"+*systemUsername+"/.kube/config", *marshalledYamlValue, 0644)
		checkError(&error)
	}
}

func checkClusterExistance(existedClusterInfo *any, downloadedClusterInfo *any) bool {
	existedClusterInfoBytes, error := json.Marshal(*existedClusterInfo)
	checkError(&error)

	downloadedClusterInfoBytes, error := json.Marshal(*downloadedClusterInfo)
	checkError(&error)

	return bytes.Equal(existedClusterInfoBytes, downloadedClusterInfoBytes)
}

func backupFile(systemUsername *string) {
	backupconfig := make(map[string]any)

	openedExistedConfig, error := os.ReadFile("/home/" + *systemUsername + "/.kube/config")
	checkError(&error)

	yaml.Unmarshal(openedExistedConfig, backupconfig)

	marshalledBackupYamlValue := marshalYaml(backupconfig)

	error = os.WriteFile("/home/"+*systemUsername+"/.kube/backupconfig", *marshalledBackupYamlValue, 0644)
	checkError(&error)
}

func fetchFiles(path *string) *[]fs.DirEntry {
	foundFiles, error := os.ReadDir(*path)
	checkError(&error)

	return &foundFiles
}

func checkError(foundError *error) {
	if *foundError != nil {
		panic(*foundError)
	}
}

func marshalYaml(rawYaml any) *[]byte {
	marshalledYaml, error := yaml.Marshal(rawYaml)
	checkError(&error)
	return &marshalledYaml
}
