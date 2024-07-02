// Goroutine အများကြီး ပြီးဆုံးတာကို စောင့်ဆိုင်းဖို့
// *wait group* ကို သုံးနိုင်ပါတယ်။

package main

import (
	"fmt"
	"sync"
	"time"
)

// ဒါက goroutine တိုင်းမှာ run မယ့် function ပါ။
func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)

	// expensive ဖြစ်တဲ့ အလုပ်တစ်ခုလိုမျိုးဖြစ်ဖို့ simulate လုပ်ပီး slepp လုပ်ခိုင်းထားပါတယ်။
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {

	// ဒီ WaitGroup ကို ဒီမှာ စတင်လိုက်တဲ့ goroutine အားလုံး
	// ပြီးဆုံးတာကို စောင့်ဆိုင်းဖို့ သုံးပါတယ်။ မှတ်ချက်: WaitGroup ကို
	// function တွေဆီ တိုက်ရိုက်ပေးပို့မယ်ဆိုရင် *pointer နဲ့* ပို့သင့်ပါတယ်။
	var wg sync.WaitGroup

	// Goroutine အများကြီးကို စတင်ပြီး WaitGroup ရဲ့
	// counter ကို တစ်ခုချင်းစီ တိုးလိုက်ပါမယ်။
	for i := 1; i <= 5; i++ {
		wg.Add(1)

		// Goroutine closure တိုင်းမှာ တူညီတဲ့ `i` တန်ဖိုးကို ပြန်သုံးတာကို ရှောင်ပါ။
		// ပိုမိုအသေးစိတ်သိချင်ရင် [FAQ](https://golang.org/doc/faq#closures_and_goroutines) ကို ကြည့်ပါ။
		i := i

		// Worker ခေါ်တာကို closure တစ်ခုနဲ့ wrap ထားဖြင့် ဒီ closure က
		// ဒီ worker ပြီးဆုံးကြောင်းကို WaitGroup ကို အသိပေးပါတယ်။ ဒီနည်းနဲ့
		// worker ကိုယ်တိုင်က သူ့ကို run တဲ့ concurrency primitive တွေအကြောင်း
		// သိစရာမလိုတော့ပါဘူး။
		go func() {
			defer wg.Done()
			worker(i)
		}()
	}

	// WaitGroup counter က 0 ပြန်ဖြစ်တဲ့အထိ block လုပ်ထားပါတယ်။
	// Worker အားလုံးက သူတို့ပြီးဆုံးကြောင်း အသိပေးတဲ့အထိစောင့်မှာပါ။
	wg.Wait()

	// မှတ်ချက်: ဒီနည်းလမ်းက worker တွေကနေ error တွေကို ပြန်ပို့ဖို့
	// straightforward မဖြစ်ပါဘူး။ ပိုပြီးအဆင့်မြင့်တဲ့ use case တွေအတွက်ဆိုရင်
	// [errgroup package](https://pkg.go.dev/golang.org/x/sync/errgroup) ကို
	// consider လုပ်ပါ။
}
