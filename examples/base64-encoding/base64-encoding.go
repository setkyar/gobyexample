// Go သည် [base64 encoding/decoding](https://en.wikipedia.org/wiki/Base64) အတွက် built-in support ပေးပါသည်။

package main

// ဒီ syntax သည် `encoding/base64` package ကို ပုံမှန် `base64` အစား `b64` နာမည်နဲ့ import လုပ်ပါတယ်။
// ဒီလိုလုပ်ခြင်းဖြင့် အောက်မှာ နေရာ (space) ချွေတာနိုင်ပါတယ်။
import (
	b64 "encoding/base64"
	"fmt"
)

func main() {
	// ဒီမှာ encode/decode လုပ်မယ့် `string` ပါ။
	data := "abc123!?$*&()'-=@~"

	// Go သည် စံပြနှင့် URL-compatible base64 နှစ်မျိုးလုံးကို ထောက်ပံ့ပါတယ်။
	// ဒီမှာ စံပြ encoder ကိုသုံးပြီး encode လုပ်ပုံပြထားပါတယ်။
	// encoder သည် `[]byte` လိုအပ်သောကြောင့် ကျွန်ုပ်တို့၏ `string` ကို ထိုအမျိုးအစားသို့ ပြောင်းလဲပါသည်။
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)

	// Decoding သည် error ပြန်နိုင်ပါတယ်။ input က ကောင်းမွန်စွာ ဖွဲ့စည်းထားကြောင်း သင်မသိသေးပါက
	// စစ်ဆေးနိုင်ပါသည်။
	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))
	fmt.Println()

	// ဒီအပိုင်းသည် URL-compatible base64 format ကိုသုံးပြီး encode/decode လုပ်ပါတယ်။
	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))
}
