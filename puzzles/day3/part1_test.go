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

func TestJoltage123456789(t *testing.T) {
	got, err := joltage("123456789")

	if err != nil {
		t.Fatal(err)
	}

	want := 89
	if got != want {
		t.Fatalf("TestJoltage123456789; got %d; want %d", got, want)
	}
}

func TestJoltage123456798(t *testing.T) {
	got, err := joltage("123456798")

	if err != nil {
		t.Fatal(err)
	}

	want := 98
	if got != want {
		t.Fatalf("TestJoltage123456798; got %d; want %d", got, want)
	}
}

func TestJoltage987654321111111(t *testing.T) {
	got, err := joltage("987654321111111")

	if err != nil {
		t.Fatal(err)
	}

	want := 98
	if got != want {
		t.Fatalf("TestJoltage987654321111111; got %d; want %d", got, want)
	}
}

func TestJoltage811111111111119(t *testing.T) {
	got, err := joltage("811111111111119")

	if err != nil {
		t.Fatal(err)
	}

	want := 89
	if got != want {
		t.Fatalf("TestJoltage811111111111119; got %d; want %d", got, want)
	}
}
