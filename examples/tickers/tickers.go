// [Timer](timers) တွေက အနာဂတ်မှာ တစ်ကြိမ်တည်း တစ်ခုခုလုပ်ချင်တဲ့အခါ သုံးပါတယ် -
// _ticker_ တွေကတော့ ပုံမှန် တစ်ခုခုကို ထပ်ခါထပ်ခါလုပ်ချင်တဲ့အချိန်မှာ သုံးပါတယ်။
// ဒီမှာ ကျွန်တော်တို့ မရပ်ခိုင်းမချင်း ပုံမှန် tick လုပ်နေမယ့် ticker တစ်ခုရဲ့ ဥပမာကို ပြထားပါတယ်။

package main

import (
	"fmt"
	"time"
)

func main() {

	// Ticker တွေက timer တွေနဲ့ ဆင်တူတဲ့ နည်းလမ်းကို သုံးပါတယ်:
	// တန်ဖိုးတွေ ပို့ပေးတဲ့ channel တစ်ခုပါ။ ဒီမှာတော့ channel ပေါ်မှာ
	// `select` builtin ကိုသုံးပြီး 500ms တိုင်း ရောက်လာမယ့် တန်ဖိုးတွေကို စောင့်ပါမယ်။
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	// Ticker တွေကို timer တွေလိုပဲ ရပ်တန့်လို့ရပါတယ်။ Ticker တစ်ခုကို
	// ရပ်တန့်လိုက်ပြီဆိုရင် သူ့ရဲ့ channel ပေါ်မှာ နောက်ထပ်တန်ဖိုးတွေ လက်ခံတော့မှာ မဟုတ်ပါဘူး။
	// ကျွန်တော်တို့ရဲ့ ticker ကို 1600ms အကြာမှာ ရပ်တန့်ပါမယ်။
	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}
