package day2

import "testing"

func TestMakeArr(t *testing.T) {
	got, err := makeArr("11", "22")

	if (err != nil) {
		t.Fatal(err)
	}

    want :=  []int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22}

	for i := 0; i < len(want); i++ {
		if got[i] != want[i] {
			t.Fatalf("TestPart1; got %d; want %d for i %d", got[i], want[i], i)
		}
	}
    
}

func TestIsSilly122(t *testing.T) {
	got:= silly(112)

    want := false
    if got != want {
        t.Fatalf("TestIsSilly122; got %t; want %t", got, want)
    }
}

func TestIsSilly1188511885(t *testing.T) {
	got:= silly(1188511885)

    want := true
    if got != want {
        t.Fatalf("TestIsSilly1188511885; got %t; want %t", got, want)
    }
}
func TestIsSilly1188511880(t *testing.T) {
	got:= silly(1188511880)

    want := false
    if got != want {
        t.Fatalf("TestIsSilly1188511880; got %t; want %t", got, want)
    }
}


func TestPart1(t *testing.T) {
	got, err := Part1("11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862")

	if (err != nil) {
		t.Fatal(err)
	}

    want := "1227775554"
    if got != want {
        t.Fatalf("TestPart1; got %s; want %s", got, want)
    }
}
