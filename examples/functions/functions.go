// _Functions_ တွေက Go မှာအဓိကအခန်းမှာပါတယ်။ ကျနော်တို့
// functions အကြောင်းကို နမူနာတချို့နှင့် လေ့လာကြပါမယ်။

package main

import "fmt"

// ဒီ function မှာဆိုရင် `int` ၂ ခုကိုယူပီးတော့ သူတို့ပေါင်းလဒ်ကို int အဖြစ် return ပြန်ပေးပါ့မယ်။
func plus(a int, b int) int {

	// Go မှာ return ပြန်မယ်ဆိုပီး `return` ဆိုပီးတော့ပြောပီးတော့မှာ return ပြန်ရပါတယ်။
	// ဘာလို့လဲဆိုလျှင် last expression ကိုအလိုအလျှောက် return ပြန်မပေးလို့ပါ။
	return a + b
}

// ကျနော်တို့ parameter မှာ same type တွေဆက်တိုက်ရှိနေတာဆိုရင်
// parameter type တွေကိုဆက်တိုက်ကြေငြာစရာမလိုဘဲ နောက်ဆုံးမှ
// type ကိုကြေငြာနိုင်ပါတယ်။
func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {

	// function တွေခေါ်သုံးတာကတော့ပုံမှန်လိုပါဘဲ
	// function name ရယ် သူ့ arguments တွေရယ်ပေါ့
	// `name(args)`
	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)
}
