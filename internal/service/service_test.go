package service

import (
	"testing"
	"time"
)

type fixedTime struct{ hour int }

func (f fixedTime) Now() time.Time { return time.Date(2026, 1, 1, f.hour, 0, 0, 0, time.UTC) }

func TestGreet(t *testing.T) {
	for _, tc := range []struct {
		hour int
		want string
	}{
		{0, "Good morning Hadi"}, {11, "Good morning Hadi"},
		{12, "Good afternoon Hadi"}, {17, "Good afternoon Hadi"},
		{18, "Good night Hadi"}, {23, "Good night Hadi"},
	} {
		t.Run(tc.want+time.Date(2026, 1, 1, tc.hour, 0, 0, 0, time.UTC).Format("15"), func(t *testing.T) {
			if got := NewGreeterService(fixedTime{tc.hour}).Greet("Hadi"); got != tc.want {
				t.Fatalf("got %q, want %q", got, tc.want)
			}
		})
	}
}
