package art

import (
	"os"
	"strings"
)

func Ascii_Art(font, font_style string) string {
	//We retrieve the data from the standard.txt document and store it temporarily in a string
	data, err := os.ReadFile(font_style)
	if err != nil {
		return err.Error()
	}
	texte := string(data)

	//We recreate the .txt document in a tab table
	var tab []string
	var temp string
	for _, element := range texte {
		if element == '\n' {
			if temp != "\n" {
				tab = append(tab, temp)
				temp = ""
			}
		} else {
			temp += string(element)
		}
	}
	if temp != "\n" {
		tab = append(tab, temp)
	}

	//We separate the string from the argument every \\n in an array
	input := strings.Split(font, "\n")
	var fstring, finalstring string

	//We do as many loops as there are in the table
	for _, table := range input {
		//If the current table entry is empty, we make a line break
		if table != "" {
			//We do a double loop which will add the letters in the correct order to the string variable.
			for i := 1; i < 9; i++ {
				//We loop as many times as there are letters in the current entry in the table
				for _, char := range table {
					/*We virtually delete the first 32 ascii then add the correct line from the clone table of
					  the .txt file in the variable containing the new sentence to display*/
					ia := int(char) - 32
					if ia >= 0 {
						fstring += tab[(ia*9 + i)]
					}

				}
				//At the end of each line we do a newline
				fstring += "\n"
			}
			//We write the line then we empty the variable containing the sentence
			finalstring += fstring
			fstring = ""
		} else {
			finalstring += "\n"
		}
	}
	return finalstring
}
