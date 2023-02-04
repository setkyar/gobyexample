// Go မှာ strings, integers, floats, booleans အစရှိတဲ့ types တွေအမျိုးမျိုးရှိပါတယ်။
// ဒီမှာဥပမာတချို့ကိုကြည့်နိုင်ပါတယ်။

package main

import "fmt"

func main() {

	// Strings, + နှင့် နှစ်ခုအတူပေါင်းပီးဆက်နိုင်တယ်
	// Strings, စာသား
	fmt.Println("go" + "lang")

	// Integers နှင့် floats.
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	// Booleans, boolean operators အတိုင်းသင်မှန်းထားသလိုပါဘဲ
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
