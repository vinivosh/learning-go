package main
import "fmt"
import "math"

////////////////////////////////////////////////////////////////////////////////

func main(){
	var num_16_bit int16 = 32767
	num_16_bit = num_16_bit + 1  // Overflow error! Only show in compile time
	fmt.Println(num_16_bit)  // Shows as `-32768`

	// var num int  // Defaults to int32 or int64 depending on your system

	var float_32_bit float32 = 12345678.9
	fmt.Println(float_32_bit)  // Shows as `1.2345679e+07` (i.e. 12345679.0) because of floating point error!

	var float_64_bit float64 = 12345678.9
	fmt.Println(float_64_bit)  // Shows correctly as `1.23456789e+07` (i.e. 12345678.9)


	fmt.Println("\nMin and max sizes for all int and float types:")
	// int
    fmt.Printf("int\nmin = %d\nmax = %d\n\n", math.MinInt, math.MaxInt)
    // int8
    fmt.Printf("int8\nmin = %d\nmax = %d\n\n", math.MinInt8, math.MaxInt8)
    // int16
    fmt.Printf("int16\nmin = %d\nmax = %d\n\n", math.MinInt16, math.MaxInt16)
    // int32
    fmt.Printf("int32\nmin = %d\nmax = %d\n\n", math.MinInt32, math.MaxInt32)
    // int64
    fmt.Printf("int64\nmin = %d\nmax = %d\n\n", math.MinInt64, math.MaxInt64)

    // unsigned
    // uint
    fmt.Printf("uint\nmin = %d\nmax = %d\n\n", uint(0), uint(math.MaxUint))
    // uint8
    fmt.Printf("uint8\nmin = %d\nmax = %d\n\n", 0, math.MaxUint8)
    // uint16
    fmt.Printf("uint16\nmin = %d\nmax = %d\n\n", 0, math.MaxUint16)
    // uint32
    fmt.Printf("uint32\nmin = %d\nmax = %d\n\n", 0, math.MaxUint32)
    // uint64
    fmt.Printf("uint64\nmin = %d\nmax = %d\n\n", 0, uint64(math.MaxUint64))

	// var result = float_32_bit + num_16_bit  // Does not work!
	var result = float_32_bit + float32(num_16_bit)  // Works! Explicit casting is obligatory
	fmt.Println(result)

	var my_str string = "Hello" + ", " + "world!"
	fmt.Println(my_str)

	var my_str_multi_line string = `Hello
World!`
	fmt.Println(my_str_multi_line)

	fmt.Println(len(my_str))  // Gives the number of bytes in the string!
	fmt.Println(len("巨魔的臉"))  // Outputs 12, not 4! Because it's an unicode string, not ascii

	var int_num int
	fmt.Println(int_num)  // Outputs 0. Types have default values!
	// All ints and floats (and runes) have a default of 0
	// str defaults to "", an empty string 

	var int_num_inferred = 12  // Infers automatically what type it is (int -> int64)
	fmt.Println(int_num_inferred)
	float_num_inferred := 133.7  // Also works, even more succint
	fmt.Println(float_num_inferred)

	// Do not use inferred types everywhere, though!
	// foo := bar()  // No info on wtf return type `bar()` has!
	// var foo string = bar()  // Much better now. Keep this good practice

	const my_const string = "My value can't ever be changed!"

	// const my_const_no_value  // Doesn't work! A value MUST be set

	// Everything else from above applies to consts too! They don't generate a
	// compile error if initialized and not used, though
	const pi = 3.14159265
}