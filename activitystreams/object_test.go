package activitystreams

import (
	"reflect"
	"testing"
)

func TestObjectNew(t *testing.T) {
	var testValue = ObjectID("test")
	var testType = ArticleType

	o := ObjectNew(testType)
	o.ID = testValue

	if o.ID != testValue {
		t.Errorf("APObject Id '%v' different than expected '%v'", o.ID, testValue)
	}
	if o.Type != testType {
		t.Errorf("APObject Type '%v' different than expected '%v'", o.Type, testType)
	}

	n := ObjectNew("")
	n.ID = testValue
	if n.ID != testValue {
		t.Errorf("APObject Id '%v' different than expected '%v'", n.ID, testValue)
	}
	if n.Type != ObjectType {
		t.Errorf("APObject Type '%v' different than expected '%v'", n.Type, ObjectType)
	}

}

func TestValidGenericType(t *testing.T) {
	for _, validType := range validGenericObjectTypes {
		if !ValidObjectType(validType) {
			t.Errorf("Generic Type '%v' should be valid", validType)
		}
	}
}

func TestValidObjectType(t *testing.T) {
	var invalidType ActivityVocabularyType = "RandomType"

	if ValidObjectType(invalidType) {
		t.Errorf("APObject Type '%v' should not be valid", invalidType)
	}
	for _, validType := range validObjectTypes {
		if !ValidObjectType(validType) {
			t.Errorf("APObject Type '%v' should be valid", validType)
		}
	}
}

func TestMarshalJSON(t *testing.T) {
	m := NaturalLanguageValue{
		{
			"en", "test",
		},
		{
			"de", "test",
		},
	}
	result, err := m.MarshalJSON()
	if err != nil {
		t.Errorf("Failed marshaling '%v'", err)
	}
	mRes := "{\"de\":\"test\",\"en\":\"test\"}"
	if string(result) != mRes {
		t.Errorf("Different results '%v' vs. '%v'", string(result), mRes)
	}
	//n := NaturalLanguageValueNew()
	//result, err := n.MarshalJSON()

	s := make(map[LangRef]string)
	s["en"] = "test"
	n1 := NaturalLanguageValue{{
		"en", "test",
	}}
	result1, err1 := n1.MarshalJSON()
	if err1 != nil {
		t.Errorf("Failed marshaling '%v'", err1)
	}
	mRes1 := "\"test\""
	if string(result1) != mRes1 {
		t.Errorf("Different results '%v' vs. '%v'", string(result1), mRes1)
	}
}

func TestNaturalLanguageValue_MarshalJSON(t *testing.T) {
	p := NaturalLanguageValue{
		{
			"en", "the test",
		},
		{
			"fr", "le test",
		},
	}
	js := "{\"en\":\"the test\",\"fr\":\"le test\"}"
	out, err := p.MarshalJSON()

	if err != nil {
		t.Errorf("Error: '%s'", err)
	}
	if js != string(out) {
		t.Errorf("Different marshal result '%s', instead of '%s'", out, js)
	}
	p1 := NaturalLanguageValue{
		{
			"en", "the test",
		},
	}

	out1, err1 := p1.MarshalJSON()

	if err1 != nil {
		t.Errorf("Error: '%s'", err1)
	}
	txt := "\"the test\""
	if txt != string(out1) {
		t.Errorf("Different marshal result '%s', instead of '%s'", out1, txt)
	}
}

func TestObject_IsLink(t *testing.T) {
	o := ObjectNew(ObjectType)
	o.ID = "test"
	if o.IsLink() {
		t.Errorf("%#v should not be a valid link", o.Type)
	}
	m := ObjectNew(AcceptType)
	m.ID = "test"
	if m.IsLink() {
		t.Errorf("%#v should not be a valid link", m.Type)
	}
}

func TestObject_IsObject(t *testing.T) {
	o := ObjectNew(ObjectType)
	o.ID = "test"
	if !o.IsObject() {
		t.Errorf("%#v should be a valid object", o.Type)
	}
	m := ObjectNew(AcceptType)
	m.ID = "test"
	if !m.IsObject() {
		t.Errorf("%#v should be a valid object", m.Type)
	}
}

func TestObjectsArr_Append(t *testing.T) {
	d := make(ItemCollection, 0)

	val := Object{ID: ObjectID("grrr")}

	d.Append(val)

	if len(d) != 1 {
		t.Errorf("Objects array should have exactly an element")
	}
	if !reflect.DeepEqual(d[0], val) {
		t.Errorf("First item in object array does not match %q", val.ID)
	}
}

