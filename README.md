## brainfuck-go
[![GoDoc](https://godoc.org/github.com/ryomak/brainfuck-go?status.svg)](https://godoc.org/github.com/ryomak/brainfuck-go)
[![GoReport](https://goreportcard.com/badge/github.com/ryomak/brainfuck-go)](https://goreportcard.com/report/github.com/ryomak/brainfuck-go)
### install 
```
$ GO111MODULE=off go get github.com/ryomak/brainfuck-go/cmd/bfgo
```

### usage 
```
NAME:
   brainfuck-go

USAGE:
   bfgo [global options] command [command options] [arguments...]

VERSION:
   1.0.0

COMMANDS:
   run      run bf file
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

run command

```
NAME:
   bfgo run - run bf file

USAGE:
   bfgo run [command options] [arguments...]

OPTIONS:
   -c value, --config value  config for setting bf token
   -m value, --max value     set max order num (default: 1000)
```

### example

```
# fizzbuzz.bf
++++++[->++++> >+>+>-<<<<<]>[<++++> >+++>++++> >+++>+
++++>+++++> > > > > >++> >++<<<<<<<<<<<<<<-]<++++>+++
>-->+++>-> >--->++> > >+++++[->++>++<<]<<<<<<<<<<[->-
[> > > > > > >]>[<+++>.>.> > > >..> > >+<]<<<<<-[> > > >]>[<+
++++>.>.>..> > >+<]> > > >+<-[<<<]<[[-<<+> >]> > >+>+<<<<<
<[-> >+>+>-<<<<]<]>>[[-]<]>[> > >[>.<<.<<<]<[.<<<<]>]>.<<<<
<<<<<<<] 
```

```
$ bfgo run fizzbuzz.bf
```

### advance Usage
1. making token file
ex)token.toml
```
[Commands]
NEXT="あ"
PREV ="い"
INC  ="う"
DEC  ="え"
READ ="お"
WRITE="か"
OPEN ="き"
CLOSE="く" 
```
2. bf file
ex)hiragana.bf
```
ううううううきえあううううあ あうあうあえいいいいいくあきいううううあ あうううあううううあ あうううあう
ううううあうううううあ あ あ あ あ あううあ あうういいいいいいいいいいいいいいえくいううううあううう
あええあうううあえあ あえええあううあ あ あうううううきえあううあうういいくいいいいいいいいいいきえあえ
きあ あ あ あ あ あ あくあきいうううあかあかあ あ あ あかかあ あ あういくいいいいいえきあ あ あ あくあきいう
ううううあかあかあかかあ あ あういくあ あ あ あういえきいいいくいききえいいうあ あくあ あ あうあういいいいい
いきえあ あうあうあえいいいいくいくああききえくいくあきあ あ あきあかいいかいいいくいきかいいいいくあくあかいいいい
いいいいいいいく    
```

3. run
```
bfgo run -c token.toml hiragana.bf
```

## LICENSE
MIT
