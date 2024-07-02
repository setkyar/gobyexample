// အရင်ဥပမာမှာ [mutex](mutexes) တွေကို အသုံးပြုပြီး goroutine အများကြီးကြားထဲမှာ
// shared state ကို access လုပ်တာကို synchronize လုပ်ခဲ့ပါတယ်။ နောက်ထပ်နည်းလမ်းတစ်ခုကတော့
// goroutine တွေနဲ့ channel တွေရဲ့ built-in synchronization feature တွေကို သုံးပြီး
// တူညီတဲ့ရလဒ်ကို ရယူတာပါ။ ဒီ channel-based နည်းလမ်းက Go ရဲ့ "memory ကို communicating ဖြင့်
// မျှဝေသုံးစွဲခြင်း" နဲ့ "ဒေတာတစ်ခုချင်းစီကို တိတိကျကျ goroutine တစ်ခုက ပိုင်ဆိုင်ခြင်း" ဆိုတဲ့
// အယူအဆတွေနဲ့ ကိုက်ညီပါတယ်။

package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

// ဒီဥပမာမှာ ကျွန်တော်တို့ရဲ့ state ကို goroutine တစ်ခုတည်းက ပိုင်ဆိုင်မှာပါ။
// ဒီနည်းက ဒေတာကို concurrent access လုပ်တဲ့အခါ ဘယ်တော့မှ ပျက်စီးမသွားစေပါဘူး။
// တခြား goroutine တွေက အဲဒီ state ကို ဖတ်ချင်တာ ရေးချင်တာရှိရင်
// ပိုင်ဆိုင်တဲ့ goroutine ဆီ message ပို့ပြီး အကြောင်းပြန်ချက်ကို လက်ခံရယူရပါမယ်။
// ဒီ `readOp` နဲ့ `writeOp` `struct` တွေက အဲဒီ request တွေနဲ့
// ပိုင်ဆိုင်တဲ့ goroutine က ပြန်လည်တုံ့ပြန်နိုင်မယ့် နည်းလမ်းကို encapsulate လုပ်ထားပါတယ်။
type readOp struct {
	key  int
	resp chan int
}
type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {

	// အရင်တုန်းကလိုပဲ ကျွန်တော်တို့ လုပ်ဆောင်တဲ့ operation အရေအတွက်ကို ရေတွက်ပါမယ်။
	var readOps uint64
	var writeOps uint64

	// `reads` နဲ့ `writes` channel တွေကို တခြား goroutine တွေက
	// read နဲ့ write request တွေ ပို့ဖို့ အသုံးပြုပါမယ်။
	reads := make(chan readOp)
	writes := make(chan writeOp)

	// ဒါက `state` ကို ပိုင်ဆိုင်တဲ့ goroutine ပါ။ `state` က
	// အရင်ဥပမာကလိုပဲ map တစ်ခုဖြစ်ပေမယ့် အခုတော့ stateful goroutine ထဲမှာပဲ
	// private ဖြစ်နေပါပြီ။ ဒီ goroutine က `reads` နဲ့ `writes` channel တွေကို
	// ထပ်ခါထပ်ခါ select လုပ်ပြီး request တွေ ရောက်လာတိုင်း တုံ့ပြန်ပါတယ်။
	// တုံ့ပြန်မှုကို တောင်းဆိုထားတဲ့ operation ကို အရင်လုပ်ဆောင်ပြီးမှ
	// အောင်မြင်ကြောင်း (နဲ့ `reads` ဖြစ်ရင် တောင်းဆိုထားတဲ့တန်ဖိုး) ကို
	// response channel `resp` ပေါ်မှာ ပို့ခြင်းဖြင့် ဆောင်ရွက်ပါတယ်။
	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	// ဒါက goroutine 100 ကို စတင်ပြီး `reads` channel ကနေတဆင့်
	// state-owning goroutine ဆီ read တွေ ပို့ခိုင်းတာပါ။
	// Read တိုင်းမှာ `readOp` တစ်ခု တည်ဆောက်ပြီး `reads` channel ပေါ်က ပို့ရပါတယ်။
	// ပြီးတော့မှ ပေးထားတဲ့ `resp` channel ကနေ ရလဒ်ကို လက်ခံရယူရပါတယ်။
	for r := 0; r < 100; r++ {
		go func() {
			for {
				read := readOp{
					key:  rand.Intn(5),
					resp: make(chan int)}
				reads <- read
				<-read.resp
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// ဒီမှာတော့ အလားတူနည်းလမ်းကိုသုံးပြီး write 10 ခုကို စတင်ပါတယ်။
	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool)}
				writes <- write
				<-write.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// Goroutine တွေကို တစ်စက္ကန့်လောက် အလုပ်လုပ်ခွင့်ပေးပါ။
	time.Sleep(time.Second)

	// နောက်ဆုံးမှာ operation အရေအတွက်တွေကို ဖတ်ယူပြီး report လုပ်ပါ။
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)
}
