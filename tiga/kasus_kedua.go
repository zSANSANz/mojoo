package main

import "fmt"

func main() {
	deretPertama := 5
	deretKedua := 8
	value := 7 - 2

	selisihBilangan := deretKedua - deretPertama

	fmt.Print(deretPertama, ", ", deretKedua, ", ")
	for i := 1; i <= value; i++ {
		fmt.Print(deretKedua+selisihBilangan, ", ")
		deretKedua = deretKedua + selisihBilangan
	}

}
