// Go ရဲ့ _structs_ ဆိုတာက type မတူတဲ့ values တွေကို ပေါင်းစည်းတဲ့နေရာမှာသုံးတယ်။
// အထူးသဖြင့် ကိုယ်ပိုင် data type တွေဖန်တီးတဲ့အချိန်မှာအသုံးပြုကြတယ်။
package main

import "fmt"

// ဒီမှာဆို `person` struct type ထဲမှာ `name` နှင့် age ဆိုတဲ့ field တွေရှိတယ်။
type person struct {
	name string
	age  int
}

// ဒီ `newPerson` function မှာဆိုရင် name ပေးပီးတော့ person struct တခုတည်ဆောက်လိုက်တယ်။
func newPerson(name string) *person {
	// local variable ဖြစ်လို့ pointer ကို စိတ်ချလက်ချ return ပေးလို့ရတယ်။
	// local variable က ပျောက်မသွားဘူး။
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {

	// ဒီ syntax က struct အသစ်တခုတည်ဆောက်လိုက်တယ်။
	fmt.Println(person{"Bob", 20})

	// struct ကို initialize လုပ်နေတဲ့အချိန်မှာ name fields တွေကိုသတ်မှတ်ပေးလို့ရတယ်။
	fmt.Println(person{name: "Alice", age: 30})

	// age ကိုမသတ်မှတ်ထားဘူးဆို zero default ဖြစ်သွားမယ်။
	fmt.Println(person{name: "Fred"})

	// ဒီနေရာမှာ `&` prefix ကိုသုံးထားတာက person data ရဲ့ address ကိုပေးထားတာ။
	// data ကိုပေးထားတာမဟုတ်ဘူး။ struct ရဲ့ pointer ပေါ့။
	fmt.Println(&person{name: "Ann", age: 40})

	//  Go မှာ person struct တခုတည်ဆောက်တော့မယ်ဆိုရင် constructor function
	// တခ ုတည်ဆောက်ပီးသုံးရမယ် တိုက်ရိုက်မသုံးရဘူးဆိုပီးလက်ခံထားကြတယ်။
	fmt.Println(newPerson("Jon"))

	// struct fields တွေကို access လုပ်မယ်ဆို dot နဲ့လုပ်တယ်။
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	// struct pointer တွေသုံးမယ်ဆိုလဲ dot ကိုသုံးလို့ရတယ်။
	// pointer တွေရဲကနေ data ကိုယူတာကို GO ကနောက်ကွယ်ကနေလုပ်ပေးလိမ့်မယ်။
	sp := &s
	fmt.Println(sp.age)

	// Structs တွေက mutable (ပြောင်းလဲနိုင်တယ်)
	sp.age = 51
	fmt.Println(sp.age)
}
