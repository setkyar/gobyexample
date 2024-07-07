# Command-line arguments တွေကို စမ်းသပ်ဖို့အတွက် 
# ပထမဆုံး `go build` နဲ့ binary ကို 
# တည်ဆောက်တာ(build)က အကောင်းဆုံးဖြစ်ပါတယ်။
$ go build command-line-arguments.go
$ ./command-line-arguments a b c d
[./command-line-arguments a b c d]       
[a b c d]
c

# နောက်တစ်ဆင့်အနေနဲ့ ပိုပြီးအဆင့်မြင့်တဲ့ 
# command-line processing ကို
# flags တွေသုံးပြီး လေ့လာကြမှာ ဖြစ်ပါတယ်။
