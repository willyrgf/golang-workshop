package main

import "fmt"

/** Types
uint8       unsigned  8-bit integers (0 to 255)
uint16      unsigned 16-bit integers (0 to 65535)
uint32      unsigned 32-bit integers (0 to 4294967295)
uint64      unsigned 64-bit integers (0 to 18446744073709551615)
int8        signed  8-bit integers (-128 to 127)
int16       signed 16-bit integers (-32768 to 32767)
int32       signed 32-bit integers (-2147483648 to 2147483647)
int64       signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

float32     IEEE-754 32-bit floating-point numbers
float64     IEEE-754 64-bit floating-point numbers
complex64   complex numbers with float32 real and imaginary parts
complex128  complex numbers with float64 real and imaginary parts

uint		unsigned, either 32 or 64 bits
int      	signed, either 32 or 64 bits
uintptr  	unsigned integer large enough to store the uninterpreted bits of a pointer value

bool		the boolean data type can be one of two values, either true or false

string		a string is a sequence of one or more characters (letters, numbers, symbols)
byte        alias for uint8
rune        alias for int32

*/

func main() {
	var UnsignedInt uint = 10
	var SignedInt int64 = -10
	var OneByte byte = 'b'
	var BigFloat float64 = 0.0000001391623098102830810298
	var HelloWorld string = "Hello World"
	var ItsFalse bool = false

	fmt.Printf("UnsignedInt:\t type=%T \t value=%+v\n", UnsignedInt, UnsignedInt)
	fmt.Printf("SignedInt:\t type=%T \t value=%+v\n", SignedInt, SignedInt)
	fmt.Printf("OneByte:\t type=%T \t value=%+v\n", OneByte, OneByte)
	fmt.Printf("BigFloat:\t type=%T \t value=%.16f\n", BigFloat, BigFloat)
	fmt.Printf("HelloWorld:\t type=%T \t value=%+v\n", HelloWorld, HelloWorld)
	fmt.Printf("ItsFalse:\t type=%T \t value=%+v\n", ItsFalse, ItsFalse)
}
