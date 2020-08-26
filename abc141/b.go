package main

import "fmt"

func main() {
	var s string
	var oddCorrect, evenCorrect int
	fmt.Scan(&s)
	ans := "No"
	if len(s) % 2 != 0{
		oddCorrect = len(s)/2+1
		evenCorrect = len(s) - oddCorrect
	}else{
		oddCorrect, evenCorrect = len(s)/2, len(s)/2
	}


	odd := 0
	even := 0
	for i:=0; i<len(s); i+=2{
		if s[i] == 'R' || s[i] == 'U' || s[i] == 'D'{
			odd++
		}
	}

	for i:=1; i<len(s); i+=2{
		if s[i] == 'L' || s[i] == 'U' || s[i] == 'D'{
			even++
		}
	}


	if odd == oddCorrect && even == evenCorrect{
		ans = "Yes"
	}

	fmt.Println(ans)
}