package day3

import "testing"

func TestPart1(t *testing.T) {
	got, err := Part1("987654321111111\n811111111111119\n234234234234278\n818181911112111")

	if err != nil {
		t.Fatal(err)
	}

	want := "357"
	if got != want {
		t.Fatalf("TestPart1; got %s; want %s", got, want)
	}
}
