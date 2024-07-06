// Go မှာ `recover` built-in function ကို သုံးပြီး panic ကနေ _recover_ နိုင်ပါတယ်။
// `recover` က `panic` ကြောင့် ပရိုဂရမ် ရပ်တန့်သွားတာကို တားဆီးပြီး
// ဆက်လက်အလုပ်လုပ်စေနိုင်ပါတယ်။

// ဒါဟာ ဘယ်နေရာမှာ အသုံးဝင်သလဲဆိုတဲ့ ဥပမာ: server တစ်ခုအနေနဲ့
// client connection တစ်ခုမှာ critical error ဖြစ်တိုင်း crash ဖြစ်သွားတာမျိုး မလိုချင်ပါဘူး။
// အဲဒီအစား၊ server က အဲဒီ connection ကိုပိတ်ပြီး တခြား client တွေကို
// ဆက်လက်ဝန်ဆောင်မှုပေးချင်မှာပါ။ တကယ်တော့ Go ရဲ့ `net/http` package က
// HTTP server တွေအတွက် ဒီအတိုင်းပဲ default လုပ်ထားပါတယ်။

package main

import "fmt"

// ဒီ function က panic ဖြစ်စေပါတယ်။
func mayPanic() {
	panic("a problem")
}

func main() {
	// `recover` ကို deferred function ထဲမှာပဲ ခေါ်ရပါမယ်။
	// လက်ရှိ function မှာ panic ဖြစ်တဲ့အခါ၊ defer က activate ဖြစ်သွားပြီး
	// သူ့ထဲက `recover` call က panic ကို ဖမ်းယူ(catch) ပါလိမ့်မယ်။
	defer func() {
		if r := recover(); r != nil {
			// `recover` ရဲ့ return value က `panic` ခေါ်တဲ့အခါ ပေါ်လာတဲ့ error ပဲ ဖြစ်ပါတယ်။
			fmt.Println("Recovered. Error:\n", r)
		}
	}()

	mayPanic()

	// ဒီ code က run မှာ မဟုတ်ပါဘူး၊ ဘာလို့လဲဆိုတော့ `mayPanic` က panic ဖြစ်သွားလို့ပါ။
	// `main` ရဲ့ execution က panic ဖြစ်တဲ့နေရာမှာ ရပ်သွားပြီး
	// defer လုပ်ထားတဲ့ closure မှာ ပြန်စပါတယ်။
	fmt.Println("After mayPanic()")
}
