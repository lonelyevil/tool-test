package assets

import "testing"

func TestGenName(t *testing.T) {
	if GenName() != "Baloo dev(fooooo)" {
		t.Fatal(GenName())
	}
}
