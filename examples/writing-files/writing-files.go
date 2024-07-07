// Go တွင် ဖိုင်ရေးသားခြင်းသည် ယခင်က ကျွန်ုပ်တို့တွေ့ခဲ့သော
// ဖိုင်ဖတ်ခြင်းနှင့် ဆင်တူသော ပုံစံများကို လိုက်နာသည်။

package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// စတင်ရန်အတွက်၊ string (သို့မဟုတ် byte များ) ကို
	// ဖိုင်ထဲသို့ ထည့်သွင်းရေးသားပုံကို ဖော်ပြထားသည်။
	d1 := []byte("hello\ngo\n")
	err := os.WriteFile("/tmp/dat1", d1, 0644)
	check(err)

	// ပိုမိုအသေးစိတ်ရေးသားရန်အတွက်၊ ဖိုင်ကို ရေးသားရန် ဖွင့်ပါ။
	f, err := os.Create("/tmp/dat2")
	check(err)

	// ဖိုင်ဖွင့်ပြီးချက်ချင်း `Close` ကို defer လုပ်ခြင်းသည် ထုံးစံဖြစ်သည်။
	defer f.Close()

	// ကျနော်တို့လုပ်ချင်တဲ့အတိုင်း byte slice များကို `Write` နိုင်သည်။
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	// `WriteString` လည်းရှိပါသည်။
	n3, err := f.WriteString("writes\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n3)

	// ရေးသားထားသည်များကို တည်ငြိမ်သော သိုလှောင်မှုသို့ flush လုပ်ရန် `Sync` ကို အသုံးပြုပါ။
	f.Sync()

	// `bufio` သည် ယခင်က မြင်ခဲ့ရသော buffered reader များအပြင်
	// buffered writer များကိုလည်း ပေးစွမ်းသည်။
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n4)

	// buffer ထဲရှိ လုပ်ဆောင်ချက်အားလုံးကို အခြေခံ writer ထဲသို့
	// အသုံးချပြီးကြောင်း သေချာစေရန် `Flush` ကို အသုံးပြုပါ။
	w.Flush()
}
