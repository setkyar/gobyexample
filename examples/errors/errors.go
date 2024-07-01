// Go မှာ error တွေကို သီးခြား return value တစ်ခုအနေနဲ့ ပြန်ပေးတာက ပုံမှန်လုပ်နည်းလုပ်ဟန်ဖြစ်ပါတယ်။
// သူက Java နဲ့ Ruby လို language တွေမှာသုံးတဲ့ exception တွေနဲ့ ကွာခြားပြီး
// C မှာ တစ်ခါတစ်ရံသုံးတဲ့ ရလဒ်တစ်ခုတည်းကို error value အဖြစ်လည်း သုံးတဲ့နည်းနဲ့လည်း မတူပါဘူး။
// Go ရဲ့နည်းလမ်းက ဘယ် function တွေက error ပြန်ပေးသလဲဆိုတာ မြင်သာစေပြီး
// တခြား non-error အလုပ်တွေမှာသုံးတဲ့ go ရဲ့ construct တွေနဲ့ပဲ error တွေကို ကိုင်တွယ်လို့ရအောင် လုပ်ပေးပါတယ်။

package main

import (
	"errors"
	"fmt"
)

// convention အရ၊ error တွေဟာ နောက်ဆုံး return value ဖြစ်ပြီး
// built-in interface ဖြစ်တဲ့ `error` type ဖြစ်ပါတယ်။
func f1(arg int) (int, error) {
	if arg == 42 {

		// `errors.New` က ပေးလိုက်တဲ့ error message နဲ့အတူ
		// အခြေခံ `error` value တစ်ခုကို တည်ဆောက်ပေးပါတယ်။
		return -1, errors.New("can't work with 42")

	}

	// Error position မှာ `nil` value ရှိနေတာက
	// error မရှိဘူးဆိုတာပါ။
	return arg + 3, nil
}

// Custom type တွေကို `Error()` method ကို implement လုပ်ပီးတော့
// `error` အဖြစ် သုံးလို့ရပါတယ်။ ဒီမှာ အပေါ်က ဥပမာရဲ့ ပုံစံကွဲတစ်ခုကို ပြထားပါတယ်။
// ဒီဥပမာမှာ argument error ကို ကိုယ်စားပြုဖို့ custom type တစ်ခုကို အသုံးပြုထားပါတယ်။
type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {

		// ဒီနေရာမှာ `&argError` syntax ကို သုံးပြီး struct အသစ်တစ်ခု တည်ဆောက်ထားပါတယ်။
		// `arg` နဲ့ `prob` ဆိုတဲ့ field နှစ်ခုအတွက် တန်ဖိုးတွေ ပေးထားပါတယ်။
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {

	// အောက်က loop နှစ်ခုက ကျွန်တော်တို့ရဲ့ error-returning function နှစ်ခုစလုံးကို စမ်းသပ်ထားပါတယ်။
	// သတိထားရမှာက `if` line မှာ inline error check သုံးထားတာဟာ Go code မှာ ပုံမှန်သုံးနေကျ ပုံစံတစ်ခုဖြစ်ပါတယ်။
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	// Custom error ထဲက data ကို programmatically သုံးချင်ရင်၊
	// type assertion သုံးပြီး error ကို custom error type ရဲ့ instance အဖြစ် ရယူရပါမယ်။
	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
