package main

import "fmt"

//stack struct created

type stack struct {
	
	values []int

}
//method to push values

func (s *stack) push(val int) {
	
	s.values = append(s.values, val)

}

func main() {
	//assingining stack to empty  stack variable 
	//short variable decleration

	Mystack := stack{}

	fmt.Println(Mystack)//printing empty stack. Pushing values in the slice

	Mystack.push(10)

	Mystack.push(20)

	Mystack.push(30)

	Mystack.push(40)

	Mystack.push(50)

	Mystack.push(60)

	fmt.Println(Mystack)

	//Mystack.pop() this call go to the method pop. And deletion will work.

	Mystack.pop()

	fmt.Println(Mystack)

	Mystack.pop()

	fmt.Println(Mystack)

}

//method to remove elementes from the stack(slice)..

func (s *stack) pop()int {

//The variable for store already deleted elements to store.

//and the values is return to the pop method.

	toremove:=s.values[len(s.values)-1]

//s.values for delete elements

	s.values=s.values[:len(s.values)-1]

	return toremove

}