// String တွေကနေ နံပါတ်တွေကို parse လုပ်တာဟာ ပရိုဂရမ်အများစုမှာ အခြေခံကျပေမယ့် အဖြစ်များတဲ့ လုပ်ငန်းတစ်ခုပါ။
// Go မှာ အဲဒါကို ဘယ်လိုလုပ်မလဲဆိုတာ ဒီမှာရှိပါတယ်။

package main

// Built-in package ဖြစ်တဲ့ `strconv` က နံပါတ် parsing ကို ပေးပါတယ်။
import (
	"fmt"
	"strconv"
)

func main() {
	// `ParseFloat` မှာ ဒီ `64` က ဘယ်နှစ် bit တိကျမှုနဲ့ parse လုပ်မလဲဆိုတာကို ပြောပြပါတယ်။
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	// `ParseInt` မှာ `0` က string ကနေ အခြေကိန်းကို မှန်းဆဖို့ ဆိုလိုပါတယ်။
	// `64` က ရလဒ်က 64 bit ထဲ ဝင်ဆံ့ရမယ်လို့ သတ်မှတ်ပါတယ်။
	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)

	// `ParseInt` က ဆယ့်ခြောက်လုံးစနစ် ဖော်မတ်နဲ့ နံပါတ်တွေကို အသိအမှတ်ပြုပါတယ်။
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	// `ParseUint` လည်း ရနိုင်ပါတယ်။
	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)

	// `Atoi` က အခြေ-10 `int` parsing အတွက် အဆင်ပြေတဲ့ function တစ်ခုပါ။
	k, _ := strconv.Atoi("135")
	fmt.Println(k)

	// Parse function တွေက မှားယွင်းတဲ့ input ပေါ်မှာ error တစ်ခု ပြန်ပေးပါတယ်။
	_, e := strconv.Atoi("wat")
	fmt.Println(e)
}
