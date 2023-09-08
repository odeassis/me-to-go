package main

import "testing"

func TestHello(t *testing.T) {
  result := hello()
  expect := "Hello Word"

  if result != expect {
    t.Errorf("result '%s', expect '%s'", result, expect)
  }
}
