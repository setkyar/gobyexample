// ကျွန်တော်တို့က Go code ကို အနာဂတ်မှာ တစ်ချိန်ချိန်မှာ သို့မဟုတ်
// သတ်မှတ်ထားတဲ့ အချိန်ကြားတွေမှာ ထပ်ခါထပ်ခါ run ချင်တာ မကြာခဏ ရှိတတ်ပါတယ်။
// Go ရဲ့ built-in _timer_ နဲ့ _ticker_ feature တွေက ဒီအလုပ်နှစ်ခုစလုံးကို
// လွယ်ကူစေပါတယ်။ ပထမဆုံး timer ကိုကြည့်ပြီး
// နောက်မှ [ticker](tickers) ကို ကြည့်ပါမယ်။

package main

import (
	"fmt"
	"time"
)

func main() {

	// Timer တွေက အနာဂတ်မှာ ဖြစ်မယ့် event တစ်ခုကို ကိုယ်စားပြုပါတယ်။
	// သင်က timer ကို ဘယ်လောက်ကြာကြာ စောင့်ချင်လဲဆိုတာ ပြောပြပြီး
	// အဲဒီအချိန်ရောက်ရင် အသိပေးမယ့် channel တစ်ခုကို ရပါတယ်။
	// ဒီ timer က 2 စက္ကန့် စောင့်ပါမယ်။
	timer1 := time.NewTimer(2 * time.Second)

	// `<-timer1.C` က timer ရဲ့ channel `C` ပေါ်မှာ block လုပ်ထားပြီး
	// timer fired ဖြစ်ကြောင်း ညွှန်ပြတဲ့ တန်ဖိုးတစ်ခု ပို့လာတဲ့အထိ စောင့်ပါတယ်။
	<-timer1.C
	fmt.Println("Timer 1 fired")

	// အကယ်၍ စောင့်ဆိုင်းချင်ရုံသက်သက်ဆိုရင် `time.Sleep` ကိုလည်း သုံးနိုင်ပါတယ်။
	// Timer သုံးတာရဲ့ အကျိုးကျေးဇူးတစ်ခုကတော့ timer မ fired ခင် cancel လုပ်လို့ရတာပါ။
	// အောက်မှာ ဥပမာတစ်ခု ပြထားပါတယ်။
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	// `timer2` က တကယ်လို့ အလုပ်လုပ်မယ်ဆိုရင် လုပ်နိုင်အောင် လုံလောက်တဲ့အချိန်ပေးထားပါတယ်။
	// ဒါပေမယ့် တကယ်တမ်း ရပ်တန့်သွားပြီ (stop) ဆိုတာကိုပြတာပါ။
	time.Sleep(2 * time.Second)
}
