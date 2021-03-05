package pong2trans

import (
	"fmt"
	"strings"
	"testing"

	"github.com/flosch/pongo2/v4"
)

type testTrans struct {
	lang string
	data map[string]string
}

func (t *testTrans) Translate(in string) string {

	for k, v := range t.data {
		if v == in {
			out, ok := t.data[t.lang+"|"+strings.Split(k, "|")[1]]
			if !ok {
				return in
			}

			return out
		}
	}
	return t.data[in]
}

func newTestTrans() *testTrans {
	return &testTrans{
		lang: "zh_CN",
		data: map[string]string{
			"zh_CN|title": "你好！世界",
			"en_US|title": "Hello World",
		},
	}
}

func TestTagsTrans(t *testing.T) {
	fmt.Println(RegisterTransTag(newTestTrans()))

	tplStr := "{% trans \"Hello World\" %}"
	tpl, err := pongo2.FromString(tplStr)
	if err != nil {
		t.Fatalf("failed to parsing %q: %v", tplStr, err)
	}
	res, err := tpl.Execute(pongo2.Context{})
	if err != nil {
		t.Fatalf("failed to executing %q: %v", tplStr, err)
	}
	fmt.Println(res)
}
