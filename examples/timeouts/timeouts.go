// *Timeout* တွေဟာ ပြင်ပ resource တွေနဲ့ ချိတ်ဆက်တဲ့ ပရိုဂရမ်တွေ သို့မဟုတ်
// အလုပ်လုပ်ချိန်ကို ကန့်သတ်ဖို့လိုအပ်တဲ့ ပရိုဂရမ်တွေအတွက် အရေးကြီးပါတယ်။
// Go မှာ timeout တွေကို implement လုပ်ဖို့ channel တွေနဲ့ `select` ကြောင့်
// လွယ်ကူပြီး သပ်ရပ်ပါတယ်။

package main

import (
	"fmt"
	"time"
)

func main() {

	// ဥပမာအနေနဲ့၊ ပြင်ပ function call တစ်ခုက သူ့ရဲ့ရလဒ်ကို 2 စက္ကန့်ကြာပြီးမှ
	// `c1` ဆိုတဲ့ channel ပေါ်မှာ ပြန်ပေးတယ်ဆိုပါစို့။ သတိပြုရမှာက ဒီ channel က
	// buffered ဖြစ်ပါတယ်၊ ဒါကြောင့် goroutine ထဲက send လုပ်တာက nonblocking ဖြစ်ပါတယ်။
	// ဒါဟာ channel ကို ဘယ်တော့မှ မဖတ်ဖြစ်တဲ့အခါ goroutine leak တွေကို
	// ကာကွယ်ဖို့ အသုံးများတဲ့ pattern တစ်ခုဖြစ်ပါတယ်။
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	// ဒီမှာတော့ timeout ကို implement လုပ်တဲ့ `select` ပါ။
	// `res := <-c1` က ရလဒ်ကို စောင့်ပါတယ်၊ `<-time.After` ကတော့
	// 1 စက္ကန့် timeout ကြာပြီးမှ ပို့မယ့်တန်ဖိုးကို စောင့်ပါတယ်။
	// `select` က ပထမဆုံး အဆင်သင့်ဖြစ်တဲ့ receive ကို ရွေးချယ်တဲ့အတွက်
	// operation က ခွင့်ပြုထားတဲ့ 1 စက္ကန့်ထက် ပိုကြာရင် timeout case ကို ယူမှာဖြစ်ပါတယ်။
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	// 3 စက္ကန့်ဆိုတဲ့ ပိုကြာတဲ့ timeout ကို ခွင့်ပြုရင်တော့
	// `c2` ကနေ လက်ခံတာ အောင်မြင်ပြီး ရလဒ်ကို print ထုတ်ပါလိမ့်မယ်။
	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}
