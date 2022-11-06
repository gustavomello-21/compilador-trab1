package steps

import (
	"regexp"
)

func ProcessOutputLine(line string, outputs []string) []string {
	result, _ := regexp.MatchString(`^output`, line)

	if result == false {
		return nil
	}

	r, _ := regexp.Compile(" [a-z]* ?")

	outputs = append(outputs, r.FindString(line))

	return outputs
}
