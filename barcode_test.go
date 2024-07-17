package barcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestGenerateBarCodeNumber checks if 100 generated barcodes are 13 characters long
func TestGenerateBarCodeNumber(t *testing.T) {
	for i := 0; i < 100; i++ {
		barcode := GenerateBarCodeNumber()
		if len(barcode) != 12 {
			t.Errorf("Barcode %s is not 12 characters long", barcode)
		}
	}
}

// TestGenerateBarCodeNumber checks if 100 generated barcodes are 13 characters long
func TestGet13BarCodeNumber(t *testing.T) {
	for i := 0; i < 100; i++ {
		barcode := Get13BarCodeNumber()
		if len(barcode) != 13 {
			t.Errorf("Barcode %s is not 13 characters long", barcode)
		}
	}
}
func TestBarCodeValidity(t *testing.T) {
	msg := Get13BarCodeNumber()
	tests := []struct {
		testName   string
		msg        string
		result     string
		errMessage error
	}{
		{testName: "ok", msg: msg, result: msg, errMessage: nil},
		{testName: "Not_Ok", msg: "", result: "", errMessage: fmt.Errorf("require name to be 13 characters")},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			filename := "barcode"
			err := Encode(test.msg, filename)
			require.EqualValues(t, test.errMessage, err)
			if err == nil {
				decodedBarcodeNumber, err := Decode(filename)
				require.NoError(t, err)
				require.EqualValues(t, test.result, decodedBarcodeNumber)
			}
		})
	}
}
