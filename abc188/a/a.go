package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"sort"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func init() {
	const max = 1024 * 256
	var buf = make([]byte, max)
	sc.Buffer(buf, max)
}

func main() {
	N := StrToInt(GetNextLine(sc))
	xList := make([]int, N)
	yList := make([]int, N)
	for i := range make([]int, 2) {
		xy := GetSpaceSplitIntlist(GetNextLine(sc))
		xList[i] = xy[0]
		yList[i] = xy[1]
	}
	fmt.Println(xList, yList)
}

// Reverse 文字列を反転
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// GetNextLine buinfo.Scanのポインタを渡し、標準入力の次の行を読み込み
// ex. sc := buinfo.NewScanner(os.stdin)
//      GetNextLine(sc)
func GetNextLine(sc *bufio.Scanner) string {
	sc.Scan()
	s := sc.Text()
	return strings.TrimSpace(s)
}

// GetSpaceSplitStrList 文字列を空白区切りの文字列のリストに変換して返却
func GetSpaceSplitStrList(s string) []string {
	return strings.Split(s, " ")
}

// GetSpaceSplitIntlist 文字列を空白区切りの整数値に変換して返却
func GetSpaceSplitIntlist(s string) []int {
	strList := strings.Split(s, " ")
	return StrListToIntList(strList)
}

// StrListToIntList string型のスライスを渡してint型の配列に変換
func StrListToIntList(strList []string) (intList []int) {
	for _, str := range strList {
		str = strings.TrimRight(str, "\n")
		i := StrToInt(str)
		intList = append(intList, i)
	}
	return
}

// StrToInt string型をint型に変換
func StrToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

// Sort int型スライスの並び替え
func Sort(slice []int, order string) []int {
	sort.SliceStable(slice, func(i, j int) bool {
		if order == "desc" {
			return slice[i] > slice[j]
		} else {
			return slice[i] < slice[j]
		}
	})
	return slice
}

// FindMaxAndMin 最大値最小値を返す
func FindMaxAndMin(slice []int) (max, min int) {
	max = slice[0]
	min = slice[0]
	for _, elm := range slice {
		if elm > max {
			max = elm
		}
		if elm < min {
			min = elm
		}
	}
	return max, min
}

// GetSum 合計値を返す
func GetSum(slice []int) (sum int) {
	for _, i := range slice {
		sum += i
	}
	return sum
}

func permutation(n int, k int) *big.Int {
	v := big.NewInt(1)
	if 0 < k && k <= n {
		for i := 0; i < k; i++ {
			k := big.NewInt(int64(n - i))
			v = v.Mul(v, k)
		}
	} else if k > n {
		v = big.NewInt(0)
	}
	return v
}

func factorial(n int) *big.Int {
	return permutation(n, n-1)
}

func combination(n int, k int) *big.Int {
	child := permutation(n, k)
	mother := factorial(k)
	return child.Div(child, mother)
}

func homogeneous(n int, k int) *big.Int {
	return combination(n+k-1, k)
}
