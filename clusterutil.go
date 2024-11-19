package main

import (
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

func addCluster(path *string, systemUsername *string) {

	downloadedConfig := make(map[string]any)
	existedConfig := make(map[string]any)

	openedDownloadedConfig, error := os.ReadFile(*path)
	checkError(&error)

	if !configFileExist(systemUsername) {
		createConfigFile(systemUsername)
	} else {
		backupFile(systemUsername)
	}

	openedExistingConfig, error := os.ReadFile("/home/" + *systemUsername + "/.kube/config")
	checkError(&error)

	yaml.Unmarshal(openedDownloadedConfig, downloadedConfig)
	yaml.Unmarshal(openedExistingConfig, existedConfig)

	if len(existedConfig) != 0 {

		existedConfig["clusters"] = addBlock(existedConfig["clusters"].([]any), &downloadedConfig["clusters"].([]any)[0])
		existedConfig["users"] = addBlock(existedConfig["users"].([]any), &downloadedConfig["users"].([]any)[0])
		existedConfig["contexts"] = addBlock(existedConfig["contexts"].([]any), &downloadedConfig["contexts"].([]any)[0])

		marshalledYamlValue := marshalYaml(existedConfig)

		error = os.WriteFile("/home/"+*systemUsername+"/.kube/config", *marshalledYamlValue, 0644)
		checkError(&error)

		return
	}

	marshalledYamlValue := marshalYaml(downloadedConfig)

	error = os.WriteFile("/home/"+*systemUsername+"/.kube/config", *marshalledYamlValue, 0644)
	checkError(&error)

}

func configFileExist(systemUsername *string) bool {

	_, error := os.ReadFile("/home/" + *systemUsername + "/.kube/config")

	if error != nil {
		if strings.ContainsAny(error.Error(), "no such file or directory") {
			return false
		}
		checkError(&error)
	}

	return true
}

func createConfigFile(systemUsername *string) {
	_, error := os.Create("/home/" + *systemUsername + "/.kube/config")
	checkError(&error)
}
