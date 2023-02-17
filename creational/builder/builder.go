package builder

import "fmt"

// Permite produzir o produtos diferentes usando o mesmo processo de construção
// O Builder também é usando quando o produto é complexo e requer varias etapas
// para ser concluido, mantendo o produto privado até que seja totalmente finalizado,

// No código abaixo, vemos diferentes tipos de casas (igloo e normalHouse)
// sendo construídas por iglooBuilder e normalBuilder. Cada tipo de casa tem
// as mesmas etapas de construção. A struct opcional diretor ajuda a organizar
// o processo de construção.

func MainBuilder() {
	normalBuilder := getBuilder("normal")
	iglooBuilder := getBuilder("igloo")

	director := newDirector(normalBuilder)
	normalHouse := director.buildHouse()

	fmt.Printf("Normal House Door Type: %s\n", normalHouse.doorType)
	fmt.Printf("Normal House Window Type: %s\n", normalHouse.windowType)
	fmt.Printf("Normal House Num Floor: %d\n", normalHouse.floor)

	director.setBuilder(iglooBuilder)
	iglooHouse := director.buildHouse()

	fmt.Printf("\nIgloo House Door Type: %s\n", iglooHouse.doorType)
	fmt.Printf("Igloo House Window Type: %s\n", iglooHouse.windowType)
	fmt.Printf("Igloo House Num Floor: %d\n", iglooHouse.floor)
}
