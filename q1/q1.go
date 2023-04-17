package q1

import "fmt"

func CalculateDiscount(currentPurchase float64, purchaseHistory []float64) (float64, error) {
	var media float64
	var soma float64
	for i := 0; i < len(purchaseHistory); i++ {
		soma += purchaseHistory[i]
	}
	media = soma / float64(len(purchaseHistory))
	if currentPurchase <= 0 {
		return 0, fmt.Errorf("valor da compra inválido")
	} else if media > 1000 {
		return currentPurchase * 0.2, nil
	} else if soma > 1000 {
		return currentPurchase * 0.1, nil
	} else if len(purchaseHistory) == 0 {
		return currentPurchase * 0.1, nil
	} else if soma <= 1000 && soma > 500 {
		return currentPurchase * 0.05, nil
	} else if soma <= 500 {
		return currentPurchase * 0.02, nil
	}
	return 0, nil
}
