// `//go:embed` သည် [compiler directive](https://pkg.go.dev/cmd/compile#hdr-Compiler_Directives) တစ်ခုဖြစ်ပြီး
// ပရိုဂရမ်များကို build လုပ်သည့်အခါ မည်သည့်ဖိုင်နှင့် ဖိုဒါများကိုမဆို Go binary ထဲသို့ ထည့်သွင်းခွင့်ပြုသည်။
// embed directive အကြောင်း ပိုမိုလေ့လာလိုပါက [ဤနေရာ](https://pkg.go.dev/embed) တွင်ဖတ်ရှုနိုင်သည်။
package main

// `embed` package ကို import လုပ်ပါ။ အကယ်၍ သင်သည် ဤ package မှ exported identifiers များကို
// အသုံးမပြုပါက `_ "embed"` ဖြင့် blank import လုပ်နိုင်သည်။
import (
	"embed"
)

// `embed` directives သည် Go source file ပါဝင်သည့် directory နှင့် ဆက်စပ်သော paths များကို လက်ခံသည်။
// ဤ directive သည် ဖိုင်၏အကြောင်းအရာများကို ၎င်း၏နောက်မှ ချက်ချင်းလိုက်လာသော `string` variable ထဲသို့ embed လုပ်သည်။
//
//go:embed folder/single_file.txt
var fileString string

// သို့မဟုတ် ဖိုင်၏အကြောင်းအရာများကို `[]byte` ထဲသို့ embed လုပ်နိုင်သည်။
//
//go:embed folder/single_file.txt
var fileByte []byte

// ကျွန်ုပ်တို့သည် ဖိုင်အများအပြားကိုသော်လည်းကောင်း၊ wildcard များဖြင့် ဖိုဒါများကိုပင် embed လုပ်နိုင်သည်။
// ဤနေရာတွင် [embed.FS type](https://pkg.go.dev/embed#FS) ၏ variable ကို အသုံးပြုထားပြီး
// ၎င်းသည် ရိုးရှင်းသော virtual file system တစ်ခုကို အကောင်အထည်ဖော်သည်။
//
//go:embed folder/single_file.txt
//go:embed folder/*.hash
var folder embed.FS

func main() {

	// `single_file.txt` ၏ အကြောင်းအရာများကို ပရင့်ထုတ်ပါ။
	print(fileString)
	print(string(fileByte))

	// Embedded folder မှ ဖိုင်အချို့ကို ရယူပါ။
	content1, _ := folder.ReadFile("folder/file1.hash")
	print(string(content1))

	content2, _ := folder.ReadFile("folder/file2.hash")
	print(string(content2))
}
