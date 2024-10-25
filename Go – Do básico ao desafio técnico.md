---
marp: true
paginate: true
style: |
  section {
     background-color: #67b8e3;
  }
  img[alt~="gofooter"] {
     background-color: #67b8e3;
  }

header: '**DevFest** _Prudente 2024_'
footer: '![gofooter w:40](https://go.dev/blog/go-brand/Go-Logo/PNG/Go-Logo_White.png)'

---

<!-- _paginate: skip -->
![bg right w:600](https://raw.githubusercontent.com/egonelbre/gophers/63b1f5a9f334f9e23735c6e09ac003479ffe5df5/vector/adventure/hiking.svg)

# Go – Do básico ao desafio técnico

---
   
   ### Conteúdo Programático 

   1) Introdução ao Golang
   2) Fundamentos da Linguagem
   3) Controle de Fluxo
   4) Funções e Métodos
   5) Estruturas de Dados
   6) Manipulação de Erros
   7) Pacotes e Modularização
   8) Boas Práticas e Idiomas
   9) Trabalho com APIs e Bibliotecas Externas
   10) Projeto Teste
   11) Conclusão e Próximos Passos
   
---
<!-- header: '**DevFest** _Prudente 2024_ <br> **Introdução** '-->

# Introdução


O que é Golang?
História e Filosofia
Instalação e Configuração do Ambiente
Configuração de IDEs e Ferramentas (vscode)

---
<!-- header: '**DevFest** _Prudente 2024_ <br> **Fundamentos da Linguagem** '-->

# Fundamentos da Linguagem 


Estrutura de um Programa Go
Tipos de Dados Básicos
Variáveis e Constantes
Operadores e Expressões

---
<style scoped>
h6 {
   font-size: 0.95em;
   color: black;
   font-weight: normal;
}
</style>
## Estrutura de um Programa Go
 ###### A estrutura de um programa em Go é simples:
 ```go
   package main
   
   import "fmt" // referencia a outros pacotes

   func main() {
      fmt.Println("Hello World")
}
 ```
   ###### **Package** (pacote) em Go é um mecanismo para organizar e estruturar o código. Cada arquivo Go está associado a um pacote, que podem conter funções, tipos e variáveis.
   ###### **func main()** Ponto de entrada do programa, onde o código é executado.
---
## Tipos de Dados Básicos
Go é uma linguagem **fortemente tipada**, ou seja, cada variável deve ser declarada com um tipo específico, e esse tipo não pode ser mudado durante a execução do programa.

---
### Numéricos:
   -  **Inteiros**: int¹, int8, int16, int32, int64
   -  **Unsigned²**: uint¹, uint8, uint16, uint32, uint64
   -  **Float**: float32, float64
   -  **Complexos**: complex64, complex128
### Booleano: 
   -  bool (true/false)
### Texto:
   -  string

