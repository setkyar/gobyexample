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

	// In addition to these basic operations, slices
	// support several more that make them richer than
	// arrays. One is the builtin `append`, which
	// returns a slice containing one or more new values.
	// Note that we need to accept a return value from
	// `append` as we may get a new slice value.
	// slices က array ထက်ပိုအသုံးဝင်တဲ့ baisc operations တွေပါတယ်
	//  `append` ဆိုရင် sli
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// Slices can also be `copy`'d. Here we create an
	// empty slice `c` of the same length as `s` and copy
	// into `c` from `s`.
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// Slices support a "slice" operator with the syntax
	// `slice[low:high]`. For example, this gets a slice
	// of the elements `s[2]`, `s[3]`, and `s[4]`.
	l := s[2:5]
	fmt.Println("sl1:", l)

	// This slices up to (but excluding) `s[5]`.
	l = s[:5]
	fmt.Println("sl2:", l)

	// And this slices up from (and including) `s[2]`.
	l = s[2:]
	fmt.Println("sl3:", l)

	// We can declare and initialize a variable for slice
	// in a single line as well.
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// Slices can be composed into multi-dimensional data
	// structures. The length of the inner slices can
	// vary, unlike with multi-dimensional arrays.
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
