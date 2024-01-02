package types

import (
	"fmt"
	"testing"
)

func TestDecimal(t *testing.T) {
	n1 := Int2Decimal(1, 0)
	fmt.Println(n1)
	fmt.Println(DivideDecimal(n1, n1))
}

func TestDecimal2(t *testing.T) {
	n1 := Int2Decimal(1, 0)
	fmt.Println(n1.String())
}

func TestDecimal3(t *testing.T) {
	src := "-1.12345678901234567890"
	d, err := Text2Decimal(src)
	fmt.Println(d, err)
}

func TestDecimal4(t *testing.T) {
	src := "1.12345678901234567890"
	d, _ := Text2Decimal(src)
	a := d.StringScale(0)
	fmt.Println(a)
}
func TestDecimal5(t *testing.T) {
	src := "-1.4275782073449123754729160278752640754"
	d, _ := Text2Decimal(src)
	a := d.String()
	fmt.Println(a)
}
