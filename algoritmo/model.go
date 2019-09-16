package algoritmo

import "gonum.org/v1/gonum/mat"

type AlgoritmoStruct struct {
	matriz10x10Facebook [][]string
	matriz10x10WP       [][]string

	patron1 mat.Matrix //vector de 1 row con 1 y -1
	patron2 mat.Matrix
	D11     *mat.Dense //matriz producto de p1 x p1 traspuesta condiagonal cero
	D22     *mat.Dense
	pesoRed *mat.Dense
}
