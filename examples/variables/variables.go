// Go မှာ variable တွေကို ကြိုတင်ကြေငြာပီးတော့ compiler ကအသုံးပြုပါတယ်။
// ဥပမာ function calls တွေရဲ့ type တွေမှန်ရဲ့လားဆိုတာကိုလဲစစ်ဆေးပါတယ်။

package main

import "fmt"

func main() {
	// variable ၁ခုသို့ ၁ခုအထက်ကို `var` နှင့်ကြေငြာနိုင်ပါတယ်။
	var a = "initial"
	fmt.Println(a)

	// variable ၁ခုနှင့်အထက်ကိုလဲ တကြိမ်ထဲသုံးပီးကြငြာနိုင်ပါတယ်။
	var b, c int = 1, 2
	fmt.Println(b, c)

	// Go will infer the type of initialized variables.
	var d = true
	fmt.Println(d)

	// Variables declared without a corresponding
	// initialization are _zero-valued_. For example, the
	// zero value for an `int` is `0`.
	var e int
	fmt.Println(e)

	// The `:=` syntax is shorthand for declaring and
	// initializing a variable, e.g. for
	// `var f string = "apple"` in this case.
	// This syntax is only available inside functions.
	f := "apple"
	fmt.Println(f)
}
