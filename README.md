# Go

## References
* [Go Documentation](https://go.dev/doc/)
* [A Tour of Go](https://go.dev/tour)

## Basics
### Packages
Cualquier programa de Go se compone de `packages`, y el punto de entrada de cualquier programa es el paquete `main`. Por convención, el nombre de cualquier paquete es el mismo que el último elemento del *path* de import.
* Ejemplo:
    * Import path: `math/rand`
    * Module name: `package rand`

### Exported names (Public, Private)
En Go, un nombre contenido en un paquete, ya sea de una función, constante, etc., solo va a ser exportado si ese nombre inicia con una letra mayúscula. Básicamente, que un nombre inicie con mayúscula lo hace público. Por lo tanto, al importar un paquete, solo se pueden usar los nombres exportados.
* Ejemplo:
    * `Pi` es un nombre exportado
    * `pi` no es un nombre exportado

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
Para declarar las variables se utiliza la palabra reservada `var`, la cual, declara una lista de variables, y se puede usar la misma sintaxis de los argumentos de una función. Es decir, poner el tipo de variable al final, y si son varias del mismo tipo, solo ponerlo al final.

La declaración con `var` puede ser usada en varios niveles (scope), ya sea a nivel de paquete o de función, y puede incluir inicializadores, uno por variable. Cuando estos se presentan en la declaración, el especificador del tipo puede ser omitido, y la variable tomará el valor de ese inicializador.

Además, se puede usar el operador `:=` dentro de un scope de función para reemplazar la declaración de `var` y utilizar de forma implícita el tipo del valor asignado (derecha). Sin embargo, es importante considerar que fuera del scope de funciones, es decir, a nivel de paquete todas las declaraciones deben iniciar con alguna palabra reservada, por lo que el uso de este operador a nivel de paquete no es posible.

### Zero Values
A las variables que son declaradas sin un valor inicial explícito se les asigna un *zero value*, es decir, una valor default que depende del tipo:
* Numeric Type: `0`
* Boolean Type: `false`
* String Type: `""`

### Basic Types
Go tiene los siguientes tipos de datos, representados en categorías:
* Boolean: `bool`
* String: `string`
* Numeric: `int`, `int8`, `int16`, `int32`, `int64`
* Unsigned Numeric: `uint`, `uint8`, `uint16`, `uint32`, `uint64`, `uintptr`
* Byte: `byte` == `uint8`
* Rune: `rune` == `int32`
* Float: `float32`, `float64`
* Complex: `complex64`, `complex128`

Generalmente, los tipos que no tienen un tamaño especificado, como `int`, `unit` o `uintptr` dependen del Sistema Operativo, ya sea que su tamaño sea de 32-bit o 64-bit.

Go recomienda usar `int` cuando se requiera utilizar un valor numérico (integer), a menos que haya razonez específicas para usar otro tipo.

### Type Conversion
Para convertir un tipo a otro tipo se utiliza la expresión `T(v)`, donde `v` es el valor a convertir y `T` es el tipo al que se va a convertir.
* Ejemplo:
    ```Go
    var i int = 42
    var f float64 = float64(i)
    var u uint = uint(f)
    ```

A diferencia de C, la asignación de items de diferentes tipos en Go requiere explícitamente la conversión de los tipos. Por lo tanto, es evidente que el casteo no se hace automáticamente como en C.

### Constants
Las constantes son declaradas de forma similar a las variables, con la diferencia de que se debe utilizar la palabra reservada `const`. Estas pueden ser the tipo character, string, boolean o numeric. Y, estas siempre deben usar la palabra reservada `const`, por lo que no pueden sar operadores de asignación corta, como lo es el operador `:=`.

### Numeric Constants
Las constantes numéricas funcionan como los constantes normales, pues tienen que ser declaradas usando la palabra reservada `const`. Sin embargo, estas constantes númericas con valores de alta precisión o `high-precison values`. Aquellas en las que no se especifique su tipo tomarán su tipo de valor dependiendo del contexto en el que sean usadas. Por ejemplo, si se declara una constante numérica sin su tipo y es utilizada por una función que recibe un `int` y otra que recibe un `float64`, actuará de forma correpondiente a dichas funciones.

### Loops
#### For loop
Go solo tiene un constructor para el `for` loop. Se utiliza la palabra reservada `for`, no requiere paréntesis y sus tres componentes son separados por punto y coma `;`:
* ***Initial statement***: Esta es ejecutada antes de realizar la primera iteración del loop. Ejemplo: `i := 0`.
* ***Condition expression***: Esta es evaluada antes de cada iteración. Ejemolo: `i < 10`.
* ***Post statement***: Esta es ejecutada al final de cada iteración. Ejemplo: `i++`.

En este tipo de loops, tanto la *initial statement* como la *post statement* son opcionales, por lo que pueden ser omitidos asegurando que la *condition expression* pueda terminar el loop.

#### While loop
En Go no existe un constructor *while*, pues este es reemplazado por el `for` loop, ya que este loop permite declarar únicamente la expresión de condición que hace que el loop continúe.

Sin embargo, es importante tener en cuanto que para cualquier tipo de loop que utilice la palabra reservada `for`, si se omiten todos sus componentes se convierte en un loop infinito.

### If condition
La estructura de este tipo de condición es muy parecida a la estructura del for loop, las condiciones `if` no necesitan estar entre paréntesis, pero sí es obligatorio que estén entre llaves `{}`. Además, pueden iniciar con un *initial statement* que se realizará antes de evaluar la condición, pero es importante tener en cuenta que las variables declaradas inicialmente en la condición solamente están en al alcance del bloque if donde fueron declaradas (incluyendo `else`).

### Switch
Esta declaración funciona como en la mayoría de lenguajes y se usa la palabra reservada `switch`. Sin embargo, a diferencia de los otros lenguajes, Go únicamente ejecuta el `case` donde la condición es cumplida, y no las que siguen (como lo hacen otros lenguajes), por lo que el uso de `break` es automáticamente hecho por Go.

Go evalúa las condiciones de arriba hacia abajo, y se detiene en la condición que se cumpla, y no continuará ejecutando las condiciones restantes, como lo hacen otros lenguajes como C. También, es posible usar un `switch` sin una condición, y eso será evaluado como un `true`.

Hay dos consideraciones importantes al momento de usar un `switch` son que: el valor usado en cada `case` no deben ser valores constantes o `const`; y los valores involucrados no deben ser `int`.

### Defer
Una declaración `defer` hace que la ejecución de una función sea realizada inmediatamente antes de que la función que las rodea finalice, o en su defecto, retorne un valor. Además, este tipo de funciones utiliza un stack especial y utiliza el orden `LIFO` (Last In - First Out), lo que es crucial para entender su funcionamiento y orden de ejecución.

### Pointers
Los punteros almacenan la dirección de memoria de un valor. Es decir, que un puntero apunta a la dirección de memoria donde es almacenado un valor, pero no apunta directamente al valor. Cabe resaltar que a diferencia de C, Go no tiene aritmética de punteros.

Para trabajar con punteros hay dos operadores clave:
* `*`: Este operador indica que una variable es un puntero de cierto tipo. Por ejemplo:
    ```Go
    var p *int // Donde "p" apunta a una dirección de memoria que almacena un valor de tipo "int".
    ```
* `&`: Este operador obtiene la dirección de memoria de una variable. Por ejemplo:
    ```Go
    var i int = 42
    p = &i // Donde "&i" genera la dirección de memoria de i.
    ```
* Al trabajar con punteros, se pueden trabajar con dos valores diferentes, de distinta forma.
    * Si se quiere usar la dirección a la que apunta, se usa `p` (nombre del puntero).
    * SI se quiere usar el valor de la dirección de memoria a la que apunta, se usa `*p` (el signo `*` y el nombre del puntero).

### Structs
Una estructura es una colección de `fields` o *miembros*. Para acceder a los miembros de una estructura de usa la notación de punto `.` y el nombre exacto del miembro. Ejemplo:
```Go
type T struct {
    X int
    Y int
}
var t1 T = T{ 1, 2 }
```

Los punteros que almacenan la dirección de memoria de una estructura se comportan de manera un poco diferente. Pues, normalmente, para acceder a los miembros de una estructura mediante un puntero se utilizaría la expresión `(*p).X`, sin embargo, Go permite una expresión más sencilla `p.X`. Esto es posible ya que una estructura es la dirección de memoria base del bloque donde se almacenan todos sus miembros, y para acceder a ellos se utilizan offsets basados en el tipo de dato del miembro y del padding agregado por el compilador. Ejemplo:
```Go
(*p).X = 3 // Original form
p.Y = 4    // Simplier form
```

Al crear una variable de un tipo que es una estructura, se pueden dar tres casos:
* Se pueden especificar los valores de uno o más miembros de una estructura de forma posicional, es decir, que los valores serán asignados dependiendo del orden de declaración de los miembros. Ejemplo:
    ```Go
    v1 := T{ 5, 6 } // X = 5, y Y = 6
    ```
* Se pueden especificar los valores de uno o más miembros de una estructura usando sus nombres, lo que resulta más claro y en algunos casos, más conveniente. Ejemplo:
    ```Go
    v2 := T{ X: 7, Y: 8 } // X = 7, y Y = 8. Se le llama string literals.
    ```
* Si no se especifican los valores de uno o más miembros de una estructura, a dichos miembros se les asignará su *zero value*. Ejemplo:
    ```Go
    v3 := T{} // X = 0, y Y = 0
    ```
* **Nota**: No se pueden combinar los dos primeros casos para asignar valores usando su posición y su nombre.

### Arrays
Los arreglos en Go funcionan de forma completamente diferente a los arreglos en C. Pues en C, un array siempre se utiliza por referencia, es decir, por su dirección de memoria; mientras que en Go, un array simpre se utiliza por valor, por lo que siempre se apunta a todos los valores del array en lugar de su dirección base.

Entonces, para declarar un array se utiliza la expresión `[n]T`, donde `n` es la cantidad de elementos del array, y `T` es el tipo de valor de los elementos. El tamaño del arreglo es parte de su tipo, por lo que son de tamaño fijo. Ejemplo:
```Go
var a1 [2]int // Array de dos elementos de tipo entero
a[0], a[1] = 3, 4
```

### Slices
Un slice tiene tamaño dinámico y es una forma flexible de acceder a los elementos de un array. Un slice es formado al especificar dos índices: `low bound` y `high bound` separados por `:`. Por defecto, el zero value de `low bound` es `0`, y el del `high boud` es la longitud del array. Ejemplo:
```Go
array := [6]int{1, 2, 3, 4, 5, 6}
var s []int = array[1:4]
// Slice de 3 elementos, donde se incluye el low bound (1) pero no el high bound (4).
```

A diferencia de un array, un slice no almacena realmente los valores, sino que describe una sección de un array existente o completo, es decir, que trabaja con la referencia de un array. Por lo tanto, como un slice trabaja con la referencia de un array, al modificar el slice, el arreglo y otros slices dependientes cambian de igual forma.

Cualquier slice tiene dos características fundamentales:
* `length`: La longitud es el número de elementos que contiene el slice. Se usa la expresión `len()` para obtener la longitud.
* `capacity`: La capacidad es el número de elementos que contiene el array contando desde el primer elemento del que parte el slice. Se usa la expresión `cap()` para obtener la capacidad.

El zero value de un slice es `nil`, y un slice vacío tiene una longitud y capacidad de 0, por lo que no existe un array del cual partir.

Otra forma de crear slices es con la función integrada `make()`, pues esta se encarga de alojar un array del tamaño especificado, llenarlo de ceros y devolver un slice que apunta al array creado. Ejemplo:
```Go
a := make([]int, 0, 5)
// Esto crea un slice de elementos int, con longitud de 0 y capacidad de 5
```

Go integra la función `append()` para agregar nuevos elementos a un slice. Esta función recibe el slice al cual se le van a agregar los elementos, y recibe los elementos a agregar. El resultado de esta función es un slice que contiene los elementos originales más los nuevos elementos agregados; si la capacidad del array no puede soportar los nuevos elementos, se crea uno nuevo para poder almacenar todos los elementos y se utiliza ese nuevo array. Ejemplo:
```Go
var s []int
s = append(s, 0, 1)
s = append(s, 2, 3, 4)
```

### Range
Utilizar `range` es una forma corta de recorrer los elementos de un array, slice o map, incluso de recorrer un rango numérico. Cuando se utiliza para recorrer algún objeto iterable, se obtienen dos valores: primero, el index del elemento, y segundo, una copia del elemento.

Alguno de los dos valores que se obtienen de usar `range` puede ser ignorado usando `_` o simplemente omitiendo la variable donde se almacenaría dicho elemento.

### Maps
