# Barcode Package

The Barcode Package is a simple yet powerful library written in Golang designed to handle barcode generation and decoding. This package provides functionalities to generate a 13-number barcode string, create barcode images, and decode barcode information from images.

## Features

- **Generate 13-digit barcode strings**
- **Encode barcode strings into images**
- **Decode barcode images to retrieve information**

## Installation

To install the Barcode Package, use `go get`:

```sh
go get github.com/yourusername/barcode
```

## Usage

First, import the package in your Go file:

```go
import "github.com/yourusername/barcode"
```

### Generate Barcode String

The `GenerateBarCodeNumber` method generates a random 13-digit barcode string.

```go
barcodeString := barcode.GenerateBarCodeNumber()
fmt.Println("Generated Barcode String:", barcodeString)
```

### Encode Barcode String to Image

The `Encode` method encodes a 13-digit barcode string into an image and saves it as a PNG file.

```go
barcodeString := barcode.GenerateBarCodeNumber()
err := barcode.Encode(barcodeString, "barcode_image")
if err != nil {
    // handle error
    fmt.Println("Error:", err)
} else {
    fmt.Println("Barcode image created successfully")
}
```

### Decode Barcode Image

The `Decode` method reads a barcode image and returns the encoded information.

```go
decodedString, err := barcode.Decode("barcode_image")
if err != nil {
    // handle error
    fmt.Println("Error:", err)
} else {
    fmt.Println("Decoded Barcode String:", decodedString)
}
```

## Methods

### `GenerateBarCodeNumber`

Generates a random 13-digit barcode string.

**Signature:**

```go
func GenerateBarCodeNumber() string
```

### `Encode`

Encodes a 13-digit barcode string into an image and saves it as a PNG file.

**Signature:**

```go
func Encode(msg, filename string) error
```

**Parameters:**

- `msg` (string): The 13-digit barcode string to encode.
- `filename` (string): The name of the file (without extension) to save the barcode image.

**Returns:**

- `error`: An error if the encoding or file creation fails, otherwise nil.

### `Decode`

Reads a barcode image and decodes it into a string.

**Signature:**

```go
func Decode(filename string) (string, error)
```

**Parameters:**

- `filename` (string): The name of the file (without extension) to read the barcode image from.

**Returns:**

- `string`: The decoded barcode string.
- `error`: An error if the decoding or file reading fails, otherwise nil.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any changes.


## Author

- **Myrachanto**
- [GitHub](https://github.com/myrachanto)
- [Email](mailto:myrachanto@gmail.com)

## Acknowledgments

- Inspired by various open-source barcode libraries.
- Special thanks to the Golang community for their support and contributions.