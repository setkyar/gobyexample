# ဒီပရိုဂရမ်ကို run လိုက်ရင် panic ဖြစ်သွားပြီး၊ 
# error message နဲ့ goroutine trace တွေကို 
# ပြပေးပါလိမ့်မယ်။ ပြီးတော့ non-zero status 
# နဲ့ exit ဖြစ်သွားပါလိမ့်မယ်။

# `main` ထဲက ပထမဆုံး panic က အလုပ်လုပ်သွားတဲ့အခါ၊ 
# ပရိုဂရမ်က ကျန်တဲ့ code တွေကို မရောက်တော့ဘဲ 
# ထွက်သွားပါလိမ့်မယ်။ အကယ်၍ ပရိုဂရမ်က temp file 
# တစ်ခု ဖန်တီးဖို့ ကြိုးစားတာကို မြင်ချင်ရင်
# ပထမဆုံး panic ကို comment out လုပ်ထားလိုက်ပါ။
$ go run panic.go
panic: a problem

goroutine 1 [running]:
main.main()
	/.../panic.go:12 +0x47
...
exit status 2

# သတိပြုရမှာက တချို့ language တွေမှာ exception တွေကို
#  error အများစုကို handle လုပ်ဖို့ သုံးကြပေမယ့်၊ 
# Go မှာတော့ ဖြစ်နိုင်သမျှနေရာတိုင်းမှာ 
# error-indicating return value 
# တွေကို သုံးတာက ပုံမှန်လုပ်နည်းလုပ်ဟန် ဖြစ်ပါတယ်။
