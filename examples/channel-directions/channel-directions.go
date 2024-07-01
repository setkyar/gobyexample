// Channel တွေကို function parameter အဖြစ်သုံးတဲ့အခါ၊ အဲဒီ channel က
// တန်ဖိုးတွေကို ပို့ဖို့ (send) ပဲသုံးမှာလား၊ လက်ခံဖို့ (receive) ပဲသုံးမှာလားဆိုတာ သတ်မှတ်ပေးနိုင်ပါတယ်။
// ဒီလို တိကျစွာသတ်မှတ်ခြင်းက ပရိုဂရမ်ရဲ့ type-safety ကို မြှင့်တင်ပေးပါတယ်။

package main

import "fmt"

// ဒီ `ping` function က တန်ဖိုး (values) တွေပို့ဖို့အတွက်ပဲ သုံးမယ့် channel ကိုပဲ လက်ခံပါတယ်။
// ဒီ channel ကနေ တန်ဖိုး (values) တွေကို လက်ခံဖို့ကြိုးစားရင် compile-time error ဖြစ်ပါလိမ့်မယ်။
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// `pong` function က တန်ဖိုးတွေလက်ခံဖို့ channel တစ်ခု (`pings`) နဲ့
// တန်ဖိုးတွေပို့ဖို့ နောက်ထပ် channel တစ်ခု (`pongs`) ကို လက်ခံပါတယ်။
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
