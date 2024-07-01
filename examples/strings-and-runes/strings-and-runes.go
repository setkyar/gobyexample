// Go string ဆိုတာ read-only byte slice တစ်ခုဖြစ်ပါတယ်။ Go language နဲ့
// standard library က string တွေကို encoded [UTF-8](https://en.wikipedia.org/wiki/UTF-8) (containers of text) အဖြစ်ဖြစ်
// special အနေဖြင့်အသုံးပြုပါတယ်။
// တခြား language တွေမှာ string တွေက "character" တွေနဲ့ ဖွဲ့စည်းထားပါတယ်။
// Go မှာတော့ character ဆိုတဲ့ သဘောတရားကို `rune` လို့ခေါ်ပါတယ် -
// Unicode code point တစ်ခုကို ကိုယ်စားပြုတဲ့ integer တစ်ခုပါ။
// [ဒီ Go blog post](https://go.dev/blog/strings) က ဒီအကြောင်းအရာနှင့်ပတ်သတ်တာကိုသေချာ မိတ်ဆက်ထားပါတယ်။

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	// `s` က ထိုင်းဘာသာစကားဖြင့် "hello" လို့ရေးထားတဲ့ literal value ကို assign လုပ်ထားတဲ့
	// `string` တစ်ခုဖြစ်ပါတယ်။ Go string literal တွေဟာ UTF-8 နဲ့
	// encode လုပ်ထားတဲ့ စာသားတွေဖြစ်ပါတယ်။
	const s = "สวัสดี"

	// String တွေဟာ `[]byte` နဲ့ တူညီတဲ့အတွက်၊ ဒီကုဒ်က
	// သိမ်းဆည်းထားတဲ့ raw byte တွေရဲ့ အရေအတွက်ကို ပြပါလိမ့်မယ်။
	fmt.Println("Len:", len(s))

	// String ထဲကို indexing လုပ်ခြင်းဟာ တစ်ခုချင်းစီရဲ့ index မှာရှိတဲ့ raw byte တန်ဖိုးတွေကို ထုတ်ပေးပါတယ်။
	// ဒီ loop က `s` ထဲမှာပါတဲ့ code point တွေကို ဖွဲ့စည်းထားတဲ့ byte အားလုံးရဲ့ hex တန်ဖိုးတွေကို ထုတ်ပေးပါတယ်။
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	// String တစ်ခုထဲမှာ *rune* ဘယ်နှစ်ခုပါလဲဆိုတာ ရေတွက်ဖို့ `utf8` package ကို သုံးနိုင်ပါတယ်။
	// `RuneCountInString` ရဲ့ run-time က string ရဲ့ အရွယ်အစားပေါ်မှာ မူတည်ပါတယ်။
	// ဘာကြောင့်လဲဆိုတော့ UTF-8 rune တိုင်းကို တစ်ခုချင်းစီ decode လုပ်ရလို့ပါ။
	// ထိုင်းစာလုံးအချို့ဟာ UTF-8 code point တွေအများကြီးနဲ့ ဖွဲ့စည်းထားတာမို့
	// ဒီရေတွက်မှုရဲ့ ရလဒ်က အံ့ဩစရာ (surprising) ဖြစ်နိုင်ပါတယ်။
	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	// `range` loop က string တွေကို
	// `rune` တိုင်းကို သူ့ရဲ့ string ထဲက offset နဲ့အတူ decode လုပ်ပီး handle လုပ်ပါတယ်။
	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	// `utf8.DecodeRuneInString` function ကို အသုံးပြုပြီးလည်း
	// အထက်ကလိုမျိုး iteration လုပ်လို့ရပါတယ်။
	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width

		// ဒါက `rune` တန်ဖိုးတစ်ခုကို function ဆီသို့ pass လုပ်တာကို ပြသပါတယ်။
		examineRune(runeValue)
	}
}

func examineRune(r rune) {
	// Single quote နဲ့ ဝိုင်းထားတဲ့ value တွေက rune literals ဖြစ်ပါတယ်။
	// `rune` တန်ဖိုးတစ်ခုကို rune literal နဲ့ တိုက်ရိုက်နှိုင်းယှဉ်လို့ရပါတယ်။
	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}
