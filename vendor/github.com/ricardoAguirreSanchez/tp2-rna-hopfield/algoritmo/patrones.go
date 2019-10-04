package algoritmo

/*
	CrearMatrizLinterna: crea la matriz 10x10 para linterna verde
*/
func CrearMatrizLinterna() [][]string {

	matriz10x10 := [][]string{}
	row1 := []string{"", "", "x", "x", "x", "x", "x", "x", "", ""}
	row2 := []string{"", "", "x", "x", "x", "x", "x", "x", "", ""}
	row3 := []string{"", "", "", "", "x", "x", "", "", "", ""}
	row4 := []string{"", "", "", "x", "", "", "x", "", "", ""}
	row5 := []string{"", "", "x", "", "", "", "", "x", "", ""}
	row6 := []string{"", "", "x", "", "", "", "", "x", "", ""}
	row7 := []string{"", "", "", "x", "", "", "x", "", "", ""}
	row8 := []string{"", "", "", "", "x", "x", "", "", "", ""}
	row9 := []string{"", "", "x", "x", "x", "x", "x", "x", "", ""}
	row10 := []string{"", "", "x", "x", "x", "x", "x", "x", "", ""}

	// row1 := []string{"x", "x", "x", "", "", "", "", "", "", ""}
	// row2 := []string{"", "x", "", "", "", "", "", "", "", ""}
	// row3 := []string{"x", "x", "x", "", "", "", "", "", "", ""}
	// row4 := []string{"", "", "", "", "", "", "", "", "", ""}
	// row5 := []string{"", "", "", "", "", "", "", "", "", ""}
	// row6 := []string{"", "", "", "", "", "", "", "", "", ""}
	// row7 := []string{"", "", "", "", "", "", "", "", "", ""}
	// row8 := []string{"", "", "", "", "", "", "", "", "", ""}
	// row9 := []string{"", "", "", "", "", "", "", "", "", ""}
	// row10 := []string{"", "", "", "", "", "", "", "", "", ""}

	matriz10x10 = append(matriz10x10, row1)
	matriz10x10 = append(matriz10x10, row2)
	matriz10x10 = append(matriz10x10, row3)
	matriz10x10 = append(matriz10x10, row4)
	matriz10x10 = append(matriz10x10, row5)
	matriz10x10 = append(matriz10x10, row6)
	matriz10x10 = append(matriz10x10, row7)
	matriz10x10 = append(matriz10x10, row8)
	matriz10x10 = append(matriz10x10, row9)
	matriz10x10 = append(matriz10x10, row10)
	return matriz10x10
}

/*
	CrearMatrizFlash: crea la matriz 10x10 para flash
*/
func CrearMatrizFlash() [][]string {

	matriz10x10 := [][]string{}
	row1 := []string{"", "", "", "", "", "", "x", "x", "", ""}
	row2 := []string{"", "", "", "", "", "x", "x", "", "", ""}
	row3 := []string{"", "", "", "", "x", "x", "", "", "", ""}
	row4 := []string{"", "", "", "x", "", "x", "", "", "", ""}
	row5 := []string{"", "", "x", "x", "", "", "x", "x", "", ""}
	row6 := []string{"", "", "", "", "x", "", "x", "", "", ""}
	row7 := []string{"", "", "", "", "x", "x", "", "", "", ""}
	row8 := []string{"", "", "", "x", "x", "", "", "", "", ""}
	row9 := []string{"", "", "x", "x", "", "", "", "", "", ""}
	row10 := []string{"", "", "", "", "", "", "", "", "", ""}

	// row1 := []string{"", "", "", "", "", "x", "", "", "", ""}
	// row2 := []string{"", "", "", "", "x", "", "", "", "", ""}
	// row3 := []string{"", "", "", "x", "x", "x", "", "", "", ""}
	// row4 := []string{"", "", "", "", "", "x", "", "", "", ""}
	// row5 := []string{"", "", "", "", "x", "", "", "", "", ""}
	// row6 := []string{"", "", "", "", "", "", "", "", "", ""}
	// row7 := []string{"", "", "", "", "", "", "", "", "", ""}
	// row8 := []string{"", "", "", "", "", "", "", "", "", ""}
	// row9 := []string{"", "", "", "", "", "", "", "", "", ""}
	// row10 := []string{"", "", "", "", "", "", "", "", "", ""}

	matriz10x10 = append(matriz10x10, row1)
	matriz10x10 = append(matriz10x10, row2)
	matriz10x10 = append(matriz10x10, row3)
	matriz10x10 = append(matriz10x10, row4)
	matriz10x10 = append(matriz10x10, row5)
	matriz10x10 = append(matriz10x10, row6)
	matriz10x10 = append(matriz10x10, row7)
	matriz10x10 = append(matriz10x10, row8)
	matriz10x10 = append(matriz10x10, row9)
	matriz10x10 = append(matriz10x10, row10)
	return matriz10x10
}

