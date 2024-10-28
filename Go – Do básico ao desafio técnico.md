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

Uma das estruturas de dados mais úteis na ciência da computação é a **tabela hash**. Existem muitas implementações de tabelas hash com propriedades variadas, mas em geral elas oferecem **pesquisas**, **adições** e **exclusões rápidas**. Go fornece um tipo de mapa integrado que implementa uma tabela hash.

Resumindo, mapas são a implementação de uma tabela de KV(chave-valor) como um dicionário, as **chaves são unicas** e possuem um valor atrelados a elas.

---
### Trabalhando com mapas


Um mapa em Go pode ser definido como abaixo:
```go
map[KeyType]ValueType
```
Onde **KeyType** pode ser qualquer tipo (comparável) e **ValueType** pode ser qualquer tipo! 

---
Até mesmo outro mapa:
```go
// mapa de permissões de acesso dos usuários
    permissoes := map[string]map[string]bool{
        "user1": {
            "transacoes": true,
            "investimentos": false,
            "suporte": true,
        },
        "user2": {
            "transacoes": true,
            "investimentos": true,
            "suporte": false,
        },
    }
```
---

## Structs

Go **não é uma linguagem orientada a objetos¹** (ao menos não ao pé da letra); não possui classes, mas sim coleções de campos, as conhecidas **structs**, como em C.

```go
type Conta struct {
   Titular string
   Numero  int
   Saldo   float64
}
```

