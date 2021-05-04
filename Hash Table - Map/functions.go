package hashtablemap

import (
	"bufio"
	"fmt"
	"os"
)

type hashmap map[int][]string

func (h *hashmap) Initialize() {

	*h = hashmap{
		1: {}, 2: {}, 3: {}, 4: {}, 5: {}, 6: {}, 7: {}, 8: {}, 9: {}, 10: {}, 11: {}, 12: {}, 13: {},
		14: {}, 15: {}, 16: {}, 17: {}, 18: {}, 19: {}, 20: {}, 21: {}, 22: {}, 23: {}, 24: {}, 25: {}, 26: {},
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
		index := getIndex(firstletter)
		if index != 27 {

			h[index] = append(h[index], name)
		} else {
			fmt.Println(name + "Not a valid name")
		}
	}

}

func (h hashmap) DeleteFromHash(namearr []string) {
	for _, name := range namearr {
		firstletter := name[0]

		index := getIndex(firstletter)
		if index != 27 {
			hashslice := DeleteElement(h[index], name)
			delete(h, index)
			h[index] = hashslice
		} else {
			fmt.Println(name + "Not a valid name")
		}

	}
}
func getIndex(firstletter byte) int {
	var index byte
	if firstletter >= 65 && firstletter <= 90 {
		index = ((firstletter - 12) % 26)
	} else if firstletter >= 97 && firstletter <= 122 {
		index = ((firstletter - 18) % 26)
	} else {
		return 27
	}
	return int(index)
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
