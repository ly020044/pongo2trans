package pong2trans

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/flosch/pongo2/v4"
)

type testTrans struct {
	lang string
	data map[string]string
}

type testExport struct {
	data map[string]string
}

func (te *testExport) Export(value string) {
	te.data[value] = ""
}

func (t *testTrans) Translate(in string) string {
	out, ok := t.data[t.lang+"|"+in]
	if ok {
		return out
	}

	return in
}

func newTestTrans() *testTrans {
	return &testTrans{
		lang: "zh_CN",
		data: map[string]string{
			"zh_CN|Hello World": "你好！世界",
			"en_US|Hello World": "Hello World",
		},
	}
}

func TestTagsTrans(t *testing.T) {
	//te := &testExport{data: []string{}}
	fmt.Println(RegisterTransTag(newTestTrans(), nil))

	//tpl, err := pongo2.FromFile("test.html")
	//if err != nil {
	//	t.Fatal(err)
	//}
	//
	//_, err = tpl.Execute(pongo2.Context{})
	//if err != nil {
	//	t.Fatal(err)
	//}
	//
	//fmt.Println(te)

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

func TestTagTransExport(t *testing.T) {
	te := &testExport{data: map[string]string{}}
	fmt.Println(RegisterTransTag(newTestTrans(), te))

	tpl, err := pongo2.FromFile("test.html")
	if err != nil {
		t.Fatal(err)
	}

	_, err = tpl.Execute(pongo2.Context{})
	if err != nil {
		t.Fatal(err)
	}

	jsonRaw, err := json.Marshal(te.data)
	if err != nil {
		t.Fatalf(err.Error())
	}

	ioutil.WriteFile("en_US.json", jsonRaw, 0644)
}
