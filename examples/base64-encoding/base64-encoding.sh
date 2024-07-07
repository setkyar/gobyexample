# စံပြ(standard)နှင့် URL base64 encoders 
# များသည် string ကို 
# အနည်းငယ်ကွဲပြားသော တန်ဖိုးများအဖြစ် 
# encode လုပ်ပါသည်
# (နောက်ဆုံးတွင် `+` နှင့် `-` အသုံးပြုခြင်း)
# သို့သော် ၎င်းတို့နှစ်ခုလုံးသည် မူရင်း string သို့ 
# လိုအပ်သလို ပြန်လည် decode လုပ်နိုင်ပါသည်။
$ go run base64-encoding.go
YWJjMTIzIT8kKiYoKSctPUB+
abc123!?$*&()'-=@~

YWJjMTIzIT8kKiYoKSctPUB-
abc123!?$*&()'-=@~
