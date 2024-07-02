# ကျွန်တော်တို့ run နေတဲ့ပရိုဂရမ်က 
# အလုပ် 5 ခုကို worker တွေက
# executed နေပါတယ်။ 
# စုစုပေါင်း အလုပ်ချိန် 5 စက္ကန့်လောက် 
# ရှိသော်လည်း ပရိုဂရမ်က 
# 2 စက္ကန့်လောက်ပဲ ကြာပါတယ်။ 
# ဘာကြောင့်လဲဆိုတော့ 
# worker 3 ခုက 
# တပြိုင်နက်တည်း အလုပ်လုပ်နေလို့ပါ။
$ time go run worker-pools.go 
worker 1 started  job 1
worker 2 started  job 2
worker 3 started  job 3
worker 1 finished job 1
worker 1 started  job 4
worker 2 finished job 2
worker 2 started  job 5
worker 3 finished job 3
worker 1 finished job 4
worker 2 finished job 5

real	0m2.358s
