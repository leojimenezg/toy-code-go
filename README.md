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

De acuerdo al artículo del blog de Go [Go's Declaration Syntax](https://go.dev/blog/declaration-syntax), estas son las razones por las que su sintáxis es diferente a la de los lenguajes de la familia C:
* La sintáxis de los lenguajes de la familia C para la declaración de identificadores, como variables o funciones, **envuelve** el nombre de los items, pues el especificador de tipo está a la izquierda del nombre y otras características o valores están a su derecha. Esto, si bien es una forma inusual (aparentemente) e inteligente, en algunos casos, esta forma hacer que la legibilidad del código se dificulte bastante.
* La sintáxis de declaración en Go, toma la idea base donde los nombres están primero (izquierda) y sus especificadores de tipo después (derecha). Por lo tanto, con esta idea base, la legibilidad **left-to-right** se vuelve muy sencilla y se da naturalmente, en comparación con C, que envuelve el nombre y se vuelve difícil de leer.
* Por otro lado, los punteros son un punto de excepción, pues no puedieron hacer esa reversión de sintáxis, por lo que estos funcionan como lo hacen en C: con la notación `*` y siendo usados antes del nombre del puntero.
    * Ejemplo:
        ```Go
        var p *int
        x = *p
        ```

Como ya se mencionó, las funciones pueden retornar múltiples valores, siempre y cuando se declaren dentro de paréntesis `()`. Estos valores pueden obtener su significado posicionalmente o mediante un nombre. Cuando los valores de retorno son nombrados, son tratados como si fueran variables definidas al inicio de la función.

#### Function values
Las funciones también son valores, como cualquier otro valor (int, string, array, etc), por lo que pueden ser usados como cualquier otro valor, por ejemplo, como argumentos de otras funciones o como valores de retorno de funciones.

#### Function closure
Además, las functiones pueden ser `closures` de otras funciones, es decir, que pueden actuar como funciones anónimas que utilizan variables fuera de su alcance (función padre) y que son regresadas por su función padre. Esto provoca que el estado de las variables utilizadas dentro del *closure* que pertenecen a la función padre mantengan su estado hasta que esa *function closure* ya no se use más.

### Variables
Para declarar las variables se utiliza la palabra reservada `var`, la cual, declara una lista de variables, y se puede usar la misma sintáxis de los argumentos de una función. Es decir, poner el tipo de variable al final, y si son varias del mismo tipo, solo ponerlo al final.

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
Utilizar `range` es una forma corta de recorrer los elementos de un array, slice o map, incluso de recorrer un rango numérico. Cuando se utiliza para recorrer algún objeto iterable, se obtienen dos valores: primero, el index o llave del elemento, y segundo, una copia del elemento.

Alguno de los dos valores que se obtienen de usar `range` puede ser ignorado usando `_` o simplemente omitiendo la variable donde se almacenaría dicho elemento.

### Maps
Un `map` es básicamente un arreglo de valores que son accedidos mediante llaves en lugar de índicies, por lo que toman la forma de `key-value`.

El *zero value* de un mapa en `nil`, y cuando un mapa que tiene este valor no tiene ninguna llave y no se le pueden agregar nuevas llaves. Sin embargo, se puede utilizar la función `make`, similar a su uso en slices, pues esta función devuelve un mapa de cierto tipo inicializado y listo para ser usado.

Se pueden realizar varias operaciones con los mapas, como:
* Insertar o Actualizar un elemento del mapa:
    ```Go
    m["Answer"] = 42
    // Si la llave "Answer" existe, actualiza el valor.
    // Si la llave "Answer" no existe, crea la llave y asigna el valor.
    ```
* Obtener un elemento del mapa:
    ```Go
    value = m["Answer"]
    // Regresa el valor correspondiente a la llave "Answer".
    ```
* Eliminar un elemento del mapa:
    ```Go
    delete(m, "Answer")
    // delete(mapa, llave)
    ```
* Verificar que una llave exista en el mapa:
    ```Go
    value, ok := m["Answer"]
    // Si existe, regresa el valor y la segunda variable (ok) es true.
    // Si no existe, regresa el zero value y la segunda variable es false.
    ```

## Methods and interfaces
### Methods
Go no tiene POO y por lo tanto, no tiene clases ni objetos. Sin embargo, es posible agregar métodos a los `types`, pues un método es básicamente una función pero con un `receiver argument` especial que aparece después de la keyword `func` pero antes del nombre del método. Dicho `receiver argument` asocia el método con el tipo de ese argumento, es decir, especifica qué `type` posee el método declarado. Ejemplo:
```Go
type Struct1 struct {
    X, Y float64
}
func (s1 Struct1) Sum() float64 {
    return s1.X + s1.Y
}
```

Los métodos pueden ser declarados con cualquier `type`, siempre y cuando el `type` se encuentre definido en el mismo paquete donde se declara el método; y que no sea un `type` predefinido por el lenguaje.

Entonces, básicamente un método es una función que recibe un objeto de un tipo especificado y dicho método es asociado al tipo del objeto recibido, no al objeto como tal. Existen dos formas por las que un método puede recibir el objeto de cierto tipo:
* `Value receiver`: Cuando se recibe un objeto por valor, se está recibiendo una copia del objeto. Por lo tanto, nunca se puede trabajar ni modificar el objeto real.
* `Pointer receiver`: Cuando se recibe un objeto por referencia o puntero, se está recibiendo la dirección de memoria del objeto real. Por lo tanto, se trabaja con el objeto real y cualquier modificación será visible en otras partes que usen el mismo objeto.
* **Nota: Un método debe usar o value receivers o pointer receivers, pero no se deben combinar.**

### Interfaces
Una interfaz es un objeto de tipo `interface` que define un conjunto de declaraciones de métodos. Es decir, que una interfaz es un contrato de métodos, y cualquier tipo que implemente los métodos definidos en la interfaz puede usar el tipo interfaz.

Es importante mencionar que cualquier interfaz almacena información sobre el objeto asignado a una variable de tipo interfaz en forma de tupla, pues almacena el valor del objeto y el tipo del objeto para saber a qué método debe llamar: `(value, type)`.

Básicamente, una interfaz establece un conjunto de métodos que un tipo debe implementar para poder usar un tipo de interfaz específico. Las interfaces son implementadas cuando un tipo implementa los métodos de la interfaz, por lo que no se necesitan palabras clave o declaraciones explícitas. Además, cuando se llama a un método implementado por la interfaz, se ejecuta el método que tiene nombre específico y el tipo específico, por lo que una interfaz guarda la información del tipo que implementan sus métodos.

Cuando un objeto de tipo interfaz contiene un valor `nil` pero tiene un tipo conocido, los métodos se pueden llamar (el receiver será nil). Sin embargo, si la interfaz misma es `nil` (no tiene tipo ni valor asignado), llamar un método causará un error en tiempo de ejecución, ya que Go no sabe qué método ejecutar.

Existen también las interfaces vacías que no implementan ningún método, y son conocidas como `empty interface`. Estas pueden guardar valores de cualquier tipo.

### Type assertions
El tipo `assertion` provee acceso al tipo de dato y al valor que contiene el objeto de una interfaz. Ejemplo:
```Go
t := i.(T)
// Donde "i" es el objeto de tipo interface y "T" es un tipo (int, string, etc.).
```
En el ejemplo anterior, la assertion se asegura de que el tipo del valor de un objeto de una interfaz es de cierto tipo (T), y asigna el valor a la variable (t) si es que es verdad. En caso contrario de que el objeto no sea del tipo (T) se produce un error.

Sin embargo, existe otra forma de usar assertions para ver si el objeto de una interfaz es de un tipo específico:
```Go
t, ok := i.(T)
```
En este caso, la assertion regresa el valor del objeto del tipo especificado (T) a la variable (t) y también regresa `true` (ok) si es que es verdad; pero si el objeto de la interfaz no es del tipo especificado (T), regresa el `zero value` del tipo (t) y `false` (ok), pero no se produce ningún error.

### Type switches
Un `type` dentro de los switches es un constructo que permite realizar varias `assertions` de diferentes tipos. En estos casos, el funcionamiento es exactamente el mismo que cualquier otro `switch`, con la diferencia de que los `case` evalúan diferentes `type` en lugar de valores (como int), y se usa la palabra reservada `type` dentro de la assertion. Ejemplo:
```Go
func do(i interface{}) {
    switch v := i.(type) {
    // La variable "v" toma el valor del objeto de la interfaz y su tipo.
    case int:
        // El objeto de la interfaz es de tipo int.
    case string:
        // El objeto de la interfaz es de tipo sring.
    default:
        // El objeto de la interfaz es de otro tipo.
    }
}
```

### Stringers
Dentro del paquete `fmt` se encuentra definida la interfaz `Stringer`. Esta interfaz implementa el método `String()`, y cualquier objeto que la implemente puede describirse a sí mismo como un string. El paquete `fmt` y muchos otros buscan esta interfaz en los objetos cuando se imprime dicho objeto, por lo que no es necesario llamar al método explícitamente.
```Go
type Stringer interface {
    String() string
}
```

### Errors
Los programas en Go expresan estados de error con valores de tipo `error`. Este tipo es una interfaz ya integrada en el lenguaje e implementa el método `Error()`, similar al método implementado por la interfaz Stringer, el paquete `fmt` y otros buscan este método al momento de imprimir objetos de tipo error.
```Go
type error interface {
    Error() string
}
```

Generalmente, las funciones regresan un valor de tipo `error`, por lo que el código debería manejar y considerar los errores al verificar dicho valor es igual a `nil` o no. Ejemplo:
```Go
i, err := strconv.Atoi("42")
if err != nil {
    // La función devolvió un error (error != nil).
}
// La función no devolvió un error (error == nil).
```

### Readers
El paquete `io` especifica la interfaz `io.Reader`, que representa la lectura de un flujo de datos hasta su fin. La librería estándar de Go tiene muchas implementaciones de esta interfaz, como archivos, conexiones de red, compresores, cifrados, etc.

La interfaz `io.Reader` implementa el método `Read`, y lo que hace este método es llenar un slice de bytes con información y regresar el número de bytes con los que fue llenado el slice, también regresa un valor de tipo error. Cuando el flujo de datos termina, simplemente regresa un error de tipo `io.EOF`.
```Go
func (t T) Read(b []byte) (n int, err error)
```

### Writers
El paquete `io` también especifica la interfaz `io.Writer`, que representa la escritura de un flujo de datos hasta su fin en un cierto destino.

La interfaz `io.Writer` implementa el método `Write`, y lo que hace este método es escribir los datos desde un slice de bytes hacia un cierto destino, también regresa el número de bytes escritos y un valor de tipo error, que puede contener errores al momento de hacer la escritura.
```Go
func (t T) Write(b []byte) (n int, err error)
```

### Images
El paquete `image` define la interfaz `Image`, que sirve para representar algunas de las características de las imágenes.
```Go
type Image interface {
    ColorModel() color.Model  // Retorna el modelo de color (RGB, RGBA, etc).
    Bounds() Rectangle  // Retorna el rectángulo que define el área de la imagen.
    At(x, y int) color.Color  // Retorna el color del pixel en las coordenadas (x,y).
}
```

## Generics
### Type parameters
Las funciones en Go pueden usar `type parameters` para trabajar con múltiples tipos. Los type parameters se definen entre `[]` después del nombre de la función y antes de los argumentos.
```Go
func Index[T comparable](s []T, x T) int {
    for i, v := range s {
        if v == x {
            return i
        }
    }
    return -1
}
```

Los type parameters pueden tener constraints que limitan qué tipos pueden usar:
* `comparable`: Permite usar `==` y `!=`.
* `any`: Acepta cualquier tipo (sin restricciones).
* Constraints personalizadas: Interfaces que definen métodos requeridos.

### Generic types
También se pueden crear tipos genéricos (structs, interfaces, etc.) que funcionen con múltiples tipos:
```Go
type Stack[T any] struct {
    items []T
}

intStack := Stack[int]{items: []int{1, 2, 3}}
stringStack := Stack[string]{items: []string{"a", "b", "c"}}
```

## Concurrency
