package formulario

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetInput(c *gin.Context) Formulario {
	fmt.Println("GetInput - inicio ")
	formularioInput := Formulario{}
	formularioInput.Matriz = nuevaMatriz()
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			value := strconv.Itoa(i) + strconv.Itoa(j)
			formularioInput.Matriz[i][j] = c.Request.FormValue(value)

		}
	}
	fmt.Println("GetInput - fin ")
	return formularioInput
}

func nuevaMatriz() [][]string {
	animals := [][]string{}
	for i := 0; i < 10; i++ {
		row1 := []string{"", "", "", "", "", "", "", "", "", ""}
		animals = append(animals, row1)
	}
	return animals
}