<!-- _footer: ![gofooter w:40](https://go.dev/blog/go-brand/Go-Logo/PNG/Go-Logo_White.png) _1- O tipo **int** e **uint** sem definição de tamanho assume a quantidade de bits da arquitetura do sistema (32 ou 64) </br> &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; 2- Sem sinal, isso indica que o tipo de dado representa apenas números não negativos._ -->
---
### Tipos Derivados:
   - **array**: _Conjuntos de elementos do mesmo tipo com tamanho fixo._
   - **slice**: _Conjuntos dinâmicos de elementos(mesmo tipo), tamanho redimensionavel._
   - **struct**: _Coleção de campos._
   - **map**: _Armazena pares chave-valor(dicionário)._
   - **channel**: _Estruturas que permitem a comunicação entre goroutines, facilitando a sincronização de processos concorrentes._
   - **pointers**:_(ponteiros) referências a locais na memória._
   - **interface**:_Conjunto de métodos que devem ser implementados por tipos que desejam ser considerados como pertencentes àquela interface._
   - **func**: _Representa uma função, permitindo que funções sejam tratadas como valores, passadas como argumentos ou retornadas de outras funções._
---
## Variáveis e Constantes
### Variáveis
Possuem escopo de **pacote** (variáveis globais), **função** ou **bloco** (_condicionais, loops, etc._).

Variáveis **não inicializadas recebem o valor 'zero'** para o tipo correspondente (_por exemplo, 0 para inteiros, false para booleans_).

Variáveis em escopos internos podem realizer **shadowing** em variáveis de escopos externos, mas isso só vale dentro do bloco interno.

Podem ser declaradas com **inferência implícita**(_o tipo da variável é inferido com base no valor que está sendo atribuído._) basta usar walrus operator(_:=_).

---
### Constantes
Valores fixos que não podem ser alterados após sua definição e seguem as mesmas regras de escopo das variaveis. 

Elas podem ser do tipo:

   - Tipada: Associada a um tipo específico.
   - Não tipada: O compilador infere o tipo com base no valor.


---

## Operadores e Expressões

Go possui a maioria dos operadores comuns em outras linguagens de programação, como operadores **aritméticos**, de **comparação**, **lógicos** e também **bitwise** e de **atribuição**, possibilitando assim uma grande variedade de combinações para expressões.

---
<!-- header: '**DevFest** _Prudente 2024_ <br> **Controle de Fluxo** '-->

# Controle de Fluxo 

Estruturas Condicionais
Laços de Repetição

---
## Estruturas Condicionais
Em Go, temos condições para controle de fluxo em blocos de código, como em muitas linguagens, mas também existem algumas de características particulares¹.

As principais são o **if** e o **switch** e suas variações. 

<!-- _footer: ![gofooter w:40](https://go.dev/blog/go-brand/Go-Logo/PNG/Go-Logo_White.png) _1- A keyword_ [select](https://go.dev/ref/spec#Select_statements)  _foi criada especificamente para trabalhar com comunicação em concorrência, funcionando exatamente como um switch._   -->
---
## Laços de Repetição

Repetições em Go são definidas basicamente com uma sintaxe muito parecida com **for** como em C ou Java.
```go
for init; condition; post {
   statements
}
```
Go não possui a estrutura **while**, as repetições são variações do uso do **for** e alguns poucos casos com o uso de **goto**. No entanto, seu uso é desencorajado para repetições.



---

Os usos mais comuns do **for** são em conjunto com a cláusula **range**.
```go
for i, n := range numeros {
    fmt.Printf("Índice: %d, Valor: %d\n", i, n)
}
```

---
<!-- header: '**DevFest** _Prudente 2024_ <br> **Funções e Métodos** '-->

# Funções e Métodos 

   Definição e Chamada de Funções
   Parâmetros e Retornos

---

## Definição e Chamada de Funções

Funções são definidas pela keyword **func** seguida pelo seu _nome_ _parametros_ e _retorno_.

```go
   func Soma(valor1 int, valor2 int) int {
      // lógica
   }
```
Sua chamada é feita pelo nome seguido dos parenteses com os parametros.

```go
   Soma(valor1,valor2)
```
Os valores de retorno são capturados para variaveis.
```go
   resultado := Soma(valor1,valor2)
```
---

Em Go podemos atachar funções a **tipos** definidos pelo usuario, criando assim **metodos**.
```go
   type Pessoa struct{
      Idade int
   }

   func (p Pessoa) EMaiorDeIdade() bool {
      return p.Idade >= 18
   }

   func main() {
      pes := Pessoa{}
      pes.Idade = 21
      if pes.EMaiorDeIdade() {
         fmt.Println("É maior de idade")
      }
   }
```
---
## Parâmetros e Retornos

Os parâmetros em funções não possuem limites, assim como não têm um número mínimo (0..n). Go **não possui parâmetros opcionais nem sobrecarga de funções**. Portanto, caso precise de uma função com parâmetros diferentes, será necessário criar outra declaração que atenda às necessidades.

---

#### Retornos

Go permite que funções e métodos retornem múltiplos valores.
```go
func (p Pessoa) EMaiorDeIdade() (bool,bool) {
      return p.Idade >= 18, p.Idade >= 21
}
```
Retornos tambem podem ser nomeados
```go
func (p Pessoa) EMaiorDeIdade() (eMaiorBrasil bool,eMaiorMundo bool) {
      eMaiorBrasil = p.Idade >= 18
      eMaiorMundo = p.Idade >= 21
      return 
}
```

---
###### Sou obrigado a receber os dois valores das funções mesmo que só precise de um?
###### R.Não!

Podemos usar o simbolo **_**(underscore/underline) para ignorar o valor recebido, isto é, não atribuilo a nada.
```go
func main() {
   maiorIdadeBR , _ := p.EMaiorDeIdade()
}
```

---
<!-- header: '**DevFest** _Prudente 2024_ <br> **Estruturas de Dados** '-->

# Estruturas de Dados 

   Arrays e Slices
   Mapas (Maps)
   Estruturas (Structs)
   Interfaces

---
## Arrays e Slices

É muito comum o questinamento em Go sobre a diferença entre arrays e slices, ambos são usados nos mesmos contextos.

Simploriamente falando podemos resumi-los em:

---
* **Arrays**: Sequencia de elementos do mesmo tipo e de tamanho fixo(não pode ser redimensionado).
 ```go
...
func main() {
   var formacaoTradicional [3]int

   formacaoTradicional[0] = 4
   formacaoTradicional[1] = 4
   formacaoTradicional[2] = 2
   fmt.Println(formacaoTradicional) // Imprime [4 4 2]

   // declaração e inicialização em uma linha
   posicoes := [3]string{"Defesa", "Meio-Campo","Ataque"}
   fmt.Println(posicoes) // Imprime [Defesa Meio-Campo Ataque]
}
 ```
---
* **Slices**: Simplificando, é uma visão flexível e dinâmica dos elementos de um array.
 ```go
 ...
func main() {
    contas := []string{"Conta Corrente", "Conta Poupança", "Conta Investimento"}
    fmt.Println(contas) // Imprime: [Conta Corrente Conta Poupança Conta Investimento]

    // adicionar um novo tipo de conta ao slice
    contas = append(contas, "Conta Digital")
    fmt.Println(contas) // Imprime: [Conta Corrente Conta Poupança Conta Investimento Conta Digital]

    // criando um slice a partir de um array de taxas de juros
    var taxas = [5]float64{1.5, 2.0, 2.5, 3.0, 3.5}
    taxasAltas := taxas[2:5]  // slice das taxas 2.5 3.0 e 3.5
    fmt.Println(taxasAltas)   // Imprime: [2.0 2.5 3.0]
}
 ```
---
## Maps

---
<!-- header: '**DevFest** _Prudente 2024_ <br> **Manipulação de Erros** '-->

# Manipulação de Erros

   Tratamento de Erros em Go

---
<!-- header: '**DevFest** _Prudente 2024_ <br> **Pacotes e Modularização** '-->

# Pacotes e Modularização 

   Importação e Exportação de Pacotes
   Estrutura de Diretórios e Arquivos
   Gerenciamento de Dependências com go mod

---
<!-- header: '**DevFest** _Prudente 2024_ <br> **Boas Práticas e Idiomas** '-->

# Boas Práticas e Idiomas 

   Convenções de Nomenclatura
   Estilo de Código e Formatação
   Comentários e Documentação

---
<!-- header: '**DevFest** _Prudente 2024_ <br> **Trabalho com APIs e Bibliotecas Externas** '-->

# Trabalho com APIs e Bibliotecas Externas 

   Consumo de APIs HTTP
   Uso de Bibliotecas e Pacotes Externos
   Testes e Mocking

--- 
<!-- header: '**DevFest** _Prudente 2024_ <br> **Projeto Teste** '-->

# Projeto Teste 

   Escopo do projeto
   Desenvolvimento Guiado Basico

--- 

## [Desafio Técnico - Dev. Back-End](./Projeto%20Teste/desafio-backend.md)

---

<!-- header: '**DevFest** _Prudente 2024_ <br> **Conclusão e Próximos Passos** '-->

# Conclusão e Próximos Passos

   Revisão dos Conceitos Abordados
   Recursos e Comunidades para Aprendizado Contínuo (https://github.com/avelino/awesome-go, sites , blogs)