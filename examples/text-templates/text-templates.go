// Go က `text/template` package နဲ့ dynamic content ဖန်တီးဖို့ သို့မဟုတ်
// user ကို customized output ပြဖို့ built-in support ပေးထားပါတယ်။
// `html/template` ဆိုတဲ့ ဆင်တူ package တစ်ခုကလည်း တူညီတဲ့ API ပေးပြီး
// လုံခြုံရေးဆိုင်ရာ feature တွေ ထပ်ပါပါတယ်။ HTML generate လုပ်ဖို့အတွက် အဲဒါကို သုံးသင့်ပါတယ်။

package main

import (
	"os"
	"text/template"
)

func main() {
	// Template အသစ်တစ်ခု ဖန်တီးပြီး သူ့ရဲ့ body ကို string တစ်ခုကနေ parse လုပ်နိုင်ပါတယ်။
	// Template တွေဟာ static text နဲ့ `{{...}}` ထဲမှာ ပါတဲ့ "action" တွေရဲ့ ပေါင်းစပ်မှုဖြစ်ပါတယ်။
	// Action တွေက content ကို dynamic ထည့်သွင်းဖို့ သုံးပါတယ်။
	t1 := template.New("t1")
	t1, err := t1.Parse("Value is {{.}}\n")
	if err != nil {
		panic(err)
	}

	// နောက်တစ်နည်းအနေနဲ့ `template.Must` function ကို သုံးပြီး
	// `Parse` က error ပြန်ပေးရင် panic ဖြစ်အောင် လုပ်နိုင်ပါတယ်။
	// ဒါက global scope မှာ initialize လုပ်တဲ့ template တွေအတွက် အထူးအသုံးဝင်ပါတယ်။
	t1 = template.Must(t1.Parse("Value: {{.}}\n"))

	// Template ကို "execute" လုပ်ခြင်းအားဖြင့် သူ့ရဲ့ action တွေအတွက်
	// သတ်မှတ်ထားတဲ့ တန်ဖိုးတွေနဲ့အတူ စာသားကို ထုတ်လုပ်ပါတယ်။
	// `{{.}}` action က `Execute` ကို parameter အဖြစ်ပေးလိုက်တဲ့ တန်ဖိုးနဲ့ အစားထိုးခံရပါတယ်။
	t1.Execute(os.Stdout, "some text")
	t1.Execute(os.Stdout, 5)
	t1.Execute(os.Stdout, []string{
		"Go",
		"Rust",
		"C++",
		"C#",
	})

	// အောက်မှာ သုံးမယ့် helper function တစ်ခုပါ။
	Create := func(name, t string) *template.Template {
		return template.Must(template.New(name).Parse(t))
	}

	// Data က struct ဖြစ်နေရင် `{{.FieldName}}` action ကို သုံးပြီး
	// သူ့ရဲ့ field တွေကို access လုပ်နိုင်ပါတယ်။ Template execute လုပ်နေချိန်မှာ
	// access လုပ်နိုင်ဖို့ field တွေကို export လုပ်ထားသင့်ပါတယ်။
	t2 := Create("t2", "Name: {{.Name}}\n")

	t2.Execute(os.Stdout, struct {
		Name string
	}{"Jane Doe"})

	// Map တွေအတွက်လည်း အတူတူပါပဲ။ Map တွေမှာတော့ key နာမည်တွေရဲ့
	// အကြီးအသေးပေါ် ကန့်သတ်ချက် မရှိပါဘူး။
	t2.Execute(os.Stdout, map[string]string{
		"Name": "Mickey Mouse",
	})

	// if/else က template တွေမှာ conditional execution ပေးပါတယ်။
	// Type တစ်ခုရဲ့ default value ဖြစ်တဲ့ 0၊ string အလွတ်၊
	// nil pointer စတာတွေကို false လို့ သတ်မှတ်ပါတယ်။
	// ဒီဥပမာက template ရဲ့ နောက်ထပ် feature တစ်ခုကိုလည်း ပြသထားပါတယ်:
	// whitespace ဖြတ်ဖို့ action တွေထဲမှာ `-` သုံးတာပါ။
	t3 := Create("t3",
		"{{if . -}} yes {{else -}} no {{end}}\n")
	t3.Execute(os.Stdout, "not empty")
	t3.Execute(os.Stdout, "")

	// range block တွေက slice၊ array၊ map သို့မဟုတ် channel တွေကို
	// loop လုပ်ခွင့်ပေးပါတယ်။ Range block ထဲမှာ `{{.}}` က
	// လက်ရှိ iteration ရဲ့ item ကို ညွှန်းပါတယ်။
	t4 := Create("t4",
		"Range: {{range .}}{{.}} {{end}}\n")
	t4.Execute(os.Stdout,
		[]string{
			"Go",
			"Rust",
			"C++",
			"C#",
		})
}
