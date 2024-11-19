package main

import (
	"bytes"
	"encoding/json"

	"gopkg.in/yaml.v3"
)

func marshalYaml(rawYaml any) *[]byte {
	marshalledYaml, error := yaml.Marshal(rawYaml)
	checkError(&error)
	return &marshalledYaml
}

func compareBlocks(sourceBlock *any, dstBlocks *any) bool {
	sourceBlockBytes, error := json.Marshal(*sourceBlock)
	checkError(&error)

	dstBlocksBytes, error := json.Marshal(*dstBlocks)
	checkError(&error)

	return bytes.Equal(sourceBlockBytes, dstBlocksBytes)
}

func blockExists(sourceBlocks *[]any, dstBlock *any) bool {

	for _, sourceBlock := range *sourceBlocks {
		if compareBlocks(&sourceBlock, dstBlock) {
			return true
		}
	}

	return false
}

func addBlock(sourceBlock []any, dstBlock *any) []any {
	if !blockExists(&sourceBlock, dstBlock) {
		return append(sourceBlock, *dstBlock)
	}
	return sourceBlock
}
