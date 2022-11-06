package main

import (
	"github.com/gustavomello-21/modules/steps"
	"github.com/gustavomello-21/modules/utils"
)

var (
	fileLines []string
	inputs    []string
	outputs   []string
)

func main() {

	fileLines = utils.GetFileContent()

	for _, value := range fileLines {
		for _, a := range steps.ProcessInputLine(value, inputs) {
			inputs = append(inputs, a)
		}
		for _, b := range steps.ProcessOutputLine(value, outputs) {
			outputs = append(outputs, b)
		}
	}

	inputs = clean_array(inputs)
	outputs = clean_array(outputs)
}

func clean_array(array []string) []string {
	copy(array[0:], array[0+1:])
	array[len(array)-1] = ""
	array = array[:len(array)-1]

	return array
}
