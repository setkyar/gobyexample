// [Environment variables](https://en.wikipedia.org/wiki/Environment_variable)
// သည် [Unix ပရိုဂရမ်များသို့ configuration အချက်အလက်များ ပေးပို့ရန်အတွက်](https://www.12factor.net/config)
// universal နည်းလမ်းတစ်ခုဖြစ်သည်။
// Environment variables များကို မည်သို့ သတ်မှတ်ခြင်း၊ ရယူခြင်းနှင့် စာရင်းပြုစုခြင်း(list)တို့ကို ကြည့်ကြပါစို့။

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	// key/value တွဲတစ်ခုကို သတ်မှတ်ရန် `os.Setenv` ကို အသုံးပြုပါ။
	// key တစ်ခုအတွက် တန်ဖိုးကို ရယူရန် `os.Getenv` ကို အသုံးပြုပါ။
	// အကယ်၍ key သည် environment တွင် မရှိပါက ဤ function သည်
	// စာကြောင်းအလွတ်တစ်ခုကို ပြန်ပေးပါလိမ့်မည်။
	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	// environment ရှိ key/value တွဲအားလုံးကို စာရင်းပြုစု(list)ရန် `os.Environ` ကို အသုံးပြုပါ။
	// ဤ function သည် `KEY=value` ပုံစံဖြင့် စာကြောင်းများ၏ slice တစ်ခုကို ပြန်ပေးပါသည်။
	// သင်သည် ၎င်းတို့ကို `strings.SplitN` ဖြင့် key နှင့် value ကို ရရှိနိုင်ပါသည်။
	// ဤနေရာတွင် ကျွန်ုပ်တို့သည် key အားလုံးကို ပုံနှိပ်ဖော်ပြထားပါသည်။
	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0])
	}
}
