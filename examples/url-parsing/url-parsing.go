// URL တွေက [resource တွေကို တည်နေရာသတ်မှတ်ဖို့ ပုံစံတူနည်းလမ်း](https://adam.herokuapp.com/past/2010/3/30/urls_are_the_uniform_way_to_locate_resources/) တစ်ခုပေးပါတယ်။
// Go မှာ URL တွေကို ဘယ်လို parse လုပ်မလဲဆိုတာ ဒီမှာရှင်းပြပါမယ်။

package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {

	// ကျွန်တော်တို့က ဒီဥပမာ URL ကို parse လုပ်ပါမယ်။ ဒီထဲမှာ
	// scheme၊ authentication info၊ host၊ port၊ path၊
	// query params နဲ့ query fragment တွေ ပါဝင်ပါတယ်။
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	// URL ကို parse လုပ်ပြီး error မရှိကြောင်း သေချာအောင်လုပ်ပါ။
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	// Scheme ကို access လုပ်တာ ရိုးရှင်းပါတယ်။
	fmt.Println(u.Scheme)

	// `User` က authentication info အားလုံးပါဝင်ပါတယ်။ တစ်ခုချင်းစီရဲ့
	// တန်ဖိုးတွေအတွက် `Username` နဲ့ `Password` ကို ခေါ်သုံးပါ။
	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	p, _ := u.User.Password()
	fmt.Println(p)

	// `Host` မှာ hostname နဲ့ port နှစ်ခုလုံး ပါဝင်ပါတယ် (port ရှိရင်)။
	// သူတို့ကို ခွဲထုတ်ဖို့ `SplitHostPort` ကို သုံးပါ။
	fmt.Println(u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host)
	fmt.Println(port)

	// ဒီမှာတော့ `path` နဲ့ `#` နောက်က fragment ကို ထုတ်ယူပါတယ်။
	fmt.Println(u.Path)
	fmt.Println(u.Fragment)

	// Query params တွေကို `k=v` ပုံစံ string အနေနဲ့ ရယူဖို့ `RawQuery` ကို သုံးပါ။
	// Query params တွေကို map အဖြစ်လည်း parse လုပ်နိုင်ပါတယ်။ Parse လုပ်ထားတဲ့
	// query param map တွေက string ကနေ string တွေရဲ့ slice ဆီကို ဖြစ်ပါတယ်။
	// ဒါကြောင့် ပထမဆုံးတန်ဖိုးကိုပဲ လိုချင်ရင် `[0]` ကို index လုပ်ပါ။
	fmt.Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["k"][0])
}
