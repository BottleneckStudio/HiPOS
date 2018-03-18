package test

import (
	"testing"

	"github.com/BottleneckStudio/HiPOS/models"
)

func TestHello(t *testing.T) {
	got := models.Hello()
	want := "Hello, World"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
