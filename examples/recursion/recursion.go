// Go က <a href="https://en.wikipedia.org/wiki/Recursion_(computer_science)"><em>recursive functions</em></a> ကို support လုပ်ပါတယ်။
// Recursive function ကိုနမူနာပြနေကြ factorial နမူနာကိုတချက်ကြည့်လိုက်ကြရအောင်။

package main

import "fmt"

// ဒီ `fact` function က base case `fact(0)` ကို မရောက်မခြင်းသူ့ကိုယ်သူ ပြန်ခေါ်နေပါလိမ့်မယ်။
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(7))

	// Closures လဲ recursive ဖြစ်လို့ရတယ်၊ ဒါပေမယ့် closure function ကို defined မလုပ်ခင်
	// type ကို `var` ဆိုပီးသီးသန့်ကြေငြာပေးရမယ်။
	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}

		//  `fib` ကိုအထက်မှာကြေငြာပီးသားဖြစ်တဲ့အတွက် Go က `fib` ကိုခေါ်တဲ့အချိန်မှာ
		//  ဘယ် funtion ကိုခေါ်ရမလဲဆိုတာကိုသိတယ်။
		return fib(n-1) + fib(n-2)
	}

	fmt.Println(fib(7))
}
