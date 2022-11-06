package steps

import (
	"regexp"
)

func ProcessInputLine(line string, inputs []string) []string {
	result, _ := regexp.MatchString(`^input`, line)
	if result == false {
		return nil
	}

	r, _ := regexp.Compile(" [a-z]* ?")

	inputs = append(inputs, r.FindString(line))

	return inputs
}
