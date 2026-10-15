package config

import "testing"

func TestConfig(t *testing.T) {
	for _, tc := range []struct{ name, addr, port, want string }{
		{"default", "", "", ":8080"},
		{"port", "", "9090", ":9090"},
		{"address precedence", "127.0.0.1:7070", "9090", "127.0.0.1:7070"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Setenv("CONFIG_ADDR", tc.addr)
			t.Setenv("CONFIG_PORT", tc.port)
			if got := NewConfig().Addr; got != tc.want {
				t.Fatalf("got %q, want %q", got, tc.want)
			}
		})
	}
}
