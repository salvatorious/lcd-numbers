package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"

	lcd "github.com/salvatorious/lcd-numbers/lcd"
)

//LCDNumber is a comment
type LCDNumber struct {
	Size   int
	Number []*lcd.Character
}

func main() {
	readInput()
}

func readInput() {
	fileHandle, _ := os.Open("input.txt")
	defer fileHandle.Close()
	fileScanner := bufio.NewScanner(fileHandle)

	// var rows []LCDNumber
	rowIndex := 0
	cursorPosition := 0
	var fmtBuffer bytes.Buffer

	for fileScanner.Scan() {
		if fileScanner.Text() != "00" {
			cursorPosition = 0 // reset "cursor"
			line := strings.Split(fileScanner.Text(), " ")
			size, err := strconv.Atoi(line[0])
			if err != nil {
				panic(err)
			}

			runes := charSplit(line[1])

			for i, char := range runes {
				fmtBuffer.WriteString(lcd.IndexPrint(char, i, rowIndex, size))
			}

			rowIndex++
			// fmt.Println("Size:", size, "Number:", lcd.IndexPrint(runes[1], 0, 0, 0))
		} else {
			// return // EOF
		}
	}

	fmt.Println(fmtBuffer.String())
}

func charSplit(num string) []string {
	var characters []string
	characters = strings.Split(num, "")

	return characters
}
