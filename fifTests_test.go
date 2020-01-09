package main

import (
	"fmt"
	"testing"
)

// Los tests:

// TestOk si la entrada y salida esta ok
func TestOk(t *testing.T) {

	res, _ := tlvParser([]byte("11A05AB398765UJ102N2300"))
	fmt.Println(res)

	if res == nil {
		// t.Errorf("tlvParser([]byte('11A05AB398765UJ102N2300')) = %d; want -2", res)
		t.Fail()
	}
}

// TestVacio si la entrada esta vacia
func TestVacio(t *testing.T) {

	res, _ := tlvParser([]byte(""))
	fmt.Println(res)

	if res != nil {
		// t.Errorf("tlvParser([]byte('')) = %d; want -2", res)
		t.Fail()
	}
}

// TestNulo si la entrada esta nula
func TestNulo(t *testing.T) {

	res, _ := tlvParser([]byte(nil))
	fmt.Print(res)

	if res != nil {
		// t.Errorf("tlvParser([]byte(nil)) = %d; want nil", res)
		t.Fail()
	}

}

// TestCamposInvalidos campos son invalidos
func TestCamposInvalidos(t *testing.T) {

	res, _ := tlvParser([]byte("112A05AB398765UJ102N2300"))
	fmt.Println(res)
	if res != nil {
		// t.Errorf("tlvParser([]byte('112A05AB398765UJ102N2300')) = %d; want -2", res)
		t.Fail()
	}
}

// TestCorrupto si la entrada esta corrupta
func TestCorrupto(t *testing.T) {
	res, _ := tlvParser([]byte("1sds1sdf2A05AB398765dsfdsfUJ102N2sdf300"))
	fmt.Println(res)
	if res != nil {
		// t.Errorf("tlvParser([]byte('1sds1sdf2A05AB398765dsfdsfUJ102N2sdf300')) = %d; want -2", res)
		t.Fail()
	}
}
