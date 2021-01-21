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
	const max = 1024 * 1024 * 1024
	var buf = make([]byte, max)
	sc.Buffer(buf, max)
}

func main() {
	NK := SplitIntlist(NextStr(sc))
	N := NK[0]
	K := NK[1]
	// 二次元配列で取得しておく
	tList := [][]int64{}
	for range make([]int, N) {
		l := SplitIntlist(NextStr(sc))
		tList = append(tList, l)
	}
	// 都市のスライス　1から始まるので最初の要素は不要
	indexex := []int64{}
	for i := range make([]int, N) {
		indexex = append(indexex, int64(i))
	}
	indexex = indexex[1:]
	count := 0
	perm := Permute(indexex)
	for i := range perm {
		total := int64(0)
		for j := range perm[i] {
			from := perm[i][j]
			// 1から最初の町の距離を追加する
			if j == 0 {
				total += tList[from][0]
			}
			to := int64(0)
			isEnd := j == len(perm[i])-1
			if isEnd {
				to = 0
			} else {
				to = perm[i][j+1]
			}
			total += tList[from][to]
		}
		if total == K {
			count++
		}
	}
	fmt.Println(count)
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
func NextInt(sc *bufio.Scanner) int64 {
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
func SplitIntlist(s string) []int64 {
	strList := strings.Split(s, " ")
	return StrListToIntList(strList)
}

// SplitFloatList 文字列を空白区切りの小数値に変換して返却
func SplitFloatList(s string) []float64 {
	strList := strings.Split(s, " ")
	return StrListToFloatList(strList)
}

// StrListToIntList string型のスライスを渡してint型の配列に変換
func StrListToIntList(strList []string) (intList []int64) {
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

// StrToInt string型をint64型に変換
func StrToInt(s string) int64 {
	i, _ := strconv.Atoi(s)
	return int64(i)
}

// StrToFloat string型をfloat64型に変換
func StrToFloat(s string) float64 {
	i, _ := strconv.ParseFloat(s, 10)
	return i
}

// Sort int型スライスの並び替え
func Sort(slice []int64, order string) []int64 {
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
func FindMaxAndMin(slice []int64) (max, min int64) {
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

// Sum 合計値を返す
func Sum(slice []int64) (sum int64) {
	for _, i := range slice {
		sum += i
	}
	return sum
}

// Permutation Pの値を計算
func Permutation(n int64, k int64) *big.Int {
	v := big.NewInt(1)
	if 0 < k && k <= n {
		for i := int64(0); i < k; i++ {
			k := big.NewInt(int64(n - i))
			v = v.Mul(v, k)
		}
	} else if k > n {
		v = big.NewInt(0)
	}
	return v
}

// Factorial Fの値を計算
func Factorial(n int64) *big.Int {
	return Permutation(n, n-1)
}

// Combination Cの計算
func Combination(n int64, k int64) *big.Int {
	child := Permutation(n, k)
	mother := Factorial(k)
	return child.Div(child, mother)
}

// Homogeneous Hの計算
func Homogeneous(n int64, k int64) *big.Int {
	return Combination(n+k-1, k)
}

// Permute 順列を返す。順列は並び順も考慮する
func Permute(nums []int64) [][]int64 {
	n := factorial(len(nums))
	ret := make([][]int64, 0, n)
	permute(nums, &ret)
	return ret
}

func permute(nums []int64, ret *[][]int64) {
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

func makeCopy(nums []int64) []int64 {
	return append([]int64{}, nums...)
}

func factorial(n int) int {
	ret := 1
	for i := 2; i <= n; i++ {
		ret *= i
	}
	return ret
}