<!-- _footer: ![gofooter w:40](https://go.dev/blog/go-brand/Go-Logo/PNG/Go-Logo_White.png) _1- Considerando o conceito de paradigma de programação OO, Go não se encaixa, pois, por exemplo, não possui herança._   -->
---

Structs podem conter metodos, para isto basta sinalizar o receptor do metodo. 
```go
type Conta struct {
   Titular string
   Numero  int
   Saldo   float64
}

// assim atachamos o metodo a um "dono" no caso a struct Conta
// e seu receptor tem como alias a var "c"

//   (alias receptor)
func (c Conta) EstaNoVermelho() bool{
   return c.Saldo < 0
}
```

---

## Interfaces

Antes de discutirmos as interfaces, vamos revisitar a definição de conceito que faz Go receber muitos elogios:

- Descreve "o que algo pode fazer", mas não "como fazer".
- Conjunto de métodos (assinaturas) que representam comportamentos que um tipo pode ter.
- Para um tipo ser considerado como interface, ele deve satisfazer(implementar) todos seus metodos.

---
<style scoped>
 section {
   columns: 2;
   display: block;
 }
 
 h1 {
   column-span: all;
 }
 
 h2 {
   break-before: column;
 }
</style>
# 

## Java

```java
interface Animal {
  public void fazUmSom(); 
  public void dorme(); 
}

// Porco "implements" the Animal interface         
class Porco implements Animal {
  public void fazUmSom() {
    System.out.println("O Porco faz: wee wee");
  }
  public void dorme() {
    System.out.println("Zzz");
  }
}
```

## Go
```go 
type Animal interface {
   FazUmSom()
   Dorme()
}

// Porco "implementa" a interface Animal implicitamente
type Porco struct {} 

func (p Porco) FazUmSom() {
   fmt.Println("O Porco faz: wee wee")
}

func (p Porco) Dorme() {
   fmt.Println("Zzz")   
}
```
---

Como vimos acima todo tipo que implemente todos os métodos definidos na **interface** automaticamente "satisfaz" a interface, sem precisar declarar explicitamente essa implementação (diferente de linguagens que usam a palavra-chave implements).

###### Mas o que tem de especial nisso?
###### R. O uso de interfaces reduz o acoplamento do código, eliminando dependências externas.

---



```go
type Cobra struct{} // Cobra satisfaz a interface Animal, mesmo tendo mais metodos
func (c Cobra) FazUmSom() {} // Animal
func (c Cobra) Dorme() {} // Animal
func (c Cobra) TrocaPele() {}

type Pato struct{} // O mesmo para o pato
func (p Pato) FazUmSom() {} // Animal
func (p Pato) Dorme() {} // Animal
func (p Pato) Voa() {}
func (p Pato) Nada() {}

func fazCobraPatoPorcoDormirem(c Cobra, p Pato, prc Porco) {
   c.Dorme()
   p.Dorme()
   prc.Dorme()
}

func fazAnimaisDormirem(animais []Animal){
   for _, a := range animais {
      a.Dorme()
   }
}

func main() {
   // Ao invez de fazer uma funcao que recebe todos os tipos de animais que eu possuo, posso simplesmente
   // usar um slice de interface Animal
   c := Cobra{}
   p := Pato{}
   prc := Porco{}
   fazCobraPatoPorcoDormirem(c,p,prc)

   animais := []Animal{
      c,
      p,
      prc,
   }
   fazAnimaisDormirem(animais)
}


```

---
<!-- header: '**DevFest** _Prudente 2024_ <br> **Manipulação de Erros** '-->

# Manipulação de Erros

   Tratamento de Erros em Go

---
## Tratamento de Erros

Em Go, o tratamento de erros é direto e explícito, feito através de valores de retorno. 

Uma função normalmente retorna tanto o resultado desejado quanto um valor de erro que indica se houve algum problema. 

Diferente de outras linguagens que utilizam exceções, Go trata **erros como valores** comuns (_assunto polêmico_) que devem ser verificados e tratados logo após a execução de cada função. 

---
![bg right fit](https://external-preview.redd.it/u73cIXT24McHoXCW8KQxfpU1yfDMq5E5BD_sZUpSxLA.jpg?width=1080&crop=smart&auto=webp&s=832c8bc0a1a633b9ba3a84d42977833d730930c5)

O famoso ```err != nil ``` divide a comunidade entre: 
- Pior abordagem ja criada. 
- Melhor abordagem de erros ja feita.

---
Vamos formar nossas próprias opiniões:
```go
func sacar(saldo, valor float64) (float64, error) {
	if valor > saldo {
		return saldo, errors.New("saldo insuficiente")
	}
	return saldo - valor, nil
}

func main() {
	saldo, err := sacar(100, 150)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}
	fmt.Println("Novo saldo:", saldo)
}

```

---
<!-- header: '**DevFest** _Prudente 2024_ <br> **Pacotes e Modularização** '-->

# Pacotes e Modularização 

   Importação e Exportação de Pacotes
   Estrutura de Diretórios e Arquivos
   Gerenciamento de Dependências com go mod

---

## Importação e Exportação de Pacotes
Até agora, definimos somente o pacote **main** e importamos o pacote **fmt** sem muitas explicações.

Então vamos entender melhor:

   >Um pacote é construído a partir de um ou mais arquivos fonte que juntos declaram constantes, tipos, variáveis ​​e funções pertencentes ao pacote e que são acessíveis em todos os arquivos do mesmo pacote. Esses elementos podem ser exportados e utilizados em outro pacote.
   [Go-Spec](https://go.dev/ref/spec#Packages)

Cada arquivo .go existente  que você escreve deve começar com uma instrução **package {name}** que indica o nome do pacote do qual o arquivo faz parte.

---

Nos exemplos até agora temos todas nossas funções dentro do pacote main e somente em um unico arquivo. Mas podemos e devemos ser capazes de separar nosso código para melhorar a legibilidade.

Inicialmente podemos criar outro arquivo para segregar funções e tratamentos que viremos a utilizar

Podemos então ter :
   >   ├── pessoa.go
   >   └── main.go

Os arquivos acima se registraram com **package main** assim podemos utilizar livremente fuções, tipos, variaveis e constantes declarados em ambos pois na verdade tudo pertence ao pacote e não aos arquivos.

---
<style scoped>
 section {
   columns: 2;
   display: block;
 }
 
 h1 {
   column-span: all;
 }
 
 h2 {
   break-before: column;
 }
</style>

# 
## main.go
```go
package main

import "fmt"

func main() {
   p := Pessoa{
      Nome:        "Raimundo",
      Sobrenome:   "Nonato",
      CPF:         "9999999999",
      Sexo:        "Masculino",
      Idade:       60,
      EstadoCivil: "Casado",
   }

   // Imprime Raimundo Nonato, Casado, 60 anos de idade portador do CPF:9999999999
   fmt.Println(p.Descricao()) 
}
```

## pessoa.go
```go
package main

import "fmt"
// é muito comum em go termos arquivos pequenos e com somente uma estrutura 
// e seus metodos
type Pessoa struct {
   Nome        string
   Sobrenome   string
   CPF         string
   Sexo        string
   Idade       int
   EstadoCivil string
}

func (p Pessoa) NomeCompleto() string {
   return p.Nome + " " + p.Sobrenome
}

func (p Pessoa) Descricao() string {
   return fmt.Sprintf(
   "%s, %s, %d anos de idade portador do CPF:%s.",
   p.NomeCompleto(), 
   p.EstadoCivil, 
   p.Idade, 
   p.CPF)
}
```
---

Agora que temos mais de um arquivo dentro do projeto, o ```$ go run main.go ``` não sera suficiente, precisamos referenciar todos os arquivos do pacote : 
```$ go run *.go```.



---
### Importação de pacotes
Um dos pontos fortes da linguagem Go é sua biblioteca nativa, que cobre um amplo espectro de domínios, incluindo **redes, manipulação de arquivos, criptografia, desenvolvimento web, testes**.

Ela está disponível automaticamente com a instalação do Go e aqui temos a relação dela [Standard library](https://pkg.go.dev/std).

---

#### Visibilidade e modificadores de acesso

Em go, diferentemente de java, C# e outras linguagens não possuimos uma keyword especifica como **public** ou **private** para controlar acesso e visibilidade.

A visibilidade de um identificador (variável, função, tipo, etc.) é determinada pela primeira letra do nome.

```go
package example

// existem apenas 2 modificadores de acesso no Go; Exportado (público) e Não exportado (privado).
type Sample struct {
	CampoPublico   bool    // Campos públicos/exportados em Go têm nomes de campos em letras maiúsculas
	campoPrivado  string  // Campos privados/não exportados em Go têm a primeira letra minúscula do nome do campo
}
// fora do pacote example o campoPrivado não podera ser acessado/exibido.
```

---
####  Importando um pacote da biblioteca nativa

Para importar um pacote bastar adicionar seu nome a relação do ```import``` do arquivo go, e para referenciar basta usar o ultimo path do import (_rand_).

```go
// detalhe: ao importar um pacote da biblioteca padrão você precisa usar o caminho completo para o 
// pacote na árvore da biblioteca padrão, e não apenas o nome do pacote.
package main

import (
   "fmt"
   "math/rand" // não somente rand
)

type Pessoa struct {
   ...
}

   // aqui vamos usar o pacote math/random para criar um metodo que vamos gerar uma idade aleatoria para a struct pessoa
func (p Pessoa) GeraIdadeAleatoria(maxIdade int) int {
   return rand.Intn(maxIdade) + 1 // para não vir o 0 
}

func (p Pessoa) NomeCompleto() string {
   ...
```

---

## Estrutura de Diretórios e Arquivos

Em Go, a organização de pastas está diretamente ligada à estrutura de pacotes, sendo **cada diretório equivalente a um pacote**. Os arquivos .go na mesma pasta sempre pertencem ao mesmo pacote, ou seja não podemos ter arquivos com duas definições de package na mesma pasta.

      ├── financeiro/
      │   └── financeiro.go    // package financeiro
      ├── usuario/
      │   └── usuario.go       // package usuario
      └── main.go

---

Outro ponto de atenção é que os pacotes não podem se referenciar mutuamente, pois isso cria uma **dependência cíclica** (_import cycle_), resultando em **erro de compilação**.

      ├── financeiro/
      │   └── financeiro.go    // Importa "usuario"
      ├── usuario/
      │   └── usuario.go       // Importa "financeiro"
      └── main.go


---
<style scoped>
 section {
   columns: 2;
   display: block;
 }
 
 h1 {
   column-span: all;
 }
 
 h2 {
   break-before: column;
 }
</style>
# 

## financeiro/financeiro.go

```Go
package financeiro

import (
    "fmt"
    // importa o pacote [usuario], 
    // criando uma dependência cíclica
    "projeto/usuario" 
)

func ProcessarPagamento() {
    fmt.Println("Processando pagamento...")
    usuario.VerificarCredito()
}

func VerificaSaldo() {}

```

## usuario/usuario.go
```Go
package usuario

import (
    "fmt"
    // importa o pacote [financeiro], 
    // criando uma dependência cíclica
    "projeto/financeiro" 
)

func VerificarCredito() {
    fmt.Println("Verificando crédito...")
    financeiro.VerificaSaldo()
}

```
---

## Gerenciamento de Dependências com go mod

Assim como python possui o requirements.txt ou o JS o package.js Go tambem possui um arquivo para controlar dependências o ```go.mod```.

Este arquivo é reponsavel por deixar descrições do projeto, como:
   - Versão do Go utilizada no projeto.
   - Nome do pacote completo do projeto.
   - Dependências diretas e indiretas.


---

Aqui temos um exemplo retirado do repositório do [terraform](https://github.com/hashicorp/terraform/blob/main/go.mod):
```go
module github.com/hashicorp/terraform

go 1.23.1

require (
	github.com/Netflix/go-expect v0.0.0-20220104043353-73e0943537d2
	github.com/ProtonMail/go-crypto v1.0.0
	github.com/agext/levenshtein v1.2.3

   ...more
```

O ```module``` representa diretamente o nome do pacote principal do projeto, logo abaixo temos a versão do Go correspondente usada, e por fim a relação de dependencias diretas e posteriormente as indiretas.

---

Até agora vimos somente exemplos basicos sem o uso propriamente dito do gerenciamento de pacotes do Go.

Para sanar isso vamos seguir os passos para criar um projeto completo.

Com o uso do ```$ go mod init nome-projeto``` criaremos nosso ```go.mod``` e a partir deste momento a importação dos pacotes nos arquivos go, seguirão a nomenclatura escolhida ao executar o comando ```init```.


   > go mod init devfest/joaovitor/jsonreader

Com isso podemos ter nosso ```main.go``` e outras pastas internas como uma ```model``` ou ```entity``` para criarmos a representação de um json.

---

Algumas considerações para o gerencimaneto de dependencias com o mod:

- ```go mod init nome-projeto```: inicializa um novo módulo, criando um arquivo go.mod.
- ```go get endereco-lib-externa```: adiciona ou atualiza uma dependência no projeto.
- ```go mod tidy```: Remove dependências não utilizadas e adiciona as necessárias ao arquivo go.mod.
- Arquivo ```go.sum``` sera gerado após o uso de uma dependencia externa, para guardar o hash da versão usada no projeto.

---
<!-- header: '**DevFest** _Prudente 2024_ <br> **Boas Práticas e Idiomas** '-->

# Boas Práticas e Idiomas 

   Convenções de Nomenclatura
   Estilo de Código e Formatação
   Comentários e Documentação

---

## Convenções de Nomenclatura

Diferentemente de [Java](https://www.oracle.com/java/technologies/javase/codeconventions-introduction.html)/[C#](https://learn.microsoft.com/en-us/dotnet/csharp/fundamentals/coding-style/coding-conventions) e outras linguagens Go não possui uma convenção própriamente dita, temos sim guias de melhores praticas e o recomendado, porem não existe uma documentação explicita.

Isso trouxe duas coisas, liberdade e confusão.

O fato de suas praticas serem livres gerou uma gama de "padrões" onde em cada lugar alguem deseja criar o metodo Universal para a linguagem.

Ao longo de sua existencia varias convenções surgiram e alguns tomaram força o bastante para se manter em diferentes abordagens.

---

Aqui temos algumas das mais famosas:
 * [Effective Go](https://go.dev/doc/effective_go) (publicado no próprio Go blog)
 * [Uber Go Style Guide](https://github.com/uber-go/guide/blob/master/style.md#style) (Uber é uma grande referencia em Go)
 * [Practical Go](https://dave.cheney.net/practical-go/presentations/qcon-china.html#_guiding_principles) (Dave Cheney um dos magos da comunidade)
 * [Go Style](https://google.github.io/styleguide/go/) (by Google)

---

Agora, vamos apresentar algumas das principais """regras""" para garantir que estamos seguindo os princípios do Go:

   *  Simplicidade 
   *  Legibilidade
   *  Produtividade

---

#### Nomenclatura
##### ➡️ Variáveis

* CamelCase (umExemplo), não use snake_case.
* Nome curtos para variáveis, geralmente 3 letras (funcionam bem quando a distância entre sua declaração e o último uso é curta).
* Booleanos devem começar com ```Has, Is, Can or Allow```
* Prefira variáveis de uma **única letra para loops**.
* Não nomeie suas variáveis ​​de acordo com seus tipos (Ex: ❌accountMap).
* Declarar, mas não inicializar, prefira  ```var```
* Declarar e inicializar, use ```:=```



---

##### ➡️ Funções e Metodos

* CamelCase.
* palavras únicas para **parâmetros** e valores de **retorno**, palavras múltiplas para **funções**.

##### ➡️ Interfaces

* Interfaces quando possível devem ser verbos (ex: Receiver, Reader, UserManager)
* Uma interface grande não diz nada.

---

##### ➡️ Pacotes

* lowercase, prefira uma palavra.
* Fuja de nomes genéricos para pacotes (ex: utils, base, common).

##### ➡️ Arquivos

* lowercase e sem _ ou seja [flatcase](https://en.wikipedia.org/wiki/Naming_convention_(programming)#Examples_of_multiple-word_identifier_formats), salva exceção para arquivos de testes.

---

## Estilo de Código e Formatação

Como vimos e ja devemos estar cansados de ver, go foi feito para ser simples, isso se reflete totalmente em sua estrutura de código.

Vamos passar agora por algumas dicas e boas maneiras de se organizar em Go.

---
### Estilo de Código

Sempre que criamos um arquivo em go, na qual teremos mais de uma definição de struct, e as mesmas terão metodos, mantenha suas definições próximas ao seus receptores.

---
<style scoped>
 section {
   columns: 2;
   display: block;
 }
 
 h1 {
   column-span: all;
 }
 
 h2 {
   break-before: column;
 }
</style>

#

## ❌
```go
package example                                                  

type PessoaFisica struct{
   Id string
   Nome string
}

type PessoaJuridica struct{
   Id string
   NomeFantasia string
}

func (pf PessoaFisica) Votar() {}

func (pf PessoaFisica) DeclararImposto() {}

func (pj PessoaJuridica) EmitirNota() {}

func (pj PessoaJuridica) DeclararImposto() {}

```

## ✅
```go
package example                                                  

type PessoaFisica struct{
   Id string
   Nome string
}

func (pf PessoaFisica) Votar() {}

func (pf PessoaFisica) DeclararImposto() {}

type PessoaJuridica struct{
   Id string
   NomeFantasia string
}

func (pj PessoaJuridica) EmitirNota() {}

func (pj PessoaJuridica) DeclararImposto() {}

```

---

### Estilo de Código

Evite aninhamentos desnecessários, seja adepto do "early return". Retorne todos os erros ou caminhos tristes antes da execução padrão da função .

---
<style scoped>
 section {
   columns: 2;
   display: block;
 }
 
 h1 {
   column-span: all;
 }
 
 h2 {
   break-before: column;
 }
</style>

#

## ❌
```go
func transferirFundos(valor float64, saldoRemetente float64) error {
    if valor <= 0 {
        return fmt.Errorf("o valor deve ser maior que zero")
    } else {
        if valor > saldoRemetente {
            return fmt.Errorf("fundos insuficientes")
        } else {
            // lógica de transferência...
            return nil
        }
    }
}

```

## ✅
```go
func transferirFundos(valor float64, saldoRemetente float64) error {
    if valor <= 0 {
        return fmt.Errorf("o valor deve ser maior que zero")
    }

    if valor > saldoRemetente {
        return fmt.Errorf("fundos insuficientes")
    }

    // lógica de transferência...
    return nil
}

```

---

**Nunca** ignore erros, mágica é legal, mágica de mais é bruxaria.

---
<style scoped>
 section {
   columns: 2;
   display: block;
 }
 
 h1 {
   column-span: all;
 }
 
 h2 {
   break-before: column;
 }
</style>

#

## ❌
```go
func lerArquivo(nome string) string {
    arquivo, _ := os.Open(nome) // ignorando o erro
    defer arquivo.Close()

    conteudo, _ := io.ReadAll(arquivo) // ignorando o erro
    return string(conteudo)
}
```
## ✅
```go
func lerArquivo(nome string) (string, error) {
    arquivo, err := os.Open(nome)
    if err != nil {
        return "", fmt.Errorf("não foi possível abrir o arquivo: %w", err)
    }
    defer arquivo.Close() // se algo é aberto, deve ser fechado

    conteudo, err := io.ReadAll(arquivo)
    if err != nil {
        return "", fmt.Errorf("erro ao ler o arquivo: %w", err)
    }
    return string(conteudo), nil
}
```
---

## Formatação

Felizmente, o Go já possui um gerenciador de formatação integrado à sua instalação, o ```gofmt```. Se tudo estiver correto no uso do VSCode, ao salvar seu arquivo, ele é aplicado automaticamente.

Existem também outras ferramentas feitas pela comunidade para ajudar com estilos e detectar pequenos erros no código como [golangci-lint](https://golangci-lint.run/).

---

## Comentários e Documentação

Comentários devem ter propósito, comentar por comentar não resolve nada.

> Good code has lots of comments, bad code requires lots of comments. — Dave Thomas and Andrew Hunt
The Pragmatic Programmer

---

<style scoped>
 section {
   columns: 2;
   display: block;
 }
 
 h1 {
   column-span: all;
 }
 
 h2 {
   break-before: column;
 }
</style>

#

## ❌
```go
const (
    StatusContinue           = 100 // corresponde a continuar
    StatusSwitchingProtocols = 101 // é para retornos onde trocou o protocolo
    StatusProcessing         = 102 // processo ainda esta sendo realizado

    StatusOK                 = 200 // deu tudo certo
)

// Esta função calcula o juro composto aplicado ao montante
// ao longo de um período específico. A fórmula utilizada é:
// montante_final = montante_inicial * (1 + taxa) ^ tempo.
func calcularJuroComposto(montante float64, taxa float64, tempo int) float64 {
    return montante * math.Pow((1+taxa), float64(tempo))
}



```
## ✅
```go
const (
    StatusContinue           = 100 // RFC 7231, 6.2.1
    StatusSwitchingProtocols = 101 // RFC 7231, 6.2.2
    StatusProcessing         = 102 // RFC 2518, 10.1

    StatusOK                 = 200 // RFC 7231, 6.3.1
)

// https://pt.wikipedia.org/wiki/Juro#Juros_compostos
func calcularJuroComposto(montante float64, taxa float64, tempo int) float64 {
    return montante * math.Pow((1+taxa), float64(tempo))
}

// justificativa para decisões de design tambem são aceitaveis

// usamos um mapa para armazenar os usuários, pois a busca por ID
// deve ser O(1) em tempo, garantindo uma performance ideal 
// em sistemas com alta concorrência.
var usuarios = make(map[string]*Usuario)
```

---

Lembre-se 🔖
   > Good naming is like a good joke. If you have to explain it, it's not funny. — Dave Cheney 
   
   > The most important skill for a programmer is the ability to effectively communicate ideas. — Gastón Jorquera

   > Programs must be written for people to read, and only incidentally for machines to execute. — Hal Abelson and Gerald Sussman
   
   > Simplicity is prerequisite for reliability. — Edsger W. Dijkstra 


---


<!-- header: '**DevFest** _Prudente 2024_ <br> **Trabalho com APIs e Bibliotecas Externas** '-->

# Trabalho com APIs e Bibliotecas Externas 

   Consumo de APIs HTTP
   Uso de Bibliotecas e Pacotes Externos
   Testes e Mocking

--- 

## Consumo de APIs HTTP

Para consumir uma API REST em Go usando a biblioteca nativa ```net/http``` é muito simples.

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
   Recursos e Comunidades para Aprendizado Contínuo (sites , blogs)

---
 <style scoped>
section {
    font-size: 22px;
}
</style>

#### Leitura 📝
 * [Go blog](https://go.dev/blog/)
 * [awesome-go](https://github.com/avelino/awesome-go) (lista de varios projetos feitos em Go por tópicos)
 * [Dave Cheney blog ](https://dave.cheney.net/)
 * [victoriametrics blog](https://victoriametrics.com/blog/categories/go-@-victoriametrics/)
 * [Go weekly newsletter](https://golangweekly.com/)
 * [ardanlabs blog](https://www.ardanlabs.com/blog/)
 * [go reddit sub](https://www.reddit.com/r/golang/)
 * [carlosbecker](https://carlosbecker.com/) (BR)
 * [Willem blog](https://www.willem.dev/articles/)
 * [threedots](https://threedots.tech/post/)
 * [Uber Backend blog](https://www.uber.com/en-BR/blog/engineering/backend/) (Backend como um todo, porem artigos de go são ótimos)
 * [Mais](https://go.dev/wiki/Blogs) (uma listagem feita pela comunidade oficial do go)

 ---
 <style scoped>
section {
    font-size: 22px;
}
</style>

 
 #### Videos 📹
 * [The Go Programming Language](https://www.youtube.com/@golang/videos)
 * [justforfunc: Programming in Go](https://www.youtube.com/@JustForFunc/videos)
 * [MarioCarrion](https://www.youtube.com/@MarioCarrion)
 * [GolangSP](https://www.youtube.com/@GolangSP)
 * [package main](https://www.youtube.com/@packagemain/videos)
 * [Golang Dojo](https://www.youtube.com/@GolangDojo/videos)
 * [Gopher Academy](https://www.youtube.com/@GopherAcademy/videos)
 * [GopherConUK](https://www.youtube.com/@GopherConUK)
 * [GopherConIsrael](https://www.youtube.com/@GopherConIsrael/videos)
 * [GopherCon Europe](https://www.youtube.com/@GopherConEurope/videos)
 * [GopherCon Latam](https://www.youtube.com/@gopherconlatam/videos)
 * [MelkeyDev](https://www.youtube.com/@MelkeyDev/videos)
 * [ThePrimeTimeagen](https://www.youtube.com/@ThePrimeTimeagen/videos)

---

 #### Podcasts 🎙️

* [gotime](https://changelog.com/gotime)

 #### Livros 📚

* [100 Go Mistakes and How to Avoid Them ](https://www.amazon.com/100-Mistakes-How-Avoid-Them/dp/1617299596)
* [Concurrency in Go: Tools and Techniques for Developers](https://www.amazon.com/Concurrency-Go-Tools-Techniques-Developers-ebook/dp/B0742NH2SG?ref_=ast_author_mpb)
