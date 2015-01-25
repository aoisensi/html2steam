package html2steam

import (
	"fmt"
	"os"
	"testing"
)

func TestH2S(t *testing.T) {
	f, err := os.Open("test.html")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	body, err := Replace(f)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(body)
}
