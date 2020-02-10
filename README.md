# zjson

json friendly string builder.
it from uber zap package.

change:

* add RawMessage Field Type

note:

* default every row have LineEnding '\n' , but you can intervene through configuration

example:

```go

package main

import (
	"log"
	"time"

	"github.com/zplzpl/zjson"
)

func main() {
	encoder := zjson.NewJSONEncoder(zjson.NewProductionEncoderConfig())

	buf, err := encoder.Encode(
		zjson.String("A", "test1"),
		zjson.Time("B", time.Now()),
		zjson.String("C", "test2"),
		zjson.Int64("D", 100),
		zjson.RawMessage("E", []byte(`{}`)),
	)

	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println(string(buf.Bytes()))
}


```

output:

```json
{"A":"test1","B":"2020-02-10T22:46:14+08:00","C":"test2","D":100,"E":{}}
```

# 中文说明

这是个友好的json字符串构造器，从uber zap包中提出出来。

用途：

* 用json输出某些简短的数据时候，又有性能要求的时候，用这个简单快捷

变动：

* 增加RawMessage的字段类型

注意:

* 目前输出默认追加"\n"符号