## Introdução

Inicialmente, é bom saber um pouco do que é o ``Go``. O Golang, ou mais conhecido como ``Go``, é a linguagem de programação criada pela Google em 2009. Ela foi criada tendo como objetivo revolucionar as linguagens de programação não-orientadas à objetos, sendo ela mais utilizada na produção de back-ends de códigos.
> back-end: É o código que conecta a internet com o banco de dados, gerencia as conexões dos usuários e alimenta a aplicação web

## Pacotes

Os pacotes em golang são uma forma de conseguir coletar dados de arquivos diferentes dentro de seu código. Além de melhorar a organização do mesmo.

### Como acessar pacotes internos? 
  - Importando os pacotes dentro de seu código, da seguinte forma: 
  ```
  import (
    "<nome-do-módulo>/<nome-do-pacote>"
  )
  ```
    
  >Módulos são, de certa forma, um conjunto dos pacotes existentes de um diretório. Para gerar um, utilize o comando no root de seu projeto ``go mod init <nome-do-modulo>``
  
### Como acessar pacotes externos? 
 - com o comando:
  ```
  go get <referencia-do-pacote>
  ```


## Variáveis 
   * ``int``: número inteiro (por padrão: 0)
   * ``int8, int16, int32, int64``: número inteiro com limites de capacidade (por padrão: 0)
   * ``uint``: número inteiro não-negativo (por padrão: 0)
   * ``uint8, uint16, uint32, uint64``: número inteiro não-negativo com limites de capacidade (por padrão: 0)
   * ``float32 e float64``: número fracionado (por padrão: 0)
   * ``string``: caracter (por padrão: "")
   * ``bool``: variável de decisão (por padrão: false)
   * ``error``: variável que retorna um tipo de erro (por padrão: nil)
   * ``byte``: um alias pro ``uint8``
   * ``rune``: um alias pro ``int32``

### Como declarar variáveis em go?

São algumas as formas de declarar variáveis em ``go``, mesmo a linguagem sendo muito rígida com sua tipagem, ela ainda possue algumas excessôes:

- 1° forma: ``var num int = 1`` (declarando o tipo da variável ``com`` tipagem)
- 2° forma: ``num := 1`` (declarando o tipo da variável ``sem``  tipagem)

## Funções

As funções em golang são simples e não tem formas distintas de serem feitas, confira uma forma de fazer uma função em golang:

```
func sum(num1, num2 rune) rune {
  sum := num1 + num2
  return sum
}
```

* O termo ``func`` é a forma em que a função é declarada
* O nome ``sum`` é o nome da função
* Os parenteses após o nome da função, são seus parametros
* O nome ``rune`` é o tipo do retorno da função
  
### Funções dentro de variáveis
```
variavel := func(num1, num2 rune) rune {
  sum := num1 + num2
  return sum
}
```
## Operadores
  
### Operadores matemáticos: 
  * Adição: 1 + 1
  * Subtração: 1 - 1  
  * Multiplicação: 1 * 1  
  * Divisão: 1 / 1  
  * Resto da divisão: 5 % 2
  
### Operadores matemáticos com atribuição:
  * Variável = Variável + 2: ``vari += 2``
  * Variável = Variável - 2: ``vari -= 2``
  * Variável = Variável * 2: ``vari *= 2``
  * Variável = Variável / 2: ``vari /= 2``
  * Varíavel + 1: ``vari++``
  * Varíavel - 1: ``vari--``

### Operadores de atribuição:
  * Recebe: ``=``
  * Recebe: ``:=``
  
### Operadores de lógica:
  * Igual: ``==``
  * Diferente: ``!=``
  * Menor: ``<``
  * Maior: ``>``
  * Maior que: ``>=``
  * Menor que: ``<=``

## Structs
  ### O que são structs?
  ``Structs`` são formas de representar uma classe em golang, já que a linguagem não é orientada à objetos:
  ```
    type Pessoa struct {
      ID    uuid.UUID
      Nome  string
      Idade uint8
    }
  ```
  > NOTA: observe que o nome da struct está começando como maiúsculo. Isso se dá, pois o Go não possue outra forma de ``declarar que um atributo é público``, além desse. Caso o nome da struct fosse ``pessoa``, esse struct seria ``privado``
  
## "Herança"

Caso uma struct queira herdar atributos de outra struct, faça assim:
```
  type Pessoa struct {
    ID    uuid.UUID
    Nome  string
    Idade uint8
  }
  
  type Estudante struct {
    Pessoa
    Curso string
  }
```
  
Desta forma acima, quando um novo struct de Estudante for criado, ele terá aqueles atributos de "Pessoa" diretamente, podendo acessa-los da seguinte forma:
```
pessoa := Pessoa{
  ID: "e84d576a-c5c6-11ed-afa1-0242ac120002",
  Name: "Emanuel Vilela",
  Age: 19
}
  
estudante := Estudante{
  pessoa,
  Curso: "Sistemas de Informação"
}
  
estudante.Nome // acessando o valor direto da variável estudante
```
  
Caso no struct, o atributo "Pessoa" fosse declarado da seguinte forma: ``Pessoa Pessoa``(Primeiro, o nome do atributo, depois o tipo dele, que nesse exemplo, seria do tipo struct Pessoa)
  
A forma de acessar esse valor, seria assim:
```
estudante.Pessoa.Nome
```
  
## Ponteiros
### O que é um ponteiro?
  
É um endereço de memória de uma variável. Como assim? Bom, imagine que você possue um armazem com gavetas numeradas. Você quer inserir um produto dentro da gaveta que tem um item do mesmo que você quer inserir, só que você não tem noção de qual das gavetas você vai inserir ele. Nesse exemplo, basta conferir o número da gaveta em que está esse item. No caso do ponteiro, ele acha onde na memória do computador, que o atributo está armazenado.
  
### Como declarar?
  
```
  var num1 int = 1
  var ponteiro *int = &num1 // o ``&`` significa que o código vai pegar o endereço de memória do atributo selecionado
  
  fmt.Println(ponteiro) // vai retornar o endereço de memória
  fmt.Println(*ponteiro) // vai retornar o valor dentro do endereço na memória, processo chamado de desreferenciação
```
  
## Arrays & Slices
Quando queremos criar uma variável que tenha vários dados do mesmo tipo, utilizamos os arrays e slices. Eles podem ser parecidos, mas não são iguais!
  
Arrays tem um limite propriamente definido no código. Já os slices, não!
  
### Como declarar e atribuir valores nas posições dos arrays?
```
  var arr[4] int // declarando um array de 5 posições do tipo inteiro
  arr[0] = 1 // atribuindo um valor na posição 0
  
```
### Como declarar e atribuir valores nas posições dos arrays?
  
arr := []int{1, 2, 3 , 4, 5,...} // declarando um slice, de x posições e do tipo inteiro
arr = append(arr, 1) // adicionando o valor 1 na posição seguinte do slice

  
