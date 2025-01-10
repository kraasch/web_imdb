package main

import (
  "testing"
)

func TestImdbWebRequest(t *testing.T) {
  got := AskImdb("12 Angry Men")
  want := `
> Search title
25 titles found.
> Retrieve rating
First hit: '12 Angry Men' (1957), duration '1h36m'.
Rating: '9.0'.
`

  if got != want {
    t.Errorf("Expected %s, actual %s", want, got)
  }
}
