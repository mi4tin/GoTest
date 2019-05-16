package main



func defertest1()int{
	i:=2
	defer func() {
		i++
	}()
	return  0
}

func defertest2()(i int){
	i=2
	defer func() {
		i++
	}()
	return  0
}