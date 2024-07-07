# ပထမဦးစွာ foo subcommand ကို ခေါ်ယူသည်။
$ ./command-line-subcommands foo -enable -name=joe a1 a2
subcommand 'foo'
  enable: true
  name: joe
  tail: [a1 a2]

# ယခု bar ကို စမ်းကြည့်သည်။
$ ./command-line-subcommands bar -level 8 a1
subcommand 'bar'
  level: 8
  tail: [a1]

# သို့သော် bar သည် foo ၏ flag များကို လက်ခံမည်မဟုတ်ပါ။
$ ./command-line-subcommands bar -enable a1
flag provided but not defined: -enable
Usage of bar:
  -level int
    	level

# နောက်တစ်ဆင့်တွင် ကျွန်ုပ်တို့သည် environment variable 
# များကို ကြည့်မည်ဖြစ်သည်။ 
# ၎င်းသည် program များကို parameter ပေးရန် 
# နောက်ထပ်နည်းလမ်းတစ်ခုဖြစ်သည်။