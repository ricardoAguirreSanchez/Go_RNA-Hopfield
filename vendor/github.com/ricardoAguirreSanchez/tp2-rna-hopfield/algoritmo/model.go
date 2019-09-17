package algoritmo

import "gonum.org/v1/gonum/mat"

type PatronStruct struct {
	matriz       [][]string //matriz 10x10 ingresada
	patronVector mat.Matrix //vector de 1 row con 1 y -1
	producMatriz *mat.Dense //matriz producto de p1 x p1 traspuesta con diagonal cero
}

type AlgoritmoStruct struct {
	patronStructs []PatronStruct
	pesoRed       *mat.Dense
}
