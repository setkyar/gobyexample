// Go က JSON encoding နဲ့ decoding အတွက် built-in support ပေးထားပါတယ်။
// built-in နဲ့ custom data type တွေကနေ JSON သို့၊ JSON ကနေ ပြန်ပြောင်းတာတွေ ပါဝင်ပါတယ်။

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// အောက်က struct နှစ်ခုကို custom type တွေရဲ့ encoding နဲ့ decoding ကို သရုပ်ပြဖို့ သုံးပါမယ်။
type response1 struct {
	Page   int
	Fruits []string
}

// Export လုပ်ထားတဲ့ field တွေကိုပဲ JSON အဖြစ် encode/decode လုပ်ပါမယ်။
// Field တွေကို export လုပ်ဖို့ စာလုံးအကြီးနဲ့ စရပါမယ်။
type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {

	// ပထမဆုံး အခြေခံ data type တွေကို JSON string အဖြစ် encode လုပ်တာကို ကြည့်ကြရအောင်။
	// ဒီမှာ atomic value တွေအတွက် ဥပမာတချို့ ရှိပါတယ်။
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))
	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))
	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))
	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	// ဒီမှာတော့ slice တွေနဲ့ map တွေအတွက် ဥပမာတွေပါ။ သူတို့က
	// မျှော်လင့်ထားသလိုပဲ JSON array တွေနဲ့ object တွေအဖြစ် encode လုပ်ပါတယ်။
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))
	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	// JSON package က သင့်ရဲ့ custom data type တွေကို အလိုအလျောက် encode လုပ်ပေးနိုင်ပါတယ်။
	// သူက encode လုပ်ထားတဲ့ output ထဲမှာ export လုပ်ထားတဲ့ field တွေကိုပဲ ထည့်သွင်းပြီး
	// ပုံမှန်အားဖြင့် အဲဒီနာမည်တွေကိုပဲ JSON key အဖြစ် သုံးပါတယ်။
	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	// Struct field ကြေညာချက်တွေမှာ tag တွေ သုံးပြီး encode လုပ်ထားတဲ့ JSON key နာမည်တွေကို
	// စိတ်ကြိုက်ပြောင်းလဲနိုင်ပါတယ်။ အဲဒီလို tag တွေရဲ့ ဥပမာကို အပေါ်က `response2` အဓိပ္ပာယ်ဖွင့်ဆိုချက်မှာ ကြည့်ပါ။
	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	// အခု JSON data ကို Go value တွေအဖြစ် decode လုပ်တာကို ကြည့်ကြရအောင်။
	// ဒီမှာ အထွေထွေသုံး data structure တစ်ခုအတွက် ဥပမာတစ်ခု ရှိပါတယ်။
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	// JSON package က decode လုပ်ထားတဲ့ data ကို ထည့်နိုင်မယ့် variable တစ်ခု ပေးဖို့ လိုပါတယ်။
	// ဒီ `map[string]interface{}` က string တွေကနေ မည်သည့် data type မဆိုကို ညွှန်းတဲ့ map တစ်ခုကို ကိုင်ထားပါလိမ့်မယ်။
	var dat map[string]interface{}

	// ဒါကတော့ တကယ့် decoding လုပ်တာပါ၊ ပြီးတော့ ဆက်စပ်နေတဲ့ error တွေကိုလည်း စစ်ဆေးပါတယ်။
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	// Decode လုပ်ထားတဲ့ map ထဲက တန်ဖိုးတွေကို သုံးဖို့ဆိုရင်
	// သူတို့ကို သင့်တော်တဲ့ type အဖြစ် ပြောင်းဖို့ လိုပါလိမ့်မယ်။
	// ဥပမာ ဒီမှာ `num` ထဲက တန်ဖိုးကို မျှော်လင့်ထားတဲ့ `float64` type အဖြစ် ပြောင်းပါတယ်။
	num := dat["num"].(float64)
	fmt.Println(num)

	// Nested data တွေကို access လုပ်ဖို့ ပြောင်းလဲမှုတွေ အများကြီး လုပ်ဖို့လိုပါတယ်။
	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	// JSON ကို custom data type တွေအဖြစ်လည်း decode လုပ်နိုင်ပါတယ်။
	// ဒီနည်းက ကျွန်တော်တို့ ပရိုဂရမ်တွေမှာ type-safety ထပ်ဖြည့်ပေးပြီး
	// decode လုပ်ထားတဲ့ data ကို access လုပ်တဲ့အခါ type assertion တွေ လုပ်စရာမလိုတဲ့ အားသာချက် ရှိပါတယ်။
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	// အပေါ်က ဥပမာတွေမှာ ကျွန်တော်တို့က data နဲ့ standard out ပေါ်က JSON ကိုယ်စားပြုမှုကြားမှာ
	// byte တွေနဲ့ string တွေကို အမြဲသုံးခဲ့ပါတယ်။ ကျွန်တော်တို့က JSON encoding တွေကို
	// `os.Stdout` သို့မဟုတ် HTTP response body တွေလိုမျိုး `os.Writer` တွေဆီ တိုက်ရိုက် stream လုပ်နိုင်ပါတယ်။
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}
