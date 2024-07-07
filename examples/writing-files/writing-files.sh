# ဖိုင်ရေးသားသည့် ကုဒ်ကို စမ်းသပ်ပြီး run ကြည့်ပါ။
$ go run writing-files.go 
wrote 5 bytes
wrote 7 bytes
wrote 9 bytes

# ထို့နောက် ရေးသားထားသော ဖိုင်များ၏ 
# အကြောင်းအရာများ(content)ကို စစ်ဆေးကြည့်ပါ။
$ cat /tmp/dat1
hello
go
$ cat /tmp/dat2
some
writes
buffered

# နောက်တစ်ဆင့်အနေဖြင့် ယခုလေ့လာခဲ့သော 
# ဖိုင် I/O သဘောတရားအချို့ကို
# `stdin` နှင့် `stdout` stream များတွင် 
# အသုံးချခြင်းကို ကြည့်ကြမည်။
