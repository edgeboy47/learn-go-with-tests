package concurrency

import (
  "reflect"
  "testing"
  "time"
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

func slowWebsiteChecker(_ string) bool {
  time.Sleep(20 * time.Millisecond)
  return true
}

func BenchmarkCheckWebsites(b *testing.B) {
  urls := make([]string, 100)
  for i := 0; i < len(urls); i++ {
    urls[i] = "a url"
  }
  b.ResetTimer()

  for i := 0; i < b.N; i++ {
    CheckWebsites(slowWebsiteChecker, urls)
  }
}
