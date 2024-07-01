// _goroutine_ ဆိုတာ ပေါ့ပါးတဲ့ execution thread တစ်ခုဖြစ်ပါတယ်။

package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	// `f(s)` ဆိုတဲ့ function ရှိတယ်ထားပါတော့ ဒီမှာ ပုံမှန်အတိုင်း
	// synchronous(တပြိုင်နက်တည်း) ခေါ်တာကို ပြထားပါတယ်။
	f("direct")

	// ဒီ function ကို goroutine အနေနဲ့ ခေါ်ချင်ရင် `go f(s)` လို့သုံးပါတယ်။
	// ဒီ goroutine အသစ်က ခေါ်လိုက်တဲ့ goroutine နဲ့ တပြိုင်နက်တည်း (concurrently) အလုပ်လုပ်ပါလိမ့်မယ်။
	go f("goroutine")

	// Anonymous function ကိုလည်း goroutine အနေနဲ့ စတင်နိုင်ပါတယ်။
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// အခု ကျွန်တော်တို့ရဲ့ function ခေါ်တာ နှစ်ခုက သီးခြား goroutine တွေမှာ asynchronous (တပြိုင်နက်တည်းမဟုတ်) အနေနဲ့ အလုပ်လုပ်နေပါပြီ။
	// သူတို့ပြီးဆုံးတာကို စောင့်နေတာပါ (ပို robust approach တွေအတွက်ကတော့ [WaitGroup](waitgroups) ကို သုံးပါ)။
	time.Sleep(time.Second)
	fmt.Println("done")
}
