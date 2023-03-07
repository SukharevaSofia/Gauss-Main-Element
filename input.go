package main

//TODO: ввод матрицы с клавиатуры
import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func input() ([][]float64, [][]float64) {
	var inputSize string
	var matrixSize int
	var inputType string

	fmt.Print("\n***\nСухарева Софья \nP32131, ИСУ:334641" +
		"\nМетод Гаусса с выбором\nглавного элемента\n***\n")

	//SIZE INPUT
	fmt.Println("Введите размер матрицы. " +
		"Для ввода через файл введите \"F\", для ввода через терминал введите \"T\"")
	fmt.Print("F/T: ")
	for {
		fmt.Scan(&inputSize)
		if inputSize == "T" || inputSize == "t" || inputSize == "F" || inputSize == "f" {
			break
		}
		fmt.Println("Выберите способ ввода: \"F\" или \"T\"")
		fmt.Print("F/T: ")
	}
	if inputSize == "F" || inputSize == "f" {
		f, _ := os.Open("size.txt")
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			line := scanner.Text()
			matrixSize, _ = strconv.Atoi(line)
			break
		}

		fmt.Printf("Matrix size is %d", matrixSize)

	} else {
		fmt.Println("Размер матрицы должен быть целым числом в пределах [2, 20]. \n" +
			"При вводе дробного числа будет использована только целая его часть.")
		fmt.Print("Раззмер матрицы: ")

		for {
			fmt.Scan(&matrixSize)
			if matrixSize <= 20 && matrixSize > 0 {
				break
			}
			fmt.Println("Размер матрицы должен быть целым числом в пределах [2, 20].")
			fmt.Print("Размер матрицы: ")
		}
	}
	//MATRIX INPUT
	fmt.Println("\n*** \n Введите матрицу. " +
		"Для ввода через файл введите \"F\", для ввода через терминал введите \"T\"")
	fmt.Print("F/T: ")
	for {
		fmt.Scan(&inputType)
		if inputType == "T" || inputType == "t" || inputType == "F" || inputType == "f" {
			break
		}
		fmt.Println("Выберите способ ввода: \"F\" или \"T\"")
		fmt.Print("F/T: ")
	}

	matrix := make([][]float64, matrixSize)
	aMatrix := make([][]float64, matrixSize)
	bMatrix := make([]float64, matrixSize)
	counter := 0

	if inputType == "F" || inputType == "f" {
		f, _ := os.Open("matrix.txt")
		scanner := bufio.NewScanner(f)

		for scanner.Scan() {
			line := strings.Split(scanner.Text(), " ")
			if len(line) != (matrixSize + 1) {
				fmt.Printf("Несовпадение с указанным"+
					"размером матрицы (%d != %d).", len(line), matrixSize)
				os.Exit(1)
			}
			matrix[counter] = make([]float64, matrixSize+1)
			aMatrix[counter] = make([]float64, matrixSize)
			for j := 0; j < matrixSize+1; j++ {
				matrix[counter][j], _ = strconv.ParseFloat(line[j], 64)
				if j == matrixSize {
					bMatrix = append(bMatrix, matrix[counter][j])
				} else {
					aMatrix[counter][j] = matrix[counter][j]
				}
			}
			counter++
		}
	} else {
		var line string
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			line = scanner.Text()
			line := strings.Split(line, " ")
			if len(line) != (matrixSize + 1) {
				fmt.Printf("Несовпадение с указанным размером"+
					"матрицы (%d != %d).", len(line), matrixSize)
				os.Exit(1)
			}
			for _, el := range line {
				number, _ := regexp.MatchString("[0-9]+", el)
				fnumber, _ := regexp.MatchString("[0-9]+.[0-9]+", el)
				if !(number || fnumber) {
					fmt.Println("Неверные входные данные :(")
					os.Exit(1)
				}
			}
			matrix[counter] = make([]float64, matrixSize+1)
			aMatrix[counter] = make([]float64, matrixSize)
			for j := 0; j < matrixSize+1; j++ {
				matrix[counter][j], _ = strconv.ParseFloat(line[j], 64)
				if j == matrixSize {
					bMatrix = append(bMatrix, matrix[counter][j])
				} else {
					aMatrix[counter][j] = matrix[counter][j]
				}
			}
			counter++
			if counter == matrixSize {
				break
			}
		}

	}
	return matrix, aMatrix
}
