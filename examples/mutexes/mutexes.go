// ပြီးခဲ့တဲ့ဥပမာမှာ [atomic operation](atomic-counters) တွေသုံးပြီး ရိုးရှင်းတဲ့
// counter state ကို ဘယ်လိုစီမံခန့်ခွဲသလဲဆိုတာ တွေ့ခဲ့ပါတယ်။
// ပိုရှုပ်ထွေးတဲ့ state တွေအတွက်တော့ goroutine အများကြီးကြားထဲမှာ ဒေတာကို
// လုံခြုံစွာ access လုပ်ဖို့ [_mutex_](https://en.wikipedia.org/wiki/Mutual_exclusion) ကို သုံးနိုင်ပါတယ်။

package main

import (
	"fmt"
	"sync"
)

// Container က counter တွေရဲ့ map ကို သိမ်းထားပါတယ်။ ဒီ map ကို
// goroutine အများကြီးကနေ တပြိုင်နက်တည်း update လုပ်ချင်တဲ့အတွက်
// access ကို synchronize လုပ်ဖို့ `Mutex` တစ်ခု ထည့်ထားပါတယ်။
// သတိပြုရမှာက mutex တွေကို copy မလုပ်သင့်ပါဘူး၊ ဒါကြောင့် ဒီ `struct` ကို
// တခြားနေရာတွေဆီ ပို့မယ်ဆိုရင် pointer နဲ့ပို့သင့်ပါတယ်။
type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

func (c *Container) inc(name string) {
	// `counters` ကို access မလုပ်ခင် mutex ကို lock လုပ်ပါ။
	// function အဆုံးမှာ unlock လုပ်ဖို့ [defer](defer) statement ကို သုံးထားပါတယ်။
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

func main() {
	c := Container{
		// Mutex ရဲ့ zero value က အသုံးပြုလို့ရပါတယ်၊ ဒါကြောင့် ဒီမှာ
		// သီးခြား initialization လုပ်စရာမလိုပါဘူး။
		counters: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup

	// ဒီ function က နာမည်ပေးထားတဲ့ counter ကို loop ထဲမှာ တိုးပေးပါတယ်။
	doIncrement := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.inc(name)
		}
		wg.Done()
	}

	// Goroutine အများကြီးကို တပြိုင်နက်တည်း run တာပါ။
	// သတိပြုရမှာက သူတို့အားလုံးက တူညီတဲ့ `Container` ကိုပဲ access လုပ်နေပြီး
	// နှစ်ခုကတော့ တူညီတဲ့ counter ကိုပဲ access လုပ်နေပါတယ်။
	wg.Add(3)
	go doIncrement("a", 10000)
	go doIncrement("a", 10000)
	go doIncrement("b", 10000)

	// Goroutine တွေ ပြီးဆုံးတဲ့အထိ စောင့်ပါ
	wg.Wait()
	fmt.Println(c.counters)
}