/*
	CrearMatrizBatman: crea la matriz 10x10 para batman
*/
func CrearMatrizBatman() [][]string {

	matriz10x10 := [][]string{}
	row1 := []string{"", "", "", "", "", "", "", "", "", ""}
	row2 := []string{"", "", "", "", "", "", "", "", "", ""}
	row3 := []string{"", "", "", "", "x", "x", "", "", "", ""}
	row4 := []string{"", "x", "", "", "x", "x", "", "", "x", ""}
	row5 := []string{"x", "x", "x", "x", "x", "x", "x", "x", "x", "x"}
	row6 := []string{"x", "", "x", "", "x", "x", "", "x", "", "x"}
	row7 := []string{"", "", "", "", "x", "x", "", "", "", ""}
	row8 := []string{"", "", "", "", "", "", "", "", "", ""}
	row9 := []string{"", "", "", "", "", "", "", "", "", ""}
	row10 := []string{"", "", "", "", "", "", "", "", "", ""}

	// row1 := []string{"", "", "", "", "", "", "x", "", "", "x"}
	// row2 := []string{"", "", "", "", "", "", "x", "x", "x", "x"}
	// row3 := []string{"", "", "", "", "", "", "x", "", "", "x"}
	// row4 := []string{"", "", "", "", "", "", "", "", "", ""}
	// row5 := []string{"", "", "", "", "", "", "", "", "", ""}
	// row6 := []string{"", "", "", "", "", "", "", "", "", ""}
	// row7 := []string{"", "", "", "", "", "", "", "", "", ""}
	// row8 := []string{"", "", "", "", "", "", "", "", "", ""}
	// row9 := []string{"", "", "", "", "", "", "", "", "", ""}
	// row10 := []string{"", "", "", "", "", "", "", "", "", ""}

	matriz10x10 = append(matriz10x10, row1)
	matriz10x10 = append(matriz10x10, row2)
	matriz10x10 = append(matriz10x10, row3)
	matriz10x10 = append(matriz10x10, row4)
	matriz10x10 = append(matriz10x10, row5)
	matriz10x10 = append(matriz10x10, row6)
	matriz10x10 = append(matriz10x10, row7)
	matriz10x10 = append(matriz10x10, row8)
	matriz10x10 = append(matriz10x10, row9)
	matriz10x10 = append(matriz10x10, row10)
	return matriz10x10
}

/*
	CrearMatriz4Fantasticos: crea la matriz 10x10 para 4 fantasticos
*/
func CrearMatriz4Fantasticos() [][]string {

	matriz10x10 := [][]string{}
	row1 := []string{"", "", "", "", "", "", "", "x", "x", "x"}
	row2 := []string{"", "", "", "", "", "", "x", "", "x", "x"}
	row3 := []string{"", "", "", "", "", "x", "", "", "x", "x"}
	row4 := []string{"", "", "", "", "x", "", "", "", "x", "x"}
	row5 := []string{"", "", "", "x", "", "", "", "", "x", "x"}
	row6 := []string{"", "", "x", "", "", "", "", "", "x", "x"}
	row7 := []string{"", "x", "", "", "", "", "", "", "x", "x"}
	row8 := []string{"", "x", "x", "x", "x", "x", "x", "x", "x", "x"}
	row9 := []string{"", "", "", "", "", "", "", "", "x", "x"}
	row10 := []string{"", "", "", "", "", "", "", "", "x", "x"}

	// row1 := []string{"", "", "", "", "", "", "", "", "", ""}
	// row2 := []string{"", "", "", "", "", "", "", "", "", ""}
	// row3 := []string{"", "", "", "", "", "", "", "", "", ""}
	// row4 := []string{"", "", "", "", "", "", "", "", "", ""}
	// row5 := []string{"", "", "x", "", "", "", "", "", "", ""}
	// row6 := []string{"", "x", "x", "", "", "", "", "", "", ""}
	// row7 := []string{"x", "x", "x", "", "", "", "", "", "", ""}
	// row8 := []string{"", "", "x", "", "", "", "", "", "", ""}
	// row9 := []string{"", "", "", "", "", "", "", "", "", ""}
	// row10 := []string{"", "", "", "", "", "", "", "", "", ""}

	matriz10x10 = append(matriz10x10, row1)
	matriz10x10 = append(matriz10x10, row2)
	matriz10x10 = append(matriz10x10, row3)
	matriz10x10 = append(matriz10x10, row4)
	matriz10x10 = append(matriz10x10, row5)
	matriz10x10 = append(matriz10x10, row6)
	matriz10x10 = append(matriz10x10, row7)
	matriz10x10 = append(matriz10x10, row8)
	matriz10x10 = append(matriz10x10, row9)
	matriz10x10 = append(matriz10x10, row10)
	return matriz10x10
}

