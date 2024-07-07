// Go စံသတ်မှတ်ထားသော(standard) library သည် `net/http` package တွင်
// HTTP clients နှင့် servers အတွက် အထူးကောင်းမွန်သော ပံ့ပိုးမှုများပါဝင်သည်။
// ဤဥပမာတွင် ၎င်းကို အသုံးပြု၍ ရိုးရှင်းသော HTTP တောင်းဆိုမှု(requests)များ ပြုလုပ်မည်ဖြစ်သည်။
package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {

	// server တစ်ခုသို့ HTTP GET တောင်းဆိုမှုတစ်ခု ပြုလုပ်ပါ။ `http.Get` သည်
	// `http.Client` object တစ်ခုကို ဖန်တီးပြီး ၎င်း၏ `Get` method ကို ခေါ်ယူခြင်းထက်
	// အဆင်ပြေသော shortcut တစ်ခုဖြစ်သည်။ ၎င်းသည် အသုံးဝင်သော
	// default settings များပါရှိသည့် `http.DefaultClient` object ကို အသုံးပြုသည်။
	resp, err := http.Get("https://gobyexample.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// HTTP တုံ့ပြန်မှု အခြေအနေကို ပုံနှိပ်(print)ထုတ်ပါ။
	fmt.Println("Response status:", resp.Status)

	// တုံ့ပြန်မှု body ၏ ပထမဆုံး စာကြောင်း ၅ ကြောင်းကို ပုံနှိပ်(print)ထုတ်ပါ။
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
