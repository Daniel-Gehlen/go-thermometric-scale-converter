package main

import "fmt"

func main() {
  // Valor do ponto de ebulição da água em Kelvin
  kelvin := 373.15

  // Convertendo para Celsius
  celsius := kelvin - 273.15

  // Apresentando o resultado
  fmt.Printf("Ponto de ebulição da água em Celsius: %.2f °C\n", celsius)
}