func CrearMatrizSpiderman() [][]string {

	matriz10x10 := [][]string{}
	row1 := []string{"", "", "", "", "", "", "", "", "", ""}
	row2 := []string{"", "", "", "x", "", "", "x", "", "", ""}
	row3 := []string{"", "x", "", "", "x", "x", "", "", "x", ""}
	row4 := []string{"", "", "x", "x", "x", "x", "x", "x", "", ""}
	row5 := []string{"", "", "", "", "x", "x", "", "", "", ""}
	row6 := []string{"", "", "x", "x", "x", "x", "x", "x", "", ""}
	row7 := []string{"", "x", "", "x", "", "", "x", "", "x", ""}
	row8 := []string{"", "x", "", "x", "", "", "x", "", "x", ""}
	row9 := []string{"", "x", "", "x", "", "", "x", "", "x", ""}
	row10 := []string{"", "", "", "", "", "", "", "", "", ""}

	// row1 := []string{"", "", "", "", "", "", "", "", "", ""}
	// row2 := []string{"", "", "", "", "", "", "", "", "", ""}
	// row3 := []string{"", "", "", "", "", "", "", "", "", ""}
	// row4 := []string{"", "", "", "", "", "", "", "", "", ""}
	// row5 := []string{"", "", "", "x", "", "", "x", "", "", ""}
	// row6 := []string{"", "", "", "", "x", "x", "", "", "", ""}
	// row7 := []string{"", "", "", "", "x", "x", "", "", "", ""}
	// row8 := []string{"", "", "", "x", "", "", "x", "", "", ""}
	// row9 := []string{"", "", "", "", "", "", "", "", "", ""}
	// row10 := []string{"", "", "", "", "", "", "", "", "", ""}

	matriz10x10 = append(matriz10x10, row1)
	matriz10x10 = append(matriz10x10, row2)
	matriz10x10 = append(matriz10x10, row3)
	matriz10x10 = append(matriz10x10, row4)
	matriz10x10 = append(matriz10x10, row5)
	matriz10x10 = append(matriz10x10, row6)
	matriz10x10 = append(matriz10x10, row7)
	matriz10x10 = append(matriz10x10, row8)
	matriz10x10 = append(matriz10x10, row9)
	matriz10x10 = append(matriz10x10, row10)
	return matriz10x10
}

func CrearMatrizThor() [][]string {

	matriz10x10 := [][]string{}
	row1 := []string{"", "", "", "", "", "", "", "", "", ""}
	row2 := []string{"", "x", "x", "x", "x", "x", "x", "x", "x", ""}
	row3 := []string{"", "x", "", "", "", "", "", "", "x", ""}
	row4 := []string{"", "x", "", "", "", "", "", "", "x", ""}
	row5 := []string{"", "x", "x", "x", "x", "x", "x", "x", "x", ""}
	row6 := []string{"", "", "", "", "x", "x", "", "", "", ""}
	row7 := []string{"", "", "", "", "x", "x", "", "", "", ""}
	row8 := []string{"", "", "", "", "x", "x", "", "", "", ""}
	row9 := []string{"", "", "", "", "x", "x", "", "", "", ""}
	row10 := []string{"", "", "", "", "x", "x", "", "", "", ""}

	// row1 := []string{"", "", "", "", "", "", "", "", "", ""}
	// row2 := []string{"", "", "", "", "", "", "", "", "", ""}
	// row3 := []string{"", "", "", "", "", "", "", "", "", ""}
	// row4 := []string{"", "", "", "", "", "", "", "", "", ""}
	// row5 := []string{"", "", "", "", "", "", "", "x", "x", "x"}
	// row6 := []string{"", "", "", "", "", "", "", "x", "x", "x"}
	// row7 := []string{"", "", "", "", "", "", "", "x", "x", "x"}
	// row8 := []string{"", "", "", "", "", "", "", "", "x", ""}
	// row9 := []string{"", "", "", "", "", "", "", "", "x", ""}
	// row10 := []string{"", "", "", "", "", "", "", "", "x", ""}

	matriz10x10 = append(matriz10x10, row1)
	matriz10x10 = append(matriz10x10, row2)
	matriz10x10 = append(matriz10x10, row3)
	matriz10x10 = append(matriz10x10, row4)
	matriz10x10 = append(matriz10x10, row5)
	matriz10x10 = append(matriz10x10, row6)
	matriz10x10 = append(matriz10x10, row7)
	matriz10x10 = append(matriz10x10, row8)
	matriz10x10 = append(matriz10x10, row9)
	matriz10x10 = append(matriz10x10, row10)
	return matriz10x10
}

//vacio!!
// row1 := []string{"", "", "", "", "", "", "", "", "", ""}
// row2 := []string{"", "", "", "", "", "", "", "", "", ""}
// row3 := []string{"", "", "", "", "", "", "", "", "", ""}
// row4 := []string{"", "", "", "", "", "", "", "", "", ""}
// row5 := []string{"", "", "", "", "", "", "", "", "", ""}
// row6 := []string{"", "", "", "", "", "", "", "", "", ""}
// row7 := []string{"", "", "", "", "", "", "", "", "", ""}
// row8 := []string{"", "", "", "", "", "", "", "", "", ""}
// row9 := []string{"", "", "", "", "", "", "", "", "", ""}
// row10 := []string{"", "", "", "", "", "", "", "", "", ""}
