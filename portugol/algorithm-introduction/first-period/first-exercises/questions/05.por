programa {
  funcao inicio() {
    inteiro c
    escreva("Escreva um valor em celsius para ser convertido para fahrenheit e kelvin\n")
    leia(c)
    celsiusToFahrenheit(c)
    celsiusToKelvin(c)
  }

  funcao celsiusToFahrenheit(inteiro c){
    real f
    f = (c * 9/5) + 32
    escreva(f, "\n")
  }

  funcao celsiusToKelvin(inteiro c){
    real k
    k = c + 273.15
    escreva(k)
  }
}
