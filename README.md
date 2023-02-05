# Go by Example မြန်မာဘာသာ

Go ကို ဥပမာ program များဖြင့် သင်ကြား ပေးသော webstie
[Go by Example မြန်မာဘာသာ](https://setkyar.github.io/gobyexample)
ကိုတည်ဆောက်ရန်လိုအပ်သည့် build toolchain နှင့် ဘာသာပြန်
စာများပါဝင်ပါတယ်။

### စာကြမ်း
Go by Example site သည် `example` source files ထဲကကုဒ်
နှင့် comments တွေကိုဖတ်၊ `templates` တွေကို `public` folder
ထဲ static file အဖြစ်ပြောင်းတဲ့ပုံစံနှင့် တည်ဆောက်ထားတာပါ။
ဒီပရိုဂရမ်တည်ဆောက်ဖို့ build process တွေကို `tools` folder
ထဲမှာ တခြား `go.mod` file တွေထဲမှာတွေ့နိုင်ပါတယ်။

The built `public` directory can be served by any
static content system. The production site uses S3 and
CloudFront, for example.

### တည်ဆောက်ခြင်း

[![test](https://github.com/setkyar/gobyexample/actions/workflows/test.yml/badge.svg)](https://github.com/setkyar/gobyexample/actions/workflows/test.yml)

ဒီဝက်ဆိုဒ်ကိုတည်ဆောက်ရန်အတွက် Go ကိုထည့်သွားထားဖို့လိုပါတယ်။
အောက်ဖော်ပြပါအတိုင်းအသုံးပြုနိုင်တယ် -

```console
$ tools/build
```

စဥ်ဆက်မပြတ် တည်ဆောက်ရန်အတွက်အောက်ဖော်ပြပါအတိုင်းအသုံးပြုပါ။


```console
$ tools/build-loop
```

ဝက်ဆိုဒ်ကို ကိုယ့်စက်မှာကြည့်ရန်အောက်ဖော်ပြပါအတိုင်းအသုံးပြုပါ။

```
$ tools/serve
```

`http://127.0.0.1:8000/` ကိုသင့်ဘရောက်ဆာကနေဖွင့်ပါ။

### လွှင့်တင်ခြင်း

ဝက်ဆိုကိုလွှင့်တင်ရန်အတွက်:

```console
$ export AWS_ACCESS_KEY_ID=...
$ export AWS_SECRET_ACCESS_KEY=...
$ tools/upload
```

မြန်မာဘာသာပြန် ဝက်ဆိုဒ်ကို လွှင့်တင်ရန်အတွက် အောက်ဖော်ပြပါ command ကို
အသုံးပြုပီး HTML တွေကို public-gh မှာထည့်ပါတယ်။ ပီးသွားလျှင် Github
Pages ကိုသုံးပီးတော့ အဲ့ဒီ့ folder ကိုညွှန်းထားပါတယ်။

```console
$ tools/upload-gh.sh
```

### လိုင်စင်

This work is copyright Mark McGranaghan and licensed under a
[Creative Commons Attribution 3.0 Unported License](http://creativecommons.org/licenses/by/3.0/).

The Go Gopher is copyright [Renée French](https://reneefrench.blogspot.com/) and licensed under a
[Creative Commons Attribution 3.0 Unported License](http://creativecommons.org/licenses/by/3.0/).


### ဘာသာစကား

မြန်မာဘာသာပြန်အပြင် တခြားအောက်ဖော်ပြပါ ဘာသာစကားများဖြင့်လဲဖတ်ရှုနိုင်ပါတယ်။

* [English](https://github.com/mmcgrana/gobyexample) မူရင်း by [gobyexample](https://github.com/mmcgrana/gobyexample)
* [Chinese](https://gobyexample-cn.github.io/) by [gobyexample-cn](https://github.com/gobyexample-cn)
* [Czech](http://gobyexamples.sweb.cz/) by [martinkunc](https://github.com/martinkunc/gobyexample-cz)
* [French](http://le-go-par-l-exemple.keiruaprod.fr) by [keirua](https://github.com/keirua/gobyexample)
* [Italian](https://gobyexample.it) by the [Go Italian community](https://github.com/golangit/gobyexample-it)
* [Japanese](http://spinute.org/go-by-example) by [spinute](https://github.com/spinute)
* [Korean](https://mingrammer.com/gobyexample/) by [mingrammer](https://github.com/mingrammer)
* [Russian](https://gobyexample.com.ru/) by [badkaktus](https://github.com/badkaktus)
* [Spanish](http://goconejemplos.com) by the [Go Mexico community](https://github.com/dabit/gobyexample)
* [Ukrainian](https://butuzov.github.io/gobyexample/) by [butuzov](https://github.com/butuzov/gobyexample)

### ကျေးဇူးတင်ပါသည်

Thanks to [Jeremy Ashkenas](https://github.com/jashkenas)
for [Docco](http://jashkenas.github.io/docco/), which
inspired this project.

### FAQ

#### I found a problem with the examples; what do I do?

We're very happy to fix problem reports and accept contributions! Please submit
[an issue](https://github.com/mmcgrana/gobyexample/issues) or send a Pull Request.
See `CONTRIBUTING.md` for more details.

#### What version of Go is required to run these examples?

Given Go's strong [backwards compatibility guarantees](https://go.dev/doc/go1compat),
we expect the vast majority of examples to work on the latest released version of Go
as well as many older releases going back years.

That said, some examples show off new features added in recent releases; therefore,
it's recommended to try running examples with the latest officially released Go version
(see Go's [release history](https://go.dev/doc/devel/release) for details).

#### I'm getting output in a different order from the example. Is the example wrong?

Some of the examples demonstrate concurrent code which has a non-deterministic
execution order. It depends on how the Go runtime schedules its goroutines and
may vary by operating system, CPU architecture, or even Go version.

Similarly, examples that iterate over maps may produce items in a different order
from what you're getting on your machine. This is because the order of iteration
over maps in Go is [not specified and is not guaranteed to be the same from one
iteration to the next](https://go.dev/ref/spec#RangeClause).

It doesn't mean anything is wrong with the example. Typically the code in these
examples will be insensitive to the actual order of the output; if the code is
sensitive to the order - that's probably a bug - so feel free to report it.