# zjson

json friendly string builder.
it from uber zap package.

change:

* add RawMessage Field Type

note:

* default every row have LineEnding '\n' , but you can intervene through configuration

example:

```go

    encoder := NewJSONEncoder(NewProductionEncoderConfig())
    
    buf, err := encoder.Encode(
        String("A", "test1"),
        Time("B", time.Now()),
        String("C", "test2"),
        Int64("D", 100),
        RawMessage("E", []byte(`{}`)),
    )

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