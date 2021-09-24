# Trabajo Practico Seminario Go 2021

_Realizamos la consigna solicitda teniendo en cuenta las pautas de esta._

## Comenzando 🚀

Comenzamos creando los documentos necesarios los cuales son:
```
  * main.go contiene las funciones:
    - divideString: se le pasa por parametro un string y se lo separa en caracteres generando un array de caracteres.
    - asign: se le pasa por parametro el array de caracteres para posteriormente asignarlos a cada valor de la estructura Result.
    - main: llama a las funciones divideString, pasando un string determinado, y asign y muestra en consola el resultado retornado por el ultimo metodo.
    
  * La carpeta model que contiene:
    - entities.go contiene:
      *type Result struct {
            Type    string //los primeros dos caracteres son el tipo y puede ser del tipo TX o NN.
            Length  int //los segundos dos caracteres son el largo del valor.
            Value   string  //los siguientes N caracteres son el valor, donde N es el valor del bullet anterior.
       }
      *NewResult: dados diferentes parametros correspondientes a la estrucutura Result retorna una estructura de este tipo.
    - entities_test.go es el test del docuemnto entities.go y contiene:
      * TestNewResult: crea valores que deberían dar error para posteriormente insertarlos en la estructura. Se visualizan en pantalla mensajes de error si estos no coinciden con la estructura.
  
  * main_test.go: es el test del documento main.go y contiene:
    - TestDivideString: se le pasa por parametro un string y se lo separa en caracteres generando un array de caracteres.
    - TestAsign: se crea una estructura de casos para testear. Se inserta en la estructura Result y se crean aserciones para cada atributo de la estructura. 
  
  * out: archivo para guardar la salida del test sobre un archivo.
  
  * out.html: visualizacion de la salida que obtiene el documento out.
  
  * go.mod: generado automaticamente cuando se ejecuta el comando go.mod.init.go2021
  
  * go.sum: generado automaticamente cuando se ejecuta el comando go.mod.init.go2021

### Comandos utilizados📋

Para la creación del módulo para importar archivos internamente: go mod init go2021

Para que encuentre los manejadores de dependencias: go mod tidy

Para ejecutar los archivos: go run <nombre_archivo>.go

Para testing:
  Para saber que porcentaje se encuentra testeado y guardar la salida de test en el archivo out: go test ./... -coverprofile=out.test

  Para analizar el archivo de salida y observar lo testeado (en verde) y lo no testeado (en rojo) en html: go tool cover -html out -o out.html

## Autores ✒️

- Cecilia Carlón: [CeciliaCarlon](https://github.com/CeciliaCarlon)
- Magalí Médico: [MagaMedico](https://github.com/MagaMedico)
- Magalí Menchón: [magalimenchon](https://github.com/magalimenchon)
