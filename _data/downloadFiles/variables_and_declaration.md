# Variáveis

## Declaração

Em Go, a declaração de variáveis começa com a palavra-chave `var`, seguida do nome da variável e, opcionalmente, seu tipo.

`var nomeDaVariavel tipoDaVariavel = valorDaVariavel`

```go
var myName string = "odeassis"
```

- A declaração `var myName string` cria uma variável chamada `myName` do tipo `string`.

```go
var age int
age := 24
```

- A declaração `age := 24` cria uma variável `age` e o tipo é inferido automaticamente a partir do valor.

A declaração `var` também pode ser usada com uma lista de variáveis, onde o tipo vem por último.

```go
var (
    name string
    age int
    location string
)
```

## Atribuição de Valores

Uma declaração `var` pode incluir valores iniciais, um ou mais, que são atribuídos às variáveis no momento da declaração. A atribuição de valores a variáveis é feita usando o operador `=` ou `:=` (para declaração e atribuição de variáveis ao mesmo tempo).

```go
var (
    name = "odeassis"
    age = 24
    location = "Brasil"
)

// ou

var  name, age, location = "odeassis", 24, "Brasil"
```

Variáveis declaradas sem um valor inicial têm o valor zero **_(zero-valued)_**. Por exemplo, o valor zero para um `int` é `0`.

### Declaração Curta

Dentro de uma função, a declaração curta `:=` **_("walrus operator"/"operador morsa")_** pode ser usada no lugar de uma declaração `var` com tipos implícitos. Esse tipo de declaração também é conhecido como declaração de atribuição. No entanto, a declaração curta só pode ser usada dentro de uma função, não fora dela. Fora de uma função, cada instrução começa com uma palavra-chave (como `var`, `func` e `return`) e a declaração curta não é uma palavra-chave.

```go
func main() {
    name := "odeassis"
    age := 24
    location := "Brasil"
}
```

## Constantes

Go permite a declarar constantes usando a palavra-chave `const`. As constantes podem ser caracteres, `strings`, `booleanos` ou valores numéricos.

```go
const Pi = 3.14
const (
    StatusOK = 200
    StatusCreated = 201
    StatusAccepted = 202
)
```

## Escopo de Variáveis

As variáveis em Go têm escopo léxico, o que significa que elas são visíveis apenas dentro do bloco onde são declaradas. Variáveis globais são declaradas fora de qualquer função e têm escopo global. Variáveis locais são declaradas dentro de uma função e têm escopo local.

```go
var global = "global"

func sum (a, b int) int {
    var local = "local"
    return a + b
}

func main() {
    var local = "local"
    fmt.Println(global) // global
    fmt.Println(local) // local
}
```
