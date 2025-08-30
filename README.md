# Go

## References
* [Go Documentation](https://go.dev/doc/)
* [A Tour of Go](https://go.dev/tour)
* [How to Write Go Code](https://go.dev/doc/code)
* [The Go Programming Language Specification](https://go.dev/ref/spec)

***A Tour of Go***

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
### Goroutines
Una `goroutine` es un hilo ligero manejado por Go en `runtime`, básicamente, es una función que se ejecuta concurrentemente (al mismo tiempo) en el programa. La evaluación de los argumentos de dicha función es realizada en la goroutine actual, es decir, en runtime, pero la ejecución de la función ocurre en la `new goroutine` y el momento en que se ejecutan es completamente decidido por el `scheduler` de Go. No son ejecutadas inmediatemente después de crearlas ni después de evaluarlas, sino cuando el `scheduler` lo decide.
```Go
go f(x, y, z)  // Nueva goroutine.
```

Todas las `goroutine` son almacenadas en el mismo espacio de memoria `heap`, por lo que comparten dicho espacio y el acceso a esta memoria compartida debe ser sincronizado. Sin embargo, cada goroutine cuenta con su propio `stack` para almacenar los objetos utilizados en la función que ejecuta.

Un punto que pasa por desapercibido, es que Go, automáticamente mueve las variables del programa que se encuentran en el `stack` al `heap` cuando detecta que esas variables son utilizadas por al menos una `goroutine`, a esto se le llama `escape analysis`. Provoca un poco de overhead (uso adicional de recursos).

Ejemplos:
```Go
x := 5
// Los argumentos se evalúan antes de crear la goroutine.
go func(val int) {
    fmt.Printf("Valor en goroutine: %d\n", val)
}(x) // "x" se evalúa aquí (valor actual: 5)
x = 10 // Cambiar x no afecta la goroutine.
```

```Go
mensaje := "Hola mundo"
// La goroutine accede a la variable desde el heap.
go func() {
    fmt.Println(mensaje) // Accede a la variable compartida.
}()
```

```Go
func problemaDelLoop() {
    // Todas las goroutines ven la misma variable "i".
    for i := 0; i < 3; i++ {
        go func() {
            fmt.Printf("Goroutine ve i = %d\n", i)
        }()
    }
    // Output probable: "3 3 3" (todas ven el valor final de i).
}

func solucionDelLoop() {
    // Pasar "i" como argumento "congela" su valor.
    for i := 0; i < 3; i++ {
        go func(val int) {
            fmt.Printf("Goroutine ve val = %d\n", val)
        }(i) // "i" se evalúa aquí, congelando su valor actual.
    }
    // Output: "0 1 2" (en cualquier orden).
}
```

### Channels
Los canales son un tipo de conducto por los que se pueden enviar y recibir valores, mediante el uso del operador de canal `<-`. La información fluye en la dirección que apunte el operador de canal, y parecido a los maps y slices, los canales deben ser creados antes de ser usados.
```Go
ch := make(chan int)  // Crear un channel antes de usarlo.
ch <- v  // Enviar "v" al canal "ch".
v := <- ch  // Recibir del canal "ch" y asignar el valor a "v".
```

Por defecto, un `unbuffered channel` bloquea todo el envío y recepción de información hasta que ambas partes del canal se encuentren listas y puedan ser completadas, lo  que permite que las goroutines se sincronicen implícitamente. Además, si el otro lado del canal nunca está listo o nunca se utiliza el canal, se produce un `deadlock`, pues todas las goroutines se quedan esperando el resultado de un canal y nada se continúa ejecutando debido a esa espera.

Es imporante reconocer que un canal, con la creación por defecto (unbuffered), no puede realmente almacenar ningún elemento, sino que asigna o devuelve el valor del canal directamente a donde se esté requiriendo dicho valor.

### Buffered channels
A los canales se les puede asignar un `buffer`, es decir, que pueden ser capaces de realmente almacenar cierta cantidad de elementos. Como segundo elemento en la función `make` se puede especificar el tamaño del canal.
```Go
ch := make(chan int, 100)  // Canal de máximo 100 elementos.
```

Sin embargo, un `buffered channel` se bloquea no cuando ambas partes no estén listas, sino que se bloquea cuando está vacío al leer del canal o lleno cuando se escriba en el canal.

Los canales `unbuffered` o dond no se especifica el tamaño del buffer o elementos máximos, es lo mismo que decir que buffer es igual a 0.
```Go
ch1 := make(chan int)     // Puede almacenar 0 elementos.
ch2 := make(chan int, 0)  // Puede almacenar 0 elementos.
// ch1 == ch2
```

### Range and Close
Naturalmente, un canal tiene un `sender` y un `receiver`. Por lo tanto, un `sender` puede usar `close` en un canal, para indicar que ya no se van a enviar más elementos por ese canal. Es decir, que sirve para cerrar un canal, lo que implica que no se puede mandar o almacenar más elementos en dicho canal, sin embargo, sí es posible seguir leyendo los elementos del canal siempre y cuando no esté vacío.

Por otro lado, los receivers pueden comprobar si un canal ya no contiene elementos y ha sido cerrado al recibir un segundo parámetro al asignar un elemento del canal. Los canales nunca son cerrados automáticamente, pero es importante saber en qué casos se necesita especificar que el canal ya no va a enviar elementos y se va a cerrar (como en loops con range).
```Go
v, ok := <-ch
// Si "ok" es "true", todavía hay valores y no ha sido cerrado el canal.
// Si "ok" es "false", ya no hay valores y el canal ha sido cerrado.
```

Además, el uso de `range` en un loop con un canal, hace que se reciban los elementos de dicho canal repetidamente hasta que el canal es cerrado.
```Go
for v := range ch {
    // Trabajar con cada elemento del canal hasta que se cierre.
}
```

Si el canal no es cerrado y se usa en un loop con range, el loop nunca dejará de esperar nuevos elementos del canal, lo que causará deadlock:
```Go
ch := make(chan int)
go func() {
    ch <- 1
    ch <- 2
    // No hay "close()" en el sender del canal.
}()

// No sabe que el canal ya no va a enviar elementos y provoca deadlock.
for v := range ch {
    fmt.Println(v)
}
```

### Select
La palabra reservada `select` permite a una `goroutine` esperar el resultado de múltiples canales, por lo que detiene la ejecución de la rutina hasta que alguno de los casos pueda ser ejecutado. Sin embargo, si múltiples casos pueden ser ejecutados, elige uno aleatoriamente.

Además, se puede hacer uso de `default` junto con esta sentencia para poder ejecutar un caso cuando ningún otro caso pueda ser ejecutado y evitar que la rutina actual se bloquee.
```Go
select {
case <-c:
// Hacer algo porque se recibió algo de c.
case <-quit:
// Hacer algo porque se recibió algo de quit.
default:
// Hacer algo cuando no se reciba nada de c o quit.
}
```

Si no se usa `default`, la `goroutine` se bloqueará hasta que al menos uno de los casos pueda ejecutarse. Select evaluará un único caso: si solo uno está listo, ejecutará ese; si varios están listos, elegirá uno aleatoriamente.

***How to Write Code***

## Code organization
Go se maneja por paquetes y módulos. Un paquete es una colección de archivos de código que se encuentran en un mismo directorio y se compilan juntos. Por otro lado, un módulo es una colección de paquetes de Go relacionados que se publican juntos como una unidad.

Cada módulo y paquete debe ser identificado de alguna forma. Para esto, existe el `module path`, que es un string (nombre) que identifica de forma única al módulo. Mientras que el `import path` es la ruta que identifica a un paquete dentro de un módulo, formada por el module path más la ruta del directorio donde se encuentra el paquete dentro del módulo.

Este sistema de identificación es necesario para que Go pueda localizar y manejar correctamente todos los paquetes y módulos. Para gestionar esta información, se utiliza el archivo `go.mod`, que se crea con el comando `go mod init`. Este archivo contiene el module path, la versión de Go requerida, y las dependencias externas del módulo con sus versiones específicas. El comando `go mod tidy` actualiza automáticamente este archivo, agregando las dependencias que realmente se usan en el código y eliminando las que ya no son necesarias.

## Your first program
Go recomienda seguir los siguientes pasos para crear un proyecto simple:
1. Crear el directorio que contendrá el proyecto.
2. Crear el archivo `go.mod`.
3. Crear los archivos que se vayan a usar.
    * Se debe tener en cuenta que cualquier comando ejecutable siempre debe usar el paquete main.
4. Construir e instalar el proyecto con el comando `go install module-path`.
    * Este comando construye el proyecto al crear un archivo binario ejecutable y lo instala dentro de la ruta `$HOME/go/bin/nombre-programa`
    * La ruta de instalación depende de las variables GOPATH y GOBIN.
5. Ejecutar el archivo binario como resulte conveniente.
    * El nombre del archivo binario ejecutable generado es tomado a partir del identificador del propio módulo.

### Importing packages from your module
Algo importante por aclarar, es que en Go hay dos posibles "resultados finales". Puede ser que se creen archivos binarios ejecutables que utilicen siempre el paquete main; o que se creen librerías/paquetes que no utilicen el paquete main ni sean ejecutables directamente, sino que puedan ser utilizados por otros códigos.

Para poder crear y usar un paquete dentro del mismo módulo, se recomienda seguir los siguientes pasos:
1. Crear el directorio que actuará como el paquete dentro del módulo.
2. Crear los archivos `.go` con el código del paquete.
3. Compilar el paquete con el comando `go build` (opcional para verificación).
    * Este comando no genera ningún archivo de salida visible. Simplemente verifica que el paquete compile correctamente y almacena el resultado en el caché local de Go.
4. Importar el paquete usando el import path completo: `module-path/nombre-del-directorio`.

Las funciones, tipos, variables y constantes que empiecen con letra mayúscula son exportadas (públicas) y pueden ser usadas desde otros paquetes. Las que empiecen con letra minúscula son privadas del paquete.

### Importing packages from remote modules
El import path puede describir cómo obtener el código de un paquete externo al usar sistemas de control de versiones como Git o Mercurial. Go utiliza esta característica para automáticamente obtener los paquetes de repositorios remotos.

Una parte importante de esto, es el comando `go mod tidy`. Este comando se encarga de descargar los paquetes externos que se estén usando en el código y grabar su requerimiento y versión en el archivo `go.mod`. Además, elimina cualquier dependencia que no se esté usando en el mismo archivo.

Las dependencias son automáticamente descargadas en el subdirectorio `pkg/mod` del directorio indicado por la variable `GOPATH`. El contenido descargado para una versión específica de un módulo es compartido entre todos los módulos que requieran esa misma versión, por lo que Go marca esos archivos y directorios como read-only.

Para remover todas las dependencias descargadas, se puede usar el comando `go clean -modcache`.

## Testing
Para hacer pruebas del código, Go ofrece un pequeño framework, para usarlo se necesita el comando `go test` y el paquete estándar `testing`.

Para crear una prueba, se debe crear un archivo con terminación `_test.go` que contenga una función llamada `TestXXX` con la firma `func (t *testing.T)`. La prueba ejecuta la función, y si esta llama a alguna función de error, como `t.Error` o `t.Fail`, se considera que la prueba ha fallado.

Para ejecutar la prueba, simplemente se debe estar dentro del directorio donde se encuentra la prueba y usar el comando `go test`.

***The Go Programming Language Specification***

## Introduction
Go es un lenguaje de programación de propósito general, es compilado y de tipado estático. Diseñado para ser simple pero poderoso. Go hace garbage collection automáticamente, es muy expresivo, maneja programación concurrente. Su idea base es "menos es más", pues deliberadamente no da muchas opciones para hacer algo (como en C), pero las pocas opciones que da son más que suficientes.

## Notation
Esta sección se refiere a la sintáxis usada a lo largo de la especificación del lenguaje para poder representar y explicar varios aspectos. Por lo tanto, es importante entender los símbolos.
* `""` significa "tokens literales (keywords)".
* `|` significa "alternación (OR)".
* `{}` significa "cero o más repeticiones".
* `[]` significa "opcional".
* `()` significa "agrupación".
* `a...b` significa "conjunto de caracteres de 'a' hasta 'b'"

## Source code representation
Go representa cada letra en el código como un código único (code point) del estándar Unicode, y los codifica usando UTF-8, por lo tanto, un archivo de código fuente debe ser una secuencia válida de UTF-8, por lo que puede trabajar con múltiples lenguajes nativamente.

Esto significa que, Go, usa Unicode para representar cada caracter del código usando code points, y luego, usa UTF-8 para convertir esos códigos Unicode en bytes.

## Lexical elements
Los elementos léxicos son los "bloques de construcción" más básicos del lenguaje, es decir, son esos elementos (palabras) de Go consideradas como unidades mínimas que el compilador puede reconocer.

### Comments
Los comentarios sirven como documentación del programa dentro del mismo. Existen dos formas en que se pueden usar:
* `//` Comentarios de línea.
* `/**/` Comentarios generales.
Un comentario no puede iniciar dentro de una `rune`, de un `string literal` o de otro comentario.

### Tokens
Los tokens son básicamente todos los elementos del vocabulario de Go. Se dividen en cuatro clases: `identifiers`, `keywords`, `operators and punctuation` y `literals`.

### Semicolons
La sintáxis formal de Go usa el punto y coma `;` como finalizadores en diversas producciones. Sin embargo, la mayoría de programas peden omitir su uso siempre y cuando se sigan las siguientes dos reglas:
* Cuando la producción es dividida en múltiples tokens, un semicolon es automáticamente insertado inmediatamente después del último token.
* Para permitir producciones complejas que ocupen una sola línea, el semicolon puede ser omitido después de los signos de agrupación terminantes `)` o `}`.

### Identifiers
Los identificadores nombran entidades del programa, como variables o tipos. Estos son una secuencia de una o más letras y dígitos, donde el primer caracter del identificador debe ser una letra.
* `identifier = letter { letter | unicode_digit }`.

### Keywords
Las palabras clave son palabras propias del lenguaje y no pueden ser usadas como identificadores.
* `break`
* `case`
* `chan`
* `const`
* `continue`
* `default`
* `defer`
* `else`
* `fallthrough`
* `for`
* `func`
* `go`
* `goto`
* `if`
* `import`
* `interface`
* `map`
* `package`
* `range`
* `return`
* `select`
* `struct`
* `switch`
* `type`
* `var`

### Operators and punctuation
Existen caracteres o secuencias de caracteres que representan operadores y signos de puntiación.
* `+`
* `-`
* `*`
* `/`
* `%`
* `&`
* `|`
* `^`
* `<<`
* `>>`
* `&^`
* `+=`
* `-=`
* `*=`
* `/=`
* `%=`
* `&=`
* `|=`
* `^=`
* `<<=`
* `>>=`
* `&^=`
* `&&`
* `||`
* `<-`
* `++`
* `--`
* `==`
* `<`
* `>`
* `=`
* `!`
* `~`
* `!=`
* `<=`
* `>=`
* `:=`
* `...`
* `(`
* `[`
* `{`
* `,`
* `.`
* `)`
* `]`
* `}`
* `;`
* `:`

### Integer literals
Un entero literal es una secuencia de dígitos que representan un entero constante. Estos pueden contener un prefijo optional que establece una base no-decimal: `0b` o `0B` para binario; `0`, `0o` o `00` para octal; y `0x` o `0X` para hexadecimal. Además, para mejorar la legibilidad de los dígitos, se puede usar `_` para representar separaciones, pero, debe aparecer únicamente después del prefijo o entre los dígitos, no al inicio ni al final.

### Floating-point literals
Un punto flotante literal es una representación decimal o hexadecimal de una constante de punto flotante. Estos se componen de su parte decimal, el punto decimal, y la parte fraccionaria, incluso pueden contener representación científica mediante el uso de `e` o `E`. Además, igual que los integer literals, pueden usar `_` para mejorar su legibilidad.

### Imaginary literals
Un imaginario literal representa la parte imaginaria de una constante compleja. Consiste de un entero o punto flotante seguido por la letra minúscula `i`, por lo que el valor resultante es el valor del valor entero o flotante multiplicado por la unidad imaginaria. Sin embargo, para evitar problemas de compatibilidad, se recomienda usar `0` al inicio.

### Rune literals
Una rune literal representa a una rune constante, es decir, un valor entero que representa un punto de código de Unicode. Estas son expresadas como uno o más caracteres encerrados en comillas simples ', y cualquier caracter que aparezca dentro de ellas representa el valor de Unicode correspondiente a dicho caracter.

Una `rune` es el concepto que Go usa para representar un carácter completo. A diferencia de `char` en C (que siempre ocupa 1 byte para ASCII), una `rune` puede representar cualquier carácter Unicode, incluso aquellos que ocupan de 1 a 4 bytes en UTF-8. Internamente, `rune` es un alias para `int32`, lo que permite representar todo el rango de códigos Unicode.

### String literals
Un string literal representa un string constante obtenido a partir de concatenar una secuencia de caracteres. Existen dos formas de string literals:
* `Raw string literals`: Son una secuencia de caracteres encerrados entre backticks `\``. Cualquier caracter puede aparecer dentro de los backticks (excepto el mismo backtick). El valor de este tipo es un string formado por los caracteres sin interpretar, es decir, no se procesan caracteres especiales como escape sequences.
* `Interpreted string literals`: Son una secuencia de caracteres encerrados entre comillas dobles `""`. Cualquier caracter puede aparecer dentro de las comillas dobles (excepto newline y la misma comilla doble, a menos que sean escapados). El valor de este tipo es un string formado por los caracteres interpretados, procesando caracteres especiales como escape sequences.

## Constants
Una constante se refiere a un valor cuyo contenido es conocido en tiempo de compilación y nunca cambian de valor a lo largo de la vida del programa. Las constantes en Go pueden ser de tipo boolean, rune, integer, floating-point, complex, string, incluso sin tipo. Las constantes sin tipo pueden tomar adaptarse a distintos contextos.

El valor de cualquier constante en Go es representado por alguno de los tipos o alguna expresión/operación constanteque pueda ser realizada en tiempo de compilación. Además, el compilador puede implementar las constantes con una precisión interna mayor a la normal de runtime.

## Variables
Una variable es un espacio de memoria para almacenar un valor, y está sujeta completamente a su tipo. Una declaración de una variable reserva un espacio de memoria en runtime que es identificado por el nombre dado a dicha variable, incluso, se pueden tener espacios de memoria anónimos, que no son accesidos por su identificador, sino por un puntero que contiene su dirección de memoria.

En cuanto a los variables estructuradas (array, slice, struct), estas tienen elementos o miembros a los cuales se les asigna un espacio de memoria individual, por lo que dichos elementos o miembros actúan como variables.

La mayoría de variables tienen un tipo estático, es decir, el tipo que se les asigna en su declaración o de forma implícita. Pero, otras variables como de tipo `interface`, pueden tener un tipo dinámico, que es el tipo del objeto por el cual son implementadas.

Todas las variables tienen un `zero value`, el cual es asignado automáticamente si ningún valor es asignado a una variable ya declarada. 

## Types
Un tipo determina el conjunto de valores y operaciones que están disponibles para las variables de ese tipo. Los tipos en Go pueden ser especificados mediante su nombre (named types) o mediante su estructura literal (unnamed types).

Clasificación de tipos:
* **Predeclarados:** Tipos básicos como `int`, `string`, `bool`, etc. que vienen incluidos en el lenguaje
* **Definidos por el usuario:** Nuevos tipos creados mediante declaraciones de tipo
* **Type literals:** Sintaxis para describir tipos compuestos como []int, map[string]int, *Person

Declaraciones de tipo:
* **Named types:** Se crea un nuevo tipo basado en un tipo predeclarado
    ```Go
    type MyInt int
    ```
* **Type aliases:** Se crea un sinónimo para un tipo predeclarado, pero, son el mismo tipo
    ```Go
    type IntAlias = int
    ```

Algunos conceptos clave incluyen el **underlying type**, que es el tipo base de un named type. También está la **type identity**, que determina cuándo dos tipos son considerados idénticos, y los **method sets**, que representan el conjunto de métodos disponibles para un tipo.

Una diferencia importante es que los named types pueden tener sus propios métodos, mientras que los type aliases heredan los métodos del tipo original y no se les puede agregar métodos nuevos.

### Boolean types
Un tipo `bool` representa un par de valores booleanos denotados por las constantes predefinidas `true` y `false`.

### Numeric types
Existen varios tipos numéricos en Go, como `int`, `float` y `complex`, que representan un conjunto de valores enteros, decimales o complejos, respectivamente. A todos estos se les llama `numeric types`, y los tipos predeclarados independientes de la arquitectura son:
* `uint8`: conjunto de todos los valores unsigned 8-bit
* `uint16`: conjunto de todos los valores unsigned 16-bit
* `uint32`: conjunto de todos los valores unsigned 32-bit
* `uint64`: conjunto de todos los valores unsigned 64-bit
* `int8`: conjunto de todos los valores signed  8-bit
* `int16`: conjunto de todos los valores signed 16-bit
* `int32`: conjunto de todos los valores signed 32-bit
* `int64`: conjunto de todos los valores signed 64-bit
* `float32`: conjunto de todos los valores IEEE 754 32-bit
* `float64`: conjunto de todos los valores IEEE 754 64-bit
* `complex64`: conjunto de todos los valores complex con parte real e imaginaria float-32
* `complex128`: conjunto de todos los valores complex con parte real e imaginaria float-64
* `byte`: alias for uint8
* `rune`: alias for int32

También, hay un conjunto de tipos predeclarados cuyo tamaño es específico a la implementación:
* `uint`: either 32 or 64 bits
* `int`: same size as uint
* `uintptr`: an unsigned integer large enough to store the uninterpreted bits of a pointer value

Para evitar problemas de portabilidad, en Go todos los tipos numéricos son definidos, y por lo tanto, específicos a la arquitectura, a excepción de `byte` y `rune`. Además, las conversiones de tipos deben ser explícitas incluso cuando parezca que sean del mismo tamaño.

### String types
El tipo `string` es una secuencia contigua de bytes que representan a cada uno de los caracteres de dicho string, y puede estar vacío. Por lo tanto, el número de bytes usado para representar el string es su longitud y nunca es negativo. Además, los strings son inmutables, por lo que una vez creados no se pueden modificar.

En Go, los caracteres individuales de los strings pueden ser accedidos mediante índices, sin embargo, no es posible obtener la dirección de memoria de dichos elementos individuales. Una cosa importante a tener en cuenta es que la longitud de un string obtenida mediante la función `len()` no representa la cantidad de caracteres del string, sino los bytes utilizados para representar el string, esto debido al uso de Unicode y UTF-8.

### Array types
Un `array` es una secuencia contigua y enumerada de elementos de un solo tipo. Este tipo es llamado tipo del elemento, y el número de elementos establece la longitud del array y nunca es negativa.

La longitud es parte del tipo array, es decir, que una longitud debe ser asignada al momento de declarar un array, esta, debe evaluar a una constante no negativa de tipo entero.

Los arrays, a diferencia de C, siempre se trabajan por valor, por lo que siempre se utilizan completos. Incluso, los arrays son inicializados automáticamente con el zero value correspondiente al tipo del elemento si no se inicializa manualmente.

Por otro lado, los arrays `T` no pueden contener elementos del mismo tipo `T` o tipos que contengan elementos del tipo `T` directa o indirectamente, a menos que los tipos que los contienen sean type literals (compuestos).

### Slice types
Un `slice` describe un segmento contiguo de un array subyacente y provee acceso a una secuencia de elementos de dicho array, por lo que su tamaño puede ser dinámico. El tipo de elemento del array indica el tipo de elemento del array subyacente.

Los slices contienen tres características fundamentales:
* **Array pointer:** un puntero que contiene la referencia al primer elemento del segmento del array subyacente al cual tienen acceso.
* **Length:** el número de elementos del propio slice, y nunca es negativo.
* **Capacity:** el número máximo de elementos del segmento al cual el slice puede acceder del array subyacente.

Debido a que los slices trabajan con arrays subyacentes, múltiples slices pueden usar el mismo array pero diferentes segmentos de él. Por lo tanto, varios slices pueden representar el mismo array; pero varios arrays no pueden representar el mismo slice.

Un slice puede ser creado usando la función `make()`, la cual siempre crea un nuevo array anónimo utilizado por dicho slice.

### Struct types
Una `struct` es una secuencia de de elementos nombrados, a los cuales se les conoce como miembros, y cada uno de estos debe tener un nombre y un tipo. Dichos miembros pueden ser nombrados explícitamente o implícitamente, sin embargo, deben ser únicos, a menos que sean miembros vacíos o de relleno.

Un miembro que es declarado con un tipo pero sin un nombre explícitamente, es llamado `embedded field`, donde este actúa tanto como tipo como el nombre del miembro. Esto es válido siempre y cuando no se repitan los tipos (nombres).

La declaración de un miembro de cualquier struct puede ser seguida de una `tag` opcional, que actúa como atributo de todos los miembros que se encuentren en la declaración correspondiente a la etiqueta. Una etiqueta vacía es lo mismo que una etiqueta ausente. Para poder trabajar con las etiquetas directamente, se necesita el paquete `reflect`

Por otro lado, los structs `T` no pueden contener elementos del mismo tipo `T` o tipos que contengan elementos del tipo `T` directa o indirectamente, a menos que los tipos que los contienen sean type literals (compuestos).

### Pointer types
Un `pointer` o puntero permite almacenar la dirección de memoria de una variable de algún tipo específico, donde dicho tipo se convierte en el tipo base del puntero. El zero value de un puntero es `nil`.

Los punteros en Go utilizan los operadores `&` para obtener la dirección de una variable y `*` para desreferenciar (acceder al valor apuntado). A diferencia de C, Go no permite aritmética de punteros por razones de seguridad, por lo que no se puede incrementar, decrementar o realizar operaciones matemáticas sobre las direcciones de memoria.

Una característica importante es que Go detecta algunos casos de `nil` pointer dereference en runtime, generando un panic en lugar de comportamiento indefinido.

### Function types
Una `func` es un tipo usado para representar a las funciones, donde su zero value es `nil`.

Las funciones en Go son `first-class citizens`, por lo que pueden ser asignadas a variables, ser pasadas como parámetros y ser devueltas por otras funciones. Estas pueden regresar ninguno o múltiples valores en un mismo retorno, incluso, dichos valores retornados pueden ser nombrados.

Existen también funciones anónimas y funciones variadicas que reciben una cantidad desconocida de parámetros.

### Interface types
Una `interface` establece comportamiento en lugar de una estructura, por lo que actúa como un contrato que establece comportamientos que un tipo debe implementar para poder ser considerado del tipo de esa interfaz o implementar la propia interfaz.

Una variable de tipo interfaz puede almacenar un valor de cualquier tipo que implemente dicha interfaz. Dicho tipo se considerará del tipo interfaz. Por otro lado, el zero value de un tipo interfaz es `nil`.

Una interfaz es definida por una lista de elementos de la interfaz, donde un elemento puede ser un método (función) u otro elemento de un solo tipo. Su implementación se hace de forma implícita al simplemente implementar los elementos de la interfaz, por lo que no es necesario usar palabras clave ni implementación explícita.

#### Basic interfaces
Estas son las formas más básicas de una interfaz, donde simplemente se especifica un listado de métodos.

Los nombres de los métodos deben ser únicos y no estar en blanco. Diferentes tipos pueden implementar las mismas interfaces, incluso, pueden implementar múltiples interfaces al mismo tiempo.

Por otro lado, la interfaz vacía `interface{}` (y su alias `any`) puede contener cualquier tipo (no solo "tipos que no son de tipo interfaz"). Incluso otras interfaces pueden ser asignadas a `interface{}`.

#### Embedded interfaces
Estas son formas un poco más generales de una interfaz, donde una interfaz puede embeber o incluir otras interfaces dentro de su lista de elementos. Esto permite crear composición de comportamientos al combinar múltiples interfaces en una sola.

Cuando una interfaz embebe otras interfaces, un tipo debe implementar todos los métodos de todas las interfaces embebidas para satisfacer la interfaz compuesta. Si diferentes interfaces embebidas contienen métodos con el mismo nombre, estos deben tener la misma firma (mismos parámetros y valores de retorno) para evitar conflictos.

Las interfaces embebidas facilitan la composición y reutilización de comportamientos, permitiendo construir interfaces más complejas a partir de interfaces más simples y específicas.

#### Generic interfaces
Estas son las formas completamente generales de una interfaz, donde una interfaz puede contener elementos de tipo arbitrario `T` (genérico) o tipos `~T` que contengan un tipo `T` subyacente. Junto con métodos con parámetros de tipo arbitrario, permiten definición genérica y reutilización de código para múltiples tipos.

Las interfaces más genéricas pueden especificar diferentes restricciones de tipo usando:
* **Tipos exactos:** `T` acepta solo el tipo `T`
* **Underlying types:** `~T` acepta cualquier tipo con underlying type `T`
* **Union types:** `T1 | T2` acepta `T1` o `T2`
* **Combinaciones:** `~T1 | ~T2` acepta tipos `T` con esos underlying types

**Limitaciones importantes:**

* Una generic interface que contiene type constraints no puede ser usada como tipo de variable ordinaria, solo como constraint en funciones o tipos genéricos
* Los type constraints y los métodos pueden combinarse en la misma interfaz
* Las generic interfaces pueden embeberse en otras interfaces genéricas

### Map types
Un `map` es un grupo de elementos sin un orden específico de un mismo tipo, donde a este tipo se le llama el tipo del elemento. Cada elemento está indexado por un conjunto de `keys` únicas de otro tipo (pero todas las llaves deben ser del mismo tipo),  donde a este tipo de le llama el tipo de la llave. Por lo tanto, un mapa toma la forma de `key: value`.

El valor de un mapa sin inicializar es `nil`, y no se pueden agregar valores, pero, esto es diferente a un mapa vacío donde sí se pueden agregar nuevos valores. Por otro lado, el tipo de la llave debe poder ser comparable y tener bien definidos los operadores de comparación `(==, !=)`, por lo que las llaves no pueden ser de tipo función, mapa, slice, etc.

El número de elementos de un mapa es conocido como su **length** (longitud) y puede ser obtenida usando la función `len()`. También, un mapa puede ser creado con la función `make()` que recibe el mapa a crear y su capacidad inicial opcinal, aunque esta no limita al propio mapa.

### Channel types
Un `chan` proporciona un mecanismo para al ejecutar funciones de forma concurrente se puedan comunicar al enviar y recibir valores de un tipo específico. El valor de un canal sin inicializar es `nil`.

Para trabajar con los canales se debe utilizar el operador `<-`, pues este especifica la dirección del canal, ya sea para enviar o para recibir. Si se provee la dirección, se usará esa, sino, se considerará como bidireccional. Además, los canales pueden ser limitados a simplemente recibir o enviar.

Los canales se pueden crear usando la función `make()`, donde se recibe el canal a crear y opcionalmente la capacidad (buffer) del canal. Si no se especifica la capacidad o es igual a cero, se crea un canal unbuffered. Por otro lado, se puede usar la función `close()` para cerrar un canal. Un canal puede ser usado en múltiples funciones concurrentes sin necesidad de sincronización, pues la propia naturaleza de los canales permite este comportamiento.

Cabe mencionar que un `unbuffered channel` simplemente puede ser usado cuando ambas partes de la comunicación estén listas, es decir, cuando hay alguien que envía un valor alguien que recibe ese valor. Mientras que un `buffered channel` puede estar constantemente recibiendo y enviando valores, siempre y cuando no esté lleno o vacío, respectivamente. Además, funcionan usando el concepto FIFO (First-In, First-Out).

## Properties of types and values
### Representation of values
Todos los valores de los tipos predeclarados, arreglos y structs contienen su propia información almacenada directamente en ellos mismos, es decir, se utilizan por valor. Cuando se asignan o pasan como parámetros, siempre se crea una copia completa de toda su información, la cual es almacenada directamente en la variable de dicho tipo.

Por otro lado, tipos como punteros, funciones, slices, mapas y canales se utilizan por referencia y almacenan metadatos que referencian a su información subyacente. Por ejemplo, un slice no solo contiene una referencia al array subyacente, sino también metadatos como longitud y capacidad. Esto permite que múltiples variables puedan referenciar la misma información subyacente pero con diferentes metadatos.

El tipo interfaz tiene un comportamiento dinámico en cuanto a su representación: su valor puede ser autocontenido o referenciado, dependiendo del tipo dinámico de la interfaz (es decir, del tipo concreto que implementa la interfaz en tiempo de ejecución).

Una consecuencia importante de esta distinción es que los tipos que se usan por valor tienen un zero value que nunca es `nil` (por ejemplo, `0` para enteros, `""` para strings), mientras que el zero value de los tipos que se usan por referencia siempre es `nil`.

### Underlying types
Todos los tipos en Go tienen un tipo subyacente (underlying type) que define su estructura fundamental y representación en memoria.

Para tipos predeclarados (como `int`, `string`, `bool`), el tipo subyacente es el tipo mismo, es decir, el underlying type de `int` es `int`.

Para tipos compuestos construidos con type literals (como `[]int`, `map[string]int`, `chan bool`), el tipo subyacente es la estructura del type literal mismo.

Es importante mencionar que incluso si dos tipos diferentes usan el mismo tipo subyacente, siguen siendo dos tipos completamente diferentes, a menos que sean declarados como aliases.

En resumen, el underlying type aplica a todos los tipos individualmente y determina la forma en que se va a representar dicho tipo en la memoria y las operaciones básicas permitidas.

### Core types
Todos los tipos que no son interfaces tienen un core type, el cual es idéntico a su underlying type.

Las interfaces tienen un comportamiento especial para determinar su core type. Una interfaz tiene un core type únicamente si se cumple una de estas condiciones:
1. Todos los tipos en el type set de la interfaz comparten el mismo underlying type T, y al menos uno de esos tipos no es una interfaz. En este caso, el core type de la interfaz es T.
2. Todos los tipos en el type set son tipos de canal con el mismo element type y la misma dirección (send, receive, o bidireccional).

Si una interfaz no cumple ninguna de estas condiciones, no tiene core type, y el core type por definición nunca puede ser un tipo definido, creado o de tipo interfaz.

En resumen, el core type aplica o está diseñado principalmente para las interfaces usadas de forma genérica (type constraints) y garantiza que todos los tipos en el type set de dichas interfaces permitan realizar las mismos operaciones.

### Type identity
Esta propiedad se refiere a cuándo dos tipos son considerados exactamente el mismo tipo por el compilador de Go, determinando así si comparten idénticas características y comportamientos.

Un tipo definido por el programador (named type como `type MyInt int`) siempre es diferente a cualquier otro tipo, incluso si tienen el mismo underlying type.

Para tipos no-definidos (type literals), dos tipos son idénticos si sus estructuras son completamente equivalentes en todos los aspectos:
* **Array types:** Idénticos si el tipo de elementos es idéntico y tienen la misma longitud
* **Slice types:** Idénticos si el tipo de elementos es idéntico
* **Struct types:** Idénticos si tienen la misma secuencia de campos con nombres idénticos, tipos idénticos, tags idénticas, y el mismo estado de embedding
* **Pointer types:** Idénticos si apuntan a tipos idénticos
* **Function types:** Idénticos si tienen igual número de parámetros y resultados con tipos idénticos, y ambas son variádicas o ninguna lo es
* **Interface types:** Idénticos si definen exactamente el mismo type set
* **Map types:** Idénticos si los tipos de llaves y valores son idénticos
* **Channel types:** Idénticos si el tipo de elementos es idéntico y tienen la misma direccionalidad

Que dos tipos sean idénticos quiere decir que esos tipos son iguales en todos sus aspectos individuales relevantes para el compilador. Entonces, type identity implica que dichos tipos puedan ser tratados de la misma forma y no necesiten conversiones explícitas.

### Assignability
Esta propiedad determina cuándo el valor x de tipo V puede ser asignado a una variable de tipo T, incluso cuando los tipos no son idénticos. Esta flexibilidad permite mayor interoperabilidad en el sistema de tipos de Go.

Un valor es asignable cuando se cumple alguna de estas condiciones:
* V y T son tipos idénticos (caso más directo)
* V y T tienen underlying types idénticos, donde ninguno es un type parameter y al menos uno no es un named type
* V y T son channel types con element types idénticos, donde V es bidireccional y al menos uno no es un named type
* T es un interface type (no type parameter) y el valor x implementa T (está en el type set de T)
* x es el valor `nil` y T es un tipo que usa referencias (pointer, function, slice, map, channel, interface) y no es un type parameter
* x es una constante sin tipo que puede ser representada por el tipo T

### Representability
Esta propiedad se aplica específicamente a constantes y determina si una constante x puede ser correctamente almacenada y representada como valor de un tipo T (donde T no es un type parameter).

Una constante x es representable por tipo T cuando se cumple alguna de estas condiciones:
* x está dentro del rango de valores que T puede representar
* T es un floating-point type y x puede ser redondeado a la precisión de T sin overflow
* T es un complex type y tanto la parte real como imaginaria de x son individualmente representables por el tipo de componente subyacente de T

Básicamente esta propiedad aplica para las constantes sin tipo que determina si dicha constante puede ser asignada a una variable de cierto tipo sin causar errores de compilación o pérdida de información crítica.

### Method sets
El conjunto de métodos (method set) de un tipo determina qué métodos pueden ser llamados sobre valores de ese tipo y, crucialmente, qué interfaces puede implementar el tipo.

**Reglas fundamentales:**
* **Para un tipo definido T:** El method set incluye todos los métodos declarados con receiver de valor (t T)
* **Para un pointer type \*T:** El method set incluye todos los métodos declarados con receiver de valor (t T) y receiver de pointer (t \*T)
* **Para un interface type:** El method set de un tipo interfaz es la intersección del conjunto de métodos de cada elementos de la interfaz, es decir, solo los métodos que todos los elementos tienen
* **Para tipos embedded en structs:** Los métodos del tipo embedded se promueven al method set del struct contenedor

Un tipo T implementa una interfaz I solo si el method set de T contiene todos los métodos requeridos por I. Además, todos los tipos tienen method set (puede estar vacío), y todos los métodos deben tener un nombre único y no estar en blanco.

## Blocks
Un bloque es una secuencia (posiblemente vacía) de declaraciones y statements encerradas dentro de llaves `{}`.

Los bloques explícitos son aquellos delimitados por llaves visibles en el código, pero también existen bloques implícitos que no requieren llaves:
* **Universe block:** Engloba todo el código fuente de Go y contiene todas las declaraciones predefinidas (`int`, `string`, `nil`, `true`, etc.)
* **Package block:** Engloba todo el código fuente dentro de un paquete específico
* **File block:** Engloba todo el código fuente dentro de un archivo individual
* **Function block:** Cada función o método crea su propio bloque
* **Statement blocks:** Estructuras de control (`if`, `for`, `switch`, etc.) y bloques explícitos `{}` crean sus propios bloques

Los bloques forman una jerarquía anidada donde cada bloque interno puede acceder a declaraciones de bloques externos, pero no viceversa. Las declaraciones en bloques internos pueden hacer shadowing (ocultar) declaraciones de bloques externos con el mismo nombre. Esta estructura determina la visibilidad y duración de variables, constantes, tipos y funciones en Go.

## Declarations and scope
Una declaración relaciona un identificador (que no sea blank) a algún objeto de Go, como `const`, `type`, `type parameter`, `variable`, `function`, `label` o `package`. Cualquier identificador usado debe ser declarado y no estar en blanco, además, no se puede declarar el mismo identificador en el mismo bloque. Los identificadores a nivel paquete no pueden redeclararse.

El identificador blank `_` puede usarse como cualquier otro identificador en una declaración, pero no establece ninguna relación (no es realmente "declarado"), permitiendo descartar valores.

**Reglas de scope (Go usa lexical scoping con bloques):**
1. **Identificadores predeclarados:** Scope del universe block
2. **Constantes, tipos, variables, o funciones (pero no métodos)**: Scope del package block
3. **Nombres de paquetes importados:** Scope del file block que contiene la declaración de import
4. **Method receiver, parámetros de función, o variables de resultado:** Scope del function body
5. **Type parameter de funciones o method receivers:** Inicia después del nombre de la función y termina al final del function body
6. **Type parameter de tipos:** Inicia después del nombre del tipo y termina al final del TypeSpec
7. **Constantes o variables declarados dentro funciones:** Inicia al final del ConstSpec o VarSpec (ShortVarDecl para declaraciones cortas) y termina al final del bloque contenedor más interno
8. **Tipos declarados dentro de funciones:** Inicia en el identificador del TypeSpec y termina al final del bloque contenedor más interno

Un identificador declarado en un bloque interno puede ocultar uno del bloque exterior con el mismo nombre (shadowing). Por otro lado, la palabra clave `package` no es una declaración, sino una directiva que identifica archivos del mismo paquete y especifica el nombre por defecto al importar.

### Label scopes
Las labels son declaradas mediante el uso de `:` y son usadas generalmente en las declaraciones de `break`, `continue` y `goto`. Una label declarada debe ser usada, pues no está permitido declararla sin usarla.

A diferencia de otros identificadores, las labels tienen un scope único limitado a la función en la que se declaran y no interfieren con otros identificadores que tengan el mismo nombre (pueden coexistir sin conflicto).

### Blank identifier
El identificador en blanco es representado por el caracter `_`. Sirve como un placeholder anónimo en lugar de un identificador normal, y puede usarse múltiples veces en el mismo scope sin causar conflictos de redeclaración.

### Predeclared identifiers
Hay un total de 44 identificadores predeclarados en el bloque universal de Go. Estos son mutables, a diferencia de las keywords, pero no es recomendable hacerlo.

**Types:**
* any
* bool
* byte
* comparable
* complex64
* complex128
* error
* float32
* float64
* int
* int8
* int16
* int32
* int64
* rune
* string
* uint
* uint8
* uint16
* uint32
* uint64
* uintptr

**Constants:**
* true
* false
* iota

**Zero value:**
* nil

**Functions:**
* append
* cap
* clear
* close
* complex
* copy
* delete
* imag
* len
* make
* max
* min
* new
* panic
* print
* println
* real
* recover

### Exported identifiers
Un identificador puede ser exportado para permitir acceso desde otros paquetes. Un identificador es exportado únicamente cuando cumple ambas condiciones:
* El primer caracter es una letra mayúscula Unicode
* El identificador es declarado en el package block o corresponde a un field de struct o es un método

Los identificadores que no cumplan estas condiciones no son exportados y solo pueden usarse dentro del mismo paquete donde se declaran.

### Uniqueness of identifiers
En un conjunto de identificadores, un identificador es único cuando es diferente a cualquier otro identificador de ese conjunto.

**Dos identificadores son diferentes si:**
* Se escriben diferente (case-sensitive)
* Están en paquetes diferentes
* Están en el mismo paquete pero no son exportados (limitados a su scope local)

**Dos identificadores son iguales si:**
* Tienen la misma escritura y están en el mismo contexto accesible (mismo paquete o ambos exportados)

### Constant declarations
Una declaración de constantes relaciona una lista de identificadores con una lista de valores dados por expresiones constantes. La cantidad de identificadores debe coincidir con la cantidad de expresiones.

* **Con tipo explícito:** Si se especifica un tipo, todas las constantes adoptan ese tipo y sus expresiones deben ser asignables a él.

* **Sin tipo explícito:** Cada constante adopta el tipo individual correspondiente a su expresión respectiva.
* **Declaraciones agrupadas:** En declaraciones múltiples dentro de paréntesis, es posible omitir valores en ConstSpecs posteriores, pues automáticamente reutilizan el valor y tipo del ConstSpec anterior.

**ConstSpec:** Es un término técnico que simplemente se refiere una línea individual donde se declara una constante.

### Iota
El identificador predeclarado `iota` representa una secuencia numérica entera sin tipo para generar constantes automáticamente.

**Funcionamiento:**
* Su valor corresponde al índice del ConstSpec en la declaración, iniciando en 0
* Se incrementa automáticamente en cada nuevo ConstSpec
* Se resetea a 0 en cada nueva declaración de constantes (`const`)

**Comportamiento especial:** Múltiples usos de `iota` en el mismo ConstSpec representan el mismo valor (no se incrementa dentro de la misma línea).

### Type declarations
Una declaración de tipo relaciona un identificador (nombre del tipo) con un tipo nuevo o existente. Las declaraciones pueden ser de dos formas: alias declarations y type definitions.

#### Alias declarations
La declaración de alias (`type NewName = ExistingType`) relaciona un identificador con un tipo existente. No crea un nuevo tipo, simplemente establece una forma alternativa de referenciar el mismo tipo, siendo tratados como idénticos.

Los alias funcionan de acuerdo a su scope. Pueden usar type parameters creando *generic aliases*, los cuales deben ser instanciados al usarse.

#### Type definitions
La definición de tipo (`type NewName ExistingType`) relaciona un identificador con un tipo completamente nuevo. Crea un *defined type* que usa un tipo existente como underlying type.

**Características de defined types:**
* Son diferentes a cualquier otro tipo, incluyendo su tipo base
* Pueden tener métodos propios asociados
* No heredan métodos del tipo base
* Pueden ser *generic types* usando type parameters

Para los generic types, los métodos asociados deben declarar la misma cantidad de type parameters que el tipo definido.

### Type parameter declarations
Una lista de type parameters declara los parámetros de tipo genérico para funciones genéricas o declaraciones de tipos. Su estructura es similar a parámetros regulares, pero están encerrados en corchetes `[]` e inmediatamente después del nombre de la función/tipo y antes de sus parámetros regulares.

Todos los nombres de los type parameters en la lista que no estén en blanco deben ser únicos. Estos actúan como named types temporales que sirven como placeholders, los cuales son reemplazados por tipos concretos duranta la instanciación.

Cada type parameter tiene un type constraint correspondiente que define sus restricciones.

#### Type constraints
Un type constraint es una interfaz que define el conjunto de tipos permitidos (type set) para un type parameter y controla las operaciones disponibles sobre valores de esos tipos.

#### Satisfying a type constraint
Un tipo satisface un type constraint cuando:
* Implementa la interfaz del constraint (tiene los métodos requeridos), o
* Está incluido en el type set del constraint (listado explícitamente con `|`)

### Variable declarations
La declaración de variable con `var` crea una o más variables, relacionando identificadores con tipos específicos y valores iniciales opcionales. Cuando hay una lista de expresiones, las variables se inicializan con sus valores correspondientes, pero si no se provee inicializador, se asigna el zero value del tipo. El tipo puede especificarse explícitamente o ser omitido, en cuyo caso la variable adopta el tipo inferido de la expresión de inicialización.

### Short variable declarations
La declaración corta usa el operador `:=` sin necesidad de `var` ni tipo explícito. Solo está permitida dentro de funciones (no a nivel de package) y siempre requiere un inicializador. En asignaciones múltiples, puede redeclarar variables existentes siempre que al menos una variable sea nueva en ese scope.

### Function declarations
La declaración de función con `func` relaciona un identificador con una función. Una función tiene signature específica (parámetros y valores de retorno) y debe seguir las reglas del lenguaje, como terminar con una expresión de retorno si declara valores de retorno, o manejar todos los casos de ejecución apropiadamente.

### Method declarations
Un método es una función con un receiver, lo que significa que está asociada específicamente al tipo declarado en su receiver. La declaración de método relaciona un identificador con un método y lo asocia al tipo del receiver correspondiente.

El receiver se declara después de `func` pero antes del nombre del método, debe ser un parámetro único no variádico cuyo tipo sea un defined type o puntero a defined type (llamado receiver base type). El identificador del receiver no puede estar en blanco y debe ser único dentro de toda la signature del método. Los métodos con receiver de valor vs receiver de puntero tienen diferentes method sets, lo que afecta qué interfaces pueden implementar.

## Expressions
Una expression especifica la computación a realizar aplicando operadores y funciones a valores (operandos).

### Operands
Los operands representan los valores elementales usados en expresiones. Pueden ser identificadores que representen constantes, variables o funciones, o expresiones anidadas entre paréntesis. 

El blank identifier `_` solo puede usarse en el lado izquierdo de asignaciones para descartar valores.

### Qualified identifiers
Un qualified identifier accede a identificadores exportados de otros paquetes usando la sintaxis `package.Identifier`. Tanto el nombre del paquete como el identificador deben ser no-blank, y el identificador debe estar exportado (empezar con mayúscula) para ser accesible desde el paquete importador.

### Composite literals
Los composite literals construyen valores para tipos compuestos (structs, arrays, slices, maps) usando la sintaxis `Type{elements}`. Cada evaluación crea una nueva instancia.

La estructura de los elementos varía según el tipo: los structs pueden usar nombres de campos, los arrays pueden especificar índices, etc. Para type parameters, todos los tipos en su type set deben ser válidos para composite literals para poder usarlos.

### Function literals
Una function literal representa una función anónima que no puede declarar type parameters. Puede ser asignada a variables o invocada directamente. Las function literals son closures, lo que significa que pueden capturar y usar variables del scope externo donde se definen. Estas variables capturadas son compartidas entre la función externa y el closure, manteniéndose vivas mientras el closure exista.

### Primary expressions
Las primary expressions son los operandos fundamentales para construir expresiones más complejas (unarias y binarias). Incluyen identificadores, literals, expresiones entre paréntesis, y expresiones que acceden a elementos o campos.

### Selectors
Un selector (`expression.selector`) accede a un campo o método de un valor. La expresión debe ser primary y no puede ser un nombre de paquete. Los selectors funcionan automáticamente tanto con valores directos como con punteros, pues Go maneja implícitamente la dereferenciación cuando es necesario.

### Method expressions
Una method expression es una forma de referenciar un método de un tipo como si fuera una función regular, donde el receiver se convierte en el primer parámetro explícito de la función.

Cuando se define un método con la sintaxis `func (t T) Method(args) ReturnType { ... }`, se puede crear una method expression usando `T.Method`. Esta expresión resulta en una función que tiene la firma `func(t T, args) ReturnType`, donde el receiver original ahora es el primer parámetro explícito.

Las method expressions se diferencian de los method values. Una method expression como `T.Method` requiere que se pase la instancia como primer argumento cuando es llamada. Por el otro lado, un method value como `instance.Method` ya tiene la instancia "capturada" y no necesita que se le pase explícitamente.

En cuanto a las reglas para crear method expressions, solo se pueden crear desde el tipo que está en el method set correspondiente. Los métodos con value receivers `(t T)` permiten la sintaxis `T.Method`. Sin embargo, los métodos con pointer receivers `(t *T)` únicamente permiten `(*T).Method`, ya que estos métodos no forman parte del method set del tipo valor.

### Method values
Un method value es una forma de referenciar un método desde una instancia específica, donde el receiver queda "capturado" dentro de la función resultante.

Cuando se tiene una instancia y se referencía uno de sus métodos usando `instance.Method` (sin paréntesis), se obtiene una función que ya no requiere el receiver como parámetro. Si el método original tiene la firma `func(t T) Method(args) ReturnType`, el method value resultante tiene la firma `func(args) ReturnType`.

A diferencia de las method expressions, los method values pueden crearse desde cualquier instancia, independientemente del tipo de receiver del método. Go realiza automáticamente las conversiones necesarias entre `T` y `*T` cuando se crea el method value, permitiendo que tanto valores como punteros puedan generar method values para métodos con cualquier tipo de receiver.

### Index expressions
Una primary expression en la forma `a[x]` representa un elemento de un array, slice, string o map donde `x` es su índice o posición (para arrays, slices y strings), o clave (para maps).

**Para arrays, slices y strings:**
El índice debe ser una constante sin tipo o un valor de tipo `int`. Debe ser un número no negativo representable por el tipo `int`, y debe estar en el rango `0 <= index < len(object)`. Una constante sin tipo usada como índice toma automáticamente el tipo `int`. Si el índice está fuera del rango permitido, se produce un panic en tiempo de ejecución.
* **Arrays:** `a[x]` representa el elemento en el índice `x` del array `a`. El comportamiento es el mismo tanto si se usa un array directamente como un puntero a un array
* **Slices:** `a[x]` representa el elemento en el índice `x` del slice `a`
* **Strings:** `a[x]` representa el byte en el índice `x` del string a. Para strings, el índice debe ser una constante, y el resultado no puede ser asignado (los strings son inmutables).

**Para maps:**
El índice `x` debe ser asignable al tipo de las claves del map. Si el map contiene `x` como clave, se devuelve el valor correspondiente. Si el map es `nil` o no contiene la clave, se devuelve el zero value del tipo de valores del map.

**Para type parameters:**
La index expression `a[x]` debe ser válida para todos los tipos del type set del type parameter. Esto significa que todos los tipos en el type set deben soportar la operación de indexado con el tipo de `x`.

Para cualquier otro tipo que no sea array, slice, string, map o type parameter, el uso de index expressions no está permitido.

### Slice expressions
Las slice expressions construyen un substring o un slice (porción) a partir de un string, array, puntero a array o slice. Existen dos variantes: una forma simple que especifica un rango (inicio y fin), y una forma completa que también especifica la capacidad del slice resultante.

Para type parameters, a menos que todo su type set contenga únicamente strings, el tipo de elemento de todos los tipos en el type set debe ser el mismo.

#### Simple slice expressions
Las simple slice expressions utilizan la sintaxis `a[low:high]` para especificar un rango de elementos. El índice inferior (`low`) es inclusivo, mientras que el índice superior (`high`) es exclusivo.

Cualquiera de los dos índices puede ser omitido por conveniencia. Si se omite el índice inferior, toma el valor `0`. Si se omite el índice superior, toma el valor de la longitud del objeto (`len(object)`).

Los índices deben ser enteros no negativos. Para strings y arrays, deben cumplir `0 <= low <= high <= len(object)`. Para slices, el índice superior puede extenderse hasta la capacidad del slice: `0 <= low <= high <= cap(slice)`.

Si la slice expression es válida pero el objeto del cual se obtiene el slice es inválido, el resultado es `nil`. Todos los slices generados a partir del mismo objeto subyacente comparten el mismo array base.

#### Full slice expressions
Las full slice expressions utilizan la sintaxis `a[low:high:max]` y están disponibles para arrays, punteros a arrays y slices, pero no para strings. El tercer índice (`max`) controla la capacidad del slice resultante.

Los índices `low` y `high` funcionan igual que en las simple slice expressions. El índice max establece que la capacidad del slice resultante será `max - low`. Solo el primer índice (`low`) puede ser omitido, tomando el valor `0`.

Los índices deben cumplir `0 <= low <= high <= max <= cap(object)` y ser enteros no negativos.

### Type assertions
Una type assertion es una expresión de la forma `x.(T)` que verifica que `x` no es `nil` y que el valor almacenado en `x` es del tipo `T`. Esta operación solo es válida cuando `x` es de tipo interface.

**Comportamiento según el tipo T:**
* Si `T` no es un tipo interface, la type assertion verifica que el tipo dinámico de `x` sea idéntico a `T`.
* Si `T` es un tipo interface, la type assertion verifica que el tipo dinámico de `x` implemente la interface `T`.

Si la type assertion es válida, la expresión devuelve el valor almacenado en `x` convertido al tipo `T`. Si la assertion es inválida, se produce un panic en tiempo de ejecución. Para evitar el panic, se puede usar la forma de dos valores: `value, ok := x.(T)`. En esta variante, si la assertion es válida, `value` contiene el valor convertido y `ok` es `true`. Si la assertion es inválida, `value` es el zero value del tipo `T` y `ok` es `false`.

### Calls
