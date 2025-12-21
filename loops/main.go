package main

import "fmt"

func main() {
	// 1. Klasik For Döngüsü (C, Java tarzı)
	for i := 0; i < 5; i++ {
		fmt.Println("Sayaç:", i)
	}

	// 2. While Tarzı For (Sadece koşul yazılır)
	sayi := 1
	for sayi < 100 {
		sayi *= 2 // sayi = sayi * 2
	}
	fmt.Println("While bitti:", sayi)

	// 3. Sonsuz Döngü (Sunucu dinlerken vs. kullanılır)
	// for {
	//    fmt.Println("Bu hiç durmaz, çıkmak için break yazmalısın")
	//    break
	// }

	// 4. Range (Foreach tarzı - Diziler/Slice'lar için)
	diller := []string{"Go", "Python", "Java"} // Bu bir Slice (Liste)

	for index, deger := range diller {
		fmt.Printf("%d. dil: %s\n", index, deger)   // %d = int, %s = string
		fmt.Printf("%+v. dil: %+v\n", index, deger) // %v = every type
	}
}
