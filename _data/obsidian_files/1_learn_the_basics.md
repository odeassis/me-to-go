---
concluded:
tags:
  - PROGRAMMING-LANGUAGE
  - GOLANG
  - LEARNING
daily: 2023-11-09
updated: 2024-07-10T09:38
created: 2023-11-09T16:00:00
---

# 1.1 Basic syntax

## Comentários

- Comentários de linha única começam com `//`.
- Comentários de várias linhas são delimitados por `/*` no inicio e `*/` no final.

```go
  // Este é um comentário de linha única

  /*
  Este é um comentário
  de várias linhas
  */
```

## Operadores aritméticos

Os operadores aritméticos `+`,`-`,`*`, e `/` podem ser aplicados a inteiros, números de ponto flutuante e números complexos, mas o operador de resto `%` aplica-se somente a inteiros.
O comportamento de `%` para números negativos varia entre linguagens de programação. Em Go, o sinal do resto é sempre igual ao sinal do dividendo, portanto `-5%3` e `-5%3` são ambos iguais a `-2`. O comportamento de `/` depende de seus operandos serem inteiros, assim `5.0/4.0` e `1.25`, mas `5/4` é `1` porque a divisão de inteiros trunca o resultado na direção do valor zero.

## Operadores unários

Os operadores unários de adição (`+`) e subtração (`-`) são utilizados para modificar o sinal de um valor numérico.

#### Operador Unário de Adição (`+`)

O operador unário de adição (`+`) é na verdade redundante, já que não altera o valor numérico. Ele pode ser usado para enfatizar que um valor é positivo, mas geralmente é omitido.

```go
var a int = 5
var b int = +a //5
```

#### Operador Unário de Subtração (`-`)

O operador unário de subtração (`-`) é utilizado para inverter o sinal de um valor numérico, transformando um valor positivo em negativo e vice-versa.

```go
var a int = 5
var b int = -a // -5
```

## Operadores bit a bit

Em Go os operadores bit a bit são usados para manipular os bits individuais de números inteiros. Eles são úteis em situações onde você precisa realizar operações em um nível mais baixo, como manipular configurações de hardware ou realizar compactação de dados.

Os operadores bit a bit disponíveis em Go:

1. **E bit a bit (&)**: Retorna 1 se ambos os bits na posição correspondente forem 1.
2. **OU bit a bit (|)**: Retorna 1 se pelo menos um dos bits na posição correspondente for 1.
3. **OU exclusivo bit a bit (^)**: Retorna 1 se exatamente um dos bits na posição correspondente for 1, mas não ambos.
4. **Negação bit a bit (~)**: Inverte todos os bits de um número.

```go
// Exemplo de números binários
a := 0b1010 // 10 em decimal
b := 0b1100 // 12 em decimal

// Operação de E bit a bit
fmt.Printf("a & b = %b\n", a&b) // Saída: 1000 (8 em decimal)

// Operação de OU bit a bit
fmt.Printf("a | b = %b\n", a|b) // Saída: 1110 (14 em decimal)

// Operação de OU exclusivo bit a bit
fmt.Printf("a ^ b = %b\n", a^b) // Saída: 0110 (6 em decimal)

// Operação de negação bit a bit
fmt.Printf("^a = %b\n", ^a) // Saída: -1011 (-11 em decimal)
```

## Pacotes

- Um programa em Go é organizado em pacotes.
- O pacote principal é chamado de `main` e é o ponto de entrada de um programa executável.

```go
package main
```

## Importações

- Para usar funcionalidades de outros pacotes, você precisa importá-los.

```go
import "fmt"
```

## Funções

- Funções são blocos de código que podem ser chamados.
- A função `main` é a primeira função executada em um programa Go.

```go
func main() {
// Seu codigo aqui
}
```

## Variáveis

- Variáveis são declaradas usando a sintaxe `var nome tipo`.
- Go é uma linguagem com tipagem estática, o que significa que você deve especificar o tipo da variável.

```go
var x int
x = 10

var nome string = "odeassis"
```

### Declaração curta de variáveis

- Você pode usar `:=` para declarar e atribuir valores a variáveis de forma mais concisa.

```go
firstName := "Francisco"
age := "24"
```

## Estrutura de controle

- Go oferece estruturas de controle padrão, como `if`, `for` e `switch`.

```go
if age >= 18 {
  fmt.Println("Maior de idade")
}


for i := 0; i < 5; i++ {
  fmt.Println(i)
}

weekDay := "quarta"

switch weekDay {
  case "segunda":
	fmt.Println("Dia de trabalho")
  case "sabado", "domingo":
	fmt.Println("Fim de semana")
  default:
	fmt.Println("Outro dia")
}
```

## Arrays e Slices

- Arrays têm um tamanho fixo, enquanto slices são flexíveis.
- Para criar um slice, você pode usar a função `make`.

```go
var array [5]int
slice := []int{1,2,3,4,5}

// Criar um slice de inteiros com tamanho inicial de 3 e capacidade de 5
mySlice := make([]int,3,5)

mySlice[0] = 1
mySlice[1] = 2
mySlice[2] = 0

fmt.Println("Slice: ", mySlice)
fmt.Println("Tamanho:", len(mySlice))
fmt.Println("Capacidade:", cap(mySlice))
```

## Structs

- Go suporta structs para definir tipos de dados personalizados.

```go
type Pessoa struct {
  name string
  age int
}
```

## Funções personalizadas

- Você pode criar suas próprias funções personalizadas.

```go
func sum(a, b int) int {
  return a + b
}
```

# 1.2 Variables and declarations

## Declaração

Em Go, a declaração de variáveis começa com a palavra-chave `var`, seguida do nome da variável e, opcionalmente, seu tipo.

`var nomeDaVariavel tipoDaVariavel = valorDaVariavel`

```go
var myName string = "odeassis"
```

A declaração `var myName string` cria uma variável chamada `myName` do tipo `string`.

```go
var age int
age := 24
```

A declaração `age := 24` cria uma variável `age` e o tipo é inferido automaticamente a partir do valor.

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

Variáveis declaradas sem um valor inicial têm o valor zero **_(zero-value)_**. Por exemplo, o valor zero para um `int` é `0`.

### Declaração Curta

Dentro de uma função, a declaração curta `:=` **_("walrus operator/operador morsa")_** pode ser usada no lugar de uma declaração `var` com tipos implícitos. Esse tipo de declaração também é conhecido como declaração de atribuição. **No entanto, a declaração curta só pode ser usada dentro de uma função, não fora dela.** Fora de uma função, cada instrução começa com uma palavra-chave (como `var`, `func` e `return`) e a declaração curta não é uma palavra-chave.

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

