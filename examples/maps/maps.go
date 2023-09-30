// _Maps_ က Go ရဲ့ built-in [associative data type](https://en.wikipedia.org/wiki/Associative_array) တစ်ခုဖြစ်ပါတယ်။
// ( တခြား languages တွေမှာဆိုရင် _hashes_ သို့ _dicts_ လို့ခေါ်ပါတယ်။)
package main

import "fmt"

func main() {

	//  empty map တခုဖန်တီးတဲ့အခါမှာ `make` ကိုအသုံးပြုပါတယ်:
	// ဥပမာ - `make(map[key-type]val-type)`
	m := make(map[string]int)

	// key/value pair တခုတည်ဆောက်မယ်ဆို ပုံမှန် `name[key] = val`
	// syntax ကိုအသုံးပြုနိုင်ပါတယ်။
	m["k1"] = 7
	m["k2"] = 13

	// `fmt.Println` လိုမျိုးနှင့် map ကို print ထုတ်မယ်ဆိုလျှင်
	// သူ့ရဲ့ key/value pairs အကုန်လုံးကိုပြပါလိမ့်မယ်။
	fmt.Println("map:", m)

	// key တခုရဲ့ value ကိုရယူမယ်ဆိုရင် `name[key]` ဆိုပီးအသုံးပြုနိုင်တယ်
	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	// `len` ကို map မှာအသုံးပြုပီးတော့ key/value pairs အရေအတွက်ကိုပြနိုင်တယ်
	fmt.Println("len:", len(m))

	// `delete` ကတော့ map ကနေပီးတော့ key/value pair ကို
	// delete လုပ်တဲ့အချိန်မှာအသုံးပြုနိုင်ပါတယ်။
	delete(m, "k2")
	fmt.Println("map:", m)

	// ကျနော်တို့ map မှာရှိတဲ့ value ကိုရယူမယ်ဆိုရင် optional second return value ရှိတယ်။
	// အဲ့ကောင်ကဘာကိုဆိုလိုတာလဲဆိုရင် ကျနော်တို့ access လုပ်တဲ့ key က map ထဲမှာရှိလားမရှိလားဆိုတာကိုသိရမယ်။
	// အထူးသဖြင့်သူက ကျနော်တို့ map ထဲမှာ keys တွေရှိလား နောက် `0`` သို့ `""` ဖြစ်နေတဲ့ကောင်တွေနှင့် ခွဲခြားဖို့ဆိုရင်အသုံးဝင်တယ်။
	// အခုကျနော်တို့ case မှာဆိုရင် ကျနော်တို့ key ရဲ့ value ကိုမလိုဘူး အဲ့ဒါကြောင့်မလို့ `_`` ကိုသုံးပီးတော့ ignored လုပ်လိုက်တာ။
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	//  map ကိုတလိုင်းထဲမှာ declare လုပ်တာကော initialize လုပ်မယ်ဆို ဒီ syntax ကိုအသုံးပြုနိုင်တယ်။
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
