package day3

import "testing"

func TestPart2(t *testing.T) {
	got, err := Part2("987654321111111\n811111111111119\n234234234234278\n818181911112111")

	if err != nil {
		t.Fatal(err)
	}

	want := "3121910778619"
	if got != want {
		t.Fatalf("TestPart2; got %s; want %s", got, want)
	}
}

func TestJoltage2_987654321111111(t *testing.T) {
	got, err := joltage("987654321111111")

	if err != nil {
		t.Fatal(err)
	}

	want := 987654321111
	if got != want {
		t.Fatalf("TestJoltage2_987654321111111; got %d; want %d", got, want)
	}
}

func TestJoltage2_811111111111119(t *testing.T) {
	got, err := joltage("811111111111119")

	if err != nil {
		t.Fatal(err)
	}

	want := 811111111119
	if got != want {
		t.Fatalf("TestJoltage2_811111111111119; got %d; want %d", got, want)
	}
}

func TestJoltage2_234234234234278(t *testing.T) {
	got, err := joltage("234234234234278")

	if err != nil {
		t.Fatal(err)
	}

	want := 434234234278
	if got != want {
		t.Fatalf("TestJoltage2_234234234234278; got %d; want %d", got, want)
	}
}

func TestJoltage2_818181911112111(t *testing.T) {
	got, err := joltage("818181911112111")

	if err != nil {
		t.Fatal(err)
	}

	want := 888911112111
	if got != want {
		t.Fatalf("TestJoltage2_818181911112111; got %d; want %d", got, want)
	}
}
