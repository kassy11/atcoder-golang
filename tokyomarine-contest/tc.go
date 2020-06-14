package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	lamps := make([]int, n)
	for i:=0; i<n; i++{
		fmt.Scan(&lamps[i])
	}

	results := make([]int, n)
	for i:=0; i<k; i++{
		// ランプの数の計算
		for l:=0; l<n; l++{
			if lamps[l] >= 1{
				results[l]++
			}
			for j:=(l+1) - lamps[l] + 1; j<=(l+1)+lamps[l]; j++{
				if j>=0 && j<n{
					results[j]+=2
				}
			}
		}

		// 計算語のランプの数を置き換える
		for m:=0; m<n; m++{
			lamps[m] = results[m]
		}
	}

	for i:=0; i<n; i++{
		fmt.Println(lamps[i])
	}

}