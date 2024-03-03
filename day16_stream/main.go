// package main

//	func main() {
//		Printfln("Product: %v, Price: %v", Kayak.Name, Kayak.Price)
//	}
package main

import (
	"io"
	"strings"
)

//	func processData(reader io.Reader) {
//		b := make([]byte, 2)
//		for {
//			count, err := reader.Read(b)
//			if count > 0 {
//				Printfln("Read %v bytes: %v", count, string(b[0:count]))
//			}
//			if err == io.EOF {
//				break
//			}
//		}
//	}
//
//	func main() {
//		r := strings.NewReader("Kayak")
//		processData(r)
//	}
func processData(reader io.Reader, writer io.Writer) {
	count, err := io.Copy(writer, reader)
	if err == nil {
		Printfln("Read %v bytes", count)
	} else {
		Printfln("Error: %v", err.Error())
	}
}
func main() {
	// r := strings.NewReader("Kayak")
	// var builder strings.Builder
	// processData(r, &builder)
	// Printfln("String builder contents: %s", builder.String())
	// pipeReader, pipeWriter := io.Pipe()
	// go GenerateData(pipeWriter)
	// ConsumeData(pipeReader)

	/*

		Write first, then read later -> io.Pipe()

	*/

	// r1 := strings.NewReader("Kayak")
	// r2 := strings.NewReader("Lifejacket")
	// r3 := strings.NewReader("Canoe")
	// concatReader := io.MultiReader(r1, r2, r3)
	// // ConsumeData(concatReader)
	// var writer strings.Builder
	// teeReader := io.TeeReader(concatReader, &writer)
	// ConsumeData(teeReader)
	// Printfln("Echo data: %v", writer.String())
	/*
		r1 -> r2 -> r3 -> concatReader.
		then from concatReader -> teeReader <Man in the middle, inspector!> -> Writer
		!!! done!
	*/

	// limited := io.LimitReader(concatReader, 5)
	// ConsumeData(limited)
	// var w1 strings.Builder
	// var w2 strings.Builder
	// var w3 strings.Builder
	// combinedWriter := io.MultiWriter(&w1, &w2, &w3)
	// GenerateData(combinedWriter)
	// Printfln("Writer #1: %v", w1.String())
	// Printfln("Writer #2: %v", w2.String())
	// Printfln("Writer #3: %v", w3.String())

	//Multiple Writer......
	/*
		One data -> Many writer with the same data!!! through one multiplexer.
	*/
	var w1 strings.Builder
	var w2 strings.Builder
	var w3 strings.Builder
	combined := io.MultiWriter(&w1, &w2, &w3)
	GenerateData(combined)
	Printfln("Writer #1: %v", w1.String())
	Printfln("Writer #2: %v", w2.String())
	Printfln("Writer #3: %v", w3.String())
}
