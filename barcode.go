package barcode

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math/big"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/oned"
	"golang.org/x/exp/rand"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

// Encode takes the message to be encode to a barcode image
// since we are working with a barcode that normally works with 13
// characters we check the length of the message to be equal to 13 characters
// it creates an image
func Encode(msg, filename string) error {
	if len(msg) != 13 {
		return fmt.Errorf("require name to be 13 characters")
	}
	// Generate a random barcode number as before
	filename = fmt.Sprintf("%s.png", filename)
	writer := oned.NewCode128Writer()
	img, err := writer.Encode(msg, gozxing.BarcodeFormat_CODE_128, 500, 100, nil) // Increased size
	if err != nil {
		return fmt.Errorf("impossible to encode barcode: %s", err)
	}

	// Add text to the barcode image
	imgWithText, err := addTextToImage(img, msg)
	if err != nil {
		return fmt.Errorf("impossible to add text to barcode image: %s", err)
	}

	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("impossible to create file: %s", err)
	}
	defer file.Close()
	err = png.Encode(file, imgWithText)
	if err != nil {
		return fmt.Errorf("impossible to encode barcode in PNG: %s", err)
	}
	return nil
}

// Decode takes an file name without the .png extension and
// decodes the information on the barcode
func Decode(filename string) (string, error) {
	filename = fmt.Sprintf("%s.png", filename)
	// Open the image file
	file, err := os.Open(filename)
	if err != nil {
		return "", fmt.Errorf("failed to open the image : %s", err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return "", fmt.Errorf("failed to decode the image : %s", err)
	}

	bmp, err := gozxing.NewBinaryBitmapFromImage(img)
	if err != nil {
		return "", fmt.Errorf("fail to generate bmp : %s", err)
	}

	// Create a Code128 barcode reader
	code128Reader := oned.NewCode128Reader()

	// Decode the barcode
	result, err := code128Reader.Decode(bmp, nil)
	if err != nil {
		return "", fmt.Errorf("failed to decode the image : %s", err)
	}
	return result.GetText(), nil
}

// addTextToImage add the number at the bottom of the image
// its a helper function
func addTextToImage(img image.Image, text string) (image.Image, error) {
	bounds := img.Bounds()
	newHeight := bounds.Dy() + 30 // Increased height for more space
	newImg := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), newHeight))

	// Fill the new image with white background
	draw.Draw(newImg, newImg.Bounds(), &image.Uniform{color.White}, image.Point{}, draw.Src)

	// Draw the original barcode image onto the new image
	draw.Draw(newImg, bounds, img, image.Point{}, draw.Src)

	// Draw the text below the barcode
	col := color.Black
	point := fixed.Point26_6{
		X: (fixed.I(bounds.Dx()) - fixed.I(len(text)*7)) / 2,
		Y: fixed.I(bounds.Dy() + 20),
	}
	d := &font.Drawer{
		Dst:  newImg,
		Src:  image.NewUniform(col),
		Face: basicfont.Face7x13,
		Dot:  point,
	}
	d.DrawString(text)

	return newImg, nil
}

// generates a string of 13 numbers
// using uuid to generate unique number
// then hash it into 13 code number
func GenerateBarCodeNumber1() string {
	newUUID := uuid.New()
	// Convert the UUID to a string
	uuid := newUUID.String()
	uuid = strings.ReplaceAll(uuid, "-", "")

	// Compute MD5 hash
	hash := md5.Sum([]byte(uuid))

	// Convert hash to hexadecimal string
	hashString := hex.EncodeToString(hash[:])

	// Convert hexadecimal to decimal
	intVal, _ := new(big.Int).SetString(hashString, 16)

	code := intVal.Mod(intVal, big.NewInt(10000000000000))
	return code.String()
}
func Get13BarCodeNumber() string {
	code := GenerateBarCodeNumber1()
	if len(code) != 13 {
		fmt.Println("not long enough hit!")
		return GenerateBarCodeNumber1()
	} else {
		return code
	}
}

func GenerateBarCodeNumber() string {
	num := time.Now().UnixNano()
	rand.Seed(uint64(num))
	min := int64(100000000000)
	max := int64(999999999999)
	code := rand.Int63n(max-min+1) + min
	return strconv.FormatInt(code, 10)
}
