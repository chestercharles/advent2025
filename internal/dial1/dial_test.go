package dial1

import "testing"

func TestInitDial(t *testing.T) {
	myDial := NewDial()
    got := myDial.Read()
    want := 50
    if got != want {
        t.Fatalf("Add(2, 3) = %d; want %d", got, want)
    }
}


func TestDialTurnRight1(t *testing.T) {
	myDial := NewDial()
    myDial.TurnRight(1)
	got := myDial.Read()
    want := 51
    if got != want {
        t.Fatalf("TestDialTurnRight1; got %d; want %d", got, want)
    }
}

func TestDialTurnRight48(t *testing.T) {
	myDial := NewDial()
    myDial.TurnRight(48)
	got := myDial.Read()
    want := 98
    if got != want {
        t.Fatalf("TestDialTurnRight48; got %d; want %d", got, want)
    }
}

func TestDialTurnRight49(t *testing.T) {
	myDial := NewDial()
    myDial.TurnRight(49)
	got := myDial.Read()
    want := 99
    if got != want {
        t.Fatalf("TestDialTurnRight49; got %d; want %d", got, want)
    }
}

func TestDialTurnRight50(t *testing.T) {
	myDial := NewDial()
    myDial.TurnRight(50)
	got := myDial.Read()
    want := 0
    if got != want {
        t.Fatalf("TestDialTurnRight50; got %d; want %d", got, want)
    }
}

func TestDialTurnLeft1(t *testing.T) {
	myDial := NewDial()
    myDial.TurnLeft(1)
	got := myDial.Read()
    want := 49
    if got != want {
        t.Fatalf("TestDialTurnLeft1; got %d; want %d", got, want)
    }
}

func TestDialTurnLeft49(t *testing.T) {
	myDial := NewDial()
    myDial.TurnLeft(49)
	got := myDial.Read()
    want := 1
    if got != want {
        t.Fatalf("TestDialTurnLeft49; got %d; want %d", got, want)
    }
}

func TestDialTurnLeft50(t *testing.T) {
	myDial := NewDial()
    myDial.TurnLeft(50)
	got := myDial.Read()
    want := 0
    if got != want {
        t.Fatalf("TestDialTurnLeft50; got %d; want %d", got, want)
    }
}

func TestDialTurnLeft51(t *testing.T) {
	myDial := NewDial()
    myDial.TurnLeft(51)
	got := myDial.Read()
    want := 99
    if got != want {
        t.Fatalf("TestDialTurnLeft51; got %d; want %d", got, want)
    }
}