
# Go RNA-Hopfield

Una implementación de red neuronal Hopfield para reconocimiento de patrones en Go.

Esta aplicación implementa el algoritmo de red neuronal Hopfield para el reconocimiento y recuperación de patrones, específicamente diseñada para reconocer patrones de superhéroes.

## Características

- Implementación del algoritmo Hopfield para redes neuronales
- Interfaz web interactiva para entrada de patrones
- Reconocimiento de patrones de superhéroes (Linterna Verde, Batman, 4 Fantásticos)
- Visualización de resultados en tiempo real
- Ejecución con Go nativo o Docker
- Arquitectura web moderna con Gin framework

## Ejecutando la Aplicación

### Forma 1: Usando Go directamente

Asegúrate de tener [Go](http://golang.org/doc/install) versión 1.21 o más nueva instalada.

```sh
# Clonar el repositorio
$ git clone https://github.com/ricardoAguirreSanchez/Go_RNA-Hopfield.git
$ cd Go_RNA-Hopfield

# Descargar dependencias
$ go mod tidy

# Ejecutar directamente
$ go run main.go

# O compilar y ejecutar
$ go build -o bin/Go_RNA-Hopfield -v .
$ ./bin/Go_RNA-Hopfield
```

### Forma 2: Usando Docker

Asegúrate de tener [Docker](https://docs.docker.com/get-docker/) instalado.

```sh
# Clonar el repositorio
$ git clone https://github.com/ricardoAguirreSanchez/Go_RNA-Hopfield.git
$ cd Go_RNA-Hopfield

# Construir la imagen Docker
$ docker build -t go-rna-hopfield .

# Ejecutar el contenedor
# La aplicación usa puerto 8081 por defecto, mapeamos al puerto 8080 del host
$ docker run -d -p 8080:8081 --name rna-hopfield-app go-rna-hopfield

# Verificar que esté ejecutándose
$ docker ps

# Ver logs del contenedor
$ docker logs rna-hopfield-app

# Detener el contenedor
$ docker stop rna-hopfield-app

# Eliminar el contenedor
$ docker rm rna-hopfield-app

# Alternativa: usar puerto 8081 directamente
$ docker run -d -p 8081:8081 --name rna-hopfield-app go-rna-hopfield
```

### Acceder a la aplicación

Una vez ejecutada con cualquiera de los métodos, la aplicación estará disponible en:

**Con Go directamente:**
- **http://localhost:8081** (puerto por defecto)

**Con Docker:**
- **http://localhost:8080** (si usas `-p 8080:8081`)
- **http://localhost:8081** (si usas `-p 8081:8081`)

## Documentación

Para más información sobre el algoritmo Hopfield y redes neuronales:

- [Hopfield Network](https://en.wikipedia.org/wiki/Hopfield_network)
- [Neural Networks](https://en.wikipedia.org/wiki/Neural_network)
