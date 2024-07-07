# သင်က `exit.go` ကို `go run` နဲ့ run ရင်၊ exit က
# `go` က ဖမ်းယူပြီး ပုံနှိပ်ဖော်ပြပေးပါလိမ့်မယ်။
$ go run exit.go
exit status 3
# Binary ကို build လုပ်ပြီး run ခြင်းဖြင့် သင်က
# terminal ထဲမှာ status ကို မြင်နိုင်ပါတယ်။
$ go build exit.go
$ ./exit
$ echo $?
3
# မှတ်ချက်။ ကျွန်ုပ်တို့ program ထဲက `!` 
# က ဘယ်တော့မှ ပုံနှိပ် (print) မထွက်လာပါ။