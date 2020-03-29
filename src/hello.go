package main

import "test"

func main(){
	//name = "spark" a string is declared but not used can cause the warning and let the program not working
	for i:=0;i<5;i++{
		test.Do(i);
	}
}
