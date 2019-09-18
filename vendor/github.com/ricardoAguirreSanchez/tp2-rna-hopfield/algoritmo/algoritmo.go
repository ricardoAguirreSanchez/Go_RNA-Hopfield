package algoritmo

import (
	"fmt"

	"github.com/ricardoAguirreSanchez/tp2-rna-hopfield/formulario"
	"gonum.org/v1/gonum/mat"
)

func Aprende() *mat.Dense {

	//creamos patrones
	fmt.Println("------------CrearMatrizLinterna------------------")
	matrizLinterna := CrearMatrizLinterna()
	patron1 := crearPatron(matrizLinterna)

	PrintMatrix(matrizLinterna)

	D1 := mat.NewDense(100, 100, nil)
	D1.Product(patron1.T(), patron1)

	D11 := mat.NewDense(100, 100, nil)
	D11.Apply(diagonalEnCeros, D1)

	fmt.Println("------------CrearMatrizFlash------------------")
	matrizFlash := CrearMatrizFlash()
	patron2 := crearPatron(matrizFlash)

	PrintMatrix(matrizFlash)

	D2 := mat.NewDense(100, 100, nil)
	D2.Product(patron2.T(), patron2)

	D22 := mat.NewDense(100, 100, nil)
	D22.Apply(diagonalEnCeros, D2)

	fmt.Println("------------CrearMatrizBatman------------------")
	matrizBatman := CrearMatrizBatman()
	patron3 := crearPatron(matrizBatman)

	PrintMatrix(matrizBatman)

	D3 := mat.NewDense(100, 100, nil)
	D3.Product(patron3.T(), patron3)

	D33 := mat.NewDense(100, 100, nil)
	D33.Apply(diagonalEnCeros, D3)

	fmt.Println("------------CrearMatriz4Fantasticos------------------")
	matriz4Fantastico := CrearMatriz4Fantasticos()
	patron4 := crearPatron(matriz4Fantastico)

	PrintMatrix(matriz4Fantastico)

	D4 := mat.NewDense(100, 100, nil)
	D4.Product(patron4.T(), patron4)

	D44 := mat.NewDense(100, 100, nil)
	D44.Apply(diagonalEnCeros, D4)

	fmt.Println("------------CrearMatrizSpiderman------------------")
	matrizSpiderman := CrearMatrizSpiderman()
	patron5 := crearPatron(matrizSpiderman)

	PrintMatrix(matrizSpiderman)

	D5 := mat.NewDense(100, 100, nil)
	D5.Product(patron5.T(), patron5)

	D55 := mat.NewDense(100, 100, nil)
	D55.Apply(diagonalEnCeros, D5)

	fmt.Println("------------CrearMatrizThor------------------")
	matrizThor := CrearMatrizThor()
	patron6 := crearPatron(matrizThor)

	PrintMatrix(matrizThor)

	D6 := mat.NewDense(100, 100, nil)
	D6.Product(patron6.T(), patron6)

	D66 := mat.NewDense(100, 100, nil)
	D66.Apply(diagonalEnCeros, D6)

	fmt.Println("------------Calculamos el peso------------------")
	pesoRed2 := mat.NewDense(100, 100, nil)
	pesoRed2.Add(D11, D22)
	pesoRed3 := mat.NewDense(100, 100, nil)
	pesoRed3.Add(pesoRed2, D33)
	pesoRed4 := mat.NewDense(100, 100, nil)
	pesoRed4.Add(pesoRed3, D44)
	pesoRed5 := mat.NewDense(100, 100, nil)
	pesoRed5.Add(pesoRed4, D55)
	pesoTotal := mat.NewDense(100, 100, nil)
	pesoTotal.Add(pesoRed5, D66)

	fmt.Println("patron1:")
	matPrint(patron1)
	fmt.Println("patron2:")
	matPrint(patron2)
	fmt.Println("patron3:")
	matPrint(patron3)
	fmt.Println("patron4:")
	matPrint(patron4)
	fmt.Println("patron5:")
	matPrint(patron5)
	fmt.Println("patron6:")
	matPrint(patron6)

	return pesoTotal
}

