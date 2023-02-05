//	`for` က Go ရဲ့တခုထဲသော looping construct တစ်ခုဖြစ်ပါတယ်။
//
// တစ်ခုဖြစ်ပါတယ်။ ဒီမှာဥပမာတချို့ကိုကြည့်နိုင်ပါတယ်။
package main

import "fmt"

func main() {

	// အခြေအနေတခုထဲကိုစစ်တဲ့ အခြေခံအကျဆုံးပုံစံ
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// ပုံမှန်အသုံးပြုနေကြ စမှတ်/အခြေအနေ/နောက်တဆင့် `for` loop
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// အခြေအနေဘာမှစစ်မထားတဲ့ `for` loop ကတောက်လျှောက်
	// `break` သို `retutrn` မပြန်မချင်း ထက်ခါထက်ခါ
	// loop ပတ်နေမှာဖြစ်ပါတယ်။
	for {
		fmt.Println("loop ပတ်")
		break
	}

	// `continue` ကိုသုံးပြီးတော့ loop ရဲ့နောက်တဆင့်ကိုဆက်သွားနိုင်ပါတယ်။
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
