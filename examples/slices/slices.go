// _Slices_ data type က Go မှာတော်တော်အသုံးဝင်တယ်။
// array ထက်တောင်ပိုပီးအသုံးဝင်တယ်။

package main

import "fmt"

func main() {

	// arrays နဲ့မတူတာက slices မှာ elements တွေရဲ့ type တွေဘဲသတ်မှတ်ပေးစရာလိုတယ်။
	// ဘယ်လောက် size ဘာညာသတ်မှတ်စရာမလိုဘူး။
	// အစကထဲကဘယ်လောက် size ရှိမယ်ဘာညာသိစရာမလိုဘူး။
	// ဒီ code မှာဆိုရင် buildin `make` ကိုသုံးပီးတော့ length 3 ခုပါတဲ့ slice တခုဖန်တီးလိုက်တယ်။
	// (default အနေနဲ့ zero value သတ်မှတ်သွားပါလိမ့်မယ်)
	s := make([]string, 3)
	fmt.Println("emp:", s)

	// We can set and get just like with arrays.
	// set/get တွေကတော့ array မှာလိုပါဘဲ
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	// `len` ကတော့ သိတဲ့အတိုင်း slice ရဲ့ length ကိုပြန်ပေးပါလိမ့်မယ်။
	fmt.Println("len:", len(s))

	// slices က array ထက်ပိုအသုံးဝင်တဲ့ baisc operations တွေပါတယ်။
	// နမူနာပြောရမယ်ဆိုရင် `append` ပေါ့။ `append` မှာတခုသတိထားရမှာက
	// return value ကို accept လုပ်ရမယ်၊ ဘာလို့လဲဆိုရင်တော့ slice value အသစ်ပြန်ရချင်ရမှာမလို့ပါ။
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// slice တခုကနေပီးတော့ နောက်တခုကို copy ကူးလို့ရပါတယ်။
	// အခုဒီကုဒ်မှာဆိုရင် `make` ကိုသုံးပြီးတော့ s နဲ့ len တူတဲ့ slice တခုကိုတည်ဆောက်လိုက်ပါတယ်
	//ပီးတော့ `s` ကနေပီးတော့ `c` ထဲကို copy ကူးလိုက်တာပါ။
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// Slices မှာ "slice" operator ကိုဒီလို `slice[low:high]` syntax နှင့်သုံးပီးတော့
	// ကိုယ်လိုချင်တဲ့အပိုင်းကိုဘဲဖြတ်ပီးသုံးနိုင်ပါတယ်။ ဒီဥပမာမှာဆိုရင်ကျနော်တို့က s slice ရဲ့အစ
	// `s[2]` ကနေပီးတော့ `s[4]` ထိ item တွေကိုဘဲဖြတ်ပီးသုံးနိုင်ပါတယ်။
	l := s[2:5]
	fmt.Println("sl1:", l)

	// ဒီဥမာကတော့ အစကနေပီးတော့ `s[5]` မတိုင်ခင်ထိ ဖြတ်ပီးသုံးပြထားတာပါ။
	l = s[:5]
	fmt.Println("sl2:", l)

	// And this slices up from (and including) `s[2]`.
	l = s[2:]
	fmt.Println("sl3:", l)

	// slice ကို line တခုထဲနှင့် initialize လုပ်ပီးတော့လဲအသုံးပြုနိုင်တယ်။
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// Slices တွေကို slice ထဲထည့်ပီး multi-dimensional data structure အဖြစ်
	// အသုံးပြုနိုင်တယ်။ ပုံမှန် multi-dimensional array တွေနဲ့မတူတာကသူက
	// fixed length ဖြစ်မနေဘဲ အမျိုးမျိုးပြောင်းလဲနေနိုင်ပါတယ်။
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
