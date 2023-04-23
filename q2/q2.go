package q2

import (
	"fmt"
	"strconv"
	"strings"
)

func AverageLettersPerWord(text string) (float64, error) {
	_, err := strconv.Atoi(text)
	if len(text) == 0 || err == nil || text == " " {
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
