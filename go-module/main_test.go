package main

import "testing"

func TestSayThisIsBasic(t *testing.T) {
	actual := SayThisIsBasic()
	expect := "This is basic go"

	if actual != expect {
		t.Errorf("FAIL: expect `This is basic go`")
	}
	t.Logf("SUCCESS: expect `This is basic go`")
}
