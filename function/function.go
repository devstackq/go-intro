package function

import "errors"

func Hello(name string, age int64) (string, error) {
	if name == "Asel" && age < 18 {
		return "Hello little girl", nil
	}
	if name == "Dwayne" && age > 50 {
		return "Hi sugar daddy", nil
	}
	return "", errors.New("not found person, by name & age")
}

func Cast(value string) int {
	return 0
}

func clousere(value string) func(string) int {
	return Cast(value)
}