As variáveis em Go têm escopo [[Escopos (Função, Bloco, Léxico)#Escopo Léxico|léxico]], o que significa que elas são visíveis apenas dentro do bloco onde são declaradas. Variáveis globais são declaradas fora de qualquer função e têm escopo global. Variáveis locais são declaradas dentro de uma função e têm escopo local.

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

# 1.3 Data types

Os tipos de dados em Go são divididos em quatro categorias:

- Tipos básicos
- Tipos agregados
- Tipos referências
- Tipos interface

## 1.3.1 bool

O tipo `bool` é usado para representar valores booleanos, ou seja, valores que podem ser `true` (verdadeiro) ou `false` (falso). Booleanos são fundamentais nas tomadas de decisões em lógica condicional, declarações `switch`, declarações `if`, fluxo de controle e lógica em geral.

### Características Principais

- `bool` pode assumir apenas dois valores: `true` ou `false`.
- Variáveis do tipo `bool` podem ser declaradas e inicializadas diretamente com um dos valores booleanos.

```go
var a bool // Declaracao sem inicializacao, valor padrao é false
var b bool = true // Declaracao com inicializacao explicita
c := false // Declaracao com inferencia de tipo
```

Sempre que ver operadores relacionais (`==`, `<=`, `>=`, `!=`, `<`, `>`), o resultado da expressão será um valor booleano.

### Operadores Booleanos

Go oferece vários operadores para manipular valores booleanos:

- **E lógico (`&&`)**: Retorna `true` se ambos os operandos forem `true`.
- **OU lógico (`||`)**: Retorna `true` se pelo menos um dos operandos for `true`.
- **Não lógico (`!`)**: Inverte o valor booleano.

Os operadores `&&` e `||` têm um comportamento de _curto-circuito_: se a resposta já estiver determinada pelo valor do operando esquerdo, o operando direito não é avaliado, tornando seguro escrever expressões como:

```go
s != "" && s[0] == 'x'
```

em que `s[0]` causaria pânico se fosse aplicado a uma string vazia.

```go
var a bool
var b bool = false
c := true

fmt.Println("a && b:", a&&b)
fmt.Println("a || b:", a||b)
fmt.Println("!c:", !c)
```

## 1.3.2 int, int8/16/32/64

Go permite operações aritméticas com inteiros com sinal (signed) e sem sinal (unsigned). Há quatro tamanhos distintos para inteiros com sinal - 8, 16, 32 e 64 bits - representados pelos tipos `int8`, `int16`, `int32`, e `int64`.

| Tipo  | Capacidade                                                                                    |
| ----- | --------------------------------------------------------------------------------------------- |
| int   | 32bits em sistemas 32bits e 64bits em sistemas 64bits, mas não é um alias para int32 ou int64 |
| int8  | -128 : 127                                                                                    |
| int16 | -32768 : 32767                                                                                |
| int32 | -2147483648 : 2147483647                                                                      |
| int64 | -9223372036854775808 : 9223372036854775807                                                    |

```go
var a int8 = 20
fmt.Printf("O tipo de %d e %T\n", a, b)
```

## 1.3.3 byte

O tipo `byte` é sinônimo (_alias_) de `uint8` e enfatiza que o valor é um dado bruto, e não uma quantidade numérica pequena. A faixa de `byte` é `0` até `255` (igual `uint8`).

```go
var a byte = 255
var b uint8 = 255
\\ a e b são do mesmo tipo subjacente e podem armazenar os mesmos valore
```

Usar `byte` em vez de `uint8` pode transmitir a intenção do programador de uma forma mais clara. Quando você vê `byte` em um código, isso geralmente indica que a variável está sendo usada para representar dados brutos ou um único caractere em uma string, e não uma quantidade ou valor numérico que será manipulado matematicamente.

Por exemplo, em um contexto onde estamos lidando com dados binários, como ao ler bytes de um arquivo ou enviar dados através de uma rede, `byte` é mais comum:

```go
data := []byte{0x48, 0x65, 0x6c, 0x6c, 0x6f} // "Hello" em ASCII
```

Ao usar `byte`, a intenção é que o valor não está sendo tratado como um número pequeno que participará de operações aritméticas, mas sim como uma unidade de dados. Em outras palavras, o foco está no fato de que estamos lidando com um fragmento de dados (como um pedaço de texto ou parte de uma estrutura binária), e não com um número para cálculos.

Portanto, a escolha entre `byte` e `uint8` é mais sobre semântica e clareza do código. Usar `byte` ajuda a comunicar a intenção de que estamos trabalhando com dados binários ou caracteres, enquanto `uint8` seria usado em contextos onde realmente estamos lidando com valores numéricos pequenos.

## 1.3.4 uint, uint8/16/32/64

Unsigned significa sem sinal, não assinalado. Então o `uint` é um `int` que desconsidera o sinal.

| Tipo   | Capacidade                                                                                      |
| ------ | ----------------------------------------------------------------------------------------------- |
| uint   | 32bits em sistemas 32bits e 64bits em sistemas 64bits, mas não é um alias para uint32 ou uint64 |
| uint8  | 0 ~ 255                                                                                         |
| uint16 | 0 ~ 65535                                                                                       |
| uint32 | 0 ~ 4294967295                                                                                  |
| uint64 | 0 ~ 18446744073709551615                                                                        |

## 1.3.5 rune

O tipo `rune` é simplesmente um alias para o tipo `int32`. Isso significa que um `rune` é armazenado como um número inteiro de 32 bits com sinal. Portanto, `rune` e `int32` são equivalentes em termos de armazenamento e operações.

```go
var a rune = 'A' // Valor Unicode do caractere 'A' é 65
varb int32 = 65 // Valor numerico 65
// a e b tem o mesmo valor subjacente
```

A convenção de usar `rune` em vez de `int32` é para indicar que o valor está sendo usado como um ponto de código Unicode. Em Unicode, cada caractere é representado por um ponto de código, que é um número associado a um caractere específico.

Usar `rune` torna o código mais legível e claro, indicando que a variável representa um caractere Unicode e não apenas um número.

#### Importância de `rune`

- **Internacionalização e Localização**: O uso de `rune` é crucial para suportar caracteres de diferentes línguas e sistemas de escrita.
- **Processamento de Texto**: Em muitas aplicações que envolvem processamento de texto, é necessário manipular caracteres individuais, e `rune` facilita isso.
- **Compatibilidade Unicode**: Go trata strings como sequências de bytes `UTF-8`, e o uso de `rune` permite manipular corretamente caracteres que podem ser representados por múltiplos bytes em `UTF-8`.

## 1.3.6 float32, float64

Go oferece dois tamanhos de números de ponto flutuante: `float32` e `float64`. Suas propriedades aritméticas são governadas pelo padrão `IEEE754`, implementado por todas as `CPUs` modernas.

Um `float32` oferece em torno de seis dígitos decimais de precisão, enquanto um `float64` oferece cerca de 15 dígitos; deve-se dar preferência a `float64` na maioria dos casos, pois cálculos com `float32` acumulam erros rapidamente.

Números de pontos flutuante podem ser escritos literalmente assim:

```go
const a = 2.71828 // (aproximadamente)
```

#### Quando usar cada tipo?

**`float64`** é geralmente preferido na maioria das aplicações porque oferece uma maior precisão e uma faixa maior. Se você não tiver restrições de memória ou desempenho e precisar de alta precisão em cálculos, `float64` é a escolha padrão.

**`float32`** pode ser usado em situações onde o espaço em memória ou a largura de banda são críticos, como em grandes conjuntos de dados ou em cálculos intensivos onde a precisão ligeiramente menor não é um problema.

```go
var a float32 = 3.14159
var b float64 = 3.141592653589793
```

## 1.3.7 complex64/128

Os tipos `complex64` e `complex128` são usados para representar números complexos. Um número complexo é composto por uma parte real e uma parte imaginária, onde a parte imaginária é representada pelo sufixo `i`.

### `complex64` e `complex128`

- **`complex64`**: Este tipo utiliza dois valores `float32` para representar a parte real e a parte imaginária de um número complexo. Assim, um número complexo `complex64` ocupa 64 bits na memória (2 \* 32 bits). Ele tem uma precisão menor do que `complex128`, mas pode ser mais eficiente em termos de espaço de armazenamento e operações de CPU.
- **`complex128`**: Este tipo utiliza dois valores `float64` para representar a parte real e a parte imaginária de um número complexo. Portanto, um número complexo `complex128` ocupa 128 bits na memória (2 \* 64 bits). Ele oferece uma precisão maior e pode representar números complexos com mais detalhes e menos erros de arredondamento do que `complex64`.

#### Quando usar cada tipo?

**`complex128`** é geralmente preferido na maioria das aplicações, especialmente em cálculos científicos e engenharia, onde a precisão é crucial. Se você precisa de alta precisão em cálculos complexos ou simulações que envolvem números complexos, `complex128` é a escolha padrão.

**`complex64`** pode ser usado em situações onde o espaço em memória ou a largura de banda são críticos, ou quando a precisão ligeiramente menor não é um problema. Em muitos casos, especialmente para aplicações gerais e não críticas em termos de precisão, `complex64` pode ser suficiente e mais eficiente.

```go
var z1 complex64 = 3 + 4i
var z2 complex128 = 3 + 4i

fmt.Println("Valor de z1 (complex64):", z1)
fmt.Println("Valor de z2 (complex128):", z2)
```

Se um número de ponto flutuante literal ou um inteiro decimal literal for seguido de `i`, como `4i`
ele se torna um imaginário literal.

## 1.3.8 uintptr

O tipo `uintptr` em Go é um tipo de dados inteiro sem sinal de largura suficiente para armazenar um ponteiro. Ele é definido na biblioteca padrão de Go e seu principal uso é para interagir com chamadas de sistema ou outras interfaces de baixo nível onde você precisa trabalhar diretamente com ponteiros como números inteiros.

#### Principais Características

- O `uintptr` é utilizado principalmente para converter ponteiros em valores inteiros e vice-versa. Isso é útil quando você precisa manipular diretamente os bits de um ponteiro.
- Pode-se converter um ponteiro para `uintptr` usando a conversão explícita. No entanto, deve-se ter cuidado, pois a manipulação direta de ponteiros pode levar a comportamentos indefinidos se feita incorretamente.
- O `uintptr` é comum em pacotes que interagem com chamadas de sistema ou bibliotecas em outras linguagens que requerem o uso de ponteiros como valores inteiros.

```go
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var x int = 42
	ptr := &x
	uintPtr := uintptr(unsafe.Pointer(ptr))

	fmt.Printf("Valor de x: %d\n", x)
	fmt.Printf("Endereço de x: %p\n", ptr)
	fmt.Printf("uintptr: %x\n", uintPtr)
}
```

#### Aviso de Segurança

Manipular ponteiros diretamente com `uintptr` deve ser feito com cautela. É fácil introduzir erros que podem causar falhas no programa, comportamentos indefinidos ou até mesmo vulnerabilidades de segurança. Aqui estão algumas considerações:

1. **Portabilidade**:
   - O `uintptr` não é garantido ser o mesmo em todas as plataformas. O tamanho de um `uintptr` é o mesmo que o de um ponteiro na plataforma, mas isso pode variar entre diferentes arquiteturas de hardware.
2. **Garbage Collection (GC)**:
   - O runtime de Go usa um garbage collection (GC) que precisa rastrear todos os ponteiros para liberar memória corretamente. Se você converter um ponteiro para `uintptr` e o GC não conseguir rastrear o ponteiro, isso pode impedir que a memória seja coletada corretamente.

## 1.3.9 strings

Uma string é uma sequência imutável de bytes. Strings podem conter qualquer dado, incluindo bytes com valor 0, mas normalmente elas contêm texto legível aos seres humanos.

```go
s := "hello, world"
fmt.Println(len(s)) // 12
fmt.Println(s[0], s[7]) // 104, 119
```

A função `len` devolve o número de bytes (não de Unicode/runas) em uma string.

O n-ésimo byte de uma string não é necessariamente o n-ésimo caractere de uma `string` pois a codificação UFT-8 de um ponto de código que não seja ASCII exige dois ou mais bytes.

A operação de `sub-string` `s[i:j]` produz uma nova string constituída dos bytes da string original, começando do índice i e continuando até o byte no índice j, mas sem incluí-lo. O resultado contém j-i bytes.

```go
fmt.Println(s[0:5]) // hello
```

Podemos concatenar strings utilizando o operador `+`:

```go
str1 := "hello, "
str2 := "world!"

result := str1 + str2

fmt.Println(result) // hello, world!
```

Strings podem ser comparadas com os operadores de comparação como `==` e `<`; a comparação é feita byte a byte, portanto o resultado é a [ordem lexicográfica natural](<https://pt.wikipedia.org/wiki/Ordem_lexicogr%C3%A1fica#:~:text=Em%20matem%C3%A1tica%2C%20uma%20ordem%20lexicogr%C3%A1fica,cartesiano%20de%20dois%20conjuntos%20ordenados.&text=(a%2Cb)%20%E2%89%A4%20(,)%20e%20b%20%E2%89%A4%20b%E2%80%B2>).

String são imutáveis: s sequência de bytes contida em um valor de string jamais pode ser alterada, embora, é claro, possamos atribuir um novo valor a uma variável do tipo string. Para concatenar uma string em outra, por exemplo, podemos escrever:

```go
s := "left foot"
t := s
s += ", right foot"
```

Isso não modifica a string armazenada originalmente em s, mas faz s armazenar a nova string composta pela instrução + =; enquanto isso, t continua armazenando a string antiga.

```go
fmt.Println(s) //"left foot, right foot"
fmt.Println(t) //"left foot"
```

Como as strings são imutáveis, construções que tentam modificar os dados de uma string in-place (no próprio lugar) não são permitidas:

```go
s[o] = 'L' // erro de compilação: não é possível fazer uma atribuição a s[o]
```

#### Strings literais

Um valor de string pode ser escrito como uma _string literal_, isto é uma sequência de bytes entre aspas duplas: "Hello, 世界".

Como os arquivos-fontes do Go são sempre codificados em UTF-8 e as strings de texto de Go são convencionalmente interpretadas como UTF-8, podemos incluir pontos de código Unicode em strings literais. Em uma string literal entre aspas duplas, _sequências de escape_ que começam com uma barra invertida `\` podem ser usadas para inserir valores arbitrários de bytes na string.

```markdown
- \a "alerta" ou sino
- \b backspace
- \f form feed
- \n quebra de linha
- \r carriage return
- \t tabulação
- \v tabulacao vertical
- \' aspas simples (somente em rune literal '\')
- \" aspas duplas (somente em literais "...")
```

Uma _string literal brutal (raw literal string)_ é escrita como `...`, usando crases no lugar de aspas duplas. Em uma string literal bruta, nenhuma sequência de escape é processada; o conteúdo é interpretado literalmente, incluindo barras invertidas e quebras de linhas, portanto uma string literal bruta pode ocupar diversas linhas no código-fonte do programa.

Strings literais brutas são uma maneira conveniente de escrever expressões regulares, que tendem a ter muitas barras invertidas. Elas também são úteis para templates `HTML`, literais `JSON`, mensagens de uso de comandos e informações desse tipo que, com frequência, ocupam várias linhas.

```go
const GoUsage = `Go is a tool for
managing Go source code.

Usage:
	go command [arguments]
`
```

### Unicode

### UTF-8

### Strings e slices

### Conversões

# 1.4 For loop; Range

Na linguagem Go, um loop `for` é uma das estruturas de controle de fluxo mais comuns e flexíveis. Ele permite executar um bloco de código repetidamente enquanto uma condição específica for verdadeira. Ao contrário do que ocorre com outras linguagens de programação que têm vários _constructos_ de looping como o `while`, `do` etc., o go tem somente o loop `for`.

Um _`ForClause loop`_ é definido como tendo uma _instrução inicial_, seguida de uma _condição_ e, depois, uma _instrução de post_. Eles são organizados com a seguinte sintaxe:
`for [initial statement]; [condition]; [post statement] { [action] }`

```go
for i := 0; i < 5; i++ {
	fmt.Println(i)
}
```

- `initial statement`: é executado apenas uma vez, antes do loop começar;
- `condition`: é verificada antes de cada iteração do loop;
- `post statement`: é executado após cada iteração do loop;
- `action`: é o código que será executado em cada iteração do loop.

Também é possível excluir a instrução inicial e a instrução post da sintaxe `for` e usar apenas a condição. Trata-se de um _loop de Condição_:

```go
i := 0

for i < 5 {
	fmt.Println(i)
	i++
}
```

Às vezes, você pode não saber o número de iterações que serão necessárias para concluir uma determinada tarefa. Neste caso, é possível omitir todas as instruções e usar a palavra-chave `break` para sair da execução:

```go
for {
	if someCondition {
		break
	}
}
```

Como exemplo disso, vamos supor que estivéssemos lendo uma estrutura de tamanho indeterminado como a de um [buffer](https://golang.org/pkg/bytes/#Buffer) e não soubéssemos quando terminaríamos a leitura:

```go
...
import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	buf := bytes.NewBufferString("on\ntwo\nthree\nfour\n")

	for {
		line, err := buf.ReadString('\n')

		if err != nil {
			if err == io.EOF {
				fmt.Print(line)
				break
			}
			fmt.Println(err)
			break
		}
		fmt.Print(line)
	}
}
```

#### Loop Infinito

Um loop infinito é um loop `for` sem condição de parada. Ele continuará executando até que seja interrompido explicitamente com a palavra-chave `break`.

```go
for {
    // código a ser executado
    if someCondition {
        break
    }
}
```

#### Range

O operador `range` é usado em conjunção com o loop `for` para iterar sobre os elementos de uma slice, array ou mapa. Ele retorna dois valores: o índice e o valor do elemento.

```go
sharks := []string{"hammerhead", "great white", "dogfish", "frilled", "bullhead", "requiem"}

for index, shark := range sharks {
	fmt.Println(index, shark)
}
```

Ao usar `range` em uma fatia, ele irá sempre retornar dois valores. O primeiro valor será o índice em que a iteração atual do loop está e o segundo será o valor naquele índice. Neste caso, para a primeira iteração, o índice era `0`, e o valor era `hammerhead`.

Também podemos usar o operador `range` para preencher valores de uma fatia:

```go
integers := make([]int, 10)
fmt.Println(integers)

for i := range integers {
	integers[i] = i
}

fmt.Println(integers)
```

Ao iterar por um `map`, o range retornará a key e o valor:

```go
sameSalary := map[string]string{"name": "odeassis", "job": "web develop", "salary": "100"}

for key, value := range sameSalary {
	fmt.Println(key + ": " + value)
}
```

> [!note] >**Nota:** é importante notar que a ordem na qual um mapa retorna é aleatória. Cada vez que você executa esse programa, você pode obter um resultado diferente.

#### Loop aninhados

_Nesting_ (aninhamento) é quando temos um _constructo_ dentro de outro. Neste caso, um loop aninhado é um loop que ocorre dentro de outro loop. Eles podem ser úteis quando o que se quer é uma ação em loop sendo realizada em cada elemento de um conjunto de dados.
`for { [action] for { [action] } }`

```go
numbers := []int{1,2,3}
letters := []string{"a", "b", "c"}

for _, number := range numbers {

	for _, letter := range letters {
		fmt.Println(letter)
	}

fmt.Println(number)
}
```

Os loops `for` aninhados podem ser úteis para iterar através de itens dentro de fatias compostas por fatias. Em uma fatia composta por fatias, se usarmos apenas um loop `for`, o programa dará como resultado cada lista interna como um item:

```go
ints := [][]int {
	[]int{1,2,3},
	[]int{3,2,1},
	[]int{2,1,3},
}

for _, v := range ints {
	fmt.Println(v)
}

for _, i := range ints {
	for _, j := range i {
		fmt.Println(j)
	}
}
```

#### Dicas

- Use o operador `range` para iterar sobre slices, arrays e mapas.
- Use o loop `for` infinito com cuidado, pois pode causar loops infinitos se não for usado corretamente.
- Use loops aninhados para iterar sobre conjuntos de dados relacionados.

# 1.5 if, switch statements

A sintaxe do `if` em Go é simples e direta, mas possui algumas características e variações importantes que vale a pena explorar.

```go
x := 10

if x > 5 {
	fmt.Println("x > 4")
}
```

O `if` pode ser seguido por um `else` para definir um bloco de código alternativo a ser executado se a condição for falsa

```go
x := 10

if x > 20 {
	fmt.Println("x > 20")
} else {
	fmt.Println("x < 20")
}

z := 10

if z > 10 {
	fmt.Println("z > 10")
} else if z == 10 {
	fmt.Println("z == 10)
} else {
	fmt.Println("z < 10")
}
```

Uma característica poderosa do `if` em Go é a capacidade de declarar e inicializar variáveis dentro da própria expressão `if`. Estas variáveis estarão disponíveis apenas no escopo do bloco `if-else`.

```go
if y := 10; y > 5 {
	fmt.Println("y é maior que 5")
} else {
	fmt.Println("y não é maior que 5")
}

// fmt.Println(y) // Isso causará um erro, pois y não está disponível aqui

}
```

Diferente de algumas linguagens como C, C++ e Java, Go não utiliza parênteses ao redor da condição no `if`. No entanto, as chaves `{}` são obrigatórias para definir o bloco de código.

```go
 x := 8

// Correto
if x > 5 {
	fmt.Println("x é maior que 5")
}

// Incorreto
// if (x > 5) {
//     fmt.Println("x é maior que 5")
// }
```

## switch

Ele pode ser usado como substituto para uma sequência de `if-else` quando há múltiplas condições a serem verificadas. Uma diferença entre o `switch` em Go de outras linguagens é que em Go o `switch` só será executa o caso selecionado eliminando assim a necessidade de colocar um `break`.

A sintaxe básica de um `switch`:

```go
switch expression {
case value1:
	// bloco de codigo
case value2
	// bloc de codigo
...
default:
	// bloco de codigo
}
```

#### Sem condição (`switch` sem expressão):

Pode-se omitir a expressão do `switch`, transformando-o em uma estrutura semelhante a uma série de `if-else`. Cada `case` será avaliado como uma expressão booleana independente.

```go
switch {
case x < 0:
    fmt.Println("x é negativo")
case x == 0:
    fmt.Println("x é zero")
case x > 0:
    fmt.Println("x é positivo")
}
```

#### Múltiplos valores em um caso:

Em Go, um único `case` pode comparar a expressão do `switch` com múltiplos valores, separados por vírgulas.

```go
switch day {
case "Saturday", "Sunday":
    fmt.Println("É fim de semana")
default:
    fmt.Println("É um dia útil")
}
```

#### Tipagem e comparações:

As comparações no `switch` devem ser feitas entre valores do mesmo tipo. Caso contrário, ocorrerá um erro de compilação.

# 1.6 defer, Erros, Panic, Recover

#### defer

O `defer` é uma instrução em Go usada para adiar a execução de uma função até que a função envolvente (aquela em que o `defer` é chamado) complete sua execução. Ele é muito útil para realizar tarefas de limpeza ou finalização, como fechar arquivos, liberar recursos ou desbloquear `mutexes`.

**Funcionamento do `defer`**

1. **Execução Postergada**: A instrução `defer` adia a execução da função deferida até que a função atual retorne.
2. **Empilhamento**: Se múltiplas instruções `defer` são usadas, elas são empilhadas e executadas na ordem inversa (Last In, First Out - LIFO).
3. **Avaliação de Argumentos**: Os argumentos da função deferida são avaliados no momento em que o `defer` é chamado, não no momento da execução da função deferida.

```go
file, err := os.Open("exemple.txt")
if err != nil {
	fmt.Println(err)
	return
}

defer file.Close()
```

**Principais cenários de uso**

- **Limpeza de Recursos**: Fechar arquivos, conexões de rede, ou qualquer outro recurso que precise ser liberado.
- **Desbloqueio de Mutexes**: Garantir que um mutex seja desbloqueado após a conclusão de uma seção crítica de código.
- **Recuperação de Panics**: Utilizado juntamente com `recover` para capturar e tratar panics.
- **Logging e Monitoramento**: Medir o tempo de execução de funções ou registrar logs no final da execução.

#### Erros (Errors)

Em vez de exceções, Go usa um padrão explícito para o retorno e verificação de erros, o que torna o fluxo de controle mais claro e previsível.

Em Go, a interface `error` é usada para representar um valor de erro. A definição dessa interface é simples:

```go
type error interface {
    Error() string
}
```

Funções e métodos frequentemente retornam um valor de erro adicional para indicar se houve algum problema:

```go
// Função que pode retornar um erro
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("divisão por zero")
    }
    return a / b, nil
}

