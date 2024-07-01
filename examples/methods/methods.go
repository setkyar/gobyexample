// Go က struct types တွေအပေါ်မှာ သတ်မှတ်တဲ့ _methods_ တွေကို support လုပ်ပါတယ်။

package main

import "fmt"

type rect struct {
	width, height int
}

// ဒီ `area` method မှာ `*rect` ဆိုတဲ့ _receiver type_ ရှိပါတယ်။
func (r *rect) area() int {
	return r.width * r.height
}

// Method တွေကို pointer သို့မဟုတ် value receiver types နှစ်မျိုးစလုံးအတွက် သတ်မှတ်နိုင်ပါတယ်။
// ဒီမှာ value receiver ရဲ့ ဥပမာတစ်ခုကို ပြထားပါတယ်။
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	// ဒီမှာ ကျွန်တော်တို့ရဲ့ struct အတွက် သတ်မှတ်ထားတဲ့ method နှစ်ခုကို ခေါ်သုံးပါတယ်။
	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	// Go က method တွေကို ခေါ်တဲ့အခါ value နဲ့ pointer ကြား အလိုအလျောက် ပြောင်းလဲပေးပါတယ်။
	// သင့်အနေနဲ့ method ခေါ်တဲ့အခါ copy လုပ်တာကို ရှောင်ရှားချင်ရင် သို့မဟုတ်
	// method က receiving struct ကို ပြောင်းလဲခွင့်ပေးချင်ရင်
	// pointer receiver type ကို သုံးချင်နိုင်ပါတယ်။
	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}
