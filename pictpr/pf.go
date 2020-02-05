package main

import (
    "fmt"
//    "image"
//    _ "image/jpeg"
    "os"
		"./pictFunc"
		"os/exec"
		"strconv"
		"strings"
		"bufio"
)

func main(){
//	fmt.Println(os.Args[1])
//	fmt.Println(os.Args[2])
//	hoge := pictFunc.Pict{}
//	fuga := pictFunc.Pict{}
//	hoge.Load(os.Args[1])
//	fuga.Load(os.Args[2])

//	ans1 := pictFunc.MozDiv(hoge, 10)
//	ans2 := pictFunc.MozDiv(fuga, 10)
//	ans3 := pictFunc.Dif(ans1, ans2)
//	ans3 := pictFunc.Difv(ans1, ans2)
//	ans3 := pictFunc.Difc(ans1, ans2, 100)

//	fmt.Println(ans3)

//		mkcheck()
mkscript()
//	mkgd()

//		glauncher()

//	ans1.Save("res1.png")
//	ans2.Save("res2.png")
//	ans3.Save("res3.png")
}

func mkcheck() {
	hoge := pictFunc.Pict{}
	hoge.Load(os.Args[1])
	hoge = pictFunc.MozDiv(hoge, 10)

	ddat := make([]uint64, len(os.Args) - 2)
	vmax := uint64(0)
	dura := 0
	ccnt := 0

	for i := 0; i < len(os.Args) - 2; i++ {
		fuga := pictFunc.Pict{}
		fuga.Load(os.Args[i + 2])
		fuga = pictFunc.MozDiv(fuga, 10)

		ans := pictFunc.Difc(hoge, fuga, 150)
		ddat[i] = ans

		if vmax < ans {
			vmax = ans
		}

		hoge = fuga
	}

	for i := 0; i < len(ddat); i++ {
		if vmax / 5 < ddat[i] {
			fmt.Println(os.Args[i + 2], i - dura)
			dura = i
			ccnt ++

			cmdstr := "convert +append " + os.Args[i] + " " + os.Args[i + 1] + " " +
				os.Args[i + 2] + " " + os.Args[i + 3] + " check_" + mk4num(i+2) + ".png"

			exec.Command("sh", "-c", cmdstr).Run()
		}
	}

	fmt.Println(ccnt)
}

func mk4num (num int) string {
	res := ""

	if num < 10 {
		res = "000" + strconv.Itoa(num)
	} else if num < 100 {
		res = "00" + strconv.Itoa(num)
	} else if num < 1000 {
		res = "0" + strconv.Itoa(num)
	} else {
		res = strconv.Itoa(num)
	}

	return res
}

func mkscript() {

	fmt.Println(os.Args[1])

	dura := 0
	for i := 2; i < len(os.Args); i++ {
		a := strings.Split(os.Args[i], ".")[0]
		sv := strings.Split(a, "_")[1]

		v, _ := strconv.Atoi(sv)

		fmt.Println(v, v - dura)
		dura = v
	}

}

func mkgd() {

	dura := 0
	for i := 1; i < len(os.Args); i++ {
		a := strings.Split(os.Args[i], ".")[0]
		sv := strings.Split(a, "_")[1]

		v, _ := strconv.Atoi(sv)

		fmt.Println(v, v - dura)
		dura = v
	}
}

func gdpers() []int {
	fp, _ := os.Open(os.Args[1])
	defer fp.Close()

	rtxs := make([]string, 0)

	scanner := bufio.NewScanner(fp)

	for scanner.Scan() {
		rtx := scanner.Text()
		if rtx != " " {
			rtxs = append(rtxs, rtx)
		}
	}

	dtxs := make([]int, len(rtxs) * 2)

	for i := 0; i < len(rtxs); i++ {
		rcs := strings.Split(rtxs[i], " ")
		dtxs[2 * i], _ = strconv.Atoi(rcs[0])
		dtxs[2 * i + 1], _ = strconv.Atoi(rcs[1])
	}

	return dtxs
}




func launcher() {
	hoge := pictFunc.Pict{}
	hoge.Load(os.Args[1])
	hoge = pictFunc.MozDiv(hoge, 10)
	pnum := hoge.Width * hoge.Height

	dura := 0
	ccnt := 0

	for i := 0; i < len(os.Args) - 2; i++ {
		fuga := pictFunc.Pict{}
		fuga.Load(os.Args[i + 2])
		fuga = pictFunc.MozDiv(fuga, 10)

		ans := pictFunc.Difc(hoge, fuga, 100)

		if pnum / 2 < int(ans) {
			fmt.Println(os.Args[i + 2], "dura", i - dura)
			dura = i
			ccnt ++
		}

		hoge = fuga
	}

	fmt.Println(ccnt)
}

func glauncher() {
	hoge := pictFunc.Pict{}
	hoge.Load(os.Args[1])
	hoge = pictFunc.MozDiv(hoge, 10)
	pnum := hoge.Width * hoge.Height

	for i := 0; i < len(os.Args) - 2; i++ {
		fuga := pictFunc.Pict{}
		fuga.Load(os.Args[i + 2])
		fuga = pictFunc.MozDiv(fuga, 10)

		ans := pictFunc.Difc(hoge, fuga, 100)

		if pnum / 2 < int(ans) {
			fmt.Println(os.Args[i + 2])
		}

		hoge = fuga
	}
}

func glauncher2() {
	hoge := pictFunc.Pict{}
	hoge.Load(os.Args[1])
	hogera := pictFunc.Ave(hoge)

	for i := 0; i < len(os.Args) - 2; i++ {
		fuga := pictFunc.Pict{}
		fuga.Load(os.Args[i + 2])
		fugara := pictFunc.Ave(fuga)

		ans := pictFunc.Sdif(hogera, fugara)
		fmt.Println(i, fugara[1], fugara[2], fugara[3], ans[1] + ans[2] + ans[3])

		hogera = fugara
	}
}
