# ကျွန်တော်တို့ program ကို run တဲ့အခါ 
# အဲဒါကို `ls` နဲ့ အစားထိုးခံရပါတယ်။
$ go run execing-processes.go
total 16
drwxr-xr-x  4 mark 136B Oct 3 16:29 .
drwxr-xr-x 91 mark 3.0K Oct 3 12:50 ..
-rw-r--r--  1 mark 1.3K Oct 3 16:28 execing-processes.go

# Go မှာ ဂန္ထဝင် Unix `fork` function 
# မပါဝင်ကြောင်း သတိပြုပါ။
# သို့သော် များသောအားဖြင့် ဒါက ပြဿနာမဖြစ်ပါဘူး၊ 
# အဘယ်ကြောင့်ဆိုသော်
# goroutine တွေ စတင်ခြင်း၊ process တွေ spawn 
# လုပ်ခြင်း၊ နဲ့ process တွေကို exec လုပ်ခြင်းတို့က
# `fork` အတွက် အသုံးပြုလေ့ရှိတဲ့ 
# ကိစ္စအများစုကို လွှမ်းခြုံနိုင်လို့ပါ။