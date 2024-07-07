// တစ်ခါတစ်ရံတွင် ကျွန်ုပ်တို့၏ Go ပရိုဂရမ်များသည်
// Unix signals များကို အသိဉာဏ်ရှိစွာ ကိုင်တွယ်စေလိုပါသည်။
// ဥပမာအားဖြင့်၊ server တစ်ခုသည် SIGTERM လက်ခံရရှိသောအခါ
// သင့်တော်စွာ ပိတ်သိမ်းစေလိုခြင်း၊ သို့မဟုတ် command-line
// tool တစ်ခုသည် SIGINT လက်ခံရရှိသောအခါ input
// ပြုလုပ်ခြင်းကို ရပ်တန့်စေလိုခြင်းတို့ ဖြစ်နိုင်ပါသည်။
// အောက်ပါသည် Go တွင် channels များဖြင့်
// signals များကို ကိုင်တွယ်နည်း ဖြစ်ပါသည်။

package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	// Go signal notification သည် os.Signal တန်ဖိုးများကို
	// channel တစ်ခုပေါ်သို့ ပေးပို့ခြင်းဖြင့် အလုပ်လုပ်ပါသည်။
	// ဤ notifications များကို လက်ခံရန် channel တစ်ခု ဖန်တီးပါမည်။
	// ဤ channel သည် buffered ဖြစ်သင့်ကြောင်း သတိပြုပါ။
	sigs := make(chan os.Signal, 1)

	// signal.Notify သည် သတ်မှတ်ထားသော signals
	// များ၏ notifications များကို လက်ခံရန်
	// ပေးထားသော channel ကို မှတ်ပုံတင်ပါသည်။
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// ကျွန်ုပ်တို့သည် main function ထဲတွင် sigs မှ လက်ခံနိုင်ပါသည်၊
	// သို့သော် ဤလုပ်ဆောင်ချက်ကို သီးခြား goroutine
	// တစ်ခုတွင် မည်ကဲ့သို့ ပြုလုပ်နိုင်ကြောင်း ကြည့်ကြပါစို့၊
	// ပိုမို realistic ဖြစ်သော graceful shutdown
	// ဖြစ်စဉ်ကို သရုပ်ပြရန် ဖြစ်ပါသည်။
	done := make(chan bool, 1)

	go func() {
		// ဤ goroutine သည် signals အတွက် blocking receive ကို executes လုပ်ပါသည်။
		// ၎င်းသည် signal တစ်ခုရရှိသောအခါ ထို signal ကို ပုံနှိပ် (print) ဖော်ပြပြီး
		// ပရိုဂရမ်အား ပြီးဆုံးနိုင်ကြောင်း အသိပေးပါလိမ့်မည်။
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	// ပရိုဂရမ်သည် မျှော်လင့်ထားသော signal ရရှိသည်အထိ
	// ဤနေရာတွင် စောင့်ဆိုင်းနေပါမည်
	// (အထက်ပါ goroutine မှ done ပေါ်တွင် တန်ဖိုးတစ်ခု ပေးပို့ခြင်းဖြင့် ညွှန်ပြထားသည့်အတိုင်း)
	// ထို့နောက် ထွက်ခွာပါမည်။
	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}
