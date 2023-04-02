package vector

import "testing"

func TestMain(m *testing.M) {

}

func TestPush(t *testing.T) {
	vecInt := New[int](0, 100)
	vecString := New[string](0, 10)
	vecString.Push("This is a test string")
	vecInt.Push(16)
	vecString.Push("Hello", "there")
	vecInt.Push(5, 6, 9, 12, 15)

	if vecString.Slice[1] != "Hello" {
		t.Errorf("item at string vector index 1 is %s but should be Hello", vecString.Slice[1])
	}

	if vecInt.Slice[3] != 9 {
		t.Errorf("item at int vector index 3 is %d but should be 9", vecInt.Slice[3])
	}

	if vecString.Size() != 3 {
		t.Errorf("string vector should have a size of 3 but has a size of %d", vecString.Size())
	}

	if vecInt.Size() != 6 {
		t.Errorf("int vector should have a size of 6 but has a size of %d", vecInt.Size())
	}

}
