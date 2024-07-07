# ကျွန်ုပ်တို့ရဲ့ line filter ကို စမ်းသပ်ရန်၊ 
# ပထမဆုံး စာလုံးအသေး အနည်းငယ်ပါသည့် 
# ဖိုင်တစ်ခုကို ဖန်တီးပါမယ်။
$ echo 'hello'   > /tmp/lines
$ echo 'filter' >> /tmp/lines

# ထို့နောက် line filter ကို 
# အသုံးပြု၍ စာလုံးအကြီးများ ရယူပါမယ်။
$ cat /tmp/lines | go run line-filters.go
HELLO
FILTER