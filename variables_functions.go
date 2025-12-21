package main

import "fmt"

func topla(a int, b int) int {
	return a + b
}

// Birden fazla değer döndüren fonksiyon (Go'nun imzasıdır)
// Genelde (sonuc, hata) döndürmek için kullanılır.
func islemYap(x int) (int, string) {
	sonuc := x * 2
	bilgi := "İşlem başarılı"
	return sonuc, bilgi
}

func main() {
	// 1. Uzun Yol (var keyword'ü ile):
	// Türü açıkça belirtirsin. Fonksiyon dışında da kullanılabilir.
	var isim string = "Enes"
	var sayi int = 10

	// 2. Kısa Yol (:= operatörü ile):
	// Türü Go otomatik anlar (Type Inference).
	// SADECE fonksiyon içinde çalışır.
	mesaj := "Merhaba"
	aktifMi := true

	fmt.Println(isim, sayi, mesaj, aktifMi)

	fmt.Println(topla(5, 10))
	fmt.Println(islemYap(5))
}
