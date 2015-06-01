package cowmap

import "testing"

func TestAll(t *testing.T) {
	m := New(map[Key]Value{
		"a": "1",
		"b": "2",
		"c": "3",
	})
	m.Set("d", "4")
	if m.Get("d").(string) != "4" {
		t.Fail()
	}
	if _, ok := m.Get2("e"); ok {
		t.Fail()
	}
	m.Del("d")
	if _, ok := m.Get2("d"); ok {
		t.Fail()
	}
}
