# ဒီပရိုဂရမ်ကို run လိုက်တဲ့အခါ၊ ကျွန်တော်တို့က 
# blocking call ရဲ့ output ကို အရင်မြင်ရပြီး
# နောက်မှ goroutine နှစ်ခုရဲ့ output ကို မြင်ရပါတယ်။
# goroutine တွေရဲ့ output က
# ရောနှောနေနိုင်ပါတယ်။ ဘာကြောင့်လဲဆိုတော့ 
# goroutine တွေက Go runtime က
# တပြိုင်နက်တည်း concurrent ဖြစ်အောင် run နေလို့ပါ။.
$ go run goroutines.go
direct : 0
direct : 1
direct : 2
goroutine : 0
going
goroutine : 1
goroutine : 2
done

# နောက်တစ်ဆင့်မှာ Go ရဲ့ concurrent 
# ပရိုဂရမ်တွေမှာ goroutine တွေနဲ့အတူ
# တွဲဖက်သုံးလေ့ရှိတဲ့ channel တွေအကြောင်း
#  လေ့လာကြပါမယ်။
