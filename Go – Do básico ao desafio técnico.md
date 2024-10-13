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
---

---
<!-- header: '**DevFest** _Prudente 2024_ <br> **Controle de Fluxo** '-->

# Controle de Fluxo 

Estruturas Condicionais (if, switch)
Laços de Repetição (for)

---
<!-- header: '**DevFest** _Prudente 2024_ <br> **Funções e Métodos** '-->

# Funções e Métodos 

   Definição e Chamada de Funções
   Parâmetros e Retornos

---
<!-- header: '**DevFest** _Prudente 2024_ <br> **Estruturas de Dados** '-->

# Estruturas de Dados 

   Arrays e Slices
   Mapas (Maps)
   Estruturas (Structs)
   Interfaces

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