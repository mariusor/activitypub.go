package jsonld

import (
	"bytes"
	"strings"
	"testing"
)

func TestRef_MarshalText(t *testing.T) {
	test := "test"
	a := IRI(test)

	out, err := a.MarshalText()
	if err != nil {
		t.Errorf("Error %s", err)
	}
	if bytes.Compare(out, []byte(test)) != 0 {
		t.Errorf("Invalid result '%s', expected '%s'", out, test)
	}
}

func TestContext_MarshalJSON(t *testing.T) {
	url := "test"
	c := Context{NullTerm: IRI(url)}
	//c.Language = "en-GB"

	out, err := c.MarshalJSON()
	if err != nil {
		t.Errorf("%s", err)
	}
	if !strings.Contains(string(out), url) {
		t.Errorf("Json doesn't contain %#v, %#v", url, string(out))
	}
}
