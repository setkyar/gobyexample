// Go ရဲ့ *select* က သင့်ကို channel operation အများကြီးကို
// တပြိုင်နက်တည်း စောင့်ဆိုင်းခွင့်ပေးပါတယ်။ Goroutine တွေနဲ့ channel တွေကို
// select နဲ့ ပေါင်းစပ်အသုံးပြုခြင်းဟာ Go ရဲ့ စွမ်းအားကြီးမားတဲ့ feature တစ်ခုဖြစ်ပါတယ်။

package main

import (
	"fmt"
	"time"
)

func main() {

	// ဒီဥပမာမှာ ကျွန်တော်တို့က channel နှစ်ခုကို select လုပ်မှာပါ။
	c1 := make(chan string)
	c2 := make(chan string)

	// Channel တစ်ခုချင်းစီက တန်ဖိုးတစ်ခုကို အချိန်အနည်းငယ်ကြာပြီးမှ လက်ခံရရှိပါမယ်။
	// ဒါဟာ ဥပမာအားဖြင့် concurrent goroutine တွေမှာ အလုပ်လုပ်နေတဲ့
	// blocking RPC operation တွေကို simulate လုပ်တာပါ။
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	// ကျွန်တော်တို့က `select` ကို သုံးပြီး ဒီတန်ဖိုးနှစ်ခုစလုံးကို တပြိုင်နက်တည်း
	// စောင့်ဆိုင်းမှာဖြစ်ပြီး၊ တန်ဖိုးတစ်ခုချင်းစီ ရောက်လာတိုင်း print ထုတ်ပါမယ်။
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}
