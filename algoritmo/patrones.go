package algoritmo

/*
	CrearMatrizLinterna: crea la matriz 16x16 para linterna verde
*/
func CrearMatrizLinterna() [][]string {

	matriz20x20 := [][]string{}
	row1 := []string{"x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x"}
	row2 := []string{"x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x"}

	row3 := []string{"", "", "", "", "", "", "", "", "x", "x", "x", "", "", "", "", "", "", "", "", ""}
	row4 := []string{"", "", "", "", "", "", "", "x", "", "", "", "x", "", "", "", "", "", "", "", ""}
	row5 := []string{"", "", "", "", "", "", "x", "", "", "", "", "", "x", "", "", "", "", "", "", ""}
	row6 := []string{"", "", "", "", "", "x", "", "", "", "", "", "", "", "x", "", "", "", "", "", ""}
	row7 := []string{"", "", "", "", "x", "", "", "", "", "", "", "", "", "", "x", "", "", "", "", ""}
	row8 := []string{"", "", "", "", "x", "", "", "", "", "", "", "", "", "", "x", "", "", "", "", ""}
	row9 := []string{"", "", "", "", "x", "", "", "", "", "", "", "", "", "", "x", "", "", "", "", ""}
	row10 := []string{"", "", "", "", "x", "", "", "", "", "", "", "", "", "", "x", "", "", "", "", ""}
	row11 := []string{"", "", "", "", "x", "", "", "", "", "", "", "", "", "", "x", "", "", "", "", ""}
	row12 := []string{"", "", "", "", "x", "", "", "", "", "", "", "", "", "", "x", "", "", "", "", ""}
	row13 := []string{"", "", "", "", "x", "", "", "", "", "", "", "", "", "", "x", "", "", "", "", ""}
	row14 := []string{"", "", "", "", "x", "", "", "", "", "", "", "", "", "", "x", "", "", "", "", ""}
	row15 := []string{"", "", "", "", "", "x", "", "", "", "", "", "", "", "x", "", "", "", "", "", ""}
	row16 := []string{"", "", "", "", "", "", "x", "", "", "", "", "", "x", "", "", "", "", "", "", ""}
	row17 := []string{"", "", "", "", "", "", "", "x", "", "", "", "x", "", "", "", "", "", "", "", ""}
	row18 := []string{"", "", "", "", "", "", "", "", "x", "x", "x", "", "", "", "", "", "", "", "", ""}

	row19 := []string{"x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x"}
	row20 := []string{"x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x"}

	matriz20x20 = append(matriz20x20, row1)
	matriz20x20 = append(matriz20x20, row2)
	matriz20x20 = append(matriz20x20, row3)
	matriz20x20 = append(matriz20x20, row4)
	matriz20x20 = append(matriz20x20, row5)
	matriz20x20 = append(matriz20x20, row6)
	matriz20x20 = append(matriz20x20, row7)
	matriz20x20 = append(matriz20x20, row8)
	matriz20x20 = append(matriz20x20, row9)
	matriz20x20 = append(matriz20x20, row10)
	matriz20x20 = append(matriz20x20, row11)
	matriz20x20 = append(matriz20x20, row12)
	matriz20x20 = append(matriz20x20, row13)
	matriz20x20 = append(matriz20x20, row14)
	matriz20x20 = append(matriz20x20, row15)
	matriz20x20 = append(matriz20x20, row16)
	matriz20x20 = append(matriz20x20, row17)
	matriz20x20 = append(matriz20x20, row18)
	matriz20x20 = append(matriz20x20, row19)
	matriz20x20 = append(matriz20x20, row20)
	return matriz20x20
}