func TestRecipientsDeduplication(t *testing.T) {
	bob := PersonNew("bob")
	alice := PersonNew("alice")
	foo := OrganizationNew("foo")
	bar := GroupNew("bar")

	first := make(ItemCollection, 0)
	if len(first) != 0 {
		t.Errorf("Objects array should have exactly an element")
	}

	first.Append(bob)
	first.Append(alice)
	first.Append(foo)
	first.Append(bar)
	if len(first) != 4 {
		t.Errorf("Objects array should have exactly 4(four) elements, not %d", len(first))
	}

	first.Append(bar)
	first.Append(alice)
	first.Append(foo)
	first.Append(bob)
	if len(first) != 8 {
		t.Errorf("Objects array should have exactly 8(eight) elements, not %d", len(first))
	}

	recipientsDeduplication(&first)
	if len(first) != 4 {
		t.Errorf("Objects array should have exactly 4(four) elements, not %d", len(first))
	}

	second := make(ItemCollection, 0)
	second.Append(bar)
	second.Append(foo)

	recipientsDeduplication(&first, &second)
	if len(first) != 4 {
		t.Errorf("First Objects array should have exactly 8(eight) elements, not %d", len(first))
	}
	if len(second) != 0 {
		t.Errorf("Second Objects array should have exactly 0(zero) elements, not %d", len(second))
	}

	err := recipientsDeduplication(&first, &second, nil)
	if err != nil {
		t.Errorf("Deduplication with empty array failed")
	}
}

func TestNaturalLanguageValue_Get(t *testing.T) {
	testVal := "test"
	a := NaturalLanguageValue{{NilLangRef, testVal}}
	if a.Get(NilLangRef) != testVal {
		t.Errorf("Invalid Get result. Expected %s received %s", testVal, a.Get(NilLangRef))
	}
}

func TestNaturalLanguageValue_Set(t *testing.T) {
	testVal := "test"
	a := NaturalLanguageValue{{NilLangRef, "ana are mere"}}
	err := a.Set(LangRef("en"), testVal)
	if err != nil {
		t.Errorf("Received error when doing Set %s", err)
	}
}

func TestNaturalLanguageValue_Append(t *testing.T) {
	var a NaturalLanguageValue

	if len(a) != 0 {
		t.Errorf("Invalid initialization of %T. Size %d > 0 ", a, len(a))
	}
	langEn := LangRef("en")
	valEn := "random value"

	a.Append(langEn, valEn)
	if len(a) != 1 {
		t.Errorf("Invalid append of one element to %T. Size %d != 1", a, len(a))
	}
	if a.Get(langEn) != valEn {
		t.Errorf("Invalid append of one element to %T. Value of %q not equal to %q, but %q", a, langEn, valEn, a.Get(langEn))
	}
	langDe := LangRef("de")
	valDe := "randomisch"
	a.Append(langDe, valDe)

	if len(a) != 2 {
		t.Errorf("Invalid append of one element to %T. Size %d != 2", a, len(a))
	}
	if a.Get(langEn) != valEn {
		t.Errorf("Invalid append of one element to %T. Value of %q not equal to %q, but %q", a, langEn, valEn, a.Get(langEn))
	}
	if a.Get(langDe) != valDe {
		t.Errorf("Invalid append of one element to %T. Value of %q not equal to %q, but %q", a, langDe, valDe, a.Get(langDe))
	}
}

func TestLangRef_UnmarshalJSON(t *testing.T) {
	lang := "en-US"
	json := `"` + lang + `"`

	var a LangRef
	a.UnmarshalJSON([]byte(json))

	if string(a) != lang {
		t.Errorf("Invalid json unmarshal for %T. Expected %q, found %q", lang, lang, string(a))
	}
}

func TestNaturalLanguageValue_UnmarshalFullObjectJSON(t *testing.T) {
	langEn := "en-US"
	valEn := "random"
	langDe := "de-DE"
	valDe := "zufällig\\n"

	//m := make(map[string]string)
	//m[langEn] = valEn
	//m[langDe] = valDe

	json := `{
		"` + langEn + `": "` + valEn + `",
		"` + langDe + `": "` + valDe + `"
	}`

	var a NaturalLanguageValue
	a.Append(LangRef(langEn), valEn)
	a.Append(LangRef(langDe), valDe)
	err := a.UnmarshalJSON([]byte(json))
	if err != nil {
		t.Error(err)
	}
	for lang, val := range a {
		if val.Ref != LangRef(langEn) && val.Ref != LangRef(langDe) {
			t.Errorf("Invalid json unmarshal for %T. Expected lang %q or %q, found %q", a, langEn, langDe, lang)
		}

		if val.Ref == LangRef(langEn) && val.Value != valEn {
			t.Errorf("Invalid json unmarshal for %T. Expected value %q, found %q", a, valEn, val)
		}
		if val.Ref == LangRef(langDe) && val.Value != valDe {
			t.Errorf("Invalid json unmarshal for %T. Expected value %q, found %q", a, valDe, val)
		}
	}
}