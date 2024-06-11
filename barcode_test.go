package barcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBarCodeValidity(t *testing.T) {
	msg := GenerateBarCodeNumber()
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
			filename := "barcode.png"
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
