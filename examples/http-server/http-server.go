// `net/http` package ကို သုံးပြီး အခြေခံ HTTP server ရေးရတာ လွယ်ကူပါတယ်။
package main

import (
	"fmt"
	"net/http"
)

// `net/http` servers မှာ အခြေခံကျတဲ့ သဘောတရားက *handlers* ဖြစ်ပါတယ်။
// handler ဆိုတာ `http.Handler` interface ကို implement လုပ်ထားတဲ့ object တစ်ခုပါ။
// handler ရေးသားရန် ပုံမှန်နည်းလမ်းတစ်ခုက သင့်လျော်တဲ့ signature ရှိတဲ့ function တွေပေါ်မှာ
// `http.HandlerFunc` adapter ကို သုံးခြင်းဖြစ်ပါတယ်။
func hello(w http.ResponseWriter, req *http.Request) {

	// Handler အဖြစ် ဆောင်ရွက်တဲ့ function တွေဟာ `http.ResponseWriter` နဲ့
	// `http.Request` ကို argument အဖြစ် လက်ခံပါတယ်။ response writer ကို
	// HTTP response ဖြည့်စွက်ဖို့ သုံးပါတယ်။ ဒီနေရာမှာ ကျွန်တော်တို့ရဲ့ ရိုးရှင်းတဲ့ response က
	// "hello\n" ပဲ ဖြစ်ပါတယ်။
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

	// ဒီ handler က နည်းနည်းပိုပြီး ရှုပ်ထွေးပါတယ်။ HTTP request headers အားလုံးကို
	// ဖတ်ပြီး response body ထဲကို ပြန်ထည့်ပေးပါတယ်။
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {

	// ကျွန်တော်တို့ရဲ့ handlers တွေကို server routes တွေမှာ မှတ်ပုံတင်ဖို့
	// `http.HandleFunc` convenience function ကို သုံးပါတယ်။ ဒီ function က
	// `net/http` package ထဲက *default router* ကို set up လုပ်ပြီး
	// function တစ်ခုကို argument အဖြစ် လက်ခံပါတယ်။
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	// နောက်ဆုံးမှာ၊ port နဲ့ handler ကို `ListenAndServe` နဲ့ ခေါ်ပါတယ်။
	// `nil` က default router ကို သုံးမယ်လို့ ပြောတာပါ။
	http.ListenAndServe(":8090", nil)
}
