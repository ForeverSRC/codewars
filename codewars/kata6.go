package codewars

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

func NbMonths(startPriceOld, startPriceNew, savingperMonth int, percentLossByMonth float64) [2]int {
	priceOld := float64(startPriceOld)
	money := 0
	month := 0
	priceNew := float64(startPriceNew)

	for {
		// month begin
		total := float64(money) + priceOld
		if total-priceNew >= 0 {
			return [2]int{month, int(total - priceNew + 0.5)}
		}

		// month end
		month++
		if month > 0 && month%2 == 0 {
			percentLossByMonth += 0.5
		}

		priceOld = priceOld * (1 - percentLossByMonth/100)
		priceNew = priceNew * (1 - percentLossByMonth/100)
		money += savingperMonth

	}

}

func Thirt(n int) int {
	digit := make(map[int]int, 10)
	res := doThirt(digit, n)
	for {
		tmp := doThirt(digit, res)
		if tmp == res {
			return tmp
		} else {
			res = tmp
		}
	}

}

func doThirt(digit map[int]int, num int) int {
	res := 0
	for i := 0; num > 0; num /= 10 {
		d := num % 10
		if c, ok := digit[i]; ok {
			res += d * c
		} else {
			digit[i] = int(math.Pow10(i)) % 13
			res += digit[i] * d
		}

		i++
	}

	return res
}

func Revrot(s string, n int) string {
	l := len(s)
	if n <= 0 || l == 0 || l < n {
		return ""
	}

	cn := l / n

	slices := make([]string, 0, l/n+1)
	index := 0

	for i := 0; i < cn; i++ {
		tmp := doRevrot(s[index : index+n])
		slices = append(slices, tmp)
		index += n
	}

	return strings.Join(slices, "")

}

func doRevrot(s string) string {
	sum := 0
	for _, c := range s {
		tmp, _ := strconv.Atoi(string(c))
		sum += tmp * tmp * tmp
	}

	if sum%2 == 0 {
		return reverse(s)
	} else {
		return s[1:] + s[0:1]
	}
}

func reverse(s string) string {
	a := []rune(s)
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return string(a)
}

func FindUniq(arr []float32) float32 {
	hash := make(map[float32]int, 2)
	for _, c := range arr {
		if t, ok := hash[c]; ok {
			hash[c] = t + 1
		} else {
			hash[c] = 1
		}
	}

	ans := float32(0)
	for k, v := range hash {
		if v == 1 {
			ans = k
			break
		}
	}

	return ans

}

func High(s string) string {
	ss := strings.Split(s, " ")
	ans := ""
	smax := int32(-1)
	for _, str := range ss {
		sc := calc(str)
		if sc > smax {
			ans = str
			smax = sc
		}
	}

	return ans
}

func calc(s string) int32 {
	score := int32(0)
	for _, c := range s {
		score += c - 'a' + 1
	}

	return score

}

func ToNato(words string) string {
	nato := []string{"Alfa", "Bravo", "Charlie", "Delta", "Echo", "Foxtrot", "Golf", "Hotel", "India", "Juliett", "Kilo", "Lima", "Mike", "November", "Oscar", "Papa", "Quebec", "Romeo", "Sierra", "Tango", "Uniform", "Victor", "Whiskey", "Xray", "Yankee", "Zulu"}
	res := make([]string, 0, len(words))
	for _, c := range words {
		if c == ' ' {
			continue
		}

		if c >= 'a' && c <= 'z' {
			res = append(res, nato[c-'a'])
		} else if c >= 'A' && c <= 'Z' {
			res = append(res, nato[c-'A'])
		} else {
			res = append(res, string(c))
		}
	}

	return strings.Join(res, " ")
}

func solve(str string) int {
	max := int32(0)
	tmp := int32(0)

	for _, c := range str {
		if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' {
			if tmp > max {
				max = tmp
			}

			tmp = 0
		} else {
			tmp += c - 'a' + 1
		}
	}

	if tmp > max {
		max = tmp
	}

	return int(max)
}

