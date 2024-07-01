// ပုံမှန်အားဖြင့် channel တွေက _unbuffered_ ဖြစ်ပါတယ်။ ဆိုလိုတာက
// ပို့လိုက်တဲ့တန်ဖိုး (`chan <-`) ကို လက်ခံဖို့ (`<- chan`) အဆင်သင့်ရှိနေမှသာ
// ပို့တာကို လက်ခံမှာဖြစ်ပါတယ်။ _Buffered channel_ တွေကတော့ သတ်မှတ်ထားတဲ့
// အရေအတွက်အထိ တန်ဖိုးတွေကို လက်ခံသူမလိုပဲ လက်ခံနိုင်ပါတယ်။

package main

import "fmt"

func main() {

	// ဒီမှာ တန်ဖိုး 2 ခုအထိ buffer လုပ်နိုင်တဲ့ string channel တစ်ခုကို `make` နဲ့ဖန်တီးထားပါတယ်။
	messages := make(chan string, 2)

	// ဒီ channel က buffered ဖြစ်တဲ့အတွက် ဒီတန်ဖိုးတွေကို တပြိုင်နက်တည်း
	// လက်ခံနေတဲ့သူမရှိပဲနဲ့ channel ထဲကို ပို့လို့ရပါတယ်။
	messages <- "buffered"
	messages <- "channel"

	// နောက်ပိုင်းမှ ဒီတန်ဖိုးနှစ်ခုကို ပုံမှန်အတိုင်း လက်ခံယူလို့ရပါတယ်။
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
