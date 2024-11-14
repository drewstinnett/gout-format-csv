package csv

import (
	"testing"
)

func TestCSVFormatter(t *testing.T) {
	f := Formatter{}
	got, err := f.Format([]struct {
		Foo string
	}{
		{Foo: "bar"},
		{Foo: "baz"},
	})
	if err != nil {
		t.Error(err)
	}

	want := "Foo\nbar\nbaz\n"
	if string(got) != want {
		t.Errorf(`Did not get back expected result:
Wanted: %v
Got: %v
`, want, string(got))
	}
}

func TestCSVFormatterCustomHeader(t *testing.T) {
	got, err := Formatter{
		Headers: []string{"custom"},
	}.Format([]struct {
		Foo string
	}{
		{Foo: "bar"},
		{Foo: "baz"},
	})
	if err != nil {
		t.Error(err)
	}

	want := "custom\nbar\nbaz\n"
	if string(got) != want {
		t.Errorf(`Did not get back expected result:
Wanted: %v
Got: %v
`, want, string(got))
	}
}
