# pongo2 trans export

Reference [pongo2](https://github.com/flosch/pongo2) library

### Installation
```
go get github.com/ly020044/pongo2trans
```

### Export
```go
package main

import (
	"encoding/json"
	"io/fs"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/flosch/pongo2/v4"

	pong2trans "github.com/ly020044/pongo2trans"
)

type testExport struct {
	data map[string]string
}

func (te *testExport) Export(value string) {
	te.data[value] = ""
}

func main() {
	te := &testExport{data: map[string]string{}}
	if err := pong2trans.RegisterTransTag(nil, te); err != nil {
		log.Fatal(err)
	}

	err := filepath.WalkDir("dir", func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}

		tpl, err := pongo2.FromFile(path)
		if err != nil {
			return err
		}
		_, err = tpl.Execute(pongo2.Context{})
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	raw, err := json.Marshal(te.data)
	// raw, err := json.MarshalIndent(te.data, "", "\t")
	if err != nil {
		log.Fatal(err)
	}

	ioutil.WriteFile("test.json", raw, 0644)
}
```
