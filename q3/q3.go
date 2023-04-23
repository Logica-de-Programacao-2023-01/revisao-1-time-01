package q3

import "fmt"

func FindMinMaxAverage(numbers []int) (int, int, float64, error) {
	maior := -999999999999
	menor := 9999999999999
	var soma int
	var media float64
	if numbers == nil {
		return 0, 0, 0, fmt.Errorf("lista vazia")
	}
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
	for i := 0; i < len(numbers); i++ {
		if maior == numbers[i] {
			numbers = append(numbers[:i], numbers[i+1:]...)
			continue
		}
		if menor == numbers[i] {
			numbers = append(numbers[:i], numbers[i+1:]...)
			continue
		}
		soma += numbers[i]
	}
	media = float64(soma) / float64(len(numbers))
	return maior, menor, media, nil
}
