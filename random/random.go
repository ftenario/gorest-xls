package random

import (
  "math/rand"
  "time"
)

var r *rand.Rand
func init() {
  r = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func Generate(strlen int) string {
  const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
  result := ""

  for i:=0; i<strlen; i++ {
    index:=r.Intn(len(chars))
    result += chars[index: index + 1]
  }
  return result
}
