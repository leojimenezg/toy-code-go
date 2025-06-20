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

Además, se puede usar el operador *:=* dentro de un scope de función para reemplazar la declaración de *var* y utilizar de forma implícita el tipo del valor asignado (derecha). Sin embargo, es importante considerar que fuera del scope de funciones, es decir, a nivel de paquete todas las declaraciones deben iniciar con alguna palabra reservada, por lo que el uso de este operador a nivel de paquete no es posible.

### Zero Values
A las variables que son declaradas sin un valor inicial explícito se les asigna un *zero value*, es decir, una valor default que depende del tipo:
* Numeric Type: 0
* Boolean Type: false
* String Type: ""

### Basic Types
Go tiene los siguientes tipos de datos, representados en categorías:
* Boolean: bool
* String: string
* Numeric: int, int8, int16, int32, int64
* Unsigned Numeric: uint, uint8, uint16, uint32, uint64, uintptr
* Byte: byte == uint8
* Rune: rune == int32
* Float: float32, float64
* Complex: complex64, complex128

Generalmente, los tipos que no tienen un tamaño especificado, como *int*, *unit* o *uintptr* dependen del Sistema Operativo, ya sea que su tamaño sea de 32-bit o 64-bit.

Go recomienda usar *int* cuando se requiera utilizar un valor numérico (integer), a menos que haya razonez específicas para usar otro tipo.

### Type Conversion
Para convertir un tipo a otro tipo se utiliza la expresión *T(v)*, donde *v* es el valor a convertir y *T* es el tipo al que se va a convertir.
* Ejemplo:
    ```Go
    var i int = 42
    var f float64 = float64(i)
    var u uint = uint(f)
    ```

A diferencia de C, la asignación de items de diferentes tipos en Go requiere explícitamente la conversión de los tipos. Por lo tanto, es evidente que el casteo no se hace automáticamente como en C.

### Constants
Las constantes son declaradas de forma similar a las variables, con la diferencia de que se debe utilizar la palabra reservada *const*. Estas pueden ser the tipo character, string, boolean o numeric. Y, estas siempre deben usar la palabra reservada *const*, por lo que no pueden sar operadores de asignación corta, como lo es el operador *:=*.

### Numeric Constants
Las constantes numéricas funcionan como los constantes normales, pues tienen que ser declaradas usando la palabra reservada *const*. Sin embargo, estas constantes númericas con valores de alta precisión o *high-precison values*. Aquellas en las que no se especifique su tipo tomarán su tipo de valor dependiendo del contexto en el que sean usadas. Por ejemplo, si se declara una constante numérica sin su tipo y es utilizada por una función que recibe un *int* y otra que recibe un *float64*, actuará de forma correpondiente a dichas funciones.

### Loops
#### For loop
Go solo tiene un constructor para el *for* loop. Se utiliza la palabra reservada *for*, no requiere paréntesis y sus tres componentes son separados por punto y coma (;):
* *Initial statement*: Esta es ejecutada antes de realizar la primera iteración del loop. Ejemplo: `i := 0`.
* *Condition expression*: Esta es evaluada antes de cada iteración. Ejemolo: `i < 10`.
* *Post statement*: Esta es ejecutada al final de cada iteración. Ejemplo: `i++`.

En este tipo de loops, tanto la *initial statement* como la *post statement* son opcionales, por lo que pueden ser omitidos asegurando que la *condition expression* pueda terminar el loop.

#### While loop
En Go no existe un constructor *while*, pues este es reemplazado por el *for* loop, ya que este loop permite declarar únicamente la expresión de condición que hace que el loop continúe.

Sin embargo, es importante tener en cuanto que para cualquier tipo de loop que utilice la palabra reservada *for*, si se omiten todos sus componentes se convierte en un loop infinito.

### If condition
La estructura de este tipo de condición es muy parecida a la estructura del for loop, las condiciones *if* no necesitan estar entre paréntesis, pero sí es obligatorio que estén entre llaves {}. Además, pueden iniciar con un *initial statement* que se realizará antes de evaluar la condición, pero es importante tener en cuenta que las variables declaradas inicialmente en la condición solamente están en al alcance del bloque if donde fueron declaradas (incluyendo *else*).

### Switch
Esta declaración funciona como en la mayoría de lenguajes y se usa la palabra reservada *switch*. Sin embargo, a diferencia de los otros lenguajes, Go únicamente ejecuta el *case* donde la condición es cumplida, y no las que siguen (como lo hacen otros lenguajes), por lo que el uso de *break* es automáticamente hecho por Go.

Go evalúa las condiciones de arriba hacia abajo, y se detiene en la condición que se cumpla, y no continuará ejecutando las condiciones restantes, como lo hacen otros lenguajes como C. También, es posible usar un *switch* sin una condición, y eso será evaluado como un *true*.

Hay dos consideraciones importantes al momento de usar un *switch* son que: el valor usado en cada *case* no deben ser valores constantes o *constants*; y los valores involucrados no deben ser *integers*.

### Defer
Una declaración *defer* hace que la ejecución de una función sea realizada inmediatamente antes de que la función que las rodea finalice, o en su defecto, retorne un valor. Además, este tipo de funciones utiliza un stack especial y utiliza el orden LIFO (Last In - First Out), lo que es crucial para entender su funcionamiento y orden de ejecución.

### Pointers
