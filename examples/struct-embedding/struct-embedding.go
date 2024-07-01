// Go က struct နဲ့ interface တွေကို _embedding_ လုပ်ခြင်းအားဖြင့်
// type တွေရဲ့ _composition_ ကို ပိုမိုချောမွေ့စွာ ဖော်ပြနိုင်ပါတယ်။
// ဒါကို Go version 1.16+ မှာ မိတ်ဆက်ထားတဲ့ `//go:embed` directive နဲ့ မတူပါဘူး၊
// အဲဒီ directive က file နဲ့ folder တွေကို application binary ထဲ embed လုပ်ဖို့သုံးတာပါ။

package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

// `container` က `base` ကို _embeds_ လုပ်ထားပါတယ်။ Embedding က
// နာမည်မပါတဲ့ field တစ်ခုလိုပုံစံမျိုးပါ။
type container struct {
	base
	str string
}

func main() {

	// Struct တွေကို literal နဲ့ တည်ဆောက်တဲ့အခါ၊ embedding ကို
	// တိတိကျကျ initialize လုပ်ပေးရပါတယ်။ ဒီမှာ
	// embed လုပ်ထားတဲ့ type က field name အဖြစ်သုံးထားပါတယ်။
	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}

	// base ရဲ့ field တွေကို `co` ကနေ တိုက်ရိုက်သုံးလို့ရပါတယ်။
	// ဥပမာ `co.num`
	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	// တခြားနည်းအနေနဲ့ embed လုပ်ထားတဲ့ type နာမည်ကို သုံးပြီး
	// အပြည့်အစုံရေးလို့ရပါတယ်။
	fmt.Println("also num:", co.base.num)

	// `container` က `base` ကို embed လုပ်ထားတဲ့အတွက်၊ `base` ရဲ့ method တွေဟာ
	// `container` ရဲ့ method တွေလည်း ဖြစ်သွားပါတယ်။ ဒီမှာ `base` ကနေ
	// embed လုပ်ထားတဲ့ method ကို `co` ကနေ တိုက်ရိုက်ခေါ်သုံးထားပါတယ်။
	fmt.Println("describe:", co.describe())

	type describer interface {
		describe() string
	}

	// Method ပါတဲ့ struct တွေကို embed လုပ်ခြင်းအားဖြင့် တခြား struct တွေမှာ
	// interface implementation တွေ ရရှိနိုင်ပါတယ်။ ဒီမှာ `container` က
	// `base` ကို embed လုပ်ထားတဲ့အတွက် `describer` interface ကို
	// implement လုပ်နေတာကို တွေ့ရပါတယ်။
	var d describer = co
	fmt.Println("describer:", d.describe())
}
