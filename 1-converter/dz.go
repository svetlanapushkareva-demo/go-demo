package main

import "fmt"

func main() {
	const USD2EUR = 2
	const USD2RUB = 100

	var EURMoney float64 = 1000

	fmt.Println(EURMoney, "изначальные евро")

	USDMoney := EURMoney * USD2EUR

	fmt.Println(USDMoney, "доллары")

	RUBMoney := USDMoney * USD2RUB

	fmt.Println(RUBMoney, "рубли")
}