func AplicoBusqueda(formularioInput formulario.Formulario, peso *mat.Dense) formulario.Formulario {

	var formularioRespuesta formulario.Formulario

	patronIngresado := crearPrueba(formularioInput)
	fmt.Println("patron ingresado:")
	matPrint(patronIngresado)

	matrizResultado := testear(patronIngresado, peso)

	formularioRespuesta.Matriz = convertMatMatriz(matrizResultado)

	fmt.Println("matriz obtenida:")
	PrintMatrix(formularioRespuesta.Matriz)

	return formularioRespuesta
}

////--------------------privadas--------------

/*
	convertMatMatriz: convierte de Dense a [][]string
*/
func convertMatMatriz(pruebaAplicoFn *mat.Dense) [][]string {

	resul := nuevaMatriz()
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			posicion := j + i*10
			if int(pruebaAplicoFn.At(0, posicion)) == 1 {
				resul[i][j] = "*"
			} else {
				resul[i][j] = ""
			}
		}
	}
	return resul
}

func matPrint(X mat.Matrix) {
	fa := mat.Formatted(X, mat.Prefix(""), mat.Squeeze())
	fmt.Printf("%v\n", fa)
}

func NuevaMatriz() [][]string {
	matriz10x10Vacia := [][]string{}
	for i := 0; i < 10; i++ {
		row1 := []string{"", "", "", "", "", "", "", "", "", ""}
		matriz10x10Vacia = append(matriz10x10Vacia, row1)
	}
	return matriz10x10Vacia
}

func matrizNoEncontrada() [][]string {
	matriz10x10Vacia := [][]string{}
	for i := 0; i < 10; i++ {
		row1 := []string{"?", "?", "?", "?", "?", "?", "?", "?", "?", "?"}
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
			if matriz[i][j] == "*" {
				resul = append(resul, 1)
			} else {
				resul = append(resul, -1)
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

/*
	diagonalEnCeros: cambia la diagonal de la matriz por ceros
*/
func diagonalEnCeros(i, j int, v float64) float64 {
	if i == j {
		return 0
	} else {
		return v
	}
}

func Fn(i, j int, v float64) float64 {
	if v > 0 {
		return 1
	}
	return -1
}

/*
	testear: devuelve la matriz de 10x10 que mejor lo cumple
*/
func testear(prueba mat.Matrix, peso *mat.Dense) *mat.Dense {
	limite := 1

	var contadorEstable int

	pruebaPorPeso1 := mat.NewDense(1, 100, nil)
	pruebaPorPeso1.Product(prueba, peso)

	pruebaPorPesoFn1 := mat.NewDense(1, 100, nil)
	pruebaPorPesoFn1.Apply(Fn, pruebaPorPeso1)

	pruebaAplicoFnEstable := mat.NewDense(1, 100, nil)
	pruebaAplicoFnEstable = pruebaPorPesoFn1
	comodin := pruebaPorPesoFn1
	contadorEstable = 0

	for limite <= 9 {
		fmt.Println("Iteracion: ", limite)
		pruebaPorPeso := mat.NewDense(1, 100, nil)
		pruebaPorPeso.Product(comodin, peso)

		pruebaPorPesoFn := mat.NewDense(1, 100, nil)
		pruebaPorPesoFn.Apply(Fn, pruebaPorPeso)

		if esIgual(pruebaAplicoFnEstable, pruebaPorPesoFn) {
			contadorEstable = contadorEstable + 1
		} else {
			pruebaAplicoFnEstable = pruebaPorPesoFn
		}

		if contadorEstable >= 2 {
			fmt.Println("Se estabilizo!")
			return pruebaAplicoFnEstable
		}
		comodin = pruebaPorPesoFn
		limite = limite + 1
	}
	fmt.Println("No se estabilizo")
	return comodin
}

/*
	creaMatriz10x10Vacio: crea una matriz vacia para indicar q no encontramos coincidencia
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
		if int(C.At(0, indexI)) != 0 {
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
