package test

import "fmt"

func Do(a int){//a func in different package, when we use the func in the other package, we should let the first letter 
		// of func name be caps, or that would be not working.
	fmt.Printf("%d is the funcdo\n",a);
} 
