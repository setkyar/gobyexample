// `go` ကဲ့သို့သော command-line tools အချို့သည် မိမိတို့ပိုင် flag set များပါဝင်သော subcommand များစွာရှိသည်။
// ဥပမာအားဖြင့် `go build` နှင့် `go get` တို့သည် `go` tool ၏ မတူညီသော subcommand နှစ်ခုဖြစ်သည်။
// `flag` package သည် ရိုးရှင်းသော subcommand များကို လွယ်ကူစွာ သတ်မှတ်နိုင်စေပြီး ၎င်းတို့တွင် မိမိတို့ပိုင် flag များပါဝင်နိုင်သည်။

package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	// `NewFlagSet` function ကိုအသုံးပြု၍ subcommand တစ်ခုကို ကြေညာပြီး
	// ထို subcommand အတွက် သီးသန့် flag အသစ်များကို သတ်မှတ်သည်။
	fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
	fooEnable := fooCmd.Bool("enable", false, "enable")
	fooName := fooCmd.String("name", "", "name")

	// အခြား subcommand တစ်ခုအတွက် မတူညီသော flag များကို သတ်မှတ်နိုင်သည်။
	barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
	barLevel := barCmd.Int("level", 0, "level")

	// Subcommand ကို program ၏ ပထမဆုံး argument အဖြစ် မျှော်လင့်ထားသည်။
	if len(os.Args) < 2 {
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}

	// မည်သည့် subcommand ကို ခေါ်ယူထားသည်ကို စစ်ဆေးသည်။
	switch os.Args[1] {
	// Subcommand တစ်ခုချင်းစီအတွက် ၎င်း၏ flag များကို parse လုပ်ပြီး
	// နောက်ဆက်တွဲ positional argument များကို အသုံးပြုနိုင်သည်။
	case "foo":
		fooCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'foo'")
		fmt.Println("  enable:", *fooEnable)
		fmt.Println("  name:", *fooName)
		fmt.Println("  tail:", fooCmd.Args())
	case "bar":
		barCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'bar'")
		fmt.Println("  level:", *barLevel)
		fmt.Println("  tail:", barCmd.Args())
	default:
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}
}
