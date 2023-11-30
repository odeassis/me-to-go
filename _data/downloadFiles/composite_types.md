# Tipos compostos

`Arrays` e `Structs` são exemplos de tipos compostos; seus valores são concatenados de outros valores em memória. Arrays são homogêneos, ou seja, todos os seus elementos são do mesmo tipo. Structs são heterogêneos, ou seja, seus elementos podem ser de tipos diferentes. Slices e Maps são estruturas de dados dinâmicas isso significa que seus tamanhos podem ser alterados durante a execução do programa.

## Arrays

Um array é uma sequência de tamanho fixo de zero ou mais elementos de um tipo específico. Arrays são declarados usando a sintaxe `var nome [tamanho]tipo`. O tamanho do array é parte do seu tipo, ou seja, `[5]int` e `[10]int` são tipos diferentes. O tamanho do array deve ser um valor inteiro constante, ou uma expressão constante.

Elementos individuais de arrays são acessados com a notação convencional de índice, que varia de zero até o tamanho do array menos um. Por exemplo, o índice do primeiro elemento de um array de tamanho 5 é 0, e o índice do último elemento é 4.

```go
var x [5]int
x[0] = 10
x[1] = 20
x[2] = 30
x[3] = 40
x[4] = 50

for idx, value := range x {
    fmt.Printf("x[%d] = %d\n", idx, value)
}

for _, value =: range x {
  fmt.Printf("v = %d\n", value)
}
```

Por padrão, os elementos de um array são inicializados com o valor zero do tipo do elemento. Para arrays de inteiros, isso significa que todos os elementos são inicializados com 0. Para arrays de strings, isso significa que todos os elementos são inicializados com uma string vazia, `""`. Para arrays de booleanos, isso significa que todos os elementos são inicializados com `false`.

```go
var x [5]int
for idx, value := range x {
    fmt.Printf("x[%d] = %d\n", idx, value)
}
```

Podemos usar um array literal para inicializar um array com um conjunto de valores. Um array literal é uma lista de valores separados por vírgulas, envolvidos por chaves. Por exemplo, o array literal `{1, 2, 3, 4, 5}` inicializa um array de cinco elementos com os valores 1, 2, 3, 4 e 5.

```go
var x [5]int = [5]int{1, 2, 3, 4, 5}
```

Se usarmos reticências `...` no lugar do tamanho do array, o compilador irá inferir o tamanho do array a partir do número de elementos no array literal.

```go
var x [5]int = [...]int{1, 2, 3, 4, 5}
```

O tamanho de um array faz parte de seu tipo, então `[5]int` e `[10]int` são tipos diferentes. Por isso, arrays não podem ser redimensionados. Arrays são valores, então quando você atribui um array a uma variável ou passa um array para uma função, a cópia do array é passada, não um ponteiro para o array. O tamanho do array deve ser um valor inteiro constante, ou uma expressão constante, ou seja o tamanho do array deve ser conhecido em tempo de compilação.

```go
var x [5]int = [5]int{1, 2, 3, 4, 5}

x = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // Erro de compilação
```

Também podemos especificar uma lista de pares com índice e valor, onde o índice é usado para inicializar o elemento no índice especificado. Por exemplo, o array literal `{2: 10, 4: 20}` inicializa um array de cinco elementos com os valores 0, 0, 10, 0 e 20.

```go
var x [5]int = [5]int{2: 10, 4: 20}

type Coin int
const (
  /*
  * iota é um contador automático usado em constantes em Go.
  * iota inicia com 0 e é incrementado automaticamente
  * para cada linha subsequente de constante não inicializada.
  */
  Bitcoin Coin = iota
  Ethereum
  Cardano
)

symbol := [...]string{Bitcoin: "BTC", Ethereum: "ETH", Cardano: "ADA"}
```

Nesse formato, os índices podem estar em qualquer ordem, e os índices não especificados são inicializados com o valor zero do tipo do elemento. `r := [...]int{99:-1}` define um array r com 100 elementos todos iguais a zero, exceto o último que é igual a -1.

Se o tipo do elemento de um array for comparável, então o tipo do array também será comparável. Dois arrays são iguais se eles tiverem o mesmo tamanho e cada um dos elementos for igual. Por exemplo, `[3]int{1, 2, 3}` é igual a `[3]int{1, 2, 3}`, mas não é igual a `[4]int{1, 2, 3, 4}`.

```go
var x [5]int = [5]int{1, 2, 3, 4, 5}
var y [5]int = [5]int{1, 2, 3, 4, 5}
var z [5]int = [5]int{1, 2, 3, 4, 6}

fmt.Println(x == y) // true
fmt.Println(x == z) // false
```

A função Sum256 do pacote crypto/sha256 gera um hash de criptografia SHA256 ou digest de uma mensagem armazenada em um slice qualquer de bytes. O digest tem 256 bits, ou 32 bytes. O digest é frequentemente usado para verificar a integridade de dados transmitidos e armazenados. Por exemplo, o digest de um arquivo pode ser usado para verificar que o arquivo não foi alterado.

