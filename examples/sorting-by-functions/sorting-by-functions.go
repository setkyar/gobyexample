// တစ်ခါတလေ အက္ခရာစဉ် နံပါတ်စဉ် စသည် သဘာဝအစဉ်အလိုက် သာမက
// အခြားပုံစံမျိုးဖြင့် စဉ်ခြင်းမျိုးကိုလည်း ဆောင်ရွက်လိုပါလိမ့်မည်။ နမူနာအားဖြင့်
// ဆိုရသော် အက္ခရာစဉ်ခြင်းမျိုးမဟုတ်ပဲ အက္ခရာအလုံးအရေအတွက်အားဖြင့်
// စဉ်လိုသည်ဟု ယူဆပါစို့။ Go တွင် မိမိစဉ်လိုသည့် ပုံစံများကို အောက်ပါ
// နမူနာတွင် ကြည့်နိုင်သည်။

package main

import (
	"fmt"
	"sort"
)

// Go မှာ custom function သုံးပြီး sort လုပ်မယ်ဆိုရင် အဲဒါနဲ့
// ကိုက်တဲ့ data type တစ်ခုလိုတယ်။ ဒီမှာ `byLength` လို့ခေါ်တဲ့ custom type တစ်ခု
// ဖန်တီးထားတယ်။​ အဲဒါက တကယ်တော့ `[]string` ကို နာမည်ပြောင်းခေါ်တာမျိုး ဖြစ်တယ်။
// `[]string` ဆိုတဲ့ type ကို `byLength` ဆိုတဲ့ type အနေနဲ့ သုံးမယ်ဆိုတာမျိုး။
type byLength []string

// အပေါ်မှာ ဖန်တီးထားတဲ့ type အတွက် `sort.Interface` ရေးပေးဖို့လိုတယ်။
// `Len`, `Less`၊ နဲ့ `Swap` လို့ခေါ်တဲ့ `Interface` ၃ ခု။ အဲဒီဟာတွေ ရေးပေးထားမှ
// `sort` package ထဲမှာရှိတဲ့ generic `Sort` function ကို သုံးလို့ရမှာ။ array လိုမျိုး
// type တွေအတွက်`Len` နဲ့ `Swap` အတွက် ရေးပေးရတဲ့ကုဒ်က ဆင်တူလေ့ရှိတယ်။ sort လုပ်တဲ့
// logic ရှိတဲ့နေရာက `Less` interface ပဲ။​ အဲဒီမှာ sorting logic ကို သွားရေး
// ပေးထားရတယ်။ အခုကိစ္စမှာဆို string တိုးတာပဲဖြစ်ဖြစ် လျော့တာပဲဖြစ်ဖြစ်
// string length ကိုသုံးပြီး sort လုပ်ပေးပါဆိုပြီး ခိုင်းတာမျိုး။ အဲ့လို sort ဖြစ်ဖို့
// `len(s[i])` နဲ့ `len(s[j])` တွေကို `Less` interface မှာ သုံးထားတယ်။
func (s byLength) Len() int {
	return len(s)
}
func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

// လိုအပ်သည့် ကုဒ်များကို နေရာတကျ ရေးပြီးသည့်အခါ မူလ `fruits` slice ကို
// `byLength` အမျိုးအစားသို့ ပြောင်းလဲလိုက်ခြင်းအားဖြင့် မိမိစဉ်လိုသည့် လောဂျစ်ကို
// အကောင်အထည် ဖော်နိုင်သွားသည်။ ထို့နောက် ၎င်း `byLength` ကို စဉ်ရန်
// `sort.Sort` အသုံးပြုသည်။
func main() {
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)
}
