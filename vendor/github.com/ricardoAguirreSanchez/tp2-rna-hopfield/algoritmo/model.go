package algoritmo

import "gonum.org/v1/gonum/mat"

type PatronStruct struct {
	Matriz       [][]string //matriz 10x10 ingresada
	PatronVector mat.Matrix //vector de 1 row con 1 y -1
	ProducMatriz *mat.Dense //matriz producto de p1 x p1 traspuesta con diagonal cero
}

// type AlgoritmoStruct struct {
// 	PatronStructs []PatronStruct
// 	PesoRed       *mat.Dense
// }
