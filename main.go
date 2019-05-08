package main

import (
	"bufio"
	"fmt"
	"github.com/kniren/gota/series"
	"os"
)

/*
Go CSV:
https://www.thepolyglotdeveloper.com/2017/03/parse-csv-data-go-programming-language/

----------------------------------------------------------------------------------------

Go Gota (Dataframes & Series):
https://github.com/go-gota/gota
https://godoc.org/github.com/kniren/gota/series#Element
https://godoc.org/github.com/kniren/gota/dataframe#DataFrame.Records

----------------------------------------------------------------------------------------

GoNum:
https://www.gonum.org/
https://github.com/gonum/gonum
*/

import (
	"github.com/go-gota/gota/dataframe"
)

func main() {
	//file := getFile()
	data := csvTry("df_test.csv")
	fmt.Println(data)
	fmt.Printf("\n")
	fmt.Printf("%T \n", data)

	amount := data.Select([]string{"Amount"})
	fmt.Println(amount)
	fmt.Printf("Amount column Type: %T\n", amount)
	fmt.Printf("Amount column Data: %q\n", amount)
	fmt.Println("--------------------------------------------------------")
	//for i, a := range data.Col("Amount") {
	//	if i > 0 {
	//		b, err := strconv.ParseInt([]a)
	//		fmt.Printf("Index: %d | Amount: %q | Type: %T \n", i, a, a)
	//	}
	//}

	//add := func(s series.Series) series.Series {
	//	ints, err := s.Int()
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//	sum := 0
	//	for _, i := range ints {
	//		sum += i
	//	}
	//	return series.Ints(sum)
	//
	//}

	dmap := data.Maps()
	fmt.Println(dmap)
	amt := data.Select([]string{"Amount"})
	fmt.Println(amt)
	fmt.Printf("%T\n", amt)
	nAmt := amt.Capply(add)
	fmt.Printf("Type: %T | Data: %q \n", nAmt, nAmt)
	fAmt := nAmt.Subset([]int{0, 0})
	fmt.Printf("Type: %T | Data: %q \n", fAmt, fAmt)
}

//func getFile() string {
//	reader := bufio.NewReader(os.Stdin)
//	fmt.Printf("\n")
//	fmt.Print("Enter text: ")
//	text, _ := reader.ReadString('\n')
//	fmt.Printf("\n")
//	fmt.Printf("Text Type: %T \n", text)
//	fmt.Printf("Text Data (Pre-Trim): %q\n", text)
//	text = strings.TrimSuffix(text, "\n") // Windows line is terminated with both "\r\n"
//	text = strings.TrimSuffix(text, "\r") // Windows line is terminated with both "\r\n"
//	fmt.Printf("Text Data (Post-Trim): %q\n", text)
//	return text
//}

func csvTry(file string) dataframe.DataFrame {
	csvFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}
	reader := bufio.NewReader(csvFile)
	df := dataframe.ReadCSV(reader)
	return df
}

func add(s series.Series) series.Series {
	ints, err := s.Int()
	if err != nil {
		fmt.Println(err)
	}
	sum := 0
	for _, i := range ints {
		sum += i
	}
	return series.Ints(sum)

}
