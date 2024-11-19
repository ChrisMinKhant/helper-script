package main

import (
	"flag"
	"fmt"
	"os/exec"
	"strings"
)

var banner = `
#  $$$$$$\                            $$\                                                                                 
#  \_$$  _|                           $$ |                                                                                
#    $$ |         $$$$$$\   $$$$$$\ $$$$$$\          $$$$$$\   $$$$$$\                                                    
#    $$ |        $$  __$$\ $$  __$$\\_$$  _|        $$  __$$\ $$  __$$\                                                   
#    $$ |        $$ /  $$ |$$ /  $$ | $$ |          $$ /  $$ |$$ /  $$ |                                                  
#    $$ |        $$ |  $$ |$$ |  $$ | $$ |$$\       $$ |  $$ |$$ |  $$ |                                                  
#  $$$$$$\       \$$$$$$$ |\$$$$$$  | \$$$$  |      \$$$$$$$ |\$$$$$$  |$$\                                               
#  \______|       \____$$ | \______/   \____/        \____$$ | \______/ \__|                                              
#                $$\   $$ |                         $$\   $$ |                                                            
#                \$$$$$$  |                         \$$$$$$  |                                                            
#                 \______/                           \______/                                                             
#  $$\     $$\                                              $$\                     $$\       $$\   $$\                   
#  \$$\   $$  |                                             $$ |                    $$ |      \__|  $$ |                  
#   \$$\ $$  /$$$$$$\  $$\   $$\        $$$$$$\   $$$$$$\ $$$$$$\          $$$$$$$\ $$$$$$$\  $$\ $$$$$$\    $$$$$$$\     
#    \$$$$  /$$  __$$\ $$ |  $$ |      $$  __$$\ $$  __$$\\_$$  _|        $$  _____|$$  __$$\ $$ |\_$$  _|  $$  _____|    
#     \$$  / $$ /  $$ |$$ |  $$ |      $$ /  $$ |$$ /  $$ | $$ |          \$$$$$$\  $$ |  $$ |$$ |  $$ |    \$$$$$$\      
#      $$ |  $$ |  $$ |$$ |  $$ |      $$ |  $$ |$$ |  $$ | $$ |$$\        \____$$\ $$ |  $$ |$$ |  $$ |$$\  \____$$\     
#      $$ |  \$$$$$$  |\$$$$$$  |      \$$$$$$$ |\$$$$$$  | \$$$$  |      $$$$$$$  |$$ |  $$ |$$ |  \$$$$  |$$$$$$$  |$$\ 
#      \__|   \______/  \______/        \____$$ | \______/   \____/       \_______/ \__|  \__|\__|   \____/ \_______/ \__|
#                                      $$\   $$ |                                                                         
#                                      \$$$$$$  |                                                                         
#                                       \______/                                                                          
#  $$\   $$\ $$$$$$$$\       $$\   $$\ $$$$$$$$\       $$\   $$\ $$$$$$$$\                                                
#  $$ |  $$ |$$  _____|      $$ |  $$ |$$  _____|      $$ |  $$ |$$  _____|                                               
#  $$ |  $$ |$$ |            $$ |  $$ |$$ |            $$ |  $$ |$$ |                                                     
#  $$$$$$$$ |$$$$$\          $$$$$$$$ |$$$$$\          $$$$$$$$ |$$$$$\                                                   
#  $$  __$$ |$$  __|         $$  __$$ |$$  __|         $$  __$$ |$$  __|                                                  
#  $$ |  $$ |$$ |            $$ |  $$ |$$ |            $$ |  $$ |$$ |                                                     
#  $$ |  $$ |$$$$$$$$\       $$ |  $$ |$$$$$$$$\       $$ |  $$ |$$$$$$$$\ $$\                                            
#  \__|  \__|\________|      \__|  \__|\________|      \__|  \__|\________|\__|                                           
#                                                                                                                         
#                                                                                                                         
#`

type flags struct {
	clusterConfigsPath *string
}

func main() {

	inputFlags := &flags{
		clusterConfigsPath: flag.String("p", "", "path"),
	}

	flag.Parse()

	fetchedSystemUsername := fetchSystemUsername()
	fetchedClusterConfigFiles := fetchFiles(inputFlags.clusterConfigsPath)

	for _, clusterConfigFile := range *fetchedClusterConfigFiles {
		clusterConfigFilePath := *inputFlags.clusterConfigsPath + "/" + clusterConfigFile.Name()
		addCluster(&clusterConfigFilePath, &fetchedSystemUsername)
	}

	displayBanner()
}

func fetchSystemUsername() string {
	commandOutput, error := exec.Command("whoami").Output()
	checkError(&error)
	return strings.TrimSuffix(string(commandOutput), "\n")
}

func displayBanner() {
	fmt.Println(banner)
	fmt.Println("# CLUSTERS ARE ADDED.\n")
}
