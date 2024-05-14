package injector

import "testing"

func TestNew(t *testing.T) {
	i := New()
	if i == nil {
		t.Fatal("injector is nil")
	}
	i.Register("key", "value")
	_, ok := i.Load("key")
	if !ok {
		t.Fatal("key is not found")
	}

	i = New()
	if i == nil {
		t.Fatal("injector is nil")
	}
	_, ok = i.Load("key")
	if !ok {
		t.Fatal("key is found")
	}
}

func TestInjector_Inject(t *testing.T) {
	var test struct {
		Name string `inject:"name"`
	}

	i := New()
	i.Register("name", "jesse")
	i.Inject(&test)
	if test.Name != "jesse" {
		t.Fatalf("test.Name is %s, not jesse", test.Name)
	}
}
