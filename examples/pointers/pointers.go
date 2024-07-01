// Go က pointers တွေကို support လုပ်ပါတယ်။
// Go မှာ pointer ဆိုတာက variable တစ်ခုရဲ့ memory address ကို သိမ်းထားတဲ့ variable လို့ နားလည်နိုင်ပါတယ်။
// Pointer ကိုသုံးပြီးတော့ variable ကို တိုက်ရိုက်တန်ဖိုးမပြင်ဘဲ သွယ်ဝိုက်ပြီး access လုပ်နိုင်ပါတယ်။
// Pointer တွေက variable တစ်ခုကို copy မလုပ်ဘဲ reference ကို pass လုပ်ချင်တဲ့အခါမှာ အသုံးဝင်ပါတယ်။

package main

import "fmt"

// pointer နဲ့ value သုံးတာ ဘာကွာခြားလဲဆိုတာကို ပြဖို့အတွက် zeroval ဆိုတဲ့ function နဲ့
// zeroptr ဆိုတဲ့ function တွေကို သုံးပြီး ရှင်းပြပါမယ်။ zeroval function က
// `int` parameter ကို လက်ခံပါတယ်။ ဒီမှာ ကျွန်တော်တို့က parameter မှာ value ထည့်ပေးလိုက်ပါတယ်။
// zeroval function က ival ရဲ့ copy ကို ရရှိပါတယ်။ ဒီ copy က သူ့ကိုခေါ်တဲ့
// function ထဲက ကောင်နဲ့ မတူဘူးဆိုတာကို သတိထားပါ။
func zeroval(ival int) {
	ival = 0
}

// zeroptr function ကတော့ `*int` parameter ရှိပါတယ်။ ဆိုလိုတာက
// သူက `int` pointer ကို လက်ခံတယ်လို့ ဆိုလိုပါတယ်။ function ထဲမှာရှိတဲ့ `*iptr` code က
// pointer ကို သူ့ရဲ့ memory address ကနေ လက်ရှိတန်ဖိုးဆီကို *dereference* လုပ်ပါတယ်။
// Dereference လုပ်ထားတဲ့ pointer ကို တန်ဖိုးသတ်မှတ်ပေးခြင်းဟာ
// ညွှန်းထားတဲ့ address မှာရှိတဲ့ တန်ဖိုးကို ပြောင်းလဲစေပါတယ်။
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	// `&i` syntax က `i` ရဲ့ memory address ကိုပေးပါတယ်။
	// တနည်းအားဖြင့် `i` ရဲ့ pointer ပါ။
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	// Pointer တွေကို Print ထုတ်ကြည့်လို့ရပါတယ်။
	fmt.Println("pointer:", &i)
}