/*
	CrearMatrizFlash: crea la matriz 10x10 para flash
*/
func CrearMatrizFlash() [][]string {

	matriz20x20 := [][]string{}
	row1 := []string{"", "", "", "", "", "", "", "", "", "", "x", "x", "x", "x", "x", "x", "x", "", "", ""}
	row2 := []string{"", "", "", "", "", "", "", "", "", "x", "", "", "", "", "", "", "", "x", "", ""}
	row3 := []string{"", "", "", "", "", "", "", "", "x", "", "", "", "", "", "", "", "x", "", "", ""}
	row4 := []string{"", "", "", "", "", "", "", "x", "", "", "", "", "", "", "", "x", "", "", "", ""}
	row5 := []string{"", "", "", "", "", "", "x", "", "", "", "", "", "", "", "x", "", "", "", "", ""}
	row6 := []string{"", "", "", "", "", "x", "", "", "", "", "", "", "", "x", "x", "x", "x", "x", "", ""}
	row7 := []string{"", "", "", "", "x", "", "", "", "", "", "", "", "", "", "", "", "x", "", "", ""}
	row8 := []string{"", "", "", "x", "x", "x", "x", "x", "x", "x", "", "", "", "", "x", "", "", "", "", ""}
	row9 := []string{"", "", "", "", "", "", "", "", "", "", "", "x", "", "", "", "x", "", "", "", ""}
	row10 := []string{"", "", "", "", "", "", "", "", "", "", "x", "", "", "", "x", "", "", "", "", ""}
	row11 := []string{"", "", "", "", "", "", "", "", "", "x", "", "", "", "x", "", "", "", "", "", ""}
	row12 := []string{"", "", "", "", "", "", "", "", "x", "", "", "", "x", "", "", "", "", "", "", ""}
	row13 := []string{"", "", "", "", "", "", "", "x", "", "", "x", "", "", "", "", "", "", "", "", ""}
	row14 := []string{"", "", "", "", "", "", "x", "", "", "x", "", "", "", "", "", "", "", "", "", ""}
	row15 := []string{"", "", "", "", "", "x", "", "", "x", "", "", "", "", "", "", "", "", "", "", ""}
	row16 := []string{"", "", "", "", "x", "", "x", "", "", "", "", "", "", "", "", "", "", "", "", ""}
	row17 := []string{"", "", "", "x", "x", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""}
	row18 := []string{"", "", "x", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""}
	row19 := []string{"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""}
	row20 := []string{"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""}

	matriz20x20 = append(matriz20x20, row1)
	matriz20x20 = append(matriz20x20, row2)
	matriz20x20 = append(matriz20x20, row3)
	matriz20x20 = append(matriz20x20, row4)
	matriz20x20 = append(matriz20x20, row5)
	matriz20x20 = append(matriz20x20, row6)
	matriz20x20 = append(matriz20x20, row7)
	matriz20x20 = append(matriz20x20, row8)
	matriz20x20 = append(matriz20x20, row9)
	matriz20x20 = append(matriz20x20, row10)
	matriz20x20 = append(matriz20x20, row11)
	matriz20x20 = append(matriz20x20, row12)
	matriz20x20 = append(matriz20x20, row13)
	matriz20x20 = append(matriz20x20, row14)
	matriz20x20 = append(matriz20x20, row15)
	matriz20x20 = append(matriz20x20, row16)
	matriz20x20 = append(matriz20x20, row17)
	matriz20x20 = append(matriz20x20, row18)
	matriz20x20 = append(matriz20x20, row19)
	matriz20x20 = append(matriz20x20, row20)
	return matriz20x20
}

/*
	CrearMatrizBatman: crea la matriz 10x10 para batman
*/
func CrearMatrizBatman() [][]string {

	matriz20x20 := [][]string{}
	row1 := []string{"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""}
	row2 := []string{"", "", "", "", "", "x", "", "", "", "", "", "", "", "x", "", "", "", "", "", ""}
	row3 := []string{"", "", "", "", "x", "", "x", "", "", "", "", "", "", "x", "x", "", "", "", "", ""}
	row4 := []string{"", "", "", "x", "", "", "x", "", "x", "", "x", "", "x", "", "", "x", "", "", "", ""}
	row5 := []string{"", "", "x", "", "", "", "x", "x", "x", "x", "x", "x", "", "", "", "x", "", "", "", ""}
	row6 := []string{"", "", "x", "", "", "", "", "", "", "", "", "", "", "", "", "", "x", "", "", ""}
	row7 := []string{"", "", "x", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "x", "", ""}
	row8 := []string{"", "", "x", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "x", "", ""}
	row9 := []string{"", "", "x", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "x", "", ""}
	row10 := []string{"", "", "x", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "x", "", ""}
	row11 := []string{"", "", "x", "", "", "", "", "", "x", "x", "", "", "", "", "", "", "x", "", "", ""}
	row12 := []string{"", "", "x", "", "", "", "", "x", "", "", "x", "", "", "", "", "", "x", "", "", ""}
	row13 := []string{"", "", "x", "", "x", "", "x", "", "", "", "", "x", "", "x", "", "x", "", "", "", ""}
	row14 := []string{"", "", "", "x", "", "x", "", "", "", "", "", "", "x", "", "x", "", "", "", "", ""}
	row15 := []string{"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""}
	row16 := []string{"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""}
	row17 := []string{"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""}
	row18 := []string{"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""}
	row19 := []string{"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""}
	row20 := []string{"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""}

	matriz20x20 = append(matriz20x20, row1)
	matriz20x20 = append(matriz20x20, row2)
	matriz20x20 = append(matriz20x20, row3)
	matriz20x20 = append(matriz20x20, row4)
	matriz20x20 = append(matriz20x20, row5)
	matriz20x20 = append(matriz20x20, row6)
	matriz20x20 = append(matriz20x20, row7)
	matriz20x20 = append(matriz20x20, row8)
	matriz20x20 = append(matriz20x20, row9)
	matriz20x20 = append(matriz20x20, row10)
	matriz20x20 = append(matriz20x20, row11)
	matriz20x20 = append(matriz20x20, row12)
	matriz20x20 = append(matriz20x20, row13)
	matriz20x20 = append(matriz20x20, row14)
	matriz20x20 = append(matriz20x20, row15)
	matriz20x20 = append(matriz20x20, row16)
	matriz20x20 = append(matriz20x20, row17)
	matriz20x20 = append(matriz20x20, row18)
	matriz20x20 = append(matriz20x20, row19)
	matriz20x20 = append(matriz20x20, row20)
	return matriz20x20
}

/*
	CrearMatriz4Fantasticos: crea la matriz 10x10 para 4 fantasticos
*/
func CrearMatriz4Fantasticos() [][]string {

	matriz20x20 := [][]string{}
	row1 := []string{"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""}
	row2 := []string{"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""}
	row3 := []string{"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""}
	row4 := []string{"", "", "", "", "", "", "", "", "", "", "", "x", "", "", "", "", "", "", "", ""}
	row5 := []string{"", "", "", "", "", "", "", "", "", "", "x", "x", "", "", "", "", "", "", "", ""}
	row6 := []string{"", "", "", "", "", "", "", "", "", "", "x", "x", "", "", "", "", "", "", "", ""}
	row7 := []string{"", "", "", "", "", "", "", "", "", "x", "", "x", "", "", "", "", "", "", "", ""}
	row8 := []string{"", "", "", "", "", "", "", "", "x", "", "", "x", "", "", "", "", "", "", "", ""}
	row9 := []string{"", "", "", "", "", "", "", "x", "", "", "", "x", "", "", "", "", "", "", "", ""}
	row10 := []string{"", "", "", "", "", "", "x", "", "", "", "", "x", "", "", "", "", "", "", "", ""}
	row11 := []string{"", "", "", "", "", "x", "", "", "", "", "", "x", "", "", "", "", "", "", "", ""}
	row12 := []string{"", "", "", "", "x", "", "", "", "", "", "", "x", "x", "", "", "", "", "", "", ""}
	row13 := []string{"", "", "", "", "x", "", "", "", "", "", "", "x", "x", "", "", "", "", "", "", ""}
	row14 := []string{"", "", "", "", "x", "", "", "", "", "", "", "x", "x", "", "", "", "", "", "", ""}
	row15 := []string{"", "", "", "", "x", "x", "x", "x", "x", "", "", "x", "x", "", "", "", "", "", "", ""}
	row16 := []string{"", "", "", "", "", "", "", "", "x", "", "", "x", "", "", "", "", "", "", "", ""}
	row17 := []string{"", "", "", "", "", "", "", "", "x", "", "", "x", "", "", "", "", "", "", "", ""}
	row18 := []string{"", "", "", "", "", "", "", "", "x", "", "", "x", "", "", "", "", "", "", "", ""}
	row19 := []string{"", "", "", "", "", "", "", "", "x", "x", "x", "x", "", "", "", "", "", "", "", ""}
	row20 := []string{"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""}

	matriz20x20 = append(matriz20x20, row1)
	matriz20x20 = append(matriz20x20, row2)
	matriz20x20 = append(matriz20x20, row3)
	matriz20x20 = append(matriz20x20, row4)
	matriz20x20 = append(matriz20x20, row5)
	matriz20x20 = append(matriz20x20, row6)
	matriz20x20 = append(matriz20x20, row7)
	matriz20x20 = append(matriz20x20, row8)
	matriz20x20 = append(matriz20x20, row9)
	matriz20x20 = append(matriz20x20, row10)
	matriz20x20 = append(matriz20x20, row11)
	matriz20x20 = append(matriz20x20, row12)
	matriz20x20 = append(matriz20x20, row13)
	matriz20x20 = append(matriz20x20, row14)
	matriz20x20 = append(matriz20x20, row15)
	matriz20x20 = append(matriz20x20, row16)
	matriz20x20 = append(matriz20x20, row17)
	matriz20x20 = append(matriz20x20, row18)
	matriz20x20 = append(matriz20x20, row19)
	matriz20x20 = append(matriz20x20, row20)
	return matriz20x20
}

func CrearMatrizSpiderman() [][]string {

	matriz20x20 := [][]string{}
	row1 := []string{"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""}
	row2 := []string{"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""}
	row3 := []string{"", "", "", "", "x", "", "x", "", "", "", "x", "", "x", "", "", "", "", "", "", ""}
	row4 := []string{"", "", "", "", "x", "", "x", "", "", "", "x", "", "x", "", "", "", "", "", "", ""}
	row5 := []string{"", "", "", "", "x", "", "x", "", "", "", "x", "", "x", "", "", "", "", "", "", ""}
	row6 := []string{"", "", "", "", "x", "", "x", "", "", "", "x", "", "x", "", "", "", "", "", "", ""}
	row7 := []string{"", "", "", "", "x", "", "x", "", "", "", "x", "", "x", "", "", "", "", "", "", ""}
	row8 := []string{"", "", "", "", "", "x", "x", "x", "x", "x", "x", "", "", "", "", "", "", "", "", ""}
	row9 := []string{"", "", "", "", "", "", "", "x", "x", "x", "", "", "", "", "", "", "", "", "", ""}
	row10 := []string{"", "", "", "", "", "", "", "x", "x", "", "", "", "", "", "", "", "", "", "", ""}
	row11 := []string{"", "", "", "", "", "", "", "x", "x", "", "", "", "", "", "", "", "", "", "", ""}
	row12 := []string{"", "", "", "x", "x", "x", "", "x", "x", "", "x", "x", "x", "", "", "", "", "", "", ""}
	row13 := []string{"", "", "x", "", "", "", "x", "x", "x", "x", "", "", "", "x", "", "", "", "", "", ""}
	row14 := []string{"", "", "x", "", "", "x", "x", "x", "x", "x", "x", "", "", "x", "", "", "", "", "", ""}
	row15 := []string{"", "", "x", "", "x", "", "", "x", "x", "", "", "x", "", "x", "", "", "", "", "", ""}
	row16 := []string{"", "", "x", "", "x", "", "", "", "", "", "", "x", "", "x", "", "", "", "", "", ""}
	row17 := []string{"", "", "x", "", "x", "", "", "", "", "", "", "x", "", "x", "", "", "", "", "", ""}
	row18 := []string{"", "", "x", "", "", "x", "", "", "", "", "x", "", "", "x", "", "", "", "", "", ""}
	row19 := []string{"", "", "", "x", "", "x", "", "", "", "", "x", "", "", "", "", "", "", "", "", ""}
	row20 := []string{"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""}

	matriz20x20 = append(matriz20x20, row1)
	matriz20x20 = append(matriz20x20, row2)
	matriz20x20 = append(matriz20x20, row3)
	matriz20x20 = append(matriz20x20, row4)
	matriz20x20 = append(matriz20x20, row5)
	matriz20x20 = append(matriz20x20, row6)
	matriz20x20 = append(matriz20x20, row7)
	matriz20x20 = append(matriz20x20, row8)
	matriz20x20 = append(matriz20x20, row9)
	matriz20x20 = append(matriz20x20, row10)
	matriz20x20 = append(matriz20x20, row11)
	matriz20x20 = append(matriz20x20, row12)
	matriz20x20 = append(matriz20x20, row13)
	matriz20x20 = append(matriz20x20, row14)
	matriz20x20 = append(matriz20x20, row15)
	matriz20x20 = append(matriz20x20, row16)
	matriz20x20 = append(matriz20x20, row17)
	matriz20x20 = append(matriz20x20, row18)
	matriz20x20 = append(matriz20x20, row19)
	matriz20x20 = append(matriz20x20, row20)
	return matriz20x20
}

func CrearMatrizThor() [][]string {

	matriz20x20 := [][]string{}
	row1 := []string{"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""}
	row2 := []string{"", "", "", "", "", "", "", "", "", "", "", "x", "", "", "", "", "", "", "", ""}
	row3 := []string{"", "", "", "", "", "", "", "", "", "", "x", "", "x", "", "", "", "", "", "", ""}
	row4 := []string{"", "", "", "", "", "", "", "", "", "x", "", "", "", "x", "", "", "", "", "", ""}
	row5 := []string{"", "", "", "", "", "", "", "", "x", "", "", "", "", "", "x", "", "", "", "", ""}
	row6 := []string{"", "", "", "", "", "", "", "", "", "x", "", "", "", "", "", "x", "", "", "", ""}
	row7 := []string{"", "", "", "", "", "", "", "", "", "", "x", "", "", "", "", "", "x", "", "", ""}
	row8 := []string{"", "", "", "", "", "", "", "", "", "", "", "x", "", "", "", "", "", "x", "", ""}
	row9 := []string{"", "", "", "", "", "", "", "", "", "", "", "x", "x", "", "", "", "", "x", "", ""}
	row10 := []string{"", "", "", "", "", "", "", "", "", "", "", "x", "x", "x", "", "", "", "x", "", ""}
	row11 := []string{"", "", "", "", "", "", "", "", "", "", "x", "x", "x", "", "x", "", "x", "", "", ""}
	row12 := []string{"", "", "", "", "", "", "", "", "", "x", "x", "x", "", "", "", "x", "", "", "", ""}
	row13 := []string{"", "", "", "", "", "", "", "", "x", "x", "x", "", "", "", "", "", "", "", "", ""}
	row14 := []string{"", "", "", "", "", "", "", "", "x", "x", "", "", "", "", "", "", "", "", "", ""}
	row15 := []string{"", "", "", "", "", "", "", "x", "", "", "", "", "", "", "", "", "", "", "", ""}
	row16 := []string{"", "", "", "", "", "", "x", "", "", "", "", "", "", "", "", "", "", "", "", ""}
	row17 := []string{"", "", "", "", "", "x", "", "", "", "", "", "", "", "", "", "", "", "", "", ""}
	row18 := []string{"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""}
	row19 := []string{"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""}
	row20 := []string{"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""}

	matriz20x20 = append(matriz20x20, row1)
	matriz20x20 = append(matriz20x20, row2)
	matriz20x20 = append(matriz20x20, row3)
	matriz20x20 = append(matriz20x20, row4)
	matriz20x20 = append(matriz20x20, row5)
	matriz20x20 = append(matriz20x20, row6)
	matriz20x20 = append(matriz20x20, row7)
	matriz20x20 = append(matriz20x20, row8)
	matriz20x20 = append(matriz20x20, row9)
	matriz20x20 = append(matriz20x20, row10)
	matriz20x20 = append(matriz20x20, row11)
	matriz20x20 = append(matriz20x20, row12)
	matriz20x20 = append(matriz20x20, row13)
	matriz20x20 = append(matriz20x20, row14)
	matriz20x20 = append(matriz20x20, row15)
	matriz20x20 = append(matriz20x20, row16)
	matriz20x20 = append(matriz20x20, row17)
	matriz20x20 = append(matriz20x20, row18)
	matriz20x20 = append(matriz20x20, row19)
	matriz20x20 = append(matriz20x20, row20)
	return matriz20x20
}
