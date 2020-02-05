package main

import (
    "fmt"
//    "image"
//    _ "image/jpeg"
    "os"
		"./pictFunc"
)

func main(){
	fmt.Println(len(os.Args))

	hogera := pictFunc.Pict{}
	hogera.Load(os.Args[1])

	var ps []pictFunc.Pict
	ps = make([]pictFunc.Pict, len(os.Args) - 1)

	for i := 0; i < len(ps); i++ {
		fmt.Printf("num : %d, ", i + 1);
		ps[i] = pictFunc.Pict{}
		ps[i].Load(os.Args[i + 1])
	}

	fmt.Println(len(ps))

	ans := pictFunc.ExposureAve(ps)
	ans.Save("hoge.png")

//	ans := pictFunc.MyGaus(hogera, 101)
//	ans := pictFunc.Mylpls_pr(hogera)
//	ans.Save(os.Args[2])
}
