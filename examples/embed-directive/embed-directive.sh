# ဤဥပမာကို run ရန် အောက်ပါ commands များကို အသုံးပြုပါ။
# (မှတ်ချက်: go playground ၏ ကန့်သတ်ချက်ကြောင့်
# ဤဥပမာကို သင့်ကွန်ပျူတာတွင်သာ run နိုင်ပါသည်။)
$ mkdir -p folder
$ echo "hello go" > folder/single_file.txt
$ echo "123" > folder/file1.hash
$ echo "456" > folder/file2.hash

$ go run embed-directive.go
hello go
hello go
123
456

