// Version 1.18 ကစပြီး Go မှာ _generics_ (သို့) _type parameters_ လို့လည်းခေါ်တဲ့
// feature ကိုထည့်သွင်း support လုပ်ထားပါတယ်။

package main

import "fmt"

// Generic function ရဲ့ ဥပမာအနေနဲ့၊ `MapKeys` က map ရဲ့ဘယ်လို type မျိုးကိုမဆို လက်ခံပြီး
// သူ့ရဲ့ key တွေကို slice အဖြစ် ပြန်ပေးပါတယ်။
// ဒီ function မှာ type parameter နှစ်ခုရှိပါတယ် - `K` နဲ့ `V`။
// `K` မှာ `comparable` ဆိုတဲ့ _constraint_ ရှိပါတယ်၊ ဆိုလိုတာက
// ဒီ type ရဲ့ တန်ဖိုးတွေကို `==` နဲ့ `!=` operator တွေနဲ့ နှိုင်းယှဉ်လို့ရပါတယ်။
// ဒါက Go မှာ map key တွေအတွက် လိုအပ်ချက်တစ်ခုဖြစ်ပါတယ်။
// `V` မှာတော့ `any` constraint ရှိပါတယ်၊ ဆိုလိုတာက သူ့ကို ဘာမှ ကန့်သတ်မထားဘူးဆိုတာပါ
// (`any` က `interface{}` ရဲ့ alias ဖြစ်ပါတယ်)။
func MapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

// Generic type ရဲ့ ဥပမာအနေနဲ့၊ `List` ရဲ့ values တွေက any type ဖြစ်တဲ့
// singly-linked list တစ်ခုဖြစ်ပါတယ်။
type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

// ပုံမှန် type တွေမှာလိုပဲ generic type တွေမှာလည်း method တွေ သတ်မှတ်လို့ရပါတယ်၊
// ဒါပေမယ့် type parameter တွေကို ထည့်ထားဖို့လိုပါတယ်။ Type က `List[T]` ဖြစ်ပြီး `List` မဟုတ်ပါဘူး။
func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) GetAll() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func main() {
	var m = map[int]string{1: "2", 2: "4", 4: "8"}

	// Generic function တွေကို ခေါ်သုံးတဲ့အခါ၊ _type inference_ ကို အားကိုးနိုင်ပါတယ်။
	// `MapKeys` ကို ခေါ်တဲ့အခါ `K` နဲ့ `V` အတွက် type တွေကို သတ်မှတ်ပေးစရာမလိုတာကို သတိထားပါ -
	// compiler က သူ့ဘာသာ အလိုအလျောက် ခန့်မှန်းပေးပါတယ်။
	fmt.Println("keys:", MapKeys(m))

	// ... ဒါပေမယ့် လိုချင်ရင်တော့ တိတိကျကျ သတ်မှတ်ပေးလို့လည်း ရပါတယ်။
	_ = MapKeys[int, string](m)

	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	fmt.Println("list:", lst.GetAll())
}
