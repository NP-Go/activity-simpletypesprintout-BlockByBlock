package main

//Insert variables declarations here based on activity

func tellMeTypes() {
	//insert code here to print out types of variables
	text := "The following isthe account information."
	firstName := "Luke"
	age := 20
	height := 1.72
	fmt.Printf(ā%T %T %T %T\nā, text, firstName, age, height)
}

func main() {
	tellMeTypes()
}