```go
import (
  "crypto/sha256"
  "fmt"
)

func main() {
  c1 := sha256.Sum256([]byte("x"))
  c2 := sha256.Sum256([]byte("X"))

  /*
  * %x - base 16, com letras minúsculas para a-f
  * %t - booleano
  * %T - tipo
  */
  fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)

}
```

Quando uma função é chamada, uma cópia de cada argumento é passada para os parâmetros da função. Para arrays, isso significa que uma cópia do array é passada para a função. Se a função modificar o array, as mudanças não serão visíveis fora da função, porque a função está trabalhando com uma cópia do array. Passar arrays grandes dessa maneira pode ser ineficiente e, em alguns casos, impraticável.

Mas claro que podemos passar um ponteiro para um array, e a função pode modificar o array original. Por exemplo, a função `zero` recebe um ponteiro para um array e atribui zero a cada elemento do array.

```go

func zero(ptr *[32]byte) {
  *ptr = [32]byte{}
}

func one(ptr *[32]byte) {
  for i := range ptr {
    ptr[i] = 1
  }
}
```

Usar um ponteiro pra um array é eficiente e permite a que função chamada altere a variável de quem chamou, mas o arrays continuam sendo um tipo fixo, ou seja, não podemos alterar o tamanho de um array. Arrays raramente são usados como parâmetros de funções em Go. Em vez disso, slices são usados.

## Slices

Slice representam uma sequência de elementos de um tipo específico. Um slice é declarado usando a sintaxe `var nome []tipo`. Um slice é uma estrutura de dados dinâmica que pode crescer ou encolher durante a execução do programa. A capacidade de um slice é o tamanho do array subjacente, que é o número de elementos no array a partir do primeiro elemento no slice. A função `cap` retorna a capacidade de um slice.

```go
var x []int
fmt.Println(len(x), cap(x)) // 0 0
```

Um slice tem um ponteiro para um array subjacente, um tamanho e uma capacidade. O tamanho do slice é o número de elementos no slice, e a capacidade é o número de elementos no array subjacente, contando a partir do primeiro elemento no slice. O tamanho e a capacidade de um slice s podem ser obtidos usando as expressões `len(s)` e `cap(s)`. O ponteiro aponta para o primeiro elemento do array que é acessível através do slice, que não é necessariamente o primeiro elemento do array. O primeiro elemento do slice pode ser acessado usando a expressão `s[0]`, e o último elemento do slice pode ser acessado usando a expressão `s[len(s)-1]`.

```go
var x []int = []int{1, 2, 3, 4, 5}
fmt.Println(len(x), cap(x)) // 5 5
```

Vários slices podem compartilhar o mesmo array subjacente. Por exemplo, o slice `x` compartilha o mesmo array subjacente com o slice `y`.

```go
var x []int = []int{1, 2, 3, 4, 5}
var y []int = x[1:3]
fmt.Println(len(y), cap(y)) // 2 4
```

O operador de slice `s[i:j]` cria um slice de um slice. O novo slice aponta para o mesmo array subjacente que o slice original, mas o novo slice acessa apenas os elementos do array no intervalo `i` até `j-1`. O novo slice tem tamanho `j-i` e capacidade `cap(s)-i`. Se `i` for omitido, o slice começará no primeiro elemento do array subjacente. Se `j` for omitido, o slice terminará no último elemento do array subjacente. Se ambos `i` e `j` forem omitidos, o slice será igual ao slice original.

Slice além de cap(s) causa um erro de tempo de execução, mas além de len(s) estende o slice, aumentando seu tamanho.

Como uma fatia contém um ponteiro para um array, passar uma fatia para uma função significa que a função pode modificar os elementos do array subjacente que o slice aponta. Por exemplo, a função `zero` recebe um slice e atribui zero a cada elemento do array subjacente.

```go
func zero(ptr []int) {
  for i := range ptr {
    ptr[i] = 0
  }
}
```

Um slice literal é como um array literal, mas sem o tamanho. Um slice literal é uma lista de valores separados por vírgulas, envolvidos por chaves. Isso cria implicitamente uma variável de array subjacente do tamanho correto e um slice que aponta para o array. Por exemplo, o slice literal `{1, 2, 3, 4, 5}` cria um array de cinco elementos e um slice que aponta para os cinco elementos. O slice tem tamanho 5 e capacidade 5.

Diferente dos arrays, slices não são comparáveis, então não podemos usar o operador `==` para testar se dois slices são iguais. A única exceção é o slice nulo, que é igual a `nil`. O valor zero de um tipo slice é `nil`. Para testar se dois slices são iguais, devemos escrever uma função que itera sobre os elementos de cada slice e testa se eles são iguais.

```go
func equal(x, y []int) bool {
  if len(x) != len(y) {
    return false
  }

  for i := range x {
    if x[i] != y[i] {
      return false
    }
  }

  return true
}


var x []int = []int{1, 2, 3, 4, 5}
var y []int = []int{1, 2, 3, 4, 5}
var z []int = []int{1, 2, 3, 4, 6}

fmt.Println(equal(x, y)) // true
fmt.Println(equal(x, z)) // false
```

