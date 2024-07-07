# ဤပရိုဂရမ်ကို ကျွန်ုပ်တို့ run သောအခါ၊ ၎င်းသည် 
# signal တစ်ခုကို စောင့်ဆိုင်းရန် block လုပ်ထားပါမည်။
# ctrl-C ကို နှိပ်ခြင်းဖြင့် (terminal က ^C 
# အဖြစ် ပြသပါသည်) ကျွန်ုပ်တို့သည် SIGINT signal ကို 
# ပို့နိုင်ပြီး၊ ထိုအခါ ပရိုဂရမ်သည် "interrupt" 
# ဟု ပုံနှိပ်ဖော်ပြပြီး ထွက်ခွာသွားပါမည်။
$ go run signals.go
awaiting signal
^C
interrupt
exiting
