// _Interface_ တွေဆိုတာ method signature တွေကို ပေါင်းပြီး အမည်ပေးထားတာပါ။

package main

import (
	"fmt"
	"math"
)

// ဒီမှာ ဂျီသြမေတီပုံသဏ္ဍာန်တွေအတွက် အခြေခံ interface တစ်ခုကို သတ်မှတ်ထားပါတယ်။
type geometry interface {
	area() float64
	perim() float64
}

// ဥပမာအနေနဲ့ ဒီ interface ကို `rect` နဲ့ `circle` type တွေမှာ
// implement လုပ်ပါ့မယ်။
type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

// Go မှာ interface တစ်ခုကို implement လုပ်ဖို့ဆိုရင် interface ထဲက method အားလုံးကို
// implementလုပ်ဖို့ပဲ လိုပါတယ်။ ဒီမှာ `rect` အတွက် `geometry` interface ကို implement လုပ်ထားပါတယ်။
func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// `circle` အတွက် implement လုပ်ထားတာပါ
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// Variable တစ်ခုက interface type ဖြစ်နေရင် အဲဒီ interface ထဲက method တွေကို ခေါ်သုံးလို့ရပါတယ်။
// ဒီမှာ `measure` ဆိုတဲ့ generic function တစ်ခုက ဒီအချက်ကို အသုံးချပြီး
// မည်သည့် `geometry` function မဆို အလုပ်လုပ်နိုင်အောင် ရေးထားပါတယ်။
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	// `circle` နဲ့ `rect` struct type နှစ်ခုစလုံးက `geometry` interface ကို
	// implement လုပ်ထားတဲ့အတွက် သူတို့ရဲ့ instance တွေကို `measure` function ရဲ့
	// argument အဖြစ် သုံးလို့ရပါတယ်။
	measure(r)
	measure(c)
}
