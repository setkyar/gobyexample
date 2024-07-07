// ယခင်ဥပမာမှာ ရိုးရှင်းတဲ့ HTTP server တစ်ခုတည်ဆောက်တာကို ကြည့်ခဲ့ပြီးပါပြီ။
// HTTP server တွေဟာ context.Context ကိုသုံးပြီး cancel လုပ်တာကို
// သရုပ်ပြဖို့ အသုံးဝင်ပါတယ်။ Context က deadline တွေ၊ cancel signal တွေနဲ့
// တခြား request-scoped value တွေကို API boundary တွေနဲ့ goroutine တွေကြား
// သယ်ဆောင်ပေးပါတယ်။
package main

import (
	"fmt"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {

	// `context.Context` တစ်ခုကို `net/http` စက်ယန္တရား(machinery)က request တိုင်းအတွက်
	// ဖန်တီးပေးပြီး `Context()` method နဲ့ ရယူနိုင်ပါတယ်။
	ctx := req.Context()
	fmt.Println("server: hello handler started")
	defer fmt.Println("server: hello handler ended")

	// Client ဆီ ပြန်စာမပို့ခင် စက္ကန့်အနည်းငယ် စောင့်ပါမယ်။ ဒါဟာ server က
	// အလုပ်လုပ်နေတာကို simulate လုပ်တာလို့ ယူဆနိုင်ပါတယ်။ အလုပ်လုပ်နေစဉ်မှာ
	// context ရဲ့ `Done()` channel ကို စောင့်ကြည့်နေပြီး အလုပ်ကို cancel လုပ်ပြီး
	// အမြန်ဆုံး ပြန်သွားဖို့ signal ကို စောင့်ကြည့်နေပါတယ်။
	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "hello\n")
	case <-ctx.Done():
		// context ရဲ့ `Err()` method က `Done()` channel
		// ဘာကြောင့် ပိတ်သွားတယ်ဆိုတာကို
		// ရှင်းပြတဲ့ error တစ်ခုကို ပြန်ပေးပါတယ်။
		err := ctx.Err()
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}

func main() {

	// အရင်ကလိုပဲ handler ကို "/hello" route မှာ register လုပ်ပြီး
	// server စတင်လည်ပတ်စေပါတယ်။
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8090", nil)
}