func main() {
    result, err := divide(4, 0)
    if err != nil {
        fmt.Println("Erro:", err)
    } else {
        fmt.Println("Resultado:", result)
    }
}
```

#### Pânico (Panic)

O `panic` em Go é uma forma de tratar situações inesperadas e severas que impedem o programa de continuar. Quando uma função chama `panic`, a execução da função é interrompida imediatamente e o controle é passado para a função `defer` mais próxima (se houver), seguida pelo término da execução do programa.

Um exemplo de uso do `panic`:

```go
func main() {
    defer fmt.Println("Defer após o panic")

    fmt.Println("Antes do panic")
    panic("Um problema grave ocorreu!")
    fmt.Println("Esta linha não será executada")
}
```

#### Recuperação (Recover)

A função `recover` é utilizada para recuperar de uma situação de `panic`. Ela deve ser chamada dentro de uma função `defer` para capturar o valor passado para o `panic` e permitir que o programa continue executando.

Exemplo de uso de `recover`:

```go
package main

import "fmt"

func main() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recuperado do pânico:", r)
        }
    }()

    fmt.Println("Antes do panic")
    panic("Um problema grave ocorreu!")
    fmt.Println("Esta linha não será executada")
}
```

O `recover` dentro da função `defer` captura o `panic` e o programa consegue imprimir a mensagem de recuperação antes de terminar.

#### Práticas Comuns e Boas Práticas

1. **Erros como valores:** Sempre que possível, trate erros explicitamente retornando-os como valores adicionais nas funções e métodos.
2. **Pânico apenas para situações excepcionais:** Utilize `panic` apenas em casos excepcionais, como erros de programação ou falhas irrecuperáveis. Não use `panic` para controle de fluxo normal.
3. **Recuperação com cuidado:** Use `recover` para capturar `panic` apenas onde for seguro e faça isso de forma controlada para evitar comportamentos inesperados.

# 1.7 Conditionals

# 1.8 Functions

Em Go, uma função é um bloco de código que pode ser executado várias vezes em diferentes partes de um programa. As funções são úteis para organizar o código, reduzir a repetição e tornar o código mais modular e reutilizável.

Uma função em Go é definida usando a palavra-chave `func` seguida pelo nome da função, parâmetros entre parênteses e o corpo da função entre chaves. Aqui está um exemplo:

```go
func saudacao(nome string) {
 fmt.Println("Olá, " + nome + "!")
}
```

#### Multiple Returns

No Go, uma função pode retornar vários valores. Isto é útil quando uma função precisa retornar vários resultados ou quando um erro precisa ser retornado junto com um resultado.

Para retornar vários valores, separe-os com vírgulas. Aqui está um exemplo:

```go
func multipleReturns() (int, string, erro) {
 retorne 1, "olá", nil
}
```

Esta função retorna três valores: um `int`, uma `string` e um `error`. O valor `error` é `nil` neste exemplo, mas em um cenário do mundo real, pode conter uma mensagem de erro.

Para chamar esta função e recuperar os múltiplos retornos, use a seguinte sintaxe:

```go
resultado1, resultado2, err := multipleReturns()
```

#### Named Returns

Named returns, também conhecidos como parâmetros de resultado nomeados, permitem atribuir nomes aos valores de retorno de uma função. Isso torna o código mais legível e fácil de entender.

Para usar retornos nomeados, atribua um nome a cada valor de retorno na declaração da função. Aqui está um exemplo:

```go
func multipleNamedReturns() (resultado1 int, resultado2 string, err erro) {
 resultado1 = 1
 resultado2 = "olá"
 err = zero
 retornar
}
```

Neste exemplo, a função retorna três valores: `result1` (um `int`), `result2` (uma `string`) e `err` (um `error`).

Ao chamar esta função, você pode usar os retornos nomeados para acessar os valores individuais:

```go
resultado1, resultado2, err := multipleNamedReturns()
fmt.Println(resultado1, resultado2, err)
```

O uso de retornos nomeados torna o código mais legível e fácil de entender, especialmente ao retornar vários valores.

# 1.9 Packages, imports and exports

Um pacote (package) no Go é uma coleção de arquivos Go que são compilados juntos. Cada arquivo Go deve começar com uma declaração de pacote:

```go
package name_of_package
```

O nome do pacote é usado para agrupar funções, tipos e variáveis relacionadas. Por exemplo, a biblioteca padrão do Go contém pacotes como `fmt` (formatação de E/S), `net/http` (HTTP client e server), entre outros.

Para usar funcionalidades de outros pacotes, utilizamos a declaração `import`. Isso permite que reutilizemos código de bibliotecas padrão ou de outros pacotes que criamos ou instalamos.

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, Word!")
}
```

