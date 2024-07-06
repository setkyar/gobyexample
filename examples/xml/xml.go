// Go က `encoding.xml` package နဲ့ XML နဲ့ XML-လို format တွေအတွက်
// built-in ထောက်ပံ့မှုပေးထားပါတယ်။

package main

import (
	"encoding/xml"
	"fmt"
)

// Plant ကို XML နဲ့ map လုပ်ပါမယ်။ JSON ဥပမာတွေလိုပဲ၊ field tag တွေက
// encoder နဲ့ decoder အတွက် ညွှန်ကြားချက်တွေ ပါဝင်ပါတယ်။ ဒီမှာ ကျွန်တော်တို့က
// XML package ရဲ့ အထူး feature တချို့ကို သုံးပါတယ်: `XMLName` field နာမည်က
// ဒီ struct ကို ကိုယ်စားပြုတဲ့ XML element ရဲ့ နာမည်ကို သတ်မှတ်ပေးပါတယ်။
// `id,attr` က `Id` field က nested element မဟုတ်ဘဲ XML
// *attribute* တစ်ခုဖြစ်တယ်လို့ ဆိုလိုပါတယ်။
type Plant struct {
	XMLName xml.Name `xml:"plant"`
	Id      int      `xml:"id,attr"`
	Name    string   `xml:"name"`
	Origin  []string `xml:"origin"`
}

func (p Plant) String() string {
	return fmt.Sprintf("Plant id=%v, name=%v, origin=%v",
		p.Id, p.Name, p.Origin)
}
func main() {
	coffee := &Plant{Id: 27, Name: "Coffee"}
	coffee.Origin = []string{"Ethiopia", "Brazil"}

	// ကျွန်တော်တို့ရဲ့ plant ကို ကိုယ်စားပြုတဲ့ XML ကို ထုတ်ပေးပါတယ်။
	// ပိုပြီး လူဖတ်လို့ရတဲ့ output ထုတ်ပေးဖို့ `MarshalIndent` ကို သုံးပါတယ်။
	out, _ := xml.MarshalIndent(coffee, " ", "  ")
	fmt.Println(string(out))

	// Output ထဲကို အထွေထွေသုံး XML header ထည့်ဖို့၊ တိုက်ရိုက် ပေါင်းထည့်လိုက်ပါ။
	fmt.Println(xml.Header + string(out))

	// XML ပါတဲ့ byte stream တစ်ခုကို data structure အဖြစ် parse လုပ်ဖို့
	// `Unmarshal` ကို သုံးပါတယ်။ XML က ပုံပျက်နေရင် သို့မဟုတ် Plant နဲ့
	// map မလုပ်နိုင်ရင် ရှင်းလင်းတဲ့ error တစ်ခု ပြန်ပေးပါလိမ့်မယ်။
	var p Plant
	if err := xml.Unmarshal(out, &p); err != nil {
		panic(err)
	}
	fmt.Println(p)
	tomato := &Plant{Id: 81, Name: "Tomato"}
	tomato.Origin = []string{"Mexico", "California"}

	// `parent>child>plant` field tag က encoder ကို `plant` အားလုံးကို
	// `<parent><child>...` အောက်မှာ nest လုပ်ခိုင်းပါတယ်
	type Nesting struct {
		XMLName xml.Name `xml:"nesting"`
		Plants  []*Plant `xml:"parent>child>plant"`
	}
	nesting := &Nesting{}
	nesting.Plants = []*Plant{coffee, tomato}
	out, _ = xml.MarshalIndent(nesting, " ", "  ")
	fmt.Println(string(out))
}
