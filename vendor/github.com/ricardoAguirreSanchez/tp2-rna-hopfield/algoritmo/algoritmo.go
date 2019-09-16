package algoritmo

import (
	"fmt"

	"github.com/ricardoAguirreSanchez/tp2-rna-hopfield/formulario"
	"gonum.org/v1/gonum/mat"
)

func Aprende() AlgoritmoStruct {
	//creamos patrones
	fmt.Println("------------Creamos patrones------------------")
	matriz10x10Facebook := [][]string{}
	row1 := []string{"", "", "", "", "", "", "", "", "", ""}
	row2 := []string{"", "", "", "", "*", "*", "*", "", "", ""}
	row3 := []string{"", "", "", "", "*", "*", "*", "", "", ""}
	row4 := []string{"", "", "", "", "*", "*", "", "", "", ""}
	row5 := []string{"", "", "", "*", "*", "*", "*", "", "", ""}
	row6 := []string{"", "", "", "*", "*", "*", "*", "", "", ""}
	row7 := []string{"", "", "", "", "*", "*", "", "", "", ""}
	row8 := []string{"", "", "", "", "*", "*", "", "", "", ""}
	row9 := []string{"", "", "", "", "*", "*", "", "", "", ""}
	row10 := []string{"", "", "", "", "*", "*", "", "", "", ""}

	matriz10x10Facebook = append(matriz10x10Facebook, row1)
	matriz10x10Facebook = append(matriz10x10Facebook, row2)
	matriz10x10Facebook = append(matriz10x10Facebook, row3)
	matriz10x10Facebook = append(matriz10x10Facebook, row4)
	matriz10x10Facebook = append(matriz10x10Facebook, row5)
	matriz10x10Facebook = append(matriz10x10Facebook, row6)
	matriz10x10Facebook = append(matriz10x10Facebook, row7)
	matriz10x10Facebook = append(matriz10x10Facebook, row8)
	matriz10x10Facebook = append(matriz10x10Facebook, row9)
	matriz10x10Facebook = append(matriz10x10Facebook, row10)

	patron1 := crearPatron(matriz10x10Facebook)

	fmt.Println("----------------------------------------------")
	matriz10x10WP := [][]string{}

	// Create three string slices.
	row1 = []string{"", "", "", "*", "*", "*", "*", "", "", ""}
	row2 = []string{"", "", "*", "", "", "", "", "*", "", ""}
	row3 = []string{"", "*", "", "*", "", "", "", "", "*", ""}
	row4 = []string{"", "*", "*", "*", "", "", "", "", "*", ""}
	row5 = []string{"", "*", "*", "", "", "", "", "", "*", ""}
	row6 = []string{"", "*", "", "*", "", "", "*", "*", "*", ""}
	row7 = []string{"", "*", "", "", "*", "*", "*", "", "*", ""}
	row8 = []string{"*", "*", "", "", "", "", "", "", "*", ""}
	row9 = []string{"*", "*", "*", "", "", "", "", "*", "", ""}
	row10 = []string{"", "", "", "*", "*", "*", "*", "", "", ""}

	// Append string slices to outer slice.
	matriz10x10WP = append(matriz10x10WP, row1)
	matriz10x10WP = append(matriz10x10WP, row2)
	matriz10x10WP = append(matriz10x10WP, row3)
	matriz10x10WP = append(matriz10x10WP, row4)
	matriz10x10WP = append(matriz10x10WP, row5)
	matriz10x10WP = append(matriz10x10WP, row6)
	matriz10x10WP = append(matriz10x10WP, row7)
	matriz10x10WP = append(matriz10x10WP, row8)
	matriz10x10WP = append(matriz10x10WP, row9)
	matriz10x10WP = append(matriz10x10WP, row10)
	patron2 := crearPatron(matriz10x10WP)
	fmt.Println("------------------Fin de crear patrones----------------------------")

	D1 := mat.NewDense(100, 100, nil)
	D1.Product(patron1.T(), patron1)
	// println("patron1' * patron1")
	// matPrint(D1)

	D11 := mat.NewDense(100, 100, nil)
	D11.Apply(diagonalEnCeros, D1)
	// println("patron1 con diagonal ceros:")
	// matPrint(D11)

	D2 := mat.NewDense(100, 100, nil)
	D2.Product(patron2.T(), patron2)

	D22 := mat.NewDense(100, 100, nil)
	D22.Apply(diagonalEnCeros, D2)

	pesoRed := mat.NewDense(100, 100, nil)
	pesoRed.Add(D11, D22)
	// println("pesoRed:")
	// matPrint(pesoRed)
	return AlgoritmoStruct{matriz10x10Facebook, matriz10x10WP, patron1, patron2, D11, D22, pesoRed}
}

