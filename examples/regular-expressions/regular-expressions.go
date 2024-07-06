// Go က [regular expressions](https://en.wikipedia.org/wiki/Regular_expression) အတွက် built-in ထောက်ပံ့မှုပေးထားပါတယ်။
// ဒီမှာ Go မှာ အသုံးများတဲ့ regexp-related လုပ်ငန်းတွေရဲ့ ဥပမာတချို့ ရှိပါတယ်။
package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {
	// ဒါက pattern တစ်ခုက string တစ်ခုနဲ့ ကိုက်ညီမှု ရှိမရှိ စမ်းသပ်ပါတယ်။
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	// အပေါ်မှာ ကျွန်တော်တို့က string pattern ကို တိုက်ရိုက်သုံးခဲ့ပါတယ်၊ ဒါပေမယ့်
	// တခြား regexp လုပ်ငန်းတွေအတွက် သင့်အနေနဲ့ optimize လုပ်ထားတဲ့ `Regexp` struct တစ်ခုကို
	// `Compile` လုပ်ဖို့ လိုအပ်ပါလိမ့်မယ်။
	r, _ := regexp.Compile("p([a-z]+)ch")

	// ဒီ struct တွေမှာ method အများကြီး ရနိုင်ပါတယ်။ ဒါက
	// အရင်က ကျွန်တော်တို့ မြင်ခဲ့တဲ့လိုမျိုး match test တစ်ခုပါ။
	fmt.Println(r.MatchString("peach"))

	// ဒါက regexp နဲ့ကိုက်ညီတာကို ရှာဖွေပါတယ်။
	fmt.Println(r.FindString("peach punch"))

	// ဒါကလည်း ပထမဆုံး ကိုက်ညီတာကို ရှာပါတယ်၊ ဒါပေမယ့် ကိုက်ညီတဲ့
	// စာသားအစား ကိုက်ညီတဲ့နေရာရဲ့ အစနဲ့အဆုံး index တွေကို ပြန်ပေးပါတယ်။
	fmt.Println("idx:", r.FindStringIndex("peach punch"))

	// `Submatch` မျိုးကွဲတွေက pattern တစ်ခုလုံးနဲ့ကိုက်ညီမှုအကြောင်း သတင်းအချက်အလက်ရော
	// အဲဒီကိုက်ညီမှုထဲက submatch တွေအကြောင်း သတင်းအချက်အလက်ပါ ထည့်သွင်းပါတယ်။
	// ဥပမာ ဒါက `p([a-z]+)ch` နဲ့ `([a-z]+)` နှစ်ခုလုံးအတွက် သတင်းအချက်အလက် ပြန်ပေးပါလိမ့်မယ်။
	fmt.Println(r.FindStringSubmatch("peach punch"))

	// ဒီနည်းတူပဲ ဒါက ကိုက်ညီမှုတွေနဲ့ submatch တွေရဲ့ index တွေအကြောင်း သတင်းအချက်အလက် ပြန်ပေးပါလိမ့်မယ်။
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))

	// ဒီ function တွေရဲ့ `All` မျိုးကွဲတွေက ပထမတစ်ခုတည်းမဟုတ်ဘဲ
	// input ထဲက ကိုက်ညီမှုအားလုံးကို အကျုံးဝင်ပါတယ်။ ဥပမာ
	// regexp နဲ့ကိုက်ညီတာအားလုံးကို ရှာဖို့ပါ။
	fmt.Println(r.FindAllString("peach punch pinch", -1))

	// ဒီ `All` မျိုးကွဲတွေက အပေါ်မှာ ကျွန်တော်တို့ မြင်ခဲ့တဲ့
	// တခြား function တွေအတွက်လည်း ရနိုင်ပါတယ်။
	fmt.Println("all:", r.FindAllStringSubmatchIndex(
		"peach punch pinch", -1))

	// ဒီ function တွေရဲ့ ဒုတိယ argument အနေနဲ့ အပေါင်းကိန်း integer တစ်ခု ပေးခြင်းအားဖြင့်
	// ကိုက်ညီမှုအရေအတွက်ကို ကန့်သတ်နိုင်ပါတယ်။
	fmt.Println(r.FindAllString("peach punch pinch", 2))

	// အပေါ်က ဥပမာတွေမှာ ကျွန်တော်တို့က string argument တွေ သုံးခဲ့ပြီး
	// `MatchString` လိုမျိုး နာမည်တွေ သုံးခဲ့ပါတယ်။ ကျွန်တော်တို့က
	// `[]byte` argument တွေကိုလည်း ပေးနိုင်ပြီး function နာမည်ကနေ `String` ကို ဖြုတ်နိုင်ပါတယ်။
	fmt.Println(r.Match([]byte("peach")))

	// Regular expression တွေနဲ့ global variable တွေ ဖန်တီးတဲ့အခါ
	// `Compile` ရဲ့ မျိုးကွဲဖြစ်တဲ့ `MustCompile` ကို သုံးနိုင်ပါတယ်။
	// `MustCompile` က error ပြန်မပေးဘဲ panic ဖြစ်စေပါတယ်၊ ဒါက global variable တွေအတွက်
	// ပိုပြီး လုံခြုံစိတ်ချရစေပါတယ်။
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println("regexp:", r)

	// `regexp` package ကို string တစ်ခုရဲ့ တစ်စိတ်တစ်ပိုင်းကို
	// တခြားတန်ဖိုးတွေနဲ့ အစားထိုးဖို့လည်း သုံးနိုင်ပါတယ်။
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	// `Func` မျိုးကွဲက ကိုက်ညီတဲ့စာသားကို ပေးထားတဲ့ function တစ်ခုနဲ့
	// ပြောင်းလဲခွင့်ပေးပါတယ်။
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}
