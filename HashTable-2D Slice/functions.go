package hashtable

import (
	"fmt"
)

//GetName to get single/Multiple input from user
func GetName() []string {
	/*Another way to read input from user preferable when input has more than one word
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	*/
	var namearr []string
	var name string
	var choice, elecount int
	fmt.Println("For Single name Press 0,else 1")
	fmt.Scanln(&choice)
	if choice == 0 {
		fmt.Println("Enter Name")
		fmt.Scanf("%s\n", &name)
		namearr = append(namearr, name)
	} else {
		fmt.Println("Enter no of names")
		fmt.Scanf("%d\n", &elecount)
		fmt.Println("Enter names")
		for i := 0; i < elecount; i++ {
			fmt.Scanf("%s\n", &name)
			namearr = append(namearr, name)
		}
	}
	return namearr
}

//InsertInHash to insert a name in hash table.
//Names starting with "A" should be in hash 1,"B" should be in hash 2,"C" should be in hash 3
func InsertInHash(namearr []string, hash [][]string) [][]string {
	for _, add := range namearr {
		firstletter := add[0]
		switch {
		case firstletter == 'A' || firstletter == 'a':
			hash[0] = append(hash[0], add)
		case firstletter == 'B' || firstletter == 'b':
			hash[1] = append(hash[1], add)
		case firstletter == 'C' || firstletter == 'c':
			hash[2] = append(hash[2], add)
		default:
			hash[3] = append(hash[3], add)
		}
	}
	return hash
}

//DeleteFromHash to delete single/Multiple entry in hash table
func DeleteFromHash(namearr []string, hash [][]string) [][]string {
	for _, del := range namearr {
		firstletter := del[0]
		switch {
		case firstletter == 'A' || firstletter == 'a':
			hash = DeleteElement(del, hash, 0)
		case firstletter == 'B' || firstletter == 'b':
			hash = DeleteElement(del, hash, 1)
		case firstletter == 'C' || firstletter == 'c':
			hash = DeleteElement(del, hash, 2)
		default:
			hash = DeleteElement(del, hash, 3)
		}
	}
	return hash
}

//DeleteElement to delete a element
func DeleteElement(del string, hash [][]string, x int) [][]string {
	a := len(hash[x])
	var i int
	for i = 0; i < a; i++ {
		if hash[x][i] == del {
			break
		} else {
			continue
		}
	}
	if i == a {
		fmt.Println("element not found")
	} else if i == (a - 1) {
		hash[x][i] = ""

	} else {
		for j := i; j < (a - 1); j++ {
			hash[x][j] = hash[x][j+1]
		}
		hash[x][a-1] = ""
	}
	return hash
}

//PrintHash to print our Hash table
func PrintHash(hashtable [][]string) {
	fmt.Println(hashtable)
}
