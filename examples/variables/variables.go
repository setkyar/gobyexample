// Go မှာ variable တွေကို ကြိုတင်ကြေငြာပီးတော့ compiler ကအသုံးပြုပါတယ်။
// ဥပမာ function calls တွေရဲ့ type တွေမှန်ရဲ့လားဆိုတာကိုလဲစစ်ဆေးပါတယ်။

package main

import "fmt"

func main() {
	// variable ၁ခုသို့ ၁ခုအထက်ကို `var` နှင့်ကြေငြာနိုင်ပါတယ်။
	var a = "စမှတ်"
	fmt.Println(a)

	// variable ၁ခုနှင့်အထက်ကိုလဲ တကြိမ်ထဲသုံးပီးကြငြာနိုင်ပါတယ်။
	var b, c int = 1, 2
	fmt.Println(b, c)

	// Go ကပထဆုံး variable ရဲ့ အမျိုးအစားကို
	// အသုံးပြုဖို့ကောက်ချက်ချပါတယ်။
	var d = true
	fmt.Println(d)

	// variable တွေကြေငြာတဲ့အချိန်မှာ တန်ဖိုးမသတ်မှတ်ထားဘူးဆိုလျှင်
	// အဲဒီ variable value ကို zero-valued လုပ်ပေးပါတယ်။ ဥပမာ
	// သုညတန်ဖိုးအတွက် `int` က `0` ဖြစ်ပါတယ်။
	var e int
	fmt.Println(e)

	// := syntax သည် variable များကို ကြေငြာခြင်းနှင့်
	// အစပြုခြင်းအတွက် အတိုကောက်အတွက်အသုံးပြုပါတယ်။
	// ဥပမာ၊ var f string = "ပန်းသီး" ဆိုတဲ့ syntax က
	// functions တွေထဲမှာအလုပ်လုပ်ပါတယ်။
	f := "ပန်းသီး"
	fmt.Println(f)
}
