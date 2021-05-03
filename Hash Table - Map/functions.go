package hashtablemap

import (
	"bufio"
	"fmt"
	"os"
)

type hashmap map[string][]string

func (h *hashmap) Initialize() {

	*h = hashmap{
		"A": {}, "B": {}, "C": {}, "D": {}, "E": {}, "F": {}, "G": {}, "H": {}, "I": {}, "J": {}, "K": {}, "L": {}, "M": {},
		"N": {}, "O": {}, "P": {}, "Q": {}, "R": {}, "S": {}, "T": {}, "U": {}, "V": {}, "W": {}, "X": {}, "Y": {}, "Z": {},
	}

}

func GetName() []string {
	var namearr []string
	var name string
	var elecount int
	fmt.Println("Enter no of names")
	fmt.Scanf("%d\n", &elecount)
	fmt.Println("Enter names")
	for i := 0; i < elecount; i++ {
		fmt.Scanf("%s\n", &name)
		namearr = append(namearr, name)
	}
	return namearr
}
func ReadFile(filename string) []string {
	var namearr []string

	file, ferr := os.Open(filename)
	if ferr != nil {
		fmt.Println("Cannot open")
		fmt.Println(ferr)
		panic(ferr)
	}
	fmt.Println("file open sucess")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		name := scanner.Text()
		namearr = append(namearr, name)

	}
	return namearr
}

func (h hashmap) InsertInHash(namearr []string) {
	for _, name := range namearr {
		firstletter := name[0]
		switch {
		case firstletter == 'A' || firstletter == 'a':
			h["A"] = append(h["A"], name)
		case firstletter == 'B' || firstletter == 'b':
			h["B"] = append(h["B"], name)
		case firstletter == 'C' || firstletter == 'c':
			h["C"] = append(h["C"], name)
		case firstletter == 'D' || firstletter == 'd':
			h["D"] = append(h["D"], name)
		case firstletter == 'E' || firstletter == 'e':
			h["E"] = append(h["E"], name)
		case firstletter == 'F' || firstletter == 'f':
			h["F"] = append(h["F"], name)
		case firstletter == 'G' || firstletter == 'g':
			h["G"] = append(h["G"], name)
		case firstletter == 'H' || firstletter == 'h':
			h["H"] = append(h["H"], name)
		case firstletter == 'I' || firstletter == 'i':
			h["I"] = append(h["I"], name)
		case firstletter == 'J' || firstletter == 'j':
			h["J"] = append(h["J"], name)
		case firstletter == 'K' || firstletter == 'k':
			h["K"] = append(h["K"], name)
		case firstletter == 'L' || firstletter == 'l':
			h["L"] = append(h["L"], name)
		case firstletter == 'M' || firstletter == 'm':
			h["M"] = append(h["M"], name)
		case firstletter == 'N' || firstletter == 'n':
			h["N"] = append(h["N"], name)
		case firstletter == 'O' || firstletter == 'o':
			h["O"] = append(h["O"], name)
		case firstletter == 'P' || firstletter == 'p':
			h["P"] = append(h["P"], name)
		case firstletter == 'Q' || firstletter == 'q':
			h["Q"] = append(h["Q"], name)
		case firstletter == 'R' || firstletter == 'r':
			h["R"] = append(h["R"], name)
		case firstletter == 'S' || firstletter == 's':
			h["S"] = append(h["S"], name)
		case firstletter == 'T' || firstletter == 't':
			h["T"] = append(h["T"], name)
		case firstletter == 'U' || firstletter == 'u':
			h["U"] = append(h["U"], name)
		case firstletter == 'V' || firstletter == 'v':
			h["V"] = append(h["V"], name)
		case firstletter == 'W' || firstletter == 'w':
			h["W"] = append(h["W"], name)
		case firstletter == 'X' || firstletter == 'x':
			h["X"] = append(h["X"], name)
		case firstletter == 'Y' || firstletter == 'y':
			h["Y"] = append(h["Y"], name)
		case firstletter == 'Z' || firstletter == 'z':
			h["Z"] = append(h["Z"], name)
		default:
			fmt.Println("Invalid Name")
		}

	}

}

