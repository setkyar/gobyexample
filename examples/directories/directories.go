// Go သည် ဖိုင်စနစ်အတွင်းရှိ _directories_ များနှင့် အလုပ်လုပ်ရန်
// အသုံးဝင်သော function များစွာ ရှိသည်။
package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// လက်ရှိ အလုပ်လုပ်နေသော directory အတွင်း
	// sub-directory အသစ်တစ်ခု ဖန်တီးပါ။
	err := os.Mkdir("subdir", 0755)
	check(err)

	// ယာယီ directories များ ဖန်တီးသောအခါ၊ ၎င်းတို့ကို ဖယ်ရှားရန်
	// `defer` သုံးခြင်းသည် ကောင်းမွန်သော အလေ့အထဖြစ်သည်။ `os.RemoveAll`
	// သည် directory tree တစ်ခုလုံးကို ဖျက်ပစ်လိမ့်မည် (`rm -rf` နှင့် အလားတူ)။
	defer os.RemoveAll("subdir")

	// ဖိုင်အလွတ်အသစ်တစ်ခု ဖန်တီးရန် Helper function။
	createEmptyFile := func(name string) {
		d := []byte("")
		check(os.WriteFile(name, d, 0644))
	}

	createEmptyFile("subdir/file1")

	// ကျွန်ုပ်တို့သည် parents ပါဝင်သော directory အဆင့်ဆင့်ကို `MkdirAll` ဖြင့်
	// ဖန်တီးနိုင်သည်။ ဤသည်မှာ command-line `mkdir -p` နှင့် ဆင်တူသည်။
	err = os.MkdirAll("subdir/parent/child", 0755)
	check(err)

	createEmptyFile("subdir/parent/file2")
	createEmptyFile("subdir/parent/file3")
	createEmptyFile("subdir/parent/child/file4")

	// `ReadDir` သည် directory အတွင်းရှိ အရာများကို စာရင်းပြုစုပြီး
	// `os.DirEntry` objects များ၏ slice ကို ပြန်ပေးသည်။
	c, err := os.ReadDir("subdir/parent")
	check(err)

	fmt.Println("Listing subdir/parent")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	// `Chdir` သည် လက်ရှိ လုပ်ငန်းလည်ပတ်နေသော directory ကို
	// ပြောင်းလဲခွင့်ပြုသည်၊ `cd` နှင့် အလားတူသည်။
	err = os.Chdir("subdir/parent/child")
	check(err)

	// ယခု ကျွန်ုပ်တို့သည် *လက်ရှိ* directory ကို စာရင်းပြုစုသောအခါ
	// `subdir/parent/child` ၏ အကြောင်းအရာများကို မြင်ရမည်။
	c, err = os.ReadDir(".")
	check(err)

	fmt.Println("Listing subdir/parent/child")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	// စတင်ခဲ့သောနေရာသို့ `cd` ဖြင့် ပြန်သွားပါ။
	err = os.Chdir("../../..")
	check(err)

	// ကျွန်ုပ်တို့သည် directory တစ်ခုကို _recursively_ လည်း
	// ကြည့်ရှု(visit_နိုင်သည်၊ ၎င်း၏ sub-directories အားလုံးအပါအဝင်ဖြစ်သည်။
	// `Walk` သည် တွေ့ရှိသော ဖိုင် သို့မဟုတ် directory တိုင်းကို
	// ကိုင်တွယ်ရန် callback function တစ်ခုကို လက်ခံသည်။
	fmt.Println("Visiting subdir")
	err = filepath.Walk("subdir", visit)
}

// `visit` သည် `filepath.Walk` မှ recursively တွေ့ရှိသော
// ဖိုင် သို့မဟုတ် directory တိုင်းအတွက် ခေါ်ဆိုခြင်းခံရသည်။
func visit(p string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	fmt.Println(" ", p, info.IsDir())
	return nil
}
