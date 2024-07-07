// _line filter_ ဆိုတာ stdin ကနေ input ကိုဖတ်ပြီး၊ process လုပ်ကာ
// ရလဒ်ကို stdout ပေါ်မှာ print ထုတ်တဲ့ ပုံစံမျိုးဖြစ်ပါတယ်။ `grep` နဲ့ `sed` တို့က
// အသုံးများတဲ့ line filter တွေဖြစ်ပါတယ်။

// ဒီနေရာမှာ Go နဲ့ရေးထားတဲ့ line filter နမူနာတစ်ခုပြထားပါတယ်။
// ဒီ code က input text အားလုံးကို စာလုံးအကြီးပြောင်းပြီး ရေးထုတ်ပေးပါတယ်။
// ဒီပုံစံကိုသုံးပြီး သင့်ရဲ့ကိုယ်ပိုင် Go line filter တွေရေးနိုင်ပါတယ်။

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Buffering မလုပ်ထားတဲ့ `os.Stdin` ကို buffered scanner နဲ့ wrap လုပ်ထားတာက
	// အဆင်ပြေတဲ့ `Scan` method ကိုပေးပါတယ်။ ဒီ method က scanner ကို နောက် token ဆီရွှေ့ပေးပါတယ်။
	// ပုံမှန်အားဖြင့် နောက် line ဆီကိုရွှေ့ပေးတာပါ။
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		// `Text()` က လက်ရှိ token ကိုပြန်ပေးပါတယ်။ ဒီနေရာမှာတော့ နောက် line ကိုပြန်ပေးတာဖြစ်ပါတယ်။
		ucl := strings.ToUpper(scanner.Text())
		// စာလုံးအကြီးပြောင်းထားတဲ့ line ကို ရေးထုတ်ပါတယ်။
		fmt.Println(ucl)
	}
	// `Scan` လုပ်နေစဉ် error ဖြစ်မဖြစ်စစ်ပါတယ်။ file အဆုံးရောက်တာကို
	// `Scan` က error အဖြစ်မသတ်မှတ်ပါဘူး။
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