No Go, a visibilidade de funções, tipos e variáveis é controlada pelo nome deles. Se o nome de um identificador (função, tipo, variável, etc.) começa com uma letra maiúscula, ele é exportado e pode ser acessado por outros pacotes. Se começar com uma letra minúscula, ele é privado ao pacote onde foi definido. Isso é útil para encapsular a lógica e esconder detalhes de implementação de outros pacotes.

```go
package my_package

// Exportada
func ExportedFunction() {
	//
}

// Nao exportada (privada)
func privateFunction() {
	//
}
```

# 1.10 Type casting

Diferente de algumas outras linguagens, Go não realiza conversões implícitas entre tipos diferentes; é necessário fazer isso explicitamente. Existem duas formas principais de conversão de tipos em Go: utilizando a função de conversão ou a interface de conversão (`type assertion`).

A forma mais comum e direta de realizar uma conversão de tipo em Go é através de funções de conversão fornecidas pela própria linguagem. Essas funções têm o mesmo nome dos tipos de dados.

```go
var intNum int = 42
var floatNum float64 = float64(intNum)
var uintNum uint = uint(floatNum)
```

Nem todas as conversões são permitidas. Por exemplo, você não pode converter diretamente uma `string` para um `int`. Você precisaria usar uma função de biblioteca padrão, como `strconv.Atoi` para fazer essa conversão.

