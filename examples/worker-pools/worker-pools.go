// ဒီဥပမာမှာ goroutine တွေနဲ့ channel တွေကို သုံးပြီး
// _worker pool_ တစ်ခုကို ဘယ်လို အကောင်အထည်ဖော်မလဲဆိုတာကို ကြည့်ကြပါမယ်။

package main

import (
	"fmt"
	"time"
)

// ဒါက worker ပါ၊ ကျွန်တော်တို့က ဒီ worker ရဲ့ instance အများကြီးကို
// တပြိုင်နက်တည်း run မှာပါ။ ဒီ worker တွေက `jobs` channel ကနေ
// အလုပ်တွေကို လက်ခံပြီး ရလဒ်တွေကို `results` channel ဆီ ပို့ပါလိမ့်မယ်။
// expensive task အလုပ်တစ်ခုကို simulate လုပ်ဖို့ task တစ်ခုအတွက်
// တစ်စက္ကန့် sleep ပါလိမ့်မယ်။
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {

	// ကျွန်တော်တို့ရဲ့ worker pool ကို သုံးဖို့ဆိုရင် သူတို့စီ အလုပ်တွေပို့ပြီး
	// သူတို့ရဲ့ ရလဒ်တွေကို စုစည်းဖို့ လိုပါတယ်။ ဒီအတွက် channel နှစ်ခု လုပ်ပါမယ်။
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// ဒါက worker သုံးခုကို စတင်ပါတယ်၊ ဒါပေမယ့် အစပိုင်းမှာ အလုပ်မရှိသေးတဲ့အတွက်
	// block ဖြစ်နေပါလိမ့်မယ်။
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// ဒီမှာ အလုပ် 5 ခု ပို့ပြီးတော့ အဲဒီ channel ကို `close` လုပ်ပါတယ်။
	// ဒါက ရှိသမျှအလုပ်အားလုံး ပို့ပြီးပြီဆိုတာကို ညွှန်ပြတာပါ။
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// နောက်ဆုံးမှာ အလုပ်အားလုံးရဲ့ ရလဒ်တွေကို စုစည်းပါတယ်။
	// ဒါက worker goroutine တွေ ပြီးဆုံးသွားပြီဆိုတာကိုလည်း သေချာစေပါတယ်။
	// Goroutine အများကြီးကို စောင့်ဆိုင်းဖို့ နောက်ထပ်နည်းလမ်းတစ်ခုကတော့
	// [WaitGroup](waitgroups) ကို သုံးတာပါ။
	for a := 1; a <= numJobs; a++ {
		<-results
	}
}
