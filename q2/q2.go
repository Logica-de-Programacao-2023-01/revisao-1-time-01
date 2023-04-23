package q2

import (
	"fmt"
	"strings"
)

func AverageLettersPerWord(text string) (float64, error) {
	if len(strings.TrimSpace(text)) == 0 {
		return 0, fmt.Errorf("texto vazio")
	} else {
		palavras := strings.Split(text, " ")
		var letras int
		for i := 0; i < len(palavras); i++ {
			letras += len(palavras[i])
		}
		media := float64(letras) / float64(len(palavras))
		return media, nil
	}
}
