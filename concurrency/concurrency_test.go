package concurrency

import (
  "reflect"
  "testing"
)

func mockWebsiteChecker(website string) bool {
  return true
}

func TestCheckWebsites(t *testing.T) {
  websites := []string {
    "https://google.com",
    "https://youtube.com",
  }

  want := map[string]bool {
    "https://google.com": true,
    "https://youtube.com": true,
  }

  

  got := CheckWebsites(mockWebsiteChecker, websites)

  if (!reflect.DeepEqual(got, want)) {
    t.Fatalf("wanted %v, got %v", want, got)
  }
}
