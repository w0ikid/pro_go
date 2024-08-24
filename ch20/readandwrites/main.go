package main

import (
	// "io"
	// "bufio"
	"fmt"
	// "io"
	"strings"
	"encoding/json"
	// "bufio"
)

// func scanFromReader(reader io.Reader, template string, vals ...interface{}) (int, error) {
// 	return fmt.Fscanf(reader, template, vals...)
// }

// func scanSingle(reader io.Reader, val interface{}) (int, error) {
// 	return fmt.Fscan(reader, val)
// }

// func writeFormatted(writer io.Writer, template string, vals ...interface{}) {
// 	fmt.Fprintf(writer, template, vals...)
// }

// func writeReplaced(writer io.Writer, str string, subs ...string) {
// 	replacer := strings.NewReplacer(subs...)
// 	replacer.WriteString(writer, str)
// }


// func processData(reader io.Reader, writer io.Writer){
// 	count, err := io.Copy(writer, reader)

// 	if err == nil {
// 		Printfln("Read %v bytes", count)
// 	} else {
// 		Printfln("Error: %v", err.Error())
// 	}
// }


func main(){
	// Printfln("Product: %v Price: %v", Kayak.Name, Kayak.Price)

	// r := strings.NewReader("Kayak")
		
	// var builder strings.Builder
	// processData(r, &builder)
	// Printfln("Strings builder contents: %s", builder.String())

	

	// pipeReader, pipeWriter := io.Pipe()
	// go GeneratedData(pipeWriter)
	// ConsumeData(pipeReader)

	// r1 := strings.NewReader("Kayak")
	// r2 := strings.NewReader("Lifejacket")
	// r3 := strings.NewReader("Canoe")

	
	// concatReader := io.MultiReader(r1, r2, r3)
	// ConsumeData(concatReader)

	// var w1 strings.Builder
	// var w2 strings.Builder
	// var w3 strings.Builder

	// combineWriter := io.MultiWriter(&w1, &w2, &w3)

	// GeneratedData(combineWriter)

	// Printfln("Writer #1: %v", w1.String())
	// Printfln("Writer #2: %v", w2.String())
	// Printfln("Writer #3: %v", w3.String())

	// r1 := strings.NewReader("Kayak")
	// r2 := strings.NewReader("Lifejacket")
	// r3 := strings.NewReader("Canoe")

	// concatReader := io.MultiReader(r1, r2, r3)

	// var writer strings.Builder
	// teeReader := io.TeeReader(concatReader, &writer)

	// ConsumeData(teeReader)
	// Printfln("Echo data: %v", writer.String())

	// text := "It was a boat. A small boat."

	// var reader io.Reader = NewCustomReader(strings.NewReader(text))

	// var writer strings.Builder
	
	// slice := make([]byte, 5)
	
	// // reader = bufio.NewReader(reader)

	// buffered := bufio.NewReader(reader)

	// for {
	// 	count, err := buffered.Read(slice)
	// 	if count > 0 {
	// 		Printfln("Buffer size: %v, buffered: %v", buffered.Size(), buffered.Buffered())
	// 	}
	// 	if err == io.EOF {
	// 		break
	// 	}
	// }
	
	// Printfln("Read data: %v", writer.String())

	// text := "It was a boat. A small boat."

	// var builder strings.Builder
	// var writer = NewCustomWriter(&builder)

	// for i := 0; true; {
	// 	end := i + 5

	// 	if end >= len(text) {
	// 		writer.Write([]byte(text[i:]))
	// 		break
	// 	}
	// 	writer.Write([]byte(text[i:end]))
	// 	i = end
	// }

	// Printfln("Written data: %v", builder.String())


	// text := "It was a boat. A small boat."

	// var builder strings.Builder

	// var writer = bufio.NewWriterSize(NewCustomWriter(&builder), 20)

	// for i := 0; true; {
	// 	end := i + 5
	// 	if end >= len(text) {
	// 		writer.Write([]byte(text[i:]))
	// 		writer.Flush()
	// 		break
	// 	}
	// 	writer.Write([]byte(text[i:end]))
	// 	i = end
	// }

	// Printfln("Written data: %v", builder.String())






	// reader := strings.NewReader("Kayak Watersports $279.00")

	// var name, category string
	// var price float64
	// scanTemplate := "%s %s $%f"

	// _, err := scanFromReader(reader, scanTemplate, &name, &category, &price)

	// if err != nil {
	// 	Printfln("Error: %v", err.Error())
	// } else {
	// 	Printfln("Name: %v", name)
	// 	Printfln("Category: %v", category)
	// 	Printfln("Price: %.2f", price)
	// }

	// reader := strings.NewReader("Kayak Watersports $279.00")

	// for {
	// 	var str string
	// 	_, err := scanSingle(reader, &str)
	// 	if err != nil {
	// 		if err != io.EOF {
	// 			Printfln("Error: %v", err.Error)
	// 		}
	// 		break
	// 	}
	// 	Printfln("Value: %v", str)
	// }


	// var writer strings.Builder
	// template := "Name: %s, Category: %s, Price: $%.2f"

	// writeFormatted(&writer, template, "Kayak", "Watersports", float64(279))

	// fmt.Println(writer.String())

	// text := "It was a boat. A small boat."
	// subs := []string{"boat", "kayak", "small", "huge"}

	// var writer strings.Builder
	// writeReplaced(&writer, text, subs...)
	// fmt.Println(writer.String())


	// 21 CHAPTER
	// ----------




	// var b bool = true
	// var str string = "Hello"
	// var fval float64 = 99.99
	// var ival int = 200
	// var pointer *int = &ival
	
	// var writer strings.Builder
	// encoder := json.NewEncoder(&writer)

	// for _, val := range []interface{} {b, str, fval, ival, pointer} {
	// 	encoder.Encode(val)
	// }

	// // json добавляет "\n" новый символ строки после кодирования каждого элемента

	// fmt.Print(writer.String())



// 	names := []string {"Kayak", "Lifejacket", "Soccer Ball"}
// 	numbers := [3]int {10, 20, 30}
// 	var byteArray [5]byte
// 	copy(byteArray[0:], []byte(names[0]))
// 	byteSlice := []byte(names[0])
	
// 	var writer strings.Builder
// 	encoder := json.NewEncoder(&writer)

// 	encoder.Encode(names)
// 	encoder.Encode(numbers)
// 	encoder.Encode(byteArray)
// 	encoder.Encode(byteSlice)

// 	fmt.Print(writer.String())


	// Кодирования карты

	// m := map[string]float64 {
	// 	"Kayak": 279,
	// 	"Lifejacket": 49.95,
	// }

	var writer strings.Builder
	encoder := json.NewEncoder(&writer)

	dp := DiscountedProduct{
		Product:  &Kayak,
		Discount: 10.5,
	}

	encoder.Encode(dp)

	// fmt.Print(writer.String())


	dp2 := DiscountedProduct{
		Discount: 12,
	}
	encoder.Encode(&dp2)

	fmt.Print(writer.String())

}