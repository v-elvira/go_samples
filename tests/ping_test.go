// -- ping/ping_test.go --
package ping

import (
	"net/http"
	"testing"
)

func TestPing(t *testing.T) {
	client := &http.Client{}
	pinger := Pinger{client}
	got := pinger.Ping("https://example.com")
	if !got {
		t.Errorf("Expected example.com to be available")
	}
	got = pinger.Ping("https://example.com/404")
	if got {
		t.Errorf("Expected example.com/404 to be unavailable")
	}
}

// go mod init github.com/v-elvira/go_samples/tests
// go test -v
