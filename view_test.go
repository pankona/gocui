package gocui

import "testing"

func requireErrNotNil(t *testing.T, err error) {
	t.Helper()
	if err == nil {
		t.Fatalf("error should NOT be nil")
	}
}

func requireErrNil(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("unexpected err returned: %v", err)
	}
}

func TestRealPosition(t *testing.T) {
	v := View{}

	_, _, err := v.realPosition(-1, -1)
	requireErrNotNil(t, err)

	rx, ry, err := v.realPosition(0, 0)
	requireErrNil(t, err)
	if rx != 0 || ry != 0 {
		t.Fatalf("unexpected result: [got] %d, %d [want] %d, %d",
			rx, ry, 0, 0)
	}

	v.viewLines = make([]viewLine, 10)
	rx, ry, err = v.realPosition(0, 0)
	requireErrNil(t, err)
	if rx != 0 || ry != 0 {
		t.Fatalf("unexpected result: [got] %d, %d [want] %d, %d",
			rx, ry, 0, 0)
	}

	linesX := 10
	v.viewLines[0].linesX = linesX

	rx, ry, err = v.realPosition(0, 0)
	requireErrNil(t, err)
	if rx != linesX || ry != 0 {
		t.Fatalf("unexpected result: [got] %d, %d [want] %d, %d",
			rx, ry, linesX, 0)
	}

	linesY := 10
	v.viewLines[0].linesY = linesX

	rx, ry, err = v.realPosition(0, 0)
	requireErrNil(t, err)
	if rx != linesX || ry != linesY {
		t.Fatalf("unexpected result: [got] %d, %d [want] %d, %d",
			rx, ry, linesX, linesY)
	}

	vx := 10
	rx, ry, err = v.realPosition(vx, 0)
	requireErrNil(t, err)
	if rx != linesX+vx || ry != linesY {
		t.Fatalf("unexpected result: [got] %d, %d [want] %d, %d",
			rx, ry, linesX+vx, linesY)
	}

	vy := 9
	rx, ry, err = v.realPosition(vx, vy)
	requireErrNil(t, err)
	if rx != vx || ry != 0 {
		t.Fatalf("unexpected result: [got] %d, %d [want] %d, %d",
			rx, ry, vx, 0)
	}

	/*
		vy = 10
		rx, ry, err = v.realPosition(0, vy)
		requireErrNil(t, err)
		if rx != vx || ry != 0 {
			t.Fatalf("unexpected result: [got] %d, %d [want] %d, %d",
				rx, ry, 0, 0)
		}
	*/
}
