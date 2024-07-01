// Channel တစ်ခုကို _ပိတ်လိုက်ခြင်း (Closing)_ က အဲဒီ channel ပေါ်မှာ နောက်ထပ်တန်ဖိုးတွေ
// ပို့တော့မှာ မဟုတ်ဘူးဆိုတာကို ပြောလိုက်တာပါ။ ဒါက channel ကနေ လက်ခံနေသူတွေဆီကို
// ပြီးဆုံးကြောင်း အသိပေးဖို့ အသုံးဝင်ပါတယ်။

package main

import "fmt"

// ဒီဥပမာမှာ `main()` goroutine ကနေ worker goroutine ဆီကို
// လုပ်ဆောင်ရမယ့် အလုပ်တွေကို ဆက်သွယ်ပေးပို့ဖို့ `jobs` channel ကို သုံးပါမယ်။
// Worker အတွက် နောက်ထပ်အလုပ်မရှိတော့တဲ့အခါ `jobs` channel ကို `close` လုပ်ပါမယ်။
func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	// ဒါက worker goroutine ပါ။ သူက `jobs` ကနေ `j, more := <-jobs` နဲ့
	// ထပ်ခါထပ်ခါ လက်ခံပါတယ်။ ဒီ special 2-value ပုံစံ လက်ခံမှုမှာ၊ `jobs` က
	// `close` ခံထားရပြီး channel ထဲက တန်ဖိုးအားလုံး လက်ခံပြီးသွားရင်
	// `more` တန်ဖိုးက `false` ဖြစ်သွားပါမယ်။ အလုပ်အားလုံး လုပ်ပြီးသွားတဲ့အခါ
	// `done` ပေါ်မှာ အသိပေး (notify) ဖို့ ဒီနည်းကို သုံးပါတယ်။
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	// ဒါက worker ဆီကို `jobs` channel ကနေ အလုပ် 3 ခု ပို့ပြီး
	// နောက်ဆုံးမှာ channel ကို ပိတ်လိုက်ပါတယ်။
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	// အရင်က တွေ့ခဲ့တဲ့ [synchronization](channel-synchronization) နည်းလမ်းကို
	// သုံးပြီး worker ကို စောင့်ပါတယ်။
	<-done
}
