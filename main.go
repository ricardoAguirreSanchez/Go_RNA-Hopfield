package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ricardoAguirreSanchez/Go_RNA-Hopfield/algoritmo"
	"github.com/ricardoAguirreSanchez/Go_RNA-Hopfield/formulario"
)

func main() {

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	peso := algoritmo.Aprende()

	router.GET("/", func(c *gin.Context) {

		formularioResultado := formulario.FormularioResultado{}
		formularioResultado.FormularioInput = formulario.Formulario{algoritmo.NuevaMatriz()}
		formularioResultado.FormularioOutput = formulario.Formulario{algoritmo.NuevaMatriz()}

		//cargo la pag con los dibujos conocidos
		formularioResultado.FormularioBase1 = formulario.Formulario{algoritmo.CrearMatrizLinterna()}
		formularioResultado.FormularioBase2 = formulario.Formulario{algoritmo.CrearMatrizFlash()}
		formularioResultado.FormularioBase3 = formulario.Formulario{algoritmo.CrearMatrizBatman()}
		formularioResultado.FormularioBase4 = formulario.Formulario{algoritmo.CrearMatriz4Fantasticos()}
		formularioResultado.FormularioBase5 = formulario.Formulario{algoritmo.CrearMatrizSpiderman()}
		formularioResultado.FormularioBase6 = formulario.Formulario{algoritmo.CrearMatrizThor()}

		c.HTML(http.StatusOK, "index.tmpl.html", formularioResultado)
	})

	router.POST("/calcular", func(c *gin.Context) {

		formularioInput := formulario.GetInput(c)
		formularioOutput := algoritmo.AplicoBusqueda(formularioInput, peso)

		//cargo la pag con resultado y el ingresado
		formularioResultado := formulario.FormularioResultado{}
		formularioResultado.FormularioInput = formularioInput
		formularioResultado.FormularioOutput = formularioOutput

		//cargo la pag con los dibujos conocidos
		formularioResultado.FormularioBase1 = formulario.Formulario{algoritmo.CrearMatrizLinterna()}
		formularioResultado.FormularioBase2 = formulario.Formulario{algoritmo.CrearMatrizFlash()}
		formularioResultado.FormularioBase3 = formulario.Formulario{algoritmo.CrearMatrizBatman()}
		formularioResultado.FormularioBase4 = formulario.Formulario{algoritmo.CrearMatriz4Fantasticos()}
		formularioResultado.FormularioBase5 = formulario.Formulario{algoritmo.CrearMatrizSpiderman()}
		formularioResultado.FormularioBase6 = formulario.Formulario{algoritmo.CrearMatrizThor()}
		c.HTML(http.StatusOK, "result.tmpl.html", formularioResultado)
	})
	router.Run(":8081")
}