Converter números em strings utilizando o método `strconv.Itoa` para números inteiros e o método `fmt.Sprint` para floats.

```go
import (
	"fmt"
	"strconv"
)

func main() {
	numero := strconv.Itoa(5)
	fmt.Printf(numero)
	fmt.Println(fmt.Sprint(421.034))
}
```

Converter strings em números utilizando o método `strconv.Atoi` para strings de números inteiros e o método `strconv.ParseFloat` para strings de floats.

```go
import (
	"fmt"
	"strconv"
)

func main() {
	numero, _ := strconv.Atoi("5")
	fmt.Println(numero)

	numero, _ := strconv.ParseFloat("421.034", 64)
	fmt.Println(numero)
}
```

#### Conversão usando type assertion

Quando estamos lidando com interfaces, é comum precisar converter de uma interface para um tipo concreto. Isso é feito utilizando `type assertion`.

```go
var i interface{} = "Hello"

s := i.(string)
fmt.Printf("value: %s - type: %T", s,s)

// converter interface{} para string com verificacao
if s, ok := i.(string); ok {
	fmt.Println(s)
} else {
	fmt.Prinln("not string")
}
```

# 1.11 Type inference

Type inference é um recurso de linguagens de programação, incluindo Go, onde o compilador deduz automaticamente o tipo de uma variável a partir do valor que lhe é atribuído, em vez de exigir que o programador declare explicitamente o tipo. Isso simplifica o código, tornando-o mais legível e reduzindo a quantidade de código que o programador precisa escrever.

