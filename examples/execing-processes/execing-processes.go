// ပြီးခဲ့တဲ့ဥပမာမှာ ကျွန်တော်တို့က [ပြင်ပလုပ်ငန်းစဉ်တွေကို စတင်ခြင်း (spawning external processes)](spawning-processes) ကိုကြည့်ခဲ့ပါတယ်။
// ဒါကို ကျွန်တော်တို့က လက်ရှိအလုပ်လုပ်နေတဲ့ Go process ကနေ ပြင်ပလုပ်ငန်းစဉ်တွေကို အသုံးပြုဖို့လိုအပ်တဲ့အခါ လုပ်ဆောင်ပါတယ်။
// တစ်ခါတစ်ရံမှာ ကျွန်တော်တို့က လက်ရှိ Go process ကို အခြားတစ်ခု (Go မဟုတ်တာလည်းဖြစ်နိုင်တယ်) နဲ့ လုံးဝအစားထိုးချင်တာမျိုးရှိပါတယ်။
// ဒါကိုလုပ်ဖို့အတွက် Go ရဲ့ ဂန္ထဝင်ဖြစ်တဲ့
// <a href="https://en.wikipedia.org/wiki/Exec_(operating_system)"><code>exec</code></a>
// function ကို အကောင်အထည်ဖော်ထားတာကို အသုံးပြုပါမယ်။

package main

import (
	"os"
	"os/exec"
	"syscall"
)

func main() {

	// ကျွန်တော်တို့ရဲ့ ဥပမာအတွက် `ls` ကို exec လုပ်ပါမယ်။ Go က ကျွန်တော်တို့ run ချင်တဲ့ binary ရဲ့
	// အပြည့်အစုံ path လိုအပ်ပါတယ်၊ ဒါကြောင့် `exec.LookPath` ကိုသုံးပြီး ရှာပါမယ်
	// (များသောအားဖြင့် `/bin/ls` ဖြစ်ပါလိမ့်မယ်)။
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}

	// `Exec` က argument တွေကို slice ပုံစံနဲ့ လိုအပ်ပါတယ်
	// (စာကြောင်းရှည်ကြီးတစ်ကြောင်းတည်းနဲ့ မရပါဘူး)။ `ls` အတွက် ပုံမှန်သုံးလေ့ရှိတဲ့ argument တချို့ပေးပါမယ်။
	// ပထမဆုံး argument က program နာမည်ဖြစ်သင့်တယ်ဆိုတာ သတိပြုပါ။
	args := []string{"ls", "-a", "-l", "-h"}

	// `Exec` က [environment variables](environment-variables) set လည်း လိုအပ်ပါတယ်။
	// ဒီနေရာမှာတော့ ကျွန်တော်တို့ရဲ့ လက်ရှိ environment variables တွေကိုပဲ ပေးလိုက်ပါတယ်။
	env := os.Environ()

	// ဒါက တကယ့် `syscall.Exec` ခေါ်တဲ့နေရာပါ။ ဒီ call က အောင်မြင်ရင်
	// ကျွန်တော်တို့ process ရဲ့ အလုပ်လုပ်ခြင်းဟာ ဒီနေရာမှာပဲ ပြီးဆုံးသွားပြီး `/bin/ls -a -l -h`
	// process နဲ့ အစားထိုးခံရပါလိမ့်မယ်။ အမှားတစ်ခုခုရှိရင်တော့ return value ရပါလိမ့်မယ်။
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
