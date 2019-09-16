package algoritmo

import (
	"fmt"

	"github.com/ricardoAguirreSanchez/tp2-rna-hopfield/formulario"
	"gonum.org/v1/gonum/mat"
)

func Aprende() AlgoritmoStruct {

	//creamos patrones
	fmt.Println("------------Creamos patrones------------------")
	matrizLinterna := CrearMatrizLinterna()
	patron1 := crearPatron(matrizLinterna)

	D1 := mat.NewDense(100, 100, nil)
	D1.Product(patron1.T(), patron1)

	D11 := mat.NewDense(100, 100, nil)
	D11.Apply(diagonalEnCeros, D1)

	fmt.Println("------------Creamos patrones------------------")
	matrizFlash := CrearMatrizFlash()
	patron2 := crearPatron(matrizFlash)

	D2 := mat.NewDense(100, 100, nil)
	D2.Product(patron2.T(), patron2)

	D22 := mat.NewDense(100, 100, nil)
	D22.Apply(diagonalEnCeros, D2)

	fmt.Println("------------Creamos patrones------------------")
	matrizBatman := CrearMatrizBatman()
	patron3 := crearPatron(matrizBatman)

	D3 := mat.NewDense(100, 100, nil)
	D3.Product(patron3.T(), patron3)

	D33 := mat.NewDense(100, 100, nil)
	D33.Apply(diagonalEnCeros, D3)

	fmt.Println("------------Creamos patrones------------------")
	matriz4Fantastico := CrearMatriz4Fantasticos()
	patron4 := crearPatron(matriz4Fantastico)

	D4 := mat.NewDense(100, 100, nil)
	D4.Product(patron4.T(), patron4)

	D44 := mat.NewDense(100, 100, nil)
	D44.Apply(diagonalEnCeros, D4)

	fmt.Println("------------Creamos patrones------------------")
	matrizSpiderman := CrearMatrizSpiderman()
	patron5 := crearPatron(matrizSpiderman)

	D5 := mat.NewDense(100, 100, nil)
	D5.Product(patron5.T(), patron5)

	D55 := mat.NewDense(100, 100, nil)
	D55.Apply(diagonalEnCeros, D5)

	fmt.Println("------------Creamos patrones------------------")
	matrizThor := CrearMatrizThor()
	patron6 := crearPatron(matrizThor)

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

	patronStructs := []PatronStruct{}
	patronStructs = append(patronStructs, PatronStruct{matrizLinterna, patron1, D11})
	patronStructs = append(patronStructs, PatronStruct{matrizLinterna, patron1, D22})
	patronStructs = append(patronStructs, PatronStruct{matrizLinterna, patron1, D33})
	patronStructs = append(patronStructs, PatronStruct{matrizLinterna, patron1, D44})
	patronStructs = append(patronStructs, PatronStruct{matrizLinterna, patron1, D55})
	patronStructs = append(patronStructs, PatronStruct{matrizLinterna, patron1, D66})

	return AlgoritmoStruct{patronStructs, pesoTotal}
}

func AplicoBusqueda(formularioInput formulario.Formulario, algoritmoStruct AlgoritmoStruct) formulario.Formulario {

	var formularioRespuesta formulario.Formulario

	prueba := crearPrueba(formularioInput)

	matrizResultado := testear(prueba, algoritmoStruct)

	formularioRespuesta.Matriz = matrizResultado

	return formularioRespuesta
}

////--------------------privadas--------------

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

func F(i, j int, v float64) float64 {
	if v > 0 {
		return 1
	}
	return -1
}

/*
	testear: devuelve la matriz de 10x10 que mejor lo cumple
*/
func testear(prueba mat.Matrix, algoStruct AlgoritmoStruct) [][]string {
	limite := 1

	var contador1, contador2, contador3, contador4, contador5, contador6 int = 0, 0, 0, 0, 0, 0

	comodin := prueba
	pruebaPorPEso := mat.NewDense(1, 100, nil)
	for limite <= 10 {
		pruebaPorPEso.Product(comodin, algoStruct.pesoRed)

		pruebaAplicoFn := mat.NewDense(1, 100, nil)
		pruebaAplicoFn.Apply(F, pruebaPorPEso)

		//controlamos si coincide con alguno
		if esIgual(pruebaAplicoFn, algoStruct.patronStructs[0].patronVector) {
			contador1 = contador1 + 1
			fmt.Println("mach con patron 1")
		} else {
			contador1 = 0
			fmt.Println("NO mach con patron 1")
		}

		if esIgual(pruebaAplicoFn, algoStruct.patronStructs[1].patronVector) {
			contador2 = contador2 + 1
			fmt.Println("mach con patron 2")
		} else {
			contador2 = 0
			fmt.Println("NO mach con patron 2")
		}

		if esIgual(pruebaAplicoFn, algoStruct.patronStructs[2].patronVector) {
			contador3 = contador3 + 1
			fmt.Println("mach con patron 3")
		} else {
			contador3 = 0
			fmt.Println("NO mach con patron 3")
		}

		if esIgual(pruebaAplicoFn, algoStruct.patronStructs[3].patronVector) {
			contador4 = contador4 + 1
			fmt.Println("mach con patron 4")
		} else {
			contador4 = 0
			fmt.Println("NO mach con patron 4")
		}

		if esIgual(pruebaAplicoFn, algoStruct.patronStructs[4].patronVector) {
			contador5 = contador5 + 1
			fmt.Println("mach con patron 5")
		} else {
			contador5 = 0
			fmt.Println("NO mach con patron 5")
		}

		if esIgual(pruebaAplicoFn, algoStruct.patronStructs[5].patronVector) {
			contador6 = contador6 + 1
			fmt.Println("mach con patron 6")
		} else {
			contador6 = 0
			fmt.Println("NO mach con patron 6")
		}

		//Revisamos la condicion de corte
		if contador1 >= 2 {
			fmt.Println("es igual al patron 1 - Fin!")
			return algoStruct.patronStructs[0].matriz
		}

		if contador2 >= 2 {
			fmt.Println("es igual al patron 2 - Fin!")
			return algoStruct.patronStructs[1].matriz
		}

		if contador3 >= 2 {
			fmt.Println("es igual al patron 3 - Fin!")
			return algoStruct.patronStructs[2].matriz
		}

		if contador4 >= 2 {
			fmt.Println("es igual al patron 4 - Fin!")
			return algoStruct.patronStructs[3].matriz
		}

		if contador5 >= 2 {
			fmt.Println("es igual al patron 5 - Fin!")
			return algoStruct.patronStructs[4].matriz
		}

		if contador6 >= 2 {
			fmt.Println("es igual al patron 6 - Fin!")
			return algoStruct.patronStructs[5].matriz
		}

		comodin = pruebaPorPEso
		limite += limite
	}

	fmt.Println("no existe ninguno - Fin :(!")
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
