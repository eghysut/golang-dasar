package main

import (
    "fmt"
)

func main() {
    // int8 
    // nilai min = -2^7 = -128
    // nilai max = 2^7 - 1 = 127
    var nilai8 int8 = 127
    fmt.Printf("tipe: %T size: %d\n", nilai8, nilai8)

    // int16
    // nilai min = -2^15 = -327678
    // nilai max = 2^15 - 1 = 32767
    var nilai16 int16 = 32767
    fmt.Printf("tipe: %T size: %d\n", nilai16, nilai16)

    // int32
    // nilai min = -2^31 = -2147483648
    // nilai max = 2^31 - 1 = 2147483647
    var nilai32 int32 = 2147483647
    fmt.Printf("tipe: %T size: %d\n", nilai32, nilai32)

    // int64
    // nilai min = -2^63 = -9223372036854775808
    // nilai max = 2^63 - 7 = 9223372036854775807
    var nilai64 int64 = 9223372036854775807
    nilai64 = int64(nilai64)
    fmt.Printf("tipe: %T size: %d\n", nilai64, nilai64)

    // uint8
    // nilai min = 0
    // nilai max = 2^8 - 1 = 255
    var unilai8 uint8 = 255 
    fmt.Printf("tipe: %T size: %d\n", unilai8, unilai8)

    // uint16
    // nilai min = 0
    // nilai max = 2^16 - 1 = 65535
    var unilai16 uint16 = 65535
    fmt.Printf("tipe: %T size: %d\n", unilai16, unilai16)

    // uint32
    // nilai min = 0
    // nilai max = 2^32 - 1 = 4294967295
    var unilai32 uint32 = 4294967295
    fmt.Printf("tipe: %T size: %d\n", unilai32, unilai32)
}
