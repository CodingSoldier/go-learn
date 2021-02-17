package main

import "fmt"

func main(){

	//// if语句判断
	//const filename = "c-01.txt"
	//file, err := ioutil.ReadFile(filename)
	//if err != nil {
	//	fmt.Println(err)
	//}else {
	//	fmt.Printf("%s\n", file)
	//}


	//// if条件里可以赋值
	//// if条件里赋值的变量作用域只在if语句中
	//const filename = "D:\\go-learn\\go01\\abase\\c-01.txt"
	//if contents, err := ioutil.ReadFile(filename); err != nil {
	//	fmt.Println(err)
	//}else{
	//	fmt.Printf("%s\n", contents)
	//}


	// switch没有break
	score := 11
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprint("错误：%d", score))
	case score > 90:
		g = "good"
	case score >=60:
		g = "一般"
	default:
		g = "bad"
	}

	fmt.Printf(g)


}
