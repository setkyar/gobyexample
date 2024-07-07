# command-line flags ပရိုဂရမ်ကို 
# စမ်းသပ်ရန်အတွက်
# ပထမဦးစွာ compile လုပ်ပြီး ထွက်လာသော 
# binary ကို တိုက်ရိုက် run သင့်သည်။
$ go build command-line-flags.go

# flag အားလုံးအတွက် တန်ဖိုးများပေး၍ 
# build လုပ်ထားသော ပရိုဂရမ်ကို စမ်းသပ်ကြည့်ပါ။
$ ./command-line-flags -word=opt -numb=7 -fork -svar=flag
word: opt
numb: 7
fork: true
svar: flag
tail: []

# flag များကို ချန်လှပ်ထားပါက ၎င်းတို့သည် 
# အလိုအလျောက် ၎င်းတို့၏ မူလ(default) တန်ဖိုးများကို 
# ယူကြောင်း သတိပြုပါ။
$ ./command-line-flags -word=opt
word: opt
numb: 42
fork: false
svar: bar
tail: []

# နောက်ဆက်တွဲ positional argument 
# များကို flag များ၏နောက်တွင် ပေးနိုင်သည်။
$ ./command-line-flags -word=opt a1 a2 a3
word: opt
...
tail: [a1 a2 a3]

# `flag` package သည် flag အားလုံးကို positional
#  argument များ၏ရှေ့တွင် ပေါ်လာရန် 
# လိုအပ်ကြောင်း သတိပြုပါ 
# (သို့မဟုတ်ပါက flag များကို positional argument 
# များအဖြစ် အဓိပ္ပာယ်ကောက်ယူမည်)။
$ ./command-line-flags -word=opt a1 a2 a3 -numb=7
word: opt
numb: 42
fork: false
svar: bar
tail: [a1 a2 a3 -numb=7]

# command-line ပရိုဂရမ်အတွက် အလိုအလျောက်ထုတ်ပေးသော 
# help စာသားကို ရရှိရန် `-h` သို့မဟုတ် `--help` 
# flag များကို အသုံးပြုပါ။
$ ./command-line-flags -h
Usage of ./command-line-flags:
 -fork=false: a bool
 -numb=42: an int
 -svar="bar": a string var
 -word="foo": a string

# `flag` package တွင် သတ်မှတ်မထားသော flag 
# တစ်ခုကို ပေးပါက ပရိုဂရမ်သည် အမှားပြ message 
# ကို ပုံနှိပ်ထုတ်ပေးပြီး help စာသားကို ထပ်ပြပေးမည်။
$ ./command-line-flags -wat
flag provided but not defined: -wat
Usage of ./command-line-flags:
...