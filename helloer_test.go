package Helloer
import "testing"

func TestSayHello(t *testing.T) {
  got := SayHello()
  expected := "hello, world\n"

  if got !=expected {
    t.Errorf("expected %d, got %d", expected, got)
  }
}