Em Go, a inferência de tipos é feita principalmente usando a declaração curta `:=`. Quando você usa essa forma de declaração, o compilador de Go determina o tipo da variável a partir do valor à direita do operador `:=`.

```go
a := 10
fmt.Printf("%T", a) // int

b := 3.14
fmt.Printf("%T", b) // float64

c := "VAI CORINTHIANS"
fmt.Printf("%T", c) // string
```

Inferência com funções

```go
func getWinner() {
	return "cotinthians"
}

func main() {
	winner := getWinner()
	fmt.Printf("%T", winner) // string
}
```

Ao permitir que o compilador/interpretador determine automaticamente os tipos, há menos chances de cometer erros ao declarar tipos manualmente. E apesar de tornar o código mais conciso, a inferência de tipos pode, em alguns casos, tornar o tipo das variáveis menos óbvio para quem lê o código. Isso pode ser mitigado com o uso de nomes de variáveis descritivos.

# 1.12 Arrays

Um array é uma sequência de tamanho fixo de zero ou mais elementos de um tipo específico. Arrays são declarados usando a sintaxe `var nome [tamanho]tipo`.

O tamanho do array é parte do seu tipo, ou seja, `[5]int` e `[10]int` são tipos diferentes. O tamanho do array deve ser um valor inteiro constante, ou uma expressão constante.

