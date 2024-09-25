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
	t.Run("compares speeds of servers, returning the url of the fastest", func(t *testing.T) {

		slowServer := makeServer(20 * time.Millisecond)
		fastServer := makeServer(0)
		defer slowServer.Close()
		defer fastServer.Close()

		slowUrl := slowServer.URL
		fastUrl := fastServer.URL

		want := fastUrl
		got, _ := Racer(slowUrl, fastUrl, 10 * time.Second)

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

	})

	t.Run("returns an error if a server doesn't response within 10s", func(t *testing.T) {
		serverA := makeServer(11 * time.Second)
		serverB := makeServer(12 * time.Second)
		defer serverA.Close()
		defer serverB.Close()
		_, err := Racer(serverA.URL, serverB.URL, 10 * time.Second)

    if err == nil {
      t.Error("expected an error but didn't get one")
    }
	})
}
