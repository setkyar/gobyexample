// Go က `printf` ပုံစံမျိုးနဲ့ string formatting အတွက် အထောက်အပံ့ကောင်းကောင်း (excellent support) ပေးထားပါတယ်။
// ဒီမှာ ပုံမှန်တွေ့ရတဲ့ string formatting လုပ်ငန်းတွေရဲ့ ဥပမာတချို့ ရှိပါတယ်။

package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {

	// Go မှာ အထွေထွေ Go value တွေကို format လုပ်ဖို့ ဒီဇိုင်းဆွဲထားတဲ့
	// printing "verb" တွေ အများကြီး ရှိပါတယ်။ ဥပမာ၊ ဒါက ကျွန်တော်တို့ရဲ့
	// `point` struct တစ်ခုကို print ထုတ်ပါတယ်။
	p := point{1, 2}
	fmt.Printf("struct1: %v\n", p)

	// Value က struct ဖြစ်နေရင်၊ `%+v` က
	// struct ရဲ့ field နာမည်တွေပါ ထည့်သွင်းပါလိမ့်မယ်။
	fmt.Printf("struct2: %+v\n", p)

	// `%#v` က value ရဲ့ Go syntax ကိုယ်စားပြုမှုကို print ထုတ်ပါတယ်။
	// ဆိုလိုတာက အဲဒီ value ကို ထုတ်လုပ်မယ့် source code အပိုင်းအစပါ။
	fmt.Printf("struct3: %#v\n", p)

	// Value တစ်ခုရဲ့ type ကို print ထုတ်ဖို့ `%T` ကို သုံးပါ။
	fmt.Printf("type: %T\n", p)

	// Boolean တွေကို format လုပ်တာ ရိုးရှင်းပါတယ်။
	fmt.Printf("bool: %t\n", true)

	// Integer တွေကို format လုပ်ဖို့ option အများကြီး ရှိပါတယ်။
	// ပုံမှန် base-10 format အတွက် `%d` ကို သုံးပါ။
	fmt.Printf("int: %d\n", 123)

	// ဒါက binary ပုံစံနဲ့ print ထုတ်ပါတယ်။
	fmt.Printf("bin: %b\n", 14)

	// ဒါက ပေးလိုက်တဲ့ integer နဲ့ ကိုက်ညီတဲ့ စာလုံးကို print ထုတ်ပါတယ်။
	fmt.Printf("char: %c\n", 33)

	// `%x` က hex encoding ပေးပါတယ်။
	fmt.Printf("hex: %x\n", 456)

	// Float တွေအတွက်လည်း formatting option အများကြီး ရှိပါတယ်။
	// အခြေခံ ဒဿမ formatting အတွက် `%f` ကို သုံးပါ။
	fmt.Printf("float1: %f\n", 78.9)

	// `%e` နဲ့ `%E` က float ကို သိပ္ပံနည်းကျ ရေးနည်း (နည်းနည်း ကွာခြားတဲ့ version နှစ်ခု) နဲ့ format လုပ်ပါတယ်။
	fmt.Printf("float2: %e\n", 123400000.0)
	fmt.Printf("float3: %E\n", 123400000.0)

	// အခြေခံ string printing အတွက် `%s` ကို သုံးပါ။
	fmt.Printf("str1: %s\n", "\"string\"")

	// Go source code မှာလိုမျိုး string တွေကို double-quote လုပ်ဖို့ `%q` ကို သုံးပါ။
	fmt.Printf("str2: %q\n", "\"string\"")

	// အပေါ်မှာ မြင်ခဲ့ရတဲ့ integer တွေလိုပဲ၊ `%x` က
	// string ကို base-16 နဲ့ ဖော်ပြပါတယ်။ Input byte တစ်ခုစီအတွက်
	// output စာလုံး နှစ်လုံးစီ ထွက်ပါတယ်။
	fmt.Printf("str3: %x\n", "hex this")

	// Pointer တစ်ခုရဲ့ ကိုယ်စားပြုမှုကို print ထုတ်ဖို့ `%p` ကို သုံးပါ။
	fmt.Printf("pointer: %p\n", &p)

	// နံပါတ်တွေကို format လုပ်တဲ့အခါ ထွက်လာမယ့် ဂဏန်းရဲ့
	// အကျယ်နဲ့ တိကျမှုကို ထိန်းချုပ်ချင်တာ မကြာခဏ လုပ်ချင်ပါလိမ့်မယ်။
	// Integer တစ်ခုရဲ့ အကျယ်ကို သတ်မှတ်ဖို့ verb ထဲက `%` နောက်မှာ
	// နံပါတ်တစ်ခု ထည့်ပါ။ ပုံမှန်အားဖြင့် ရလဒ်က ညာဘက်ကပ်ပြီး
	// space တွေနဲ့ ဖြည့်ထားပါလိမ့်မယ်။
	fmt.Printf("width1: |%6d|%6d|\n", 12, 345)

	// နံပါတ်တွေကို format လုပ်တဲ့အခါ ထွက်လာမယ့် ဂဏန်းရဲ့
	// အကျယ်နဲ့ တိကျမှုကို ထိန်းချုပ်ချင်တာ မကြာခဏ ဖြစ်ပါလိမ့်မယ်။
	// Integer တစ်ခုရဲ့ အကျယ်ကို သတ်မှတ်ဖို့ verb ထဲက `%` နောက်မှာ
	// နံပါတ်တစ်ခု ထည့်ပါ။ ပုံမှန်အားဖြင့် ရလဒ်က ညာဘက်ကပ်ပြီး
	// space တွေနဲ့ ဖြည့်ထားပါလိမ့်မယ်။
	fmt.Printf("width2: |%6.2f|%6.2f|\n", 1.2, 3.45)

	// ဘယ်ဘက်ကပ်ဖို့ `-` flag ကို သုံးပါ။
	fmt.Printf("width3: |%-6.2f|%-6.2f|\n", 1.2, 3.45)

	// String တွေကို format လုပ်တဲ့အခါလည်း အကျယ်ကို ထိန်းချုပ်ချင်နိုင်ပါတယ်။
	// အထူးသဖြင့် ဇယားလို output မှာ တန်းစီအောင် လုပ်ချင်တဲ့အခါပါ။
	// အခြေခံ ညာဘက်ကပ် အကျယ်အတွက်:
	fmt.Printf("width4: |%6s|%6s|\n", "foo", "b")

	// ဘယ်ဘက်ကပ်ဖို့ နံပါတ်တွေမှာလိုပဲ `-` flag ကို သုံးပါ။
	fmt.Printf("width5: |%-6s|%-6s|\n", "foo", "b")

	// အခုထိ ကျွန်တော်တို့ `Printf` ကို မြင်ခဲ့ကြပါတယ်။ သူက
	// format လုပ်ထားတဲ့ string ကို `os.Stdout` မှာ print ထုတ်ပါတယ်။
	// `Sprintf` ကတော့ string ကို format လုပ်ပြီး ဘယ်နေရာမှာမှ print မထုတ်ဘဲ
	// ပြန်ပေးပါတယ်။
	s := fmt.Sprintf("sprintf: a %s", "string")
	fmt.Println(s)

	// `os.Stdout` မဟုတ်တဲ့ တခြား `io.Writers` တွေမှာ format+print လုပ်ချင်ရင်
	// `Fprintf` ကို သုံးနိုင်ပါတယ်။
	fmt.Fprintf(os.Stderr, "io: an %s\n", "error")
}
