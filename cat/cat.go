package main

import (
	"fmt"
	"os"
)
func check(e error){
	if e !=nil{
		fmt.Fprintf(os.Stderr,"cat: %v  \n",e)
	}
}


func main(){
	if len(os.Args)<2{
		fmt.Println("Usage: program <filename>")
		return 
	}

	for _,filename := range os.Args[1:]{

	data,err := os.ReadFile(filename)
	check(err)
	fmt.Println(string(data))
}
}

