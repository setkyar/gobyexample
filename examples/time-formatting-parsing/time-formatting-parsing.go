// Go က pattern-based layout တွေကို သုံးပြီး အချိန် format လုပ်ခြင်းနဲ့ parse လုပ်ခြင်းကို ထောက်ပံ့ပေးပါတယ်။

package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	// ဒါက သက်ဆိုင်ရာ layout constant ကို သုံးပြီး အချိန်တစ်ခုကို RFC3339 အရ
	// format လုပ်တဲ့ အခြေခံဥပမာတစ်ခုပါ။
	t := time.Now()
	p(t.Format(time.RFC3339))

	// အချိန် parse လုပ်တာက `Format` နဲ့ တူညီတဲ့ layout တန်ဖိုးတွေကို သုံးပါတယ်။
	t1, e := time.Parse(
		time.RFC3339,
		"2012-11-01T22:08:41+00:00")
	p(t1)

	// `Format` နဲ့ `Parse` က ဥပမာပေါ် အခြေခံတဲ့ layout တွေ သုံးပါတယ်။ များသောအားဖြင့်
	// ဒီ layout တွေအတွက် `time` ကနေ constant တစ်ခုကို သုံးပါမယ်၊ ဒါပေမယ့်
	// ကိုယ်ပိုင် custom layout တွေလည်း ပေးနိုင်ပါတယ်။ Layout တွေက
	// ပေးထားတဲ့ အချိန်/string တစ်ခုကို format/parse လုပ်ရမယ့်ပုံစံကို ပြဖို့
	// `Mon Jan 2 15:04:05 MST 2006` ဆိုတဲ့ ကိုးကားအချိန်ကို သုံးရပါမယ်။
	// ဥပမာအချိန်က တိတိကျကျ ပြထားသလို ဖြစ်ရပါမယ်: နှစ်က 2006၊
	// နာရီက 15၊ အပတ်ရဲ့ရက်က တနင်္လာ၊ စသဖြင့်ပါ။
	p(t.Format("3:04PM"))
	p(t.Format("Mon Jan _2 15:04:05 2006"))
	p(t.Format("2006-01-02T15:04:05.999999-07:00"))
	form := "3 04 PM"
	t2, e := time.Parse(form, "8 41 PM")
	p(t2)

	// သီးသန့် ဂဏန်းဖော်ပြမှုတွေအတွက် အချိန်တန်ဖိုးထဲက ထုတ်ယူထားတဲ့
	// အစိတ်အပိုင်းတွေနဲ့ စံပြ string formatting ကိုလည်း သုံးနိုင်ပါတယ်။
	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	// `Parse` က ပုံစံမကျတဲ့ input ပေါ်မှာ parse လုပ်တဲ့ ပြဿနာကို
	// ရှင်းပြတဲ့ error တစ်ခု ပြန်ပေးပါလိမ့်မယ်။
	ansic := "Mon Jan _2 15:04:05 2006"
	_, e = time.Parse(ansic, "8:41PM")
	p(e)
}
