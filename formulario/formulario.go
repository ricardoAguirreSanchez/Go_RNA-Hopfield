package formulario

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetInput(c *gin.Context) Formulario {
	fmt.Println("Matriz ingresada: ")
	formularioInput := Formulario{}
	formularioInput.Matriz = nuevaMatriz()
	for i := 0; i < 16; i++ {
		for j := 0; j < 16; j++ {
			value := strconv.Itoa(i) + strconv.Itoa(j)
			formularioInput.Matriz[i][j] = c.Request.FormValue(value)
		}
	}
	PrintMatrix(formularioInput.Matriz)
	return formularioInput
}

func nuevaMatriz() [][]string {
	animals := [][]string{}
	for i := 0; i < 16; i++ {
		row1 := []string{"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""}
		animals = append(animals, row1)
	}
	return animals
}

/*
	printMatrix: printea matrices de 10 x 10
*/
func PrintMatrix(matrizLinterna [][]string) {
	matrizLinternaAux := matrizLinterna
	for i, row := range matrizLinternaAux {
		for j, value := range row {
			if value == "" {
				matrizLinternaAux[i][j] = " "
			}
		}
		fmt.Println(row)
	}
}
