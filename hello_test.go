package main
import(
	"testing"
)

func TestHello(t *testing.T){
	got:=Hello()
	want:="Hello World"
	if got != want{
		t.Errorf("Got: %q - Want: %q", got, want)
	}
}