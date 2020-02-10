package zjson

import (
	"testing"
	"time"
)

func TestJsonEncoder_Test1(t *testing.T) {

	encoder := NewJSONEncoder(NewProductionEncoderConfig())

	buf, err := encoder.Encode(
		String("A", "test1"),
		Time("B", time.Now()),
		String("C", "test2"),
		Int64("D", 100),
		RawMessage("E", []byte(`{}`)),
	)

	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(buf.Bytes()))

}
