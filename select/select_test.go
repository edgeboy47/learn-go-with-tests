package select_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func makeServer(delay time.Duration) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))

  return server
}

func TestRacer(t *testing.T) {
	slowServer := makeServer(20 * time.Millisecond)
	fastServer := makeServer(0)
  defer slowServer.Close()
  defer fastServer.Close()

	slowUrl := slowServer.URL
	fastUrl := fastServer.URL

	want := fastUrl
	got := Racer(slowUrl, fastUrl)

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

}