func (h hashmap) DeleteFromHash(namearr []string) {
	for _, name := range namearr {
		firstletter := name[0]
		switch {
		case firstletter == 'A' || firstletter == 'a':
			hashslice := DeleteElement(h["A"], name)
			delete(h, "A")
			h["A"] = hashslice
		case firstletter == 'B' || firstletter == 'b':
			hashslice := DeleteElement(h["B"], name)
			delete(h, "b")
			h["B"] = hashslice
		case firstletter == 'C' || firstletter == 'c':
			hashslice := DeleteElement(h["C"], name)
			delete(h, "C")
			h["C"] = hashslice
		case firstletter == 'D' || firstletter == 'd':
			hashslice := DeleteElement(h["A"], name)
			delete(h, "D")
			h["D"] = hashslice
		case firstletter == 'E' || firstletter == 'e':
			hashslice := DeleteElement(h["E"], name)
			delete(h, "E")
			h["E"] = hashslice
		case firstletter == 'F' || firstletter == 'f':
			hashslice := DeleteElement(h["F"], name)
			delete(h, "F")
			h["F"] = hashslice
		case firstletter == 'G' || firstletter == 'g':
			hashslice := DeleteElement(h["G"], name)
			delete(h, "G")
			h["G"] = hashslice
		case firstletter == 'H' || firstletter == 'h':
			hashslice := DeleteElement(h["H"], name)
			delete(h, "H")
			h["H"] = hashslice
		case firstletter == 'I' || firstletter == 'i':
			hashslice := DeleteElement(h["I"], name)
			delete(h, "I")
			h["I"] = hashslice
		case firstletter == 'J' || firstletter == 'j':
			hashslice := DeleteElement(h["J"], name)
			delete(h, "J")
			h["J"] = hashslice
		case firstletter == 'K' || firstletter == 'k':
			hashslice := DeleteElement(h["K"], name)
			delete(h, "K")
			h["K"] = hashslice
		case firstletter == 'L' || firstletter == 'l':
			hashslice := DeleteElement(h["L"], name)
			delete(h, "L")
			h["L"] = hashslice
		case firstletter == 'M' || firstletter == 'm':
			hashslice := DeleteElement(h["M"], name)
			delete(h, "M")
			h["M"] = hashslice
		case firstletter == 'N' || firstletter == 'n':
			hashslice := DeleteElement(h["N"], name)
			delete(h, "N")
			h["N"] = hashslice
		case firstletter == 'O' || firstletter == 'o':
			hashslice := DeleteElement(h["O"], name)
			delete(h, "O")
			h["O"] = hashslice
		case firstletter == 'P' || firstletter == 'p':
			hashslice := DeleteElement(h["P"], name)
			delete(h, "P")
			h["P"] = hashslice
		case firstletter == 'Q' || firstletter == 'q':
			hashslice := DeleteElement(h["Q"], name)
			delete(h, "Q")
			h["Q"] = hashslice
		case firstletter == 'R' || firstletter == 'r':
			hashslice := DeleteElement(h["R"], name)
			delete(h, "R")
			h["R"] = hashslice
		case firstletter == 'S' || firstletter == 's':
			hashslice := DeleteElement(h["S"], name)
			delete(h, "S")
			h["S"] = hashslice
		case firstletter == 'T' || firstletter == 't':
			hashslice := DeleteElement(h["T"], name)
			delete(h, "T")
			h["T"] = hashslice
		case firstletter == 'U' || firstletter == 'u':
			hashslice := DeleteElement(h["U"], name)
			delete(h, "U")
			h["U"] = hashslice
		case firstletter == 'V' || firstletter == 'v':
			hashslice := DeleteElement(h["V"], name)
			delete(h, "V")
			h["V"] = hashslice
		case firstletter == 'W' || firstletter == 'w':
			hashslice := DeleteElement(h["W"], name)
			delete(h, "W")
			h["W"] = hashslice
		case firstletter == 'X' || firstletter == 'x':
			hashslice := DeleteElement(h["X"], name)
			delete(h, "X")
			h["X"] = hashslice
		case firstletter == 'Y' || firstletter == 'y':
			hashslice := DeleteElement(h["Y"], name)
			delete(h, "Y")
			h["Y"] = hashslice
		case firstletter == 'Z' || firstletter == 'z':
			hashslice := DeleteElement(h["Z"], name)
			delete(h, "Z")
			h["Z"] = hashslice
		default:
			fmt.Println("Invalid Name")
		}

	}
}

func DeleteElement(hashslice []string, name string) []string {
	a := len(hashslice)
	var i int
	for i = 0; i < a; i++ {
		if hashslice[i] == name {
			break
		} else {
			continue
		}
	}
	if i == a {
		fmt.Println("element not found")
	} else if i == (a - 1) {
		hashslice[i] = ""

	} else {
		for j := i; j < (a - 1); j++ {
			hashslice[j] = hashslice[j+1]
		}
		hashslice[a-1] = ""
	}
	return hashslice

}

func (h hashmap) display() {
	fmt.Println("Our Hash table is")
	fmt.Println(h)
}