//-----------------------------

func Gcdi(x, y int) int {
	if x < 0 {
		x = -1 * x
	}

	if y < 0 {
		y = -1 * y
	}

	for x != y {
		if x > y {
			x -= y
		} else {
			y -= x
		}
	}

	return x
}

func Som(x, y int) int {
	return x + y
}
func Maxi(x, y int) int {
	if x >= y {
		return x
	} else {
		return y
	}
}
func Mini(x, y int) int {
	if x <= y {
		return x
	} else {
		return y
	}
}
func Lcmu(x, y int) int {
	if x < 0 {
		x = -1 * x
	}

	if y < 0 {
		y = -1 * y
	}

	return x * y / Gcdi(x, y)
}

type FParam func(int, int) int

func OperArray(f FParam, arr []int, init int) []int {
	n := len(arr)
	r := make([]int, n)
	r[0] = f(init, arr[0])
	for i := 1; i < n; i++ {
		r[i] = f(r[i-1], arr[i])
	}

	return r
}

//-------------------------

func PosAverage(s string) float64 {
	strs := strings.Split(s, ", ")
	n := len(strs)
	if n <= 0 {
		return 0.0
	}

	l := len(strs[0])
	num := 0
	for i := 0; i < l; i++ {
		num += calNum(strs, n, i, l)
	}

	return float64(num) / float64(n*(n-1)/2*l) * 100.0
}

func calNum(strs []string, n int, i int, l int) int {
	hash := make(map[string]int, l)
	for j := 0; j < n; j++ {
		if num, ok := hash[string(strs[j][i])]; ok {
			hash[string(strs[j][i])] = num + 1
		} else {
			hash[string(strs[j][i])] = 1
		}
	}

	sum := 0
	for _, v := range hash {
		if v > 1 {
			sum += v * (v - 1) / 2
		}
	}

	return sum

}

type SortableIntSlice []int

func (s SortableIntSlice) Len() int           { return len(s) }
func (s SortableIntSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s SortableIntSlice) Less(i, j int) bool { return s[i] < s[j] }

func SolveC(arr []int) int {
	ss := SortableIntSlice(arr)
	sort.Sort(ss)
	ans := 0
	for ss[1] > 0 && ss[2] > 0 {
		ans++
		ss[2] -= 1
		ss[1] -= 1
		sort.Sort(ss)
	}

	return ans
}

func Race(v1, v2, g int) [3]int {
	if v1 >= v2 {
		return [3]int{-1, -1, -1}
	}

	delta := v2 - v1
	gs := 3600 * g
	sec := gs / delta
	if gs%delta > 0 {
		sec += 1
	}

	h := sec / 3600
	m := (sec % 3600) / 60
	s := (sec % 3600) % 60

	return [3]int{h, m, s}

}

func LongestConsec(strarr []string, k int) string {
	n := len(strarr)
	_, s := calLength(0, n, strarr, k)
	return s
}

func calLength(index int, n int, strarr []string, l int) (int, string) {
	if index >= n {
		return 0, ""
	}
	ans := 0
	in := index

	if l == 1 {
		for i := index; i < n; i++ {
			if len(strarr[i]) > ans {
				ans = len(strarr[i])
				in = i
			}
		}

		return ans, strarr[in]
	}

	st := ""
	for i := index; i < n; i++ {
		leng, s := calLength(i+1, n, strarr, l-1)
		le := len(strarr[i]) + leng
		if le > ans {
			ans = le
			st = strarr[i] + s
		}

	}

	return ans, st
}

func Multiple3And5(number int) int {
	if number < 0 {
		return 0
	}

	sum := 0

	for i := 0; i < number; i++ {
		if i%15 == 0 || i%5 == 0 || i%3 == 0 {
			sum += i
		}

	}

	return sum

}
