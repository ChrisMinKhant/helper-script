package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

type command struct {
	namespace    *string
	serviceName  *string
	dstContext   *string
	dstNamespace *string
}

func main() {
	commandInstance := &command{
		namespace:    flag.String("n", "", "namespace"),
		serviceName:  flag.String("s", "", "service name"),
		dstContext:   flag.String("dc", "", "destination context"),
		dstNamespace: flag.String("dn", "", "destination namespace"),
	}

	flag.Parse()

	if !validateCommandInstance(*commandInstance) {
		panic("You provided invalid input!")
	}

	commandOutput, _ := exec.Command("whoami").Output()
	systemUsername := strings.TrimSuffix(string(commandOutput), "\n")

	createdServiceFile := createServiceFile("/home/" + systemUsername + "/" + *commandInstance.serviceName + ".yaml")

	reterivedServiceYaml, error := exec.Command("kubectl", "-n", *commandInstance.namespace, "get", "svc", *commandInstance.serviceName, "-o", "yaml").Output()

	if error != nil {
		panic("Fetched error at reteriving service file ::: " + error.Error())
	}

	yamlValue := make(map[string]any)

	yaml.Unmarshal(reterivedServiceYaml, yamlValue)

	delete(yamlValue["metadata"].(map[any]any), "creationTimestamp")
	delete(yamlValue["metadata"].(map[any]any), "resourceVersion")
	delete(yamlValue["metadata"].(map[any]any), "uid")
	delete(yamlValue["spec"].(map[any]any), "clusterIP")
	delete(yamlValue["spec"].(map[any]any), "clusterIPs")
	delete(yamlValue["spec"].(map[any]any), "internalTrafficPolicy")
	delete(yamlValue["spec"].(map[any]any), "ipFamilyPolicy")
	delete(yamlValue["spec"].(map[any]any), "ipFamilies")
	delete(yamlValue, "status")

	yamlByteValue, error := yaml.Marshal(yamlValue)

	if error != nil {
		panic("Fetched error at marshalling yaml file ::: " + error.Error())
	}

	_, error = createdServiceFile.Write(yamlByteValue)

	if error != nil {
		panic("Fetched error at writing service file to local ::: " + error.Error())
	}

	_, error = exec.Command("kubectl", "config", "use-context", *commandInstance.dstContext).Output()

	if error != nil {
		panic("Fetched error at kubectl switch context ::: " + error.Error())
	}
	serviceFilePath := filepath.Join("/home", systemUsername, *commandInstance.serviceName+".yaml")

	commandOutput, error = exec.Command("kubectl", "apply", "-f", serviceFilePath, "-n", *commandInstance.dstNamespace).Output()

	if error != nil {
		panic("Fetched error at kubectl applly service file ::: " + error.Error())
	}

	fmt.Println(string(commandOutput))
	fmt.Printf("Your service file had been applied!")
}

func validateCommandInstance(commandInstance command) bool {

	if *commandInstance.namespace != "" {
		if *commandInstance.serviceName != "" {
			if *commandInstance.dstNamespace != "" {
				return true
			}

			return false
		}

		return false
	}

	return false
}

func createServiceFile(path string) *os.File {
	createdServiceFile, error := os.Create(path)

	if error != nil {
		panic("Fetched error at creating service file ::: " + error.Error())
	}

	return createdServiceFile
}
