# Server ကို နောက်ခံမှာ run ပါ။
$ go run context-in-http-servers.go &

# Client က `/hello` ကို request 
# လုပ်တာကို simulate လုပ်ပါ။
# စလိုက်ပြီး မကြာခင်မှာပဲ Ctrl+C နှိပ်ပြီး 
# cancel လုပ်တာကို signal ပေးပါ။
$ curl localhost:8090/hello
server: hello handler started
^C
server: context canceled
server: hello handler ended
