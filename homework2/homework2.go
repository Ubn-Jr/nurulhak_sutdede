// Bir kişinin alışveriş sepetini doldurup kasaya gittiği zaman toplam ödeyeceği ücreti bulan kod.
package main

import "fmt"

func main() {

	//Müşterinin alışveriş sepetini aşağıdaki gibi doldurduğu varsayılmıştır.
	var shoppingCartProducts [5]string = [5]string{"Egg", "Car", "Chocolate", "Milk", "Meat"}
	var shoppingCartProductsPiece [5]float32 = [5]float32{12, 1, 5, 1, 0.8}

	fmt.Println("Shopping Cart Total Price =", calculateTotalShoppingCartPrice(shoppingCartProducts, shoppingCartProductsPiece))
}

//Alışveriş sepetindeki ürünlerin fiyatını hesaplar ve her bir ürünün alınan miktara göre toplam fiyatın yazdırır.
//Sepetteki bütün ürünler için toplam fiyatı geri döndürür.
//return totalShoppingCartPrice float32
func calculateTotalShoppingCartPrice(shoppingCartProducts [5]string, shoppingCartProductsPiece [5]float32) float32 {
	var purchasedProductTotalPrice float32
	var totalShoppingCartPrice float32
	//Mağazadaki ürünler ve fiyatları
	var menu [3]string = [3]string{"Toy", "Junk Food", "Food"}
	var product [3][4]string = [3][4]string{{"Ball", "Car", "Train"}, {"Chips", "Chocolate", "Icecream"}, {"Meat", "Milk", "Egg"}}
	var productPrice [3][4]float32 = [3][4]float32{{10, 8, 15}, {1, 1.5, 2}, {11, 2.3, 0.3}}

	for i := 0; i < len(shoppingCartProducts); i++ {
		for j := 0; j < len(menu); j++ {
			for k := 0; k < len(product); k++ {
				if shoppingCartProducts[i] == product[j][k] {
					purchasedProductTotalPrice = shoppingCartProductsPiece[i] * productPrice[j][k]
					fmt.Println(shoppingCartProducts[i], "*", shoppingCartProductsPiece[i], "=", purchasedProductTotalPrice)
					totalShoppingCartPrice += purchasedProductTotalPrice
					break
				}
			}
		}
	}

	return totalShoppingCartPrice
}
