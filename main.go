package main

import "fmt"

func main(){
	mybill :=newBill("aleen's bill")
	mybill.updateTip(10)
	mybill.addItem("pie",3.55)
	mybill.addItem("coffe",4.55)
	fmt.Println(mybill.format())
}