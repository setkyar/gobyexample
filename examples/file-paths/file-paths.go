// `filepath` package သည် operating systems အကြား portable ဖြစ်သော
// နည်းလမ်းဖြင့် *file paths* များကို parse လုပ်ရန်နှင့် တည်ဆောက်ရန် functions များ
// ပေးပါသည်။ ဥပမာအားဖြင့် Linux တွင် `dir/file` နှင့် Windows တွင် `dir\file` ဖြစ်သည်။
package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {

	// `Join` ကို portable နည်းလမ်းဖြင့် paths များ တည်ဆောက်ရန် အသုံးပြုသင့်သည်။
	// ၎င်းသည် မည်သည့် arguments အရေအတွက်ကိုမဆို လက်ခံပြီး ၎င်းတို့မှ
	// အဆင့်ဆင့် path တစ်ခု တည်ဆောက်ပေးသည်။
	p := filepath.Join("dir1", "dir2", "filename")
	fmt.Println("p:", p)

	// `/`s သို့မဟုတ် `\`s များကို ကိုယ်တိုင် ပေါင်းစပ်ခြင်းအစား
	// `Join` ကို အမြဲတမ်း အသုံးပြုသင့်သည်။ portable
	// ဖြစ်စေခြင်းအပြင် `Join` သည် မလိုအပ်သော separators
	// နှင့် directory ပြောင်းလဲမှုများကို ဖယ်ရှားခြင်းဖြင့် paths
	// များကို ပုံမှန်ဖြစ်အောင် ပြုလုပ်ပေးပါသည်။
	fmt.Println(filepath.Join("dir1//", "filename"))
	fmt.Println(filepath.Join("dir1/../dir1", "filename"))

	// `Dir` နှင့် `Base` ကို path တစ်ခုကို directory နှင့်
	// file အဖြစ် ခွဲရန် အသုံးပြုနိုင်သည်။ တနည်းအားဖြင့်
	// `Split` သည် နှစ်ခုလုံးကို တစ်ကြိမ်တည်းဖြင့် ပြန်ပေးပါသည်။
	fmt.Println("Dir(p):", filepath.Dir(p))
	fmt.Println("Base(p):", filepath.Base(p))

	// path တစ်ခုသည် absolute ဖြစ်မဖြစ် စစ်ဆေးနိုင်သည်။
	fmt.Println(filepath.IsAbs("dir/file"))
	fmt.Println(filepath.IsAbs("/dir/file"))
	filename := "config.json"

	// အချို့ file အမည်များတွင် အစက်(dot .)နောက်တွင်
	// extensions များ ပါရှိသည်။ ထိုကဲ့သို့သော
	// အမည်များမှ extension ကို `Ext` ဖြင့် ခွဲထုတ်နိုင်သည်။
	ext := filepath.Ext(filename)
	fmt.Println(ext)

	// extension ကို ဖယ်ထုတ်ထားသော file ၏အမည်ကို
	// ရှာရန် `strings.TrimSuffix` ကို အသုံးပြုပါ။
	fmt.Println(strings.TrimSuffix(filename, ext))

	// `Rel` သည် *base* နှင့် *target* အကြား relative path
	// ကို ရှာဖွေပေးသည်။ target ကို base နှင့် ဆက်စပ်၍ မရနိုင်ပါက error ကို ပြန်ပေးသည်။
	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)
}
