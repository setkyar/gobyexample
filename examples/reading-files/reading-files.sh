$ echo "hello" > /tmp/dat
$ echo "go" >>   /tmp/dat
$ go run reading-files.go
hello
go
5 bytes: hello
2 bytes @ 6: go
2 bytes @ 6: go
5 bytes: hello

# နောက်တစ်ဆင့်အနေဖြင့် 
# ဖိုင်ရေးသားခြင်းကို လေ့လာကြမည်။
