package sh

import "testing"

func TestCat(t *testing.T) {
	if err := Run("echo 'test' > tmp.txt"); err != nil {
		t.Fatal((err))
	}

	result, err := RunR("cat tmp.txt")
	if err != nil {
		t.Fatal((err))
	}
	if result != "test" {
		t.Fatal("failed test: cat")
	}
	if err := Run("rm tmp.txt"); err != nil {
		t.Fatal((err))
	}
}
