// Standard library ရဲ့ `strings` package က အသုံးဝင်တဲ့ string-related function တွေ
// အများကြီး ပေးထားပါတယ်။ ဒီမှာ ဒီ package အကြောင်း ခန့်မှန်းသိနိုင်အောင်
// ဥပမာအချို့ ပြထားပါတယ်။

package main

import (
	"fmt"
	s "strings"
)

// `fmt.Println` ကို နာမည်တိုတိုနဲ့ alias လုပ်ထားပါတယ်၊
// ဘာလို့လဲဆိုတော့ အောက်မှာ အများကြီး သုံးမှာမို့ပါ။
var p = fmt.Println

func main() {

	// ဒီမှာ `strings` package ထဲက function တွေရဲ့ နမူနာတွေ ရှိပါတယ်။
	// ဒါတွေက package ထဲက function တွေဖြစ်တဲ့အတွက် string object ပေါ်က
	// method တွေ မဟုတ်ပါဘူး။ ဒါကြောင့် စစ်ဆေးချင်တဲ့ string ကို
	// ပထမဆုံး argument အနေနဲ့ ထည့်ပေးရပါတယ်။ နောက်ထပ် function တွေကို
	// [`strings`](https://pkg.go.dev/strings) package docs မှာ တွေ့နိုင်ပါတယ်။
	p("Contains:  ", s.Contains("test", "es"))
	p("Count:     ", s.Count("test", "t"))
	p("HasPrefix: ", s.HasPrefix("test", "te"))
	p("HasSuffix: ", s.HasSuffix("test", "st"))
	p("Index:     ", s.Index("test", "e"))
	p("Join:      ", s.Join([]string{"a", "b"}, "-"))
	p("Repeat:    ", s.Repeat("a", 5))
	p("Replace:   ", s.Replace("foo", "o", "0", -1))
	p("Replace:   ", s.Replace("foo", "o", "0", 1))
	p("Split:     ", s.Split("a-b-c-d-e", "-"))
	p("ToLower:   ", s.ToLower("TEST"))
	p("ToUpper:   ", s.ToUpper("test"))
}