Um slice `nil` não tem um array subjacente. Um slice `nil` tem um tamanho e capacidade zero, e não aponta para um array. Um slice `nil` é igual a `nil`, mas não é igual a um slice vazio. Um slice vazio tem um array subjacente, mas o tamanho e a capacidade do slice são zero. Um slice vazio é igual a outro slice vazio, mas não é igual a `nil`.

```go
var x []int
var y []int = []int{}
var z []int = nil

fmt.Println(x == nil) // true
fmt.Println(y == nil) // false
fmt.Println(z == nil) // true
```

Para verificar se um slice está vazio, use a expressão `len(s) == 0`. Para verificar se um slice é `nil`, use a expressão `s == nil`.

### Funções built-in

Funções built-in são funções pré-definidas que podem ser usadas em qualquer lugar em um programa. Estas funções são parte do pacote builtin e são incorporadas diretamente no compilador `Go`. Algumas funções built-in são funções de conversão de tipo, como `int` e `float64`. Outras funções built-in são funções de manipulação de slice, como `append` e `copy`.

#### make

A função built-in `make` cria um slice de um tipo específico, tamanho e capacidade. A função `make` retorna um slice inicializado que aponta para um array recém-alocado. A função `make` recebe três argumentos: o tipo do slice, o tamanho do slice e a capacidade do slice. O tipo do slice deve ser um tipo de slice válido. O tamanho e a capacidade do slice devem ser um valor inteiro não negativo. A capacidade do slice deve ser maior ou igual ao tamanho do slice. Os elementos do slice são inicializados com o valor zero do tipo do elemento. Internamente `make` cria uma variável do tipo array sem nome e retorna um slice que aponta para o array, o array é acessível apenas através do slice.

```go
var x []int = make([]int, 5, 10)
fmt.Println(len(x), cap(x)) // 5 10
```

#### append

A função `append` adiciona um ou mais elementos ao final de um slice e retorna o slice resultante. Se o slice tiver capacidade suficiente, os elementos serão adicionados ao slice existente. Se o slice não tiver capacidade suficiente, um novo array será alocado e os elementos serão copiados para o novo array antes de adicionar os novos elementos. A função `append` retorna o slice resultante, que pode ser diferente do slice original. Se o slice original tiver capacidade suficiente, o slice resultante apontará para o mesmo array subjacente que o slice original. Se o slice original não tiver capacidade suficiente, o slice resultante apontará para um novo array alocado.

A função `append` é crucial para entender o funcionamento dos slices em Go.

```go
func appendCap(x []int, y int) []int {
  var z []int
  zlen := len(x) + 1
  if zlen <= cap(x) {
    z = x[:zlen]
  } else {
    zcap := zlen
    if zcap < 2*len(x) {
      zcap = 2*len(x)
    }
    z = make([]int, zlen, zcap)
    copy(z,x)
  }
  z[len(x)] = y
  return z
}
```

A função `appendCap` recebe um slice e um inteiro e retorna um slice. A função `appendCap` verifica se o slice tem capacidade suficiente para adicionar um elemento. Se o slice tiver capacidade suficiente, o slice original é retornado com o novo elemento adicionado. Se o slice não tiver capacidade suficiente, um novo array é alocado e os elementos são copiados para o novo array antes de adicionar o novo elemento. O slice resultante pode ser diferente do slice original. Se o slice original tiver capacidade suficiente, o slice resultante apontará para o mesmo array subjacente que o slice original. Se o slice original não tiver capacidade suficiente, o slice resultante apontará para um novo array alocado.

A função `copy` copia elementos de um slice para outro slice do mesmo tipo e retorna o número de elementos copiados, que é o menor entre o tamanho dos dois slices. O primeiro argumento é o slice de destino *(dst)* e o segundo é o slice de origem *(src)*.

```go
var x []int = []int{1, 2, 3, 4, 5}
var y []int = []int{6, 7, 8, 9, 10}
copy(x, y)
fmt.Println(x) // [6 7 8 9 10]
```

Normalmente, não sabemos se uma dada chamada a append produzirá um slice que aponta para o mesmo array subjacente que o slice original ou para um novo array alocado. De modo semelhante, não devemos supor que atribuições a elementos de slice antiga se refletirão em um slice novo. É comum atribuir o resultado de uma chamada a append a um slice cujo valor passamos para append.

```go
var x []int = []int{1, 2, 3, 4, 5}
x = append(x, 6)
fmt.Println(x,) // [1 2 3 4 5 6]
```

Para usar slices corretamente é importante ter em mente que, embora os elementos do array subjacente sejam indiretos, o ponteiro da fatia, o tamanho e a capacidade não são. Portanto, se um slice for passado para uma função, as alterações feitas pela função no slice serão visíveis fora da função.
