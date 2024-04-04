package main

import (
	"fmt"
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
	fetchFile(&path, &systemUserName)
}

func fetchFile(path *string, systemUsername *string) {
	defer recover()

	downloadedConfig := make(map[string]any)
	existedConfig := make(map[string]any)
	backupconfig := make(map[string]any)

	openedDownloadedConfig, error := os.ReadFile(*path)
	checkError(&error)

	openedExistedConfig, error := os.ReadFile("/home/" + *systemUsername + "/.kube/config")
	checkError(&error)

	yaml.Unmarshal(openedDownloadedConfig, downloadedConfig)
	yaml.Unmarshal(openedExistedConfig, existedConfig)
	yaml.Unmarshal(openedExistedConfig, backupconfig)

	marshalledBackupYamlValue := marshalYaml(backupconfig)

	error = os.WriteFile("/home/"+*systemUsername+"/.kube/backupconfig", *marshalledBackupYamlValue, 0644)
	checkError(&error)

	existedConfig["clusters"] = append(existedConfig["clusters"].([]any), downloadedConfig["clusters"].([]any)[0])
	existedConfig["users"] = append(existedConfig["users"].([]any), downloadedConfig["users"].([]any)[0])
	existedConfig["contexts"] = append(existedConfig["contexts"].([]any), downloadedConfig["contexts"].([]any)[0])

	marshalledYamlValue := marshalYaml(existedConfig)

	error = os.WriteFile("/home/"+*systemUsername+"/.kube/config", *marshalledYamlValue, 0644)
	checkError(&error)

	fmt.Println("New cluster is added to config. Don't be happy, bye.")
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
