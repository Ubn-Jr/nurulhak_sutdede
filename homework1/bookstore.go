//Bu proje bir kitapçının aldığı kitabı
//kitap resimli mi? resimliyse kaç sayfada resim var
//sayfa miktarı
//kitap kapağı standart mı kaplamamı
//kitabın sayfaları saman kağıdı mı, kuşe kağıt mı, veya hamur kağıt mı
//gibi kriterlere bakarak en az 5.5 tl kar edebilmesi için kaç paraya satması gerektiğini
//bulan bir projedir. (Kitapçının kitabı aldığı fiyat yukarıdaki şartlara göre belirlenmektedir.)

package main

import "fmt"

func main() {
	//burda kitabımızla ilgili bilgileri değişkenlerimize atıyoruz
	var normalPagesCount int = 80
	var picturePagesCount int = 10
	var pageType string = "kuşe"
	var isCoverHard bool = true

	//burda hesaplama  yapan fonksiyonumuzu çağırıyoruz
	bookPrice := bookIncludeProfitPriceCalculator(normalPagesCount, picturePagesCount, pageType, isCoverHard)
	fmt.Println("5.5 tl kar etmek için satılması gereken fiyat : ", bookPrice)
}

//5.5 kar elde edilebilmesi için kitabın kaç paraya satılması gerektiğini hesaplayan fonksiyon
func bookIncludeProfitPriceCalculator(normalPagesCount int, picturePagesCount int, pageType string, isCoverHard bool) float32 {
	var totalPrice float32

	//resimsiz sayfalar için hesaplamayı yapması için hasPicture parametresini false olarak çağırıyoruz
	totalPrice = pagePriceCalculator(pageType, normalPagesCount, false)
	//resimli sayfalar için hesaplamayı yapması için hasPicture parametresini false olarak çağırıyoruz
	totalPrice += pagePriceCalculator(pageType, picturePagesCount, true)

	//kitabın kapağı sert kapak mı yoksa normal kapak mı buna bakıyoruz
	if isCoverHard {
		totalPrice += 2
	} else {
		totalPrice += 0.5
	}
	// en son totalPrice yani toplam fiyat üzerine 5.5 karımızı ekleyip fonksiyonumuzu geri dönüşüne veriyoruz.
	return totalPrice + 5.5
}

//sayfaların toplam fiyatını bulan fonksiyon
func pagePriceCalculator(pageType string, pageCount int, hasPicture bool) float32 {
	var pagePerPrice float32

	//hasPicture parametresi ile resimli sayfalar için mi hesap yapacak yoksa resimsiz sayfalar için mi yapacak bunu ayırıyoruz
	if hasPicture {
		//ve ilk if else içinde tekrar bir if else yapıyoruz ve bu seferde kağıt
		//türünün Saman mı kuşe mi yoksa hamur mu olduğuna göre sayfa başı fiyatı buluyoruz
		if pageType == "saman" {
			pagePerPrice = 0.5
		} else if pageType == "kuşe" {
			pagePerPrice = 0.8
		} else {
			pagePerPrice = 1
		}
	} else {
		//ve ilk if else içinde tekrar bir if else yapıyoruz ve bu seferde kağıt
		//türünün Saman mı kuşe mi yoksa hamur mu olduğuna göre sayfa başı fiyatı buluyoruz
		if pageType == "saman" {
			pagePerPrice = 0.1
		} else if pageType == "kuşe" {
			pagePerPrice = 0.3
		} else {
			pagePerPrice = 0.4
		}
	}
	//sayfa başına fiyat ile sayfa sayısını çarparak toplam sayfa fiyatını bulup geri döndürüyoruz
	//burada pageCount int tanımlıydı ancak biz bu değişkenimizin burada bir float32 gibi işlem yapmasını istediğimiz için float32() fonksiyonunun içine soktuk.
	return float32(pageCount) * pagePerPrice
}
