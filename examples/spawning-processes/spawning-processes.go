// တစ်ခါတလေမှာ ကျွန်တော်တို့ရဲ့ Go ပရိုဂရမ်တွေဟာ Go မဟုတ်တဲ့ တခြား
// လုပ်ငန်းစဉ် (processes) တွေကို လုပ်ဆောင်ဖို့ လိုအပ်ပါတယ်။

package main

import (
	"fmt"
	"io"
	"os/exec"
)

func main() {

	// ကျွန်တော်တို့က argument တွေ၊ input တွေ မလိုအပ်ပဲ stdout ပေါ်မှာ တစ်ခုခုကို
	// ရိုက်ထုတ်ပြတဲ့ ရိုးရှင်းတဲ့ command တစ်ခုနဲ့ စမယ်။ `exec.Command` helper က
	// ဒီ ပြင်ပ လုပ်ငန်းစဉ်ကို ကိုယ်စားပြုဖို့ object တစ်ခုကို ဖန်တီးပေးပါတယ်။
	dateCmd := exec.Command("date")
	// `Output` method က command ကို run ပြီး အဆုံးသတ်တဲ့အထိ စောင့်ပြီး
	// သူ့ရဲ့ standard output ကို စုဆောင်းပါတယ်။
	// အမှားမရှိဘူးဆိုရင် `dateOut` မှာ ရက်စွဲအချက်အလက်ပါတဲ့ bytes တွေ ပါဝင်မှာဖြစ်ပါတယ်။
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> date")
	fmt.Println(string(dateOut))
	// `Output` နဲ့ `Command` ရဲ့ တခြား method တွေဟာ command ကို run ဖို့
	// ပြဿနာရှိခဲ့ရင် (ဥပမာ - မှားယွင်းတဲ့လမ်းကြောင်း) `*exec.Error` ကို return ပြန်ပြီး၊
	// command က run ပေမယ့် non-zero return code နဲ့ ထွက်သွားခဲ့ရင်
	// `*exec.ExitError` ကို return ပြန်ပါတယ်။
	_, err = exec.Command("date", "-x").Output()
	if err != nil {
		switch e := err.(type) {
		case *exec.Error:
			fmt.Println("failed executing:", err)
		case *exec.ExitError:
			fmt.Println("command exit rc =", e.ExitCode())
		default:
			panic(err)
		}
	}

	// နောက်တစ်ဆင့်မှာ ပြင်ပလုပ်ငန်းစဉ်ရဲ့ `stdin` ပေါ်ကို data တွေ pipe လုပ်ပြီး
	// သူ့ရဲ့ `stdout` ကနေ ရလဒ်တွေကို စုဆောင်းတဲ့ နည်းနည်းပိုရှုပ်ထွေးတဲ့ ကိစ္စကို ကြည့်ကြမယ်။
	grepCmd := exec.Command("grep", "hello")

	// ဒီမှာတော့ input/output pipes တွေကို အတိအကျ ယူပြီး၊ လုပ်ငန်းစဉ်ကို စတင်၊
	// အထဲကို input တချို့ ရေးထည့်၊ ထွက်လာတဲ့ output ကို ဖတ်ယူပြီး နောက်ဆုံးမှာ
	// လုပ်ငန်းစဉ် ပြီးဆုံးတဲ့အထိ စောင့်ပါတယ်။
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()
	grepBytes, _ := io.ReadAll(grepOut)
	grepCmd.Wait()

	// အထက်ပါ ဥပမာမှာ အမှားစစ်ဆေးမှုတွေကို ချန်လှပ်ထားပါတယ်၊ ဒါပေမယ့်
	// သင်က ပုံမှန် `if err != nil` pattern ကို အားလုံးအတွက် သုံးနိုင်ပါတယ်။
	// ကျွန်တော်တို့က `StdoutPipe` ရလဒ်တွေကိုပဲ စုဆောင်းထားပေမယ့် သင်က
	// အတူတူပဲ `StderrPipe` ကိုလည်း စုဆောင်းနိုင်ပါတယ်။
	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	// မှတ်သားရန် - command တွေကို spawn လုပ်တဲ့အခါ command-line string တစ်ခုတည်းကို
	// ပေးနိုင်တာမဟုတ်ဘဲ၊ အတိအကျ သတ်မှတ်ထားတဲ့ command နဲ့ argument array ကို
	// ပေးဖို့ လိုအပ်ပါတယ်။ သင်က string တစ်ခုနဲ့ အပြည့်အစုံ command တစ်ခုကို
	// spawn လုပ်ချင်ရင် `bash` ရဲ့ `-c` option ကို သုံးနိုင်ပါတယ်။
	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))
}
