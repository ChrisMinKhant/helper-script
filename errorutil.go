package main

import "log"

func checkError(foundError *error) {
	if *foundError != nil {
		log.Panic(*foundError)
	}
}
