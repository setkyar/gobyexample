//  Go မှာ built-in _multiple return values_ ကို support လုပ်ပါတယ်။
//  ဒီ feature ကို Go မှာအတော်လေးအသုံးပြုပါတယ်၊ ဥမမာဆိုရင် function မှာ
// result ကော error ကောပြန်တာမျိုးပေါ့။

package main

import "fmt"

// ဒီ function မှာ `(int, int)` ဆိုတာက ဒီ function က `int`
// ၂ ခုကို return ပြန်တယ်ဆိုတဲ့အဓိပ္ပါယ်ပါ။
func vals() (int, int) {
	return 3, 7
}

func main() {

	// ဒီမှာ return ပြန်လာတဲ့ values ၂ ခုကို _multiple assignment_ နဲ့
	//  assign လုပ်ထားတာပေါ့။
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// return ပြန်လာတဲ့ ၂ ခုထဲကတခုကိုဘဲယူချင်တယ်ဆို
	// `_` ကိုသုံးလို့ရတယ်။
	_, c := vals()
	fmt.Println(c)
}
