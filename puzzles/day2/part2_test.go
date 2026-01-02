package day2

import "testing"


func TestIsSilly2_122(t *testing.T) {
	got:= silly2(112)

    want := false
    if got != want {
        t.Fatalf("TestIsSilly2_122; got %t; want %t", got, want)
    }
}

func TestIsSilly2_1188511885(t *testing.T) {
	got:= silly2(1188511885)

    want := true
    if got != want {
        t.Fatalf("TestIsSilly2_1188511885; got %t; want %t", got, want)
    }
}
func TestIsSilly2_1188511880(t *testing.T) {
	got:= silly2(1188511880)

    want := false
    if got != want {
        t.Fatalf("TestIsSilly2_1188511880; got %t; want %t", got, want)
    }
}

func TestIsSilly2_123123123(t *testing.T) {
	got:= silly2(123123123)

    want := true
    if got != want {
        t.Fatalf("TestIsSilly2_123123123; got %t; want %t", got, want)
    }
}
func TestIsSilly2_1212121212(t *testing.T) {
	got:= silly2(1212121212)

    want := true
    if got != want {
        t.Fatalf("TestIsSilly2_1212121212; got %t; want %t", got, want)
    }
}


func TestPart2(t *testing.T) {
	got, err := Part2("11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124")

	if (err != nil) {
		t.Fatal(err)
	}

    want := "4174379265"
    if got != want {
        t.Fatalf("TestPart2; got %s; want %s", got, want)
    }
}
