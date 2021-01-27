package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func init() {
	const max = 1024 * 1024 * 1024
	var buf = make([]byte, max)
	sc.Buffer(buf, max)
}

func main() {
	l := SplitIntlist(NextStr(sc))
	N := l[0]
	X := l[1]
	T := l[2]
	i := 0
	if N%X == 0 {
		i = N / X
	} else {
		i = N/X + 1
	}
	PrintLn(T * i)
}

// PrintLn fmt.Printlnのショート
func PrintLn(a ...interface{}) {
	fmt.Println(a...)
}

// PrintFloat float64を最大桁数まで表示
func PrintFloat(f float64) {
	fmt.Println(strconv.FormatFloat(f, 'f', -1, 64))
}

// Reverse 文字列を反転
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// NextStr buinfo.Scanのポインタを渡し、標準入力の次の行を読み込み
// ex. sc := buinfo.NewScanner(os.stdin)
//      GetNextLine(sc)
func NextStr(sc *bufio.Scanner) string {
	sc.Scan()
	s := sc.Text()
	return strings.TrimSpace(s)
}

// NextInt 次のラインをint型で得る
func NextInt(sc *bufio.Scanner) int {
	sc.Scan()
	s := sc.Text()
	return StrToInt(s)
}

// NextFlaot 次のラインをfloat64型で得る
func NextFlaot(sc *bufio.Scanner) float64 {
	sc.Scan()
	s := sc.Text()
	f, _ := strconv.ParseFloat(s, 10)
	return f
}

// SplitStrList 文字列を空白区切りの文字列のリストに変換して返却
func SplitStrList(s string) []string {
	return strings.Split(s, " ")
}

// SplitIntlist 文字列を空白区切りの整数値に変換して返却
func SplitIntlist(s string) []int {
	strList := strings.Split(s, " ")
	return StrListToIntList(strList)
}

// SplitFloatList 文字列を空白区切りの小数値に変換して返却
func SplitFloatList(s string) []float64 {
	strList := strings.Split(s, " ")
	return StrListToFloatList(strList)
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

// StrListToFloatList string型のスライスを渡してint型の配列に変換
func StrListToFloatList(strList []string) (floatList []float64) {
	for _, str := range strList {
		str = strings.TrimRight(str, "\n")
		i := StrToFloat(str)
		floatList = append(floatList, i)
	}
	return
}

// StrToInt string型をint型に変換
func StrToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return int(i)
}

// StrToFloat string型をfloat64型に変換
func StrToFloat(s string) float64 {
	i, _ := strconv.ParseFloat(s, 10)
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

// Min 最小値を算出
func Min(xs ...int) int {
	min := xs[0]
	for _, x := range xs[1:] {
		if min > x {
			min = x
		}
	}
	return min
}

// Max 最大値を算出
func Max(xs ...int) int {
	max := xs[0]
	for _, x := range xs[1:] {
		if max < x {
			max = x
		}
	}
	return max
}

// Sum 合計値を返す
func Sum(slice []int) (sum int) {
	for _, i := range slice {
		sum += i
	}
	return sum
}

// Permutation Pの値を計算
func Permutation(n int, k int) int {
	v := 1
	if 0 < k && k <= n {
		for i := 0; i < k; i++ {
			k := n - i
			v = v * k
		}
	} else if k > n {
		v = 0
	}
	return v
}

// Permute 順列を返す。順列は並び順も考慮する
func Permute(nums []int) [][]int {
	n := factorial(len(nums))
	ret := make([][]int, 0, n)
	permute(nums, &ret)
	return ret
}

func permute(nums []int, ret *[][]int) {
	*ret = append(*ret, makeCopy(nums))

	n := len(nums)
	p := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		p[i] = i
	}
	for i := 1; i < n; {
		p[i]--
		j := 0
		if i%2 == 1 {
			j = p[i]
		}

		nums[i], nums[j] = nums[j], nums[i]
		*ret = append(*ret, makeCopy(nums))
		for i = 1; p[i] == 0; i++ {
			p[i] = i
		}
	}
}

func makeCopy(nums []int) []int {
	return append([]int{}, nums...)
}

func factorial(n int) int {
	ret := 1
	for i := 2; i <= n; i++ {
		ret *= i
	}
	return ret
}

// MakeCopy コピーを作成
func MakeCopy(nums []int) []int {
	return append([]int{}, nums...)
}

// Factorial Fの値を計算
func Factorial(n int) int {
	return Permutation(n, n-1)
}

// Combination Cの計算
func Combination(n int, k int) int {
	child := Permutation(n, k)
	mother := Factorial(k)
	return child / mother
}

// Homogeneous Hの計算
func Homogeneous(n int, k int) int {
	return Combination(n+k-1, k)
}
