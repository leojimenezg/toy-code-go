# Go

## References
* [Go Documentation](https://go.dev/doc/)
* [A Tour of Go](https://go.dev/tour)

## Basics
### Packages
Cualquier programa de Go se compone de *packages*, y el punto de entrada de cualquier programa es el paquete *main*. Por convención, el nombre de cualquier paquete es el mismo que el último elemento del *path* de import.
* Ejemplo:
    * Import path: "math/rand"
    * Module name: "package rand"

### Exported names (Public, Private)
En Go, un nombre contenido en un paquete, ya sea de una función, constante, etc., solo va a ser exportado si ese nombre inicia con una letra mayúscula. Básicamente, que un nombre inicie con mayúscula lo hace público. Por lo tanto, al importar un paquete, solo se pueden usar los nombres exportados.
* Ejemplo:
    * Pi es un nombre exportado
    * pi no es un nombre exportado

### Functions
Las funciones pueden tomar cero o más argumentos, y su funcionamiento es lo usual. Sin embargo, es importante tener en cuenta que en los argumentos, primero se pone el nombre del argumento y luego el tipo de dato, contrario a lo usual.

De acuerdo al artículo del blog de Go [Go's Declaration Syntax](https://go.dev/blog/declaration-syntax), estas son las razones por las que su sintaxis es diferente a la de los lenguajes de la familia C:
* La sintaxis de los lenguajes de la familia C para la declaración de identificadores, como variables o funciones, **envuelve** el nombre de los items, pues el especificador de tipo está a la izquierda del nombre y otras características o valores están a su derecha. Esto, si bien es una forma inusual (aparentemente) e inteligente, en algunos casos, esta forma hacer que la legibilidad del código se dificulte bastante.
* La sintaxis de declaración en Go, toman la idea base donde los nombres están primero (izquierda) y sus especificadores de tipo después (derecha). Por lo tanto, con esta idea base, la legibilidad left-to-right se vuelve muy sencilla y se da naturalmente, en comparación con C, que envuelve el nombre y se vuelve difícil de leer.
* Por otro lado, los punteros son un punto de excepción, pues no puedieron hacer esa reversión de sintaxis, por lo que estos funcionan como lo hacen en C: con la notación * y siendo usados antes del nombre del puntero.
    * Ejemplo:
        ```Go
        var p *int
        x = *p
        ```

Como ya se mencionó, las funciones pueden retornar múltiples valores, siempre y cuando se declaren dentro de paréntesis (). Estos valoes pueden obtener su significado posicionalmente o mediante un nombre.
Cuando los valores de retorno son nombrados, son tratados como si fueran variables definidas al inicio de la función.

### Variables
Para declarar las variables se utiliza la palabra reservada *var*, la cual, declara una lista de variables, y se puede usar la misma sintaxis de los argumentos de una función. Es decir, poner el tipo de variable al final, y si son varias del mismo tipo, solo ponerlo al final.

La declaración con *var* puede ser usada en varios niveles (scope), ya sea a nivel de paquete o de función, y puede incluir inicializadores, uno por variable. Cuando estos se presentan en la declaración, el especificador del tipo puede ser omitido, y la variable tomará el valor de ese inicializador.

Además, se puede usar el operador *:=* dentro de un scope de función para reemplazar la declaración de *var* y utilizar de forma implícita el tipo de valor. Sin embargo, es importante considerar que fuera del scope de funciones, es decir, a nivel de paquete todas las declaraciones deben iniciar con alguna palabra clave, por lo que el uso de este operador a nivel de paquete no es posible.
