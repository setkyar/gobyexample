// [ယခင်](range) ဥပမာမှာ `for` နဲ့ `range` က အခြေခံ data structure တွေကို
// ဘယ်လို iterate လုပ်သလဲဆိုတာ တွေ့ခဲ့ပါတယ်။ ဒီ syntax ကိုပဲ channel ကနေ
// လက်ခံရရှိတဲ့ တန်ဖိုးတွေကို iterate လုပ်ဖို့လည်း သုံးနိုင်ပါတယ်။

package main

import "fmt"

func main() {

	// `queue` channel ထဲက တန်ဖိုး 2 ခုကို iterate လုပ်ပါမယ်။
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	// ဒီ `range` က `queue` ကနေ လက်ခံရရှိတဲ့ element တိုင်းကို iterate လုပ်ပါတယ်။
	// အပေါ်မှာ channel ကို `close` လုပ်ထားတဲ့အတွက် တန်ဖိုး 2 ခုကို
	// လက်ခံရရှိပြီးတဲ့နောက် iteration က ရပ်တန့်သွားပါတယ်။
	for elem := range queue {
		fmt.Println(elem)
	}
}
