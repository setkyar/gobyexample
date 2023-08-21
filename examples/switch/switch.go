// _Switch statements_ တွေက expression တွေကိုသုံးပြီး conditionals တွေကိုစစ်တာပေါ့။
package main

import (
	"fmt"
	"time"
)

func main() {

	// ဒီမှာ ဥပမာ အခြေခံ `switch` တခုအသုံးပြုပုံ
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// expression နှစ်ခု သုံးခုစစ်ဖို့ comma သုံးပြီးစစ်နိုင်ပါတယ်။
	// ကျနော်တို့ဒီဥပမာဆိုရင် မဖြစ်မနေသုံးဖို့တော့မလိုတဲ့
	// `default` ကိုသုံးထားပါတယ်။
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	// ကျနော်တို့ `switch` ကို expression မပါဘဲနှင့်သုံးမယ်ဆိုရင်
	// if/else logic လိုမျိုးသုံးလို့ရတယ်။ ဒီဥပမာမှာဆိုရင် ကျနော်တို့စစ်တဲ့
	// case က constant (ပုံသေ) မဟုတ်ဘူးဆိုတာကိုပြထားတယ်။
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	// ကျနော်တို့ `switch` ကို types တွေအတွက်လဲအသုံးပြုတယ်။ သူကတော့
	// type ကိုဘဲစစ်လိမ့်မယ်။ ဒီဥပမာမှာဆိုရင် varaible t က i ရဲ့
	// type ကိုစစ်တာပေါ့။
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
