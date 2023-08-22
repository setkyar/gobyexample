//  Go မှာ _array_ length ဘယ်လောက်ရှိလဲဆိုတာသတ်မှတ်ပေးရတယ်။
// ပုံမှန် Go code တွေမှာ [slices](slices) တွေကိုအသုံးပြုကြတာများတယ်။
//  တချို့ special scenarios တွေမှာဆို arrays ကိုအသုံးပြုကြတယ်။

package main

import "fmt"

func main() {

	// ဒီမှာကျနော်တို့ array `a` ကိုကြေငြာထားတယ်၊ အဲ့ array ထဲမှာ `int` 5 ခုရှိတယ်။
	// array ထဲမှာပါတဲ့ elements တွေရဲ့ type နဲ့ length ကတော့ array ရဲ့ type ဖြစ်ပါတယ်။
	// array ရဲ့ default value က 0 ပါ။ array ထဲမှာဘာမှသတ်မှတ်ထားတာမရှိဘူးဆိုရင်
	// default 0 တခါထဲထည့်ပေးလိုက်မှာပါ။
	var a [5]int
	fmt.Println("emp:", a)

	// ကျနော်တို့ array ကို array index နဲ့ထောက်ပီးတော့ value ထည့်ပေးလို့ရတယ်။
	// `array[index] = value` ဒီလို syntax ကိုသုံးပြီးတော့ value ထည့်ပေးလို့ရသလို
	// `array[index]` ဒီလို syntax ကိုသုံးပြီးတော့ value ကိုရှာနိုင်ပါတယ်။
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	// builtin `len` ကတော့သိတဲ့အတိုင်း array ရဲ့ length ကိုရနိုင်ပါတယ်။
	fmt.Println("len:", len(a))

	// ဒီ syntax ကိုသုံးပီးတော့ array ကို declare/initialize
	// တကြောင်းထဲနှင့်သတ်မှတ်နိုင်ပါတယ်။
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// Array type တွေက one-dimensional၊ ဒါပေမယ့် multi-dimensional
	// data structure တွေကိုလဲအသုံးပြုနိုင်ပါတယ်။
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
