// Go offers excellent support for string formatting in
// the `printf` tradition. Here are some examples of
// common string formatting tasks.

package main

import "fmt"
import "os"
import "strconv"

type point struct {
	x, y int
}

func main() {

	// Go offers several printing "verbs" designed to
	// format general Go values. For example, this prints
	// an instance of our `point` struct.
	p := point{1, 2}
	fmt.Printf("%v\n", p)

	// If the value is a struct, the `%+v` variant will
	// include the struct's field names.
	fmt.Printf("%+v\n", p)

	// The `%#v` variant prints a Go syntax representation
	// of the value, i.e. the source code snippet that
	// would produce that value.
	fmt.Printf("%#v\n", p)

	// To print the type of a value, use `%T`.
	fmt.Printf("%T\n", p)

	// Formatting booleans is straight-forward.
	fmt.Printf("%t\n", true)

	// There are many options for formatting integers.
	// Use `%d` for standard, base-10 formatting.
	fmt.Printf("%d\n", 123)

	// This prints a binary representation.
	fmt.Printf("%b\n", 14)

	// This prints the character corresponding to the
	// given integer.
	fmt.Printf("%c\n", 33)

	// `%x` provides hex encoding.
	fmt.Printf("%x\n", 456)

	// There are also several formatting options for
	// floats. For basic decimal formatting use `%f`.
	fmt.Printf("%f\n", 78.9)

	// `%e` and `%E` format the float in (slightly
	// different versions of) scientific notation.
	fmt.Printf("%e\n", 123400000.0)
	fmt.Printf("%E\n", 123400000.0)

	// For basic string printing use `%s`.
	fmt.Printf("%s\n", "\"string\"")

	// To double-quote strings as in Go source, use `%q`.
	fmt.Printf("%q\n", "\"string\"")

	// As with integers seen earlier, `%x` renders
	// the string in base-16, with two output characters
	// per byte of input.
	fmt.Printf("%x\n", "hex this")

	// To print a representation of a pointer, use `%p`.
	fmt.Printf("%p\n", &p)

	// When formatting numbers you will often want to
	// control the width and precision of the resulting
	// figure. To specify the width of an integer, use a
	// number after the `%` in the verb. By default the
	// result will be right-justified and padded with
	// spaces.
	fmt.Printf("|%6d|%6d|\n", 12, 345)

	// You can also specify the width of printed floats,
	// though usually you'll also want to restrict the
	// decimal precision at the same time with the
	// width.precision syntax.
	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)

	// To left-justify, use the `-` flag.
	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)

	// You may also want to control width when formatting
	// strings, especially to ensure that they align in
	// table-like output. For basic right-justified width.
	fmt.Printf("|%6s|%6s|\n", "foo", "b")

	// To left-justify use the `-` flag as with numbers.
	fmt.Printf("|%-6s|%-6s|\n", "foo", "b")

	// So far we've seen `Printf`, which prints the
	// formatted string to `os.Stdout`. `Sprintf` formats
	// and returns a string without printing it anywhere.
	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)

	// You can format+print to `io.Writers` other than
	// `os.Stdout` using `Fprintf`.
	fmt.Fprintf(os.Stderr, "an %s\n", "error")

	bs := []byte{}
	str := "this is a test!"
	c := len(str)
	bs = strconv.AppendInt(bs, int64(c), 10)
	bs = strconv.AppendQuote(bs, str)
	bs = strconv.AppendInt(bs, 123, 10)
	bs = strconv.AppendBool(bs, true)
	bs = strconv.AppendFloat(bs, 1.23, 'f', -1, 64)
	fmt.Println("string:", string(bs), " len:", len(bs))

	fmt.Println("bool: ", strconv.FormatBool(true))
	fmt.Println("2 int: ", strconv.FormatInt(123, 10))
	fmt.Println("36 int: ", strconv.FormatInt(123, 2))
	fmt.Println("2 uint: ", strconv.FormatUint(123, 10))
	fmt.Println("36 uint: ", strconv.FormatUint(123, 2))
	fmt.Println("float b: ", strconv.FormatFloat(1.230, 'b', -1, 64))
	fmt.Println("float E: ", strconv.FormatFloat(1.230, 'E', -1, 64))
	fmt.Println("float e: ", strconv.FormatFloat(1.230, 'e', -1, 64))
	fmt.Println("float f: ", strconv.FormatFloat(1.2301231, 'f', 6, 64))
	fmt.Println("float G: ", strconv.FormatFloat(1.23012321, 'G', 6, 64))
	fmt.Println("float g: ", strconv.FormatFloat(1.230123213, 'g', 6, 64))

	reader := strings.NewReader(str)
	b, _ := reader.ReadByte()
	fmt.Println(strconv.Atoi(string(b)))
}
