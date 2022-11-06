package steps

func ProcessLine(line string, inputs, outputs []string) ([]string, []string) {
	inputs = ProcessInputLine(line, inputs)
	outputs = ProcessOutputLine(line, outputs)

	return inputs, outputs
}
