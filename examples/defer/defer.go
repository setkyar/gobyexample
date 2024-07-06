// _Defer_ ကို ပရိုဂရမ်ရဲ့ လုပ်ဆောင်မှုအတွင်းမှာ function call တစ်ခုကို နောက်ပိုင်းမှ
// လုပ်ဆောင်စေဖို့ သုံးပါတယ်။ အများအားဖြင့် cleanup လုပ်ဖို့ ရည်ရွယ်ချက်နဲ့ သုံးလေ့ရှိပါတယ်။
// တခြားဘာသာစကားတွေမှာ `ensure` နဲ့ `finally` သုံးတဲ့နေရာမျိုးတွေမှာ Go က `defer` ကို
// မကြာခဏ သုံးပါတယ်။

package main

import (
	"fmt"
	"os"
)

// ဖိုင်တစ်ခု ဖန်တီးပြီး၊ ဖိုင်အထဲကို ရေးမယ်၊ ပြီးရင် ပိတ်ပစ်မယ်ဆိုပါစို့။
// `defer` ကိုသုံးပြီး ဒီလိုလုပ်လို့ရပါတယ်။
func main() {

	// `createFile` နဲ့ file object ရရှိပြီးတာနဲ့ ချက်ချင်း
	// အဲဒီဖိုင်ကို ပိတ်မယ့် `closeFile` ကို defer လုပ်ထားလိုက်ပါတယ်။
	// ဒါက လက်ရှိ function (`main`) ပြီးဆုံးချိန်မှာ၊
	// `writeFile` ပြီးသွားတဲ့နောက်မှာ အလုပ်လုပ်ပါလိမ့်မယ်။
	f := createFile("/tmp/defer.txt")
	defer closeFile(f)
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")

}

func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()
	// Deferred function ထဲမှာတောင် ဖိုင်ပိတ်တဲ့အခါ
	// error စစ်ဖို့ အရေးကြီးပါတယ်။
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