```go
var x [5]int = [5]int{1, 2, 3, 4, 5}

x = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // Erro de compilação
```

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

Por padrão, os elementos de um array são inicializados com o [[#1.2 Variables and declarations#Atribuição de Valores|valor zero]] do tipo do elemento. Para arrays de inteiros, isso significa que todos os elementos são inicializados com 0. Para arrays de strings, isso significa que todos os elementos são inicializados com uma string vazia, `""`. Para arrays de booleanos, isso significa que todos os elementos são inicializados com `false`.

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

Se usarmos reticências (`...`) no lugar do tamanho do array, o compilador irá inferir o tamanho do array a partir do número de elementos no array literal.

```go
var x [5]int = [...]int{1, 2, 3, 4, 5}
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

#### Arrays como parâmetros para funções

Quando uma função é chamada, uma cópia de cada argumento é passada para os parâmetros da função. Para arrays, isso significa que uma cópia do array é passada para a função. Se a função modificar o array, as mudanças não serão visíveis fora da função, porque a função está trabalhando com uma cópia do array. Passar arrays grandes dessa maneira pode ser ineficiente e, em alguns casos, impraticável.

Mas claro que podemos passar um ponteiro para um array, e a função pode modificar o array original. Por exemplo, a função `zero` recebe um ponteiro para um array e atribui zero a cada elemento do array.

```go
func zero(ptr *[32]byte) {
  *ptr = [32]byte{}
}
```

Usar um ponteiro pra um array é eficiente e permite a que função chamada altere a variável de quem chamou, mas o arrays continuam sendo um tipo fixo, ou seja, não podemos alterar o tamanho de um array. Arrays raramente são usados como parâmetros de funções em Go. Em vez disso, slices são usados.

# 1.13 Slices

Slices são uma estrutura de dados poderosa e flexível que proporcionam uma maneira de trabalhar com sequências de elementos, oferecendo uma interface de alto nível para a manipulação de arrays subjacentes. Slices são, essencialmente, referências a arrays ou partes de arrays, o que permite operações como redimensionamento e partição de forma eficiente.

Um slice é definido por três componentes:

1. **Pointer**: Um ponteiro para o primeiro elemento do array que é acessado pelo slice.
2. **Length**: O número de elementos no slice.
3. **Capacity**: O número de elementos no array subjacente, começando do primeiro elemento no slice.

```go
var x []int
fmt.Println(len(x), cap(x)) // 0 0
```

O tamanho e a capacidade de um slice x podem ser obtidos usando as expressões `len(x)` e `cap(x)`. O ponteiro aponta para o primeiro elemento do array que é acessível através do slice, que **não é necessariamente o primeiro elemento do array**. O primeiro elemento do slice pode ser acessado usando a expressão `x[0]`, e o último elemento do slice pode ser acessado usando a expressão `x[len(x)-1]`.

Vários slices podem compartilhar o mesmo array subjacente. Por exemplo, o slice `x` compartilha o mesmo array subjacente com o slice `y`.

```go
var x []int = []int{1, 2, 3, 4, 5}
var y []int = x[1:3]
fmt.Println(len(y), cap(y)) // 2 4
```

O operador de slice `s[i:j]` cria um slice de um slice. O novo slice aponta para o mesmo array subjacente que o slice original, mas o novo slice acessa apenas os elementos do array no intervalo `i` até `j-1`. O novo slice tem tamanho `j-i` e capacidade `cap(s)-i`. Se `i` for omitido, o slice começará no primeiro elemento do array subjacente. Se `j` for omitido, o slice terminará no último elemento do array subjacente. Se ambos `i` e `j` forem omitidos, o slice será igual ao slice original.

Slice além de `cap(s)` causa um erro de tempo de execução, mas além de `len(s)` estende o slice, aumentando seu tamanho.

Como um slice contém um ponteiro para um array, passar um slice para uma função significa que a função pode modificar os elementos do array subjacente que o slice aponta. Por exemplo, a função `zero` recebe um slice e atribui zero a cada elemento do array subjacente.

```go
func zero(ptr []int) {
  for i := range ptr {
    ptr[i] = 0
  }
}
```

Um slice literal é como um array literal, mas sem o tamanho. Um slice literal é uma lista de valores separados por vírgulas, envolvidos por chaves. Isso cria implicitamente uma variável de array subjacente do tamanho correto e um slice que aponta para o array. Por exemplo, o slice literal `{1, 2, 3, 4, 5}` cria um array de cinco elementos e um slice que aponta para os cinco elementos. O slice tem tamanho 5 e capacidade 5.

Diferente dos arrays, slices não são comparáveis, então não podemos usar o operador ` ==` para testar se dois slices são iguais. A única exceção é o slice nulo, que é igual a `nil`.

O valor zero de um tipo slice é `nil`. Para testar se dois slices são iguais, devemos escrever uma função que itera sobre os elementos de cada slice e testa se eles são iguais.

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

#### Operações com Slices

**Append**
A função `append` permite adicionar elementos a um slice, redimensionando-o conforme necessário.

```go
slice := []int{1,2,3}
slice = append(slice, 4, 5)

fmt.Println(slice) // [1 2 3 4 5]
```

**Copy**
A função `copy` permite copiar elementos de um slice para outro.

```go
src := []int{1,2,3}
dest := make([]int, len(src))
copy(dest, src)

fmt.Printl(dest)
```

**Slicing**
Você pode criar um novo slice a partir de um existente, especificando um intervalo.

```go
slice := []int{1,2,3,4,5}
newSlice := slice[1:4] // novo slice de slice[1] a slice[3]

fmt.Println(newSlice) // [2 3 4]
```

#### Cuidados ao usar slices

- **Compartilhamento de Arrays**: Slices compartilham o mesmo array subjacente, o que significa que mudanças em um slice podem afetar outros slices que compartilham o mesmo array.
- **Cópias de Slices**: Copiar um slice para outro não copia os elementos, mas sim os ponteiros para o array subjacente.
- **Capacidade e Redimensionamento**: Ao usar `append`, se a capacidade do slice for excedida, um novo array é alocado e os elementos são copiados, o que pode ter implicações de desempenho.

# 1.14 Maps

Maps são uma das estruturas de dados mais poderosas e versáteis em Go. Eles são implementados como tabelas hash, permitindo uma coleção não ordenada de pares chave-valor, onde cada chave é única. A grande vantagem das tabelas hash é a eficiência: a recuperação, atualização ou remoção de valores ocorre, em média, em tempo constante, independentemente do tamanho do map.

Todas as chaves em um map devem ser do mesmo tipo, assim como todos os valores. No entanto, as chaves e os valores não precisam ser do mesmo tipo entre si. As chaves devem ser comparáveis usando `==`, o que permite ao map verificar se uma chave já existe. Não há restrições sobre os tipos de valores.

O valor zero de um tipo map é `nil`, representando uma referência a nenhuma tabela hash.

Você pode criar um map usando a função built-in `make`:

```go
ages := make(map[string]int)
```

Alternativamente, você pode usar um map literal para inicializar um map com alguns pares chave-valor:

```go
ages := map[string]int{
	"alice": 32,
	"charles": 23,
}
```

Elementos em um map são acessados e modificados usando a notação de indexação:

```go
ages["alice"] = 32
```

Para remover um elemento, use a função built-in `delete`:

```go
delete(ages, "alice")
```

Todas essas operações são seguras mesmo se o elemento não estiver no map. Se você buscar uma chave que não está presente, o map retornará o valor zero para o tipo do valor.

> [!warning]
> Um elemento de map não é uma variável, e não podemos obter seu endereço.

A maioria das operações com maps, incluindo buscas, `delete`, `len` e loops range, é executada de forma segura em uma referência `nil` de map, pois ela se comporta como um map vazio. Porém, fazer um armazenamento em um map `nil` causa pânico.

> [!note]
> Você deve alocar o map antes de poder armazenar valores nele.

Acessar um elemento do map por meio de indexação sempre produz um valor. Se a chave está presente no map, obterá o valor correspondente; caso contrário, o valor zero para o tipo do elemento será obtido.

```go
if age, ok := ages["bob"]; !ok {
	//
}
```

Usar indexação em um map nesse contexto produz dois valores; o segundo valor é um booleano que informa se o elemento estava presente. A variável booleana geralmente é chamada de `ok`, em especial se ela for usada imediatamente em uma condição if.

Como ocorre com os slices, os maps não podem ser comparados uns com os outros; a única comparação permitida é um nil.

# 1.15 make()

A função `make()` uma função embutida (_built-in_) usada para alocar e inicializar objetos de tipos slice, map e channel. Diferentemente de `new()`, que apenas aloca memória, `make()` inicializa a estrutura de dados, preparando-a para uso. Vamos explorar a função `make()` em detalhes, incluindo sua sintaxe.

```go
make(t Type, size ...IntegerType) Type
```

- **t**: O tipo do objeto a ser criado (slice, map ou channel).
- **size**: Um ou dois inteiros que especificam o tamanho inicial e, no caso dos slices, a capacidade (opcional).

#### Uso de `make()` com Slices

Quando usada para criar slices, `make()` aloca um array subjacente e retorna um slice que faz referência a esse array. É possível especificar o comprimento (length) e a capacidade (capacity) do slice.

```go
slice := make([]int, 5) // Cria um slice com comprimento e capacidade de 5
fmt.Println(slice) // [0 0 0 0 0]
fmt.Printf("len: %d - cap: %d\n", len(slice), cap(slice)) // len: 5 - cap: 5

sliece2 := make([]int, 5,6) // Cria um slice com comprimento 5 e capacidade 6
fmt.Println(slice2) // [0 0 0 0 0]
fmt.Printf("len: %d - cap: %d\n", len(slice2), cap(slice2)) // len: 5 - cap: 6
```

#### Uso de `make()` com Maps

Quando usada para criar mapas (maps), `make()` inicializa e retorna um mapa vazio.

```go
m := make(map[string]int) // Cria um mapa vazio com chave string e valor int
m["a"] = 1
m["b"] = 2

fmt.Println(m) // map[a:1, b:2]
```

#### Uso de `make()` com Channels

Quando usada para criar canais (channels), `make()` inicializa e retorna um canal. É possível criar canais bufferizados especificando a capacidade.

```go
ch := make(chan int) // Cria um canal não bufferizado
go func() {
	ch <- 1
}()

fmt.Println(<-ch) // 1

ch2 := make(chan int, 2) // Cria um canal bufferizado com capacidade 2
ch2 <- 1
ch2 <- 2

fmt.Println(<-ch) // 1
fmt.Println(<-ch) // 2
```

#### Comparação entre `make()` e `new()`

- **`make()`**: Usada apenas para slices, maps e channels. Aloca e inicializa o objeto, retornando um valor inicializado do tipo (slice, map ou channel).
- **`new()`**: Usada para alocar memória para qualquer tipo. Retorna um ponteiro para o valor alocado.

```go
p := new(int)   // Aloca memória para um int e retorna um ponteiro para ele
fmt.Println(*p) // Output: 0
*p = 2
fmt.Println(*p) // Output: 2
```

#### Detalhes Internos

**Slices**
Quando `make()` é usada para slices, um array subjacente é alocado. O comprimento e a capacidade podem ser diferentes, permitindo o crescimento do slice até a capacidade definida sem realocar o array.

**Maps**
`make()` inicializa um mapa vazio e pronto para uso. Isso evita `nil` maps, que causariam um panic ao tentar adicionar elementos.

**Channels**
`make()` inicializa canais com a capacidade opcional. Canais bufferizados permitem armazenamento temporário de valores até a capacidade definida, enquanto canais não bufferizados bloqueiam até que o valor seja lido.

# 1.16 Structs

Uma `struct` é um tipo de dado agregado que agrupa zero ou mais valores nomeados de tipos quaisquer como uma única entidade.

```go
type Employee struct {
	ID int
	Name string
	Address string
	DoB time.Time
	Position string
	Salary float64
	ManagerID int
}
```

Os campos individuais são acessados como a notação de ponto:

```go
var employee1 Employee

employee1.Name
```

Campos normalmente são escritos um por linha, com o seu nome antes do tipo, no entanto campos consecutivos de mesmo tipo podem ser combinados:

```go
type Employee struct {
	ID int
	Name, Address string
	...
}
```

A ordem dos campos é significativa para a identidade do tipo. Se tivéssemos combinados também a declaração do campo `Position` ou trocado a ordem de `Name` e `Address`, estaríamos definido um tipo diferente de estrutura. Normalmente, só combinamos as declarações de campos relacionados.

Um tipo nomeado `S` para struct não pode declarar um campo do mesmo tipo `S`: um valor agregado não pode conter a si mesmo. Porém, `S` pode declarar um campo do tipo ponteiro `*S`, que nos permite criar estruturas de dados recursivas como listas ligadas e árvores.

## 1.16.1 Estruturas literais
