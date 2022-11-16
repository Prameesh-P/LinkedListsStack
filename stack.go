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
	Mystack := stack{}
	fmt.Println(Mystack)
	Mystack.push(10)
	Mystack.push(20)
	Mystack.push(30)
	Mystack.push(40)
	Mystack.push(50)
	Mystack.push(60)
	fmt.Println(Mystack)
	Mystack.pop()
	fmt.Println(Mystack)
	Mystack.pop()
	fmt.Println(Mystack)
}
func (s *stack) pop()int  {
	toremove:=s.values[len(s.values)-1]
	s.values=s.values[:len(s.values)-1]
	return toremove
}