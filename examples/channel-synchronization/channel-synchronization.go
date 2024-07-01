// ကျွန်တော်တို့က channel တွေကို goroutine တွေကြား execution ကို
// synchronize (ချိန်ကိုက်) လုပ်ဖို့ သုံးနိုင်ပါတယ်။ ဒီဥပမာမှာ goroutine တစ်ခု ပြီးဆုံးတာကို
// စောင့်ဆိုင်းဖို့ blocking receive ကို သုံးပြထားပါတယ်။
// Goroutine အများကြီး ပြီးဆုံးတာကို စောင့်မယ်ဆိုရင်တော့
// [WaitGroup](waitgroups) သုံးတာက ပိုသင့်တော်နိုင်ပါတယ်။

package main

import (
	"fmt"
	"time"
)

// ဒါက goroutine အနေနဲ့ run မယ့် function ပါ။ `done` channel က
// ဒီ function ရဲ့ အလုပ်ပြီးပြီဆိုတာကို တခြား goroutine ကို အသိပေးဖို့ သုံးမှာပါ။
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	// အလုပ်ပြီးပြီဆိုတာ အသိပေးဖို့ value တခု ပို့လိုက်ပါတယ်။
	done <- true
}

func main() {

	// Worker goroutine တစ်ခု စတင်ပါတယ်၊ သူ့ကို အသိပေး (notify) ဖို့ channel ကိုလည်း ပေးလိုက်ပါတယ်။
	done := make(chan bool, 1)
	go worker(done)

	// Worker က channel ပေါ်မှာ အသိပေး (nofity) တဲ့အထိ block လုပ်ထားပါတယ်။
	<-done
}
