package main

import (
	"fmt"
	"math"
)

func gauss(matrix, aMatrix [][]float64) ([]float64, []float64, [][]float64, int) {

	timesSwitched := 0

	//Приведение к треугольному виду
	for i := 0; i < len(aMatrix)-1; i++ {
		curMax := 0.0
		curMaxInd := 0
		for j := i; j < len(aMatrix); j++ {
			if math.Abs(matrix[j][i]) > math.Abs(curMax) {
				curMax = math.Abs(matrix[j][i])
				curMaxInd = j
			}
		}
		for w := 0; w < len(matrix[i]); w++ {
			if matrix[i][w] != matrix[curMaxInd][w] {
				timesSwitched++
				break
			}
		}
		matrix[i], matrix[curMaxInd] = matrix[curMaxInd], matrix[i]
		b := matrix[i][i]
		for k := i + 1; k < len(matrix); k++ {
			a := matrix[k][i]
			for t := i; t < len(matrix); t++ {
				matrix[k][t] -= matrix[i][t] * a / b
			}
		}
	}

	fmt.Println("\nТреугольная матрица :")
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix); j++ {
			aMatrix[i][j] = matrix[i][j]
		}
		fmt.Println(matrix[i])
	}

	answ := make([]float64, len(matrix))
	var sum float64
	for i := len(matrix) - 1; i >= 0; i-- {
		if i == (len(matrix) - 1) {

			answ[i] = matrix[i][i+1] / matrix[i][i]
		} else {
			sum = matrix[i][len(matrix)]
			for j := i + 1; j < len(matrix); j++ {
				sum -= answ[j] * matrix[i][j]
			}
			answ[i] = sum / matrix[i][i]
		}
	}

	dif := make([]float64, len(matrix))
	for i := 0; i < len(matrix); i++ {
		sum = 0
		for j := 0; j < len(matrix); j++ {
			sum += matrix[i][j] * answ[j]
		}
		dif[i] = sum - matrix[i][len(matrix)]
	}

	return answ, dif, aMatrix, timesSwitched
}
