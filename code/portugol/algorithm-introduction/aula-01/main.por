programa {
  funcao inicio() {
    real b1, b2, b3, b4, result

    escreva("Digite a nota do primeiro bimestre: ")
    leia(b1)
    escreva("Digite a nota do segundo bimestre: ")
    leia(b2)
    escreva("Digite a nota do terceiro bimestre: ")
    leia(b3)
    escreva("Digite a nota do quarto bimestre: ")
    leia(b4)

    result = (b1+b2+b3+b4)/4

    escreva("média: ", result)

    se (result <= 6){
      escreva("\nInfelizmente o aluno não conseguiu passar")
    } 
    senao {
      escreva("\nO aluno passou de ano!!!")
    }
  }
}
