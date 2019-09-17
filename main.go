package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/ricardoAguirreSanchez/tp2-rna-hopfield/algoritmo"

	"github.com/ricardoAguirreSanchez/tp2-rna-hopfield/formulario"
)

func main() {
	port := os.Getenv("PORT")
	// port := "8080"

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	algoritmoStruct := algoritmo.Aprende()

	router.GET("/", func(c *gin.Context) {

		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	router.POST("/calcular", func(c *gin.Context) {

		formularioInput := formulario.GetInput(c)

		formularioOutput := algoritmo.AplicoBusqueda(formularioInput, algoritmoStruct)

		fmt.Printf("%+v\n", formularioOutput)

		//cargo la pag con resultado
		formularioResultado := formulario.FormularioResultado{formularioInput, formularioOutput}
		c.HTML(http.StatusOK, "result.tmpl.html", formularioResultado)
	})
	router.Run(":" + port)
}
