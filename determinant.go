package main

import (
	"fmt"
)

func determinant(matrix [][]float64, timesSwitched int) float64 {
	if len(matrix) == 1 {
		fmt.Println(matrix[0][0])
		return 0
	} else {
		det := 1.0
		for i := 0; i < len(matrix); i++ {
			det *= matrix[i][i]
		}
		if timesSwitched%2 != 0 {
			det *= -1
		}
		return det
	}
}
