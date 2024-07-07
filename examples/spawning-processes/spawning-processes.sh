# ဖန်တီးထားတဲ့ ပရိုဂရမ်တွေက ထွက်လာတဲ့ output ဟာ 
# ကျွန်တော်တို့ command-line ကနေ 
# တိုက်ရိုက် run သလိုပဲ ဖြစ်ပါတယ်။
$ go run spawning-processes.go 
> date
Thu 05 May 2022 10:10:12 PM PDT
# date မှာ `-x` flag မရှိတဲ့အတွက် 
# အမှားပြ message နဲ့အတူ 
# non-zero return code နဲ့ ထွက်သွားပါလိမ့်မယ်။
command exited with rc = 1
> grep hello
hello grep
> ls -a -l -h
drwxr-xr-x  4 mark 136B Oct 3 16:29 .
drwxr-xr-x 91 mark 3.0K Oct 3 12:50 ..
-rw-r--r--  1 mark 1.3K Oct 3 16:28 spawning-processes.go