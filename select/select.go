package select_test

import (
	"net/http"
	"time"
)

func measureResponseTime(url string) time.Duration {
	start := time.Now()
  http.Get(url)
  return time.Since(start)
}

func Racer(url1, url2 string) (winner string) {
  aDuration := measureResponseTime(url1)
  bDuration := measureResponseTime(url2)

  if aDuration < bDuration {
    return url1
  }

  return url2
}
