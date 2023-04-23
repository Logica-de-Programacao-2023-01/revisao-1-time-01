package q3

import "fmt"

func FindMinMaxAverage(numbers []int) (int, int, float64, error) {
	if len(numbers) == 0 {
		return 0, 0, 0, fmt.Errorf("Lista vazia")
	}
	maior := -999999999999
	menor := 9999999999999
	var soma int
	var media float64
	if len(numbers) == 1 {
		maior = numbers[0]
		menor = numbers[0]
		media = float64(numbers[0])
		return maior, menor, media, nil
	}
	for i := 0; i < len(numbers); i++ {
		if maior < numbers[i] {
			maior = numbers[i]
		}
		if menor > numbers[i] {
			menor = numbers[i]
		}
	}
	j := 0
	for i := 0; i < len(numbers); i++ {
		if maior == numbers[i] {
			continue
		}
		if menor == numbers[i] {
			continue
		}
		soma += numbers[i]
		numbers[j] = numbers[i]
		j++
	}
	media = float64(soma) / float64(len(numbers)-2)
	return maior, menor, media, nil
}
