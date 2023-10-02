// Go က [_anonymous functions_](https://en.wikipedia.org/wiki/Anonymous_function) တွေကို support လုပ်ပါတယ်၊
// ဒီ Anonymous (အမည်မပါသော) function တွေသုံးပီးတော့ <a href="https://en.wikipedia.org/wiki/Closure_(computer_science)"><em>closures</em></a> တွေကို
// တည်ဆောက်တယ်။ Anonymous function တွေက function တွေကိုအမည်မပါဘဲသုံးချင်တဲ့အချိန်မှာအတော်လေးအသုံးဝင်တယ်။

package main

import "fmt"

// ဒီ `intSeq` ဆိုတဲ့ function ကနောက်ထက် function တခုကို return ပြန်ထားတယ်၊
// အဲ့ function က `intSeq` ရဲ့ body ထဲမှာအမည်မပါတဲ့ function တခုသတ်မှတ်ထားတယ်။
// အဲ့မှာအမည်မပါဘဲသတ်မှတ်ထားတဲ့ function က `i` variable ကို _close over_ လုပ်ပီးတော့
// closure တခုတည်ဆောက်လိုက်တယ်။
// ဒါကဘာကိုဆိုလိုတာလဲဆိုရင် `intSeq` ကိုခေါ်ပီးတော့ Anonymous function ကို
// return ပြန်ပီးတာတောင် return ပြန်လိုက်တဲ့ Anonymous function က `intSeq` ရဲ့ `i` variable ကို access
// ရနေသေးတယ်ဆိုတာကိုပြောချင်တာပါ။

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	// ကျနော်တို့ `intSeq` ကိုခေါ်ပီးတော့ return ပြန်လိုက်တဲ့ Anonymous function ကို
	// `nextInt` ဆိုတဲ့ variable ထဲမှာထည့်သွင်းထားတယ်၊ ဒီ function ကသူ့ထဲက value ဖြစ်တဲ့
	// i ကို capture လုပ်ပီးတော့ `nextInt` ကိုခေါ်တိုင်း အဲ့ဒီ့ i တန်ဖိုးကို update လုပ်ပေးတယ်။
	nextInt := intSeq()

	// closure ရဲ့ကျိုးသတ်ရောက်မှူကိုမြင်ဖို့ `nextInt` ကိုအနည်းငယ်ခေါ်ကြည့်လိုက်ကြရအောင်
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// ဒီ function တခုထဲကိုဘဲဒီ state သတ်ရောက်နေတယ်ဆိုသေချာဖို့ရာအတွက်
	// အသစ်တခုတည်ဆောက်ပီး စမ်းကြည့်လိုက်ကြရအောင်။
	newInts := intSeq()
	fmt.Println(newInts())
}