func AplicoBusqueda(formularioInput formulario.Formulario, algoritmoStruct AlgoritmoStruct) formulario.Formulario {

	var formularioRespuesta formulario.Formulario

	prueba := crearPrueba(formularioInput)

	matrizResultado := testear(prueba, algoritmoStruct)

	formularioRespuesta.Matriz = matrizResultado

	return formularioRespuesta
}

////--------------------privadas--------------

// /*
// 	convertMatrizAFormulario: convierte de la matriz a la estructura del form q se usa
// */
// func convertMatrizAFormulario(matrizRespuesta mat.Matrix) formulario.Formulario {
// 	formulario := formulario.Formulario{}
// 	formulario.Matriz = nuevaMatriz()

// 	for i := 0; i < 10; i++ {
// 		for j := 0; j < 10; j++ {
// 			value := matrizRespuesta.At(i, j)
// 			formulario.Matriz[i][j] = strconv.Itoa(int(value))
// 		}
// 	}
// 	return formulario
// }

func matPrint(X mat.Matrix) {
	fa := mat.Formatted(X, mat.Prefix(""), mat.Squeeze())
	fmt.Printf("%v\n", fa)
}

func nuevaMatriz() [][]string {
	matriz10x10Vacia := [][]string{}
	for i := 0; i < 10; i++ {
		row1 := []string{"", "", "", "", "", "", "", "", "", ""}
		matriz10x10Vacia = append(matriz10x10Vacia, row1)
	}
	return matriz10x10Vacia
}

/*
	crearRow: crea una row de 1 y -1
*/
func crearRow(matriz [][]string) []float64 {
	resul := []float64{}
	for i := range matriz {
		for j := range matriz[i] {
			if matriz[i][j] == "" {
				resul = append(resul, -1)
			} else {
				resul = append(resul, 1)
			}
		}
	}
	return resul
}

/*
	crearPatron:devuelve la row correspondiente a la matriz 10 x 10
*/
func crearPatron(matriz10x10 [][]string) mat.Matrix {

	A := mat.NewDense(1, 100, nil)
	rowTEst := crearRow(matriz10x10)
	A.SetRow(0, rowTEst)
	return A
}

func diagonalEnCeros(i, j int, v float64) float64 {
	if i == j {
		return 0
	} else {
		return v
	}
}

func F(i, j int, v float64) float64 {
	if v > 0 {
		return 1
	}
	return -1
}

func testear(prueba mat.Matrix, algoStruct AlgoritmoStruct) [][]string {
	limite := 1
	contador1 := 0
	contador2 := 0

	comodin := prueba
	pruebaPorPEso := mat.NewDense(1, 100, nil)
	for limite <= 10 {
		pruebaPorPEso.Product(comodin, algoStruct.pesoRed)

		pruebaAplicoFn := mat.NewDense(1, 100, nil)
		pruebaAplicoFn.Apply(F, pruebaPorPEso)

		if esIgual(pruebaAplicoFn, algoStruct.patron1) {
			contador1 = contador1 + 1
			fmt.Println("mach con patron 1")
		} else {
			contador1 = 0
			fmt.Println("NO mach con patron 1")
		}

		if esIgual(pruebaAplicoFn, algoStruct.patron2) {
			contador2 = contador2 + 1
			fmt.Println("mach con patron 2")
		} else {
			contador2 = 0
			fmt.Println("NO mach con patron 2")
		}

		if contador1 >= 2 {
			//fin
			fmt.Println("es igual al patron 1 - Fin!")
			return algoStruct.matriz10x10Facebook
		}

		if contador2 >= 2 {
			//fin
			fmt.Println("es igual al patron 2 - Fin!")
			return algoStruct.matriz10x10WP
		}
		comodin = pruebaPorPEso
		limite += limite
	}

	fmt.Println("no existe ninguno - Fin!")
	return nuevaMatriz()
}

/*
	creaPatronVacio: crea una matriz vacia para indicar q no encontramos coincidencia
*/
func creaMatriz10x10Vacio() mat.Matrix {

	v := make([]float64, 100)
	for i := 0; i < 100; i++ {
		v[i] = float64(0)
	}
	// Create a new matrix
	A := mat.NewDense(10, 10, v)
	return A
}

/*
	esIgual: busca si es igual a un patron que tenemos en la base de conocimiento
*/
func esIgual(pruebaAplicoFn *mat.Dense, patron1 mat.Matrix) bool {
	resul := true

	C := mat.NewDense(1, 100, nil)
	C.Sub(pruebaAplicoFn, patron1)

	count := 100
	//buscamos si hay todos en cero
	for indexI := 0; indexI < count; indexI++ {
		if C.At(0, indexI) != 0 {
			return false
		}
	}
	return resul
}

/*
	crearPrueba: a partir de una matriz de "" y "*", lo pasa a una row de 1 y -1
*/
func crearPrueba(formu formulario.Formulario) mat.Matrix {
	A := mat.NewDense(1, 100, nil)
	rowTEst := crearRow(formu.Matriz)
	A.SetRow(0, rowTEst)
	return A
}
