package main

import "fmt"

type palindrome struct {
	str string
	resultStart int
	resultLength int
}

func main() {
	var name string
	fmt.Println("Enter the string")
	fmt.Scanln(&name)
	var pname palindrome
	pname.str=name
	strlength:=len(name)
	for start:=0;start< strlength;start++{
		(&pname).expandRange(start,start)
		(&pname).expandRange(start,start+1)
	}
	fmt.Println(pname.resultLength,pname.resultStart)
	fmt.Println(len(name))
	output:= pname.str[pname.resultStart:pname.resultStart+pname.resultLength]
	fmt.Println(output)

}

func (p *palindrome) expandRange(start int,end int){

	for ((start>=0 && end<len(p.str)) && (p.str[start]==p.str[end])){	
		start--;
		end++;
	}
	if (p.resultLength<end-start-1){
		p.resultStart=start+1;
		p.resultLength=end-start-1;
	}

 }