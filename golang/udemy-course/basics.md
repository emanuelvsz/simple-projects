## Variáveis 
   * ``int``: número inteiro (por padrão: 0)
   * ``int8, int16, int32, int64``: número inteiro com limites de capacidade (por padrão: 0)
   * ``uint``: número inteiro não-negativo (por padrão: 0)
   * ``uint8, uint16, uint32, uint64``: número inteiro não-negativo com limites de capacidade (por padrão: 0)
   * ``float32 e float64``: número fracionado (por padrão: 0)
   * ``string``: caracter (por padrão: "")
   * ``bool``: variável de decisão (por padrão: false)
   * ``error``: variável que retorna um tipo de erro (por padrão: nil)

### Como declarar variáveis em go?

São algumas as formas de declarar variáveis em ``go``, mesmo a linguagem sendo muito rígida com sua tipagem, ela ainda possue algumas excessôes:

- 1° forma: ``var num int = 1`` (declarando o tipo da variável ``com`` tipagem)
- 2° forma: ``num := 1`` (declarando o tipo da variável ``sem``  tipagem)

