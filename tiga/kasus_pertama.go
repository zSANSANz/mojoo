package main

import "fmt"

func main() {
	deretPertama := 2
	deretKedua := 4
	value := 5 - 2

	selisihBilangan := deretKedua - deretPertama

	fmt.Print(deretPertama, ", ", deretKedua, ", ")
	for i := 1; i <= value; i++ {
		fmt.Print(deretKedua+selisihBilangan, ", ")
		deretKedua = deretKedua + selisihBilangan
	}

}
