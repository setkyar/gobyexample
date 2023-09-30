//  _range_ က data structures တော်တော်များများရဲ့ elements တွေကို
// iterate လုပ်တဲ့နေရာမှာသုံးတာပါ။ `range` ကိုကျနော်တို့လေ့လာထားတဲ့
// data structures တွေမှာဘယ်လိုသုံးမလဲဆိုတာာအောက်မှာတချက်ကြည့်လိုက်ရအောင်။

package main

import "fmt"

func main() {
	// ဒီမှာကျနော်တို့ slice ထဲမှာရှိတဲ့နံပါတ်တွေကို `range` ကို သုံးပီး ပေါင်းပြထားပါတယ်။
	// Array တွေလဲအဲ့လိုဘဲ အသုံးပြုပြီးပေါင်းနိုင်ပါတယ်။

	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	// arrays နှင့် slices တွေကို `range` ဖြင့်အသုံးပြုရင် index နှင့် value နှစ်ခုလုံးပါပါတယ်။
	// အပေါ်ကနမူနာမှာဆိုရင် ကျနော်တို့ index မလိုတဲ့အတွက် `_` ကိုသုံးပီးတော့ ignored လုပ်ခဲ့ပါတယ်။
	// တခါတလေကျရင်တော့ indexs တွေကိုအသုံးပြုဖို့လိုအပ်ပါတယ်။
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// `range` ကို map မှာအသုံးပြုပီးတော့ key/value pairs ကို iterate လုပ်ပြတာပါ။
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// `range` ကိုသုံးပီးတော့ map ထဲမှာရှိတဲ့ keys တွေကိုဘဲ iterate လုပ်နိုင်ပါတယ်။
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// `range` ကိုသုံးပီးတော့ string ထဲမှာရှိတဲ့ unicode code points တွေကို iterate လုပ်နိုင်ပါတယ်။
	// ပထမ value ကတော့ character `rune` ရဲ့ starting byte index ပါပဲ။ ဒါတိယ value ကတော့ character `rune` ပါ။
	// အသေးစိတ်ကို [Strings and Runes](strings-and-runes)  မှာကြည့်နိုင်ပါတယ်။
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
