// ပရိုဂရမ်တွေမှာ အဖြစ်များတဲ့ လိုအပ်ချက်တစ်ခုကတော့ [Unix epoch](https://en.wikipedia.org/wiki/Unix_time) ကနေစပြီး
// ကြာခဲ့တဲ့ စက္ကန့်၊ မီလီစက္ကန့်၊ သို့မဟုတ် နာနိုစက္ကန့် အရေအတွက်ကို ရယူဖို့ပါ။
// Go မှာ ဒါကို ဘယ်လိုလုပ်ရမလဲဆိုတာ ဒီမှာရှိပါတယ်။

package main

import (
	"fmt"
	"time"
)

func main() {

	// Unix epoch ကနေစပြီး ကြာခဲ့တဲ့အချိန်ကို စက္ကန့်၊ မီလီစက္ကန့်၊ သို့မဟုတ်
	// နာနိုစက္ကန့်အလိုက် ရယူဖို့ `time.Now` ကို `Unix`၊ `UnixMilli` သို့မဟုတ် `UnixNano` နဲ့
	// တွဲသုံးပါတယ်။
	now := time.Now()
	fmt.Println(now)

	fmt.Println(now.Unix())
	fmt.Println(now.UnixMilli())
	fmt.Println(now.UnixNano())

	// သင့်အနေနဲ့ epoch ကနေစပြီး ကြာခဲ့တဲ့ စက္ကန့် သို့မဟုတ် နာနိုစက္ကန့် integer တွေကို
	// သက်ဆိုင်ရာ `time` အဖြစ် ပြောင်းလဲနိုင်ပါတယ်။
	fmt.Println(time.Unix(now.Unix(), 0))
	fmt.Println(time.Unix(0, now.UnixNano()))
}
