// Unit testing သည် အခြေခံကျသော Go ပရိုဂရမ်များ ရေးသားရာတွင် အရေးကြီးသော အပိုင်းတစ်ခုဖြစ်သည်။ `testing` package က
// unit test များရေးသားရန် လိုအပ်သော tools များကို ပေးပြီး `go test` command က test များကို run ပေးသည်။

// သရုပ်ပြရန်အတွက် ဤ code သည် `main` package တွင်ရှိသော်လည်း မည်သည့် package တွင်မဆို ဖြစ်နိုင်သည်။ Testing code သည်
// အများအားဖြင့် ၎င်းစစ်ဆေးသော code နှင့် package နှင့်အတး တည်ရှိလေ့ရှိသည်။
package main

import (
	"fmt"
	"testing"
)

// ကျွန်ုပ်တို့သည် ကိန်းဂဏန်း အနည်းဆုံးတန်ဖိုးရှာသော ရိုးရှင်းသောအကောင်အထည်ဖော်မှုကို စမ်းသပ်မည်ဖြစ်သည်။ ပုံမှန်အားဖြင့်
// ကျွန်ုပ်တို့စစ်ဆေးနေသော code သည် `intutils.go` ကဲ့သို့သော အမည်ရှိ source file တွင်ရှိမည်ဖြစ်ပြီး ၎င်းအတွက် test file မှာ
// `intutils_test.go` ဟု အမည်ပေးမည်ဖြစ်သည်။
func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// `Test` ဖြင့်စသော အမည်ရှိ function တစ်ခုရေးခြင်းဖြင့် test တစ်ခုကို ဖန်တီးသည်။
func TestIntMinBasic(t *testing.T) {
	ans := IntMin(2, -2)
	if ans != -2 {
		// `t.Error*` သည် test အောင်မြင်မှုမရှိကြောင်း report လုပ်သော်လည်း test ကို ဆက်လက်အကောင်အထည်ဖော် (execute)မည်။
		// `t.Fatal*` သည် test အောင်မြင်မှုမရှိကြောင်း report လုပ်ပြီး test ကို ချက်ချင်းရပ်တန့်စေမည်။
		t.Errorf("IntMin(2, -2) = %d; want -2", ans)
	}
}

// Test များရေးသားခြင်းသည် ထပ်ခါတလဲလဲဖြစ်နိုင်သောကြောင့် *table-driven style* ကို အသုံးပြုရန် ထုံးစံဖြစ်သည်။
// ဤနည်းတွင် test input များနှင့် မျှော်မှန်းထားသော output များကို ဇယားတစ်ခုတွင် စာရင်းပြုစုထားပြီး loop တစ်ခုက
// ၎င်းတို့ကိုလိုက်လံစစ်ဆေးကာ test logic ကို အကောင်အထည်ဖော်သည်။
func TestIntMinTableDriven(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{0, 1, 0},
		{1, 0, 0},
		{2, -2, -2},
		{0, -1, -1},
		{-1, 0, -1},
	}
	for _, tt := range tests {
		// t.Run သည် "subtest" များကို run နိုင်စေသည်။ ဇယားထဲရှိ entry တစ်ခုစီအတွက် တစ်ခုစီဖြစ်သည်။
		// ၎င်းတို့ကို `go test -v` run သောအခါ သီးခြားစီပြသသည်။
		testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := IntMin(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

// Benchmark test များသည် အများအားဖြင့် `_test.go` file များတွင်ရှိပြီး `Benchmark` ဖြင့်စသော အမည်ရှိကြသည်။
// `testing` runner သည် benchmark function တစ်ခုစီကို အကြိမ်ပေါင်းများစွာ run ပြီး တိကျသောတိုင်းတာမှုရရှိသည်အထိ
// `b.N` ကို တိုးမြှင့်သွားသည်။
func BenchmarkIntMin(b *testing.B) {
	// ပုံမှန်အားဖြင့် benchmark သည် ကျွန်ုပ်တို့ benchmark လုပ်နေသော function ကို `b.N` အကြိမ် loop ပတ်၍ run သည်။
	for i := 0; i < b.N; i++ {
		IntMin(1, 2)
	}
}
