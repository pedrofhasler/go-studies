package main

import (
	"fmt"
	"os"
	"strconv"
)

func main()  {
	entrada := os.Args[1:]
	numeros := make([]int, len(entrada))

	for i, n := range entrada {
		numero, err := strconv.Atoi(n)
		if err != nil {
			fmt.Printf("%s não é um número válido")
			os.Exit(1)
		}
		numeros[i] = numero
	}

	fmt.Println(quicksort(numeros))
}

func quicksort(numeros []int) []int {
	if len(numeros) <= 1 {
		return numeros
	}

	n := make([]int, len(numeros))
	copy(n, numeros)

	indicePivot := len(n)/2
	pivot := n[indicePivot]
	n = append(n[:indicePivot], n[indicePivot+1:]...)

	menores, maiores := particionar(n, pivot)

	return append(
		append(quicksort(menores), pivot),
		quicksort(maiores)...
	)
}

func particionar(numeros []int, pivot int) (menores []int, maiores []int) {
	for _, n := range numeros{
		if n <= pivot {
			menores = append(menores,n)
		} else {
			maiores = append(maiores, n)
		}
	}

	return menores, maiores
}