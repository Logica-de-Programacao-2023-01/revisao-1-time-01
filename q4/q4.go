package q4

import "fmt"

func CalculateFinalPrice(basePrice float64, state string, taxCode int) (float64, error) {
	var precoFinal float64
	var frete float64
	var imposto float64

	if basePrice <= 0 {
		return 0, fmt.Errorf("preço base inválido")
	}

	switch state {
	case "SP":
		frete = 0.10
	case "RJ":
		frete = 0.15
	case "MG":
		frete = 0.20
	case "ES":
		frete = 0.25
	default:
		frete = 0.30
	}

	switch taxCode {
	case 1:
		imposto = 0.05
	case 2:
		imposto = 0.10
	case 3:
		imposto = 0.15
	default:
		return 0, fmt.Errorf("imposto não encontrado")
	}

	precoFinal = basePrice + basePrice*frete + basePrice*imposto

	return precoFinal, nil
}
