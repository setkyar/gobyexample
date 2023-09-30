// [_Variadic functions_](https://en.wikipedia.org/wiki/Variadic_function)
// ဆိုတာက arguments တွေအများကြီးထက်ထက်ထည့်ပီးခေါ်နိုင်တဲ့ function ကိုပြောတာပါ။
// ဥပမာ မြင်သာအောင်ပြောရမယ်ဆိုရင် `fmt.Println` ကကျနော်တို့ပုံမှန်သုံးနေတဲ့ variadic function
// တခုဘဲဖြစ်ပါတယ်။

package main

import "fmt"

// ဒီ function မှာဆိုရင် ကြိုက်သလောက် numbers `int` တွေထည့်နိုင်တယ်။
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0

	// ဒီ function ထဲမှာဆိုရင်တော့ `nums` ရဲ့ type က `[]int` နဲ့တူပါတယ်။
	// အဲ့ဒီတော့ `len(nums)` ကိုသုံးနိုင်တယ်၊ `range` ကိုသုံးပီးတော့ iterate လုပ်နိုင်တယ်။
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {

	//  Variadic function ကိုအသုံးပြုမယ်ဆိုရင် ပုံမှန် individual
	// arguments တွေနှင့်ခေါ်ပီးအသုံးပြုနိုင်ပါတယ်။
	sum(1, 2)
	sum(1, 2, 3)

	// slice ထဲမှာ multiple args ရှိပီးသားဆိုရင်
	// အဲ့ slice တွေကိုဘဲ variadic function နှင့်
	// `func(slice...)` ဆိုပီးသုံးနိုင်ပါတယ်။
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
