// _Channel_ တွေဆိုတာ concurrent ဖြစ်နေတဲ့ goroutine တွေကို ဆက်သွယ်ပေးတဲ့ ပိုက်လိုင်းတွေပါ။
// goroutine တစ်ခုကနေ channel ထဲကို တန်ဖိုးတွေ ပို့နိုင်ပြီး
// နောက် goroutine တစ်ခုကနေ အဲဒီတန်ဖိုးတွေကို လက်ခံယူနိုင်ပါတယ်။
package main

import "fmt"

func main() {

	// `make(chan val-type)` နဲ့ channel အသစ်တစ်ခု ဖန်တီးပါတယ်။
	// Channel တွေက သူတို့သယ်ဆောင်တဲ့ တန်ဖိုးတွေရဲ့ type အလိုက် type သတ်မှတ်ချက်ရှိပါတယ်။
	messages := make(chan string)

	// `channel <-` syntax ကိုသုံးပြီး channel ထဲကို တန်ဖိုးတစ်ခု _ပို့_ ပါတယ်။
	// ဒီမှာ အပေါ်မှာဖန်တီးထားတဲ့ `messages` channel ထဲကို `"ping"` ဆိုတဲ့စကားလုံးကို
	// goroutine အသစ်တစ်ခုကနေ ပို့လိုက်ပါတယ်။
	go func() { messages <- "ping" }()

	// `<-channel` syntax က channel ကနေ တန်ဖိုးတစ်ခုကို _လက်ခံယူ(receive)_ ပါတယ်။
	// ဒီမှာ အပေါ်မှာပို့လိုက်တဲ့ `"ping"` message ကို လက်ခံယူပြီး print ထုတ်ပါမယ်။
	msg := <-messages
	fmt.Println(msg)
}
