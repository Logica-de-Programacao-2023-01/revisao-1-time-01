package q5

import "fmt"

func ConvertTemperature(temp float64, fromScale string, toScale string) (float64, error) {
	var temperaturaFinal float64

	if fromScale == "C" && toScale == "F" {
		temperaturaFinal = temp*9/5 + 32
	} else if fromScale == "C" && toScale == "K" {
		temperaturaFinal = temp + 273.15
	} else if fromScale == "F" && toScale == "C" {
		temperaturaFinal = (temp - 32) * 5 / 9
	} else if fromScale == "F" && toScale == "K" {
		temperaturaFinal = (temp-32)*5/9 + 273.15
	} else if fromScale == "K" && toScale == "C" {
		temperaturaFinal = temp - 273.15
	} else if fromScale == "K" && toScale == "F" {
		temperaturaFinal = (temp-273.15)*9/5 + 32
	} else {
		return 0, fmt.Errorf("escala inválida")
	}

	return temperaturaFinal, nil
}
