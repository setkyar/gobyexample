// Go က character, string, boolean, နှင့် numeric values
// အမျိုးမျိုးအတွက် _constants_ ကိုအသုံးပြုနိုင်ပါတယ်။

package main

import (
	"fmt"
	"math"
)

// `const` ကိုသုံးပီး constant value ကိုဖန်တီးနိုင်ပါတယ်။
const s string = "constant"

func main() {
	fmt.Println(s)

	// `const` statement က `var` statement ကြေငြာလို့ရတဲ့
	// ဘယ်နေရာမှာမဆို ကြောငြာနိုင်ပါတယ်။
	const n = 500000000

	// Constant အသုံးအနှုန်းက တိကျတဲ့ ဂဏန်းသင်္ချာ တွက်ချက်နိုင်ပါတယ်။
	const d = 3e20 / n
	fmt.Println(d)

	// ဂဏန်း constant က သေချာ type မပြောင်းထားဘူးဆိုလျှင်
	// type မရှိပါဘူး။
	fmt.Println(int64(d))

	// variable assignment သို့ function call မှာ
	// ဂဏန်းသင်္ချာတွေကို လိုအပ်လျှင် လိုအပ်သလို
	// အမျိုးအစားပြောင်းနိုင်ပါတယ်။
	// ဥပမာအနေနဲ့ `math.Sin` က `float64` ကိုလိုအပ်ပါတယ်။
	fmt.Println(math.Sin(n))
}
