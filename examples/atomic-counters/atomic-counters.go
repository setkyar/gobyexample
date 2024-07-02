// Go မှာ state ကို စီမံခန့်ခွဲဖို့ အဓိက နည်းလမ်းကတော့ channel တွေကနေ တဆင့်
// ဆက်သွယ်မှုပါ။ ဒါကို [worker pools](worker-pools) မှာ ဥပမာအနေနဲ့ တွေ့ခဲ့ပါတယ်။
// ဒါပေမယ့် state စီမံခန့်ခွဲဖို့ တခြားနည်းလမ်းအနည်းငယ်လည်း ရှိပါသေးတယ်။ ဒီမှာတော့
// goroutine အများကြီးက တပြိုင်နက်တည်း access လုပ်တဲ့ _atomic counter_ တွေအတွက်
// `sync/atomic` package ကို သုံးတာကို ကြည့်ကြပါမယ်။

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	// ကျွန်တော်တို့ရဲ့ (အမြဲတမ်း အပေါင်းကိန်းဖြစ်တဲ့) counter ကို ကိုယ်စားပြုဖို့
	// unsigned integer တစ်ခုကို သုံးပါမယ်။
	var ops uint64

	// WaitGroup က goroutine အားလုံး အလုပ်ပြီးတဲ့အထိ စောင့်ဖို့ ကူညီပါလိမ့်မယ်။
	var wg sync.WaitGroup

	// Goroutine 50 ခုကို စတင်ပါမယ်။
	// တစ်ခုချင်းစီက counter ကို တိတိကျကျ 1000 ကြိမ်စီ တိုးပါလိမ့်မယ်။
	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func() {
			for c := 0; c < 1000; c++ {
				// Counter ကို atomic ဖြစ်အောင် တိုးဖို့ `AddUint64` ကို သုံးပါတယ်။
				// `&` syntax နဲ့ ကျွန်တော်တို့ရဲ့ `ops` counter ရဲ့ memory address ကို ပေးလိုက်ပါတယ်။
				atomic.AddUint64(&ops, 1)
			}
			wg.Done()
		}()
	}

	// Goroutine အားလုံး ပြီးဆုံးတဲ့အထိ စောင့်ပါ။
	wg.Wait()

	// အခုဆိုရင် `ops` ကို access လုပ်ဖို့ လုံခြုံပါပြီ။ ဘာလို့လဲဆိုတော့
	// တခြား goroutine တွေက သူ့ကို ရေးနေတာ မရှိတော့ဘူးဆိုတာ သိလို့ပါ။
	// Atomic တွေကို update လုပ်နေချိန်မှာ safely read လုပ်လို့ရပါတယ်။
	// `atomic.LoadUint64` လို function တွေကို သုံးပြီး လုပ်လို့ရပါတယ်။
	fmt.Println("ops:", ops)
}
