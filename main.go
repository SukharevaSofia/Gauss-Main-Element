// Метод Гаусса с выбором главного элемента по столбцам
package main

import "fmt"

func main() {
	var matrix, aMatrix [][]float64
	var answ, dif []float64
	var det float64
	var timesSwitched int
	matrix, aMatrix = input()
	fmt.Println("***\nМатрица:")
	for i := 0; i < len(matrix); i++ {
		fmt.Println(matrix[i])
	}
	answ, dif, aMatrix, timesSwitched = gauss(matrix, aMatrix)

	det = determinant(aMatrix, timesSwitched)
	fmt.Println("\nОпределитель =", det)

	fmt.Println("\nРешение:")
	for i := 0; i < len(answ); i++ {
		fmt.Printf("x%d=", i+1)
		fmt.Print(answ[i])
		fmt.Print("\n")
	}
	fmt.Println("\nНевязки:")
	for i := 0; i < len(dif); i++ {
		fmt.Printf("r%d=", i+1)
		fmt.Print(dif[i])
		fmt.Print("\n")
	}

}
