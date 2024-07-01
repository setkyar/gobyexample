// Channel တွေပေါ်မှာ အခြေခံ ပို့တာ(send)နဲ့ လက်ခံတာ (receive) တွေက blocking ဖြစ်ပါတယ်။
// ဒါပေမယ့် `select` နဲ့ `default` clause ကိုသုံးပြီး `non-blocking` ပို့တာ၊
// လက်ခံတာ၊ ဒါ့အပြင် non-blocking multi-way `select` တွေကိုတောင်
// implement လုပ်နိုင်ပါတယ်။

package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	// ဒါက non-blocking လက်ခံမှုပါ။ `messages` မှာ တန်ဖိုးတစ်ခု အဆင်သင့်ရှိနေရင်
	// `select` က အဲဒီတန်ဖိုးနဲ့အတူ `<-messages` `case` ကို ယူပါလိမ့်မယ်။
	// မရှိဘူးဆိုရင်တော့ ချက်ချင်း `default` case ကို ယူပါလိမ့်မယ်။
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	// Non-blocking ပို့မှုကလည်း အပေါ်ကလိုဘဲ အလုပ်လုပ်ပါတယ်။ ဒီမှာ `msg` ကို
	// `messages` channel ဆီ ပို့လို့မရပါဘူး၊ ဘာလို့လဲဆိုတော့
	// channel မှာ buffer မရှိသလို လက်ခံမယ့်သူလည်း မရှိလို့ပါ။
	// ဒါကြောင့် `default` case ကို ရွေးချယ်ပါတယ်။
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	// `default` clause ရဲ့အပေါ်မှာ `case` အများကြီးသုံးပြီး
	// multi-way non-blocking select ကို implement လုပ်နိုင်ပါတယ်။
	// ဒီမှာ `messages` နဲ့ `signals` နှစ်ခုလုံးကနေ non-blocking လက်ခံဖို့ ကြိုးစားပါတယ်။
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
