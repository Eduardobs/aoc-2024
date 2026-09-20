package tests

import "testing"

// CheckNumbers validates the common two-part numeric result returned by most days.
func CheckNumbers(t testing.TB, got1, got2, want1, want2 int64, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
	if got1 != want1 || got2 != want2 {
		t.Fatalf("got (%d, %d), want (%d, %d)", got1, got2, want1, want2)
	}
}
