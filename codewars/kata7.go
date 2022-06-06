package codewars

import (
	"bytes"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func HighAndLow(in string) string {
	nums := strings.Split(in, " ")
	min := math.MaxInt32
	max := math.MinInt32

	for _, n := range nums {
		ni, _ := strconv.Atoi(n)
		if ni > max {
			max = ni
		}

		if ni < min {
			min = ni
		}
	}

	return fmt.Sprintf("%d %d", max, min)
}

func Divisors(n int) int {
	r := n
	ans := 0
	for l := 1; ; l++ {
		if n%l == 0 {
			r = n / l
			if l == r {
				ans++
			} else {
				ans += 2
			}
		}

		if l+1 >= r {
			break
		}
	}

	return ans
}

func PrinterError(s string) string {
	t := len(s)
	var r int
	for _, c := range s {
		if c > 'm' {
			r++
		}
	}

	return fmt.Sprintf("%d/%d", r, t)
}

func SquareSum(numbers []int) int {
	cache := make(map[int]int, 100)
	var res int
	for _, n := range numbers {
		if c, ok := cache[n]; ok {
			res += c
		} else {
			tmp := n * n
			res += tmp
			cache[n] = tmp
		}
	}

	return res
}

func MultiTable(number int) string {
	l := make([]string, 10, 10)
	var tmp int
	for i := 1; i <= 10; i++ {
		tmp += number
		l[i-1] = fmt.Sprintf("%d * %d = %d", i, number, tmp)
	}

	return strings.Join(l, "\n")
}

func GetSize(w, h, d int) [2]int {
	return [2]int{2 * (w*h + w*d + d*h), w * h * d}
}

func PartList(arr []string) string {
	n := len(arr)
	res := make([]string, n-1)
	for i := 1; i < n; i++ {
		res[i-1] = "(" + strings.Join(arr[0:i], " ") + ", " + strings.Join(arr[i:], " ") + ")"
	}

	return strings.Join(res, "")

}

func RowSumOddNumbers(n int) int {
	return n * n * n
}

func TwoToOne(s1 string, s2 string) string {
	set := make(map[rune]struct{}, 26)
	for _, r := range s1 {
		set[r] = struct{}{}
	}

	for _, r := range s2 {
		set[r] = struct{}{}
	}

	res := make([]string, 0, 26)

	for i := 'a'; i <= 'z'; i++ {

		if _, ok := set[i]; ok {
			res = append(res, string(i))
		}
	}

	return strings.Join(res, "")
}

func Solve(arr []int) []int {
	set := make(map[int]int, 10)

	for _, i := range arr {

		if v, ok := set[i]; ok {
			set[i] = v + 1
		} else {
			set[i] = 1
		}
	}

	res := make([]int, 0, len(arr))

	for _, i := range arr {
		if v := set[i]; v == 1 {
			res = append(res, i)
		} else {
			set[i] = v - 1
		}
	}

	return res

}

func NewAvg(arr []float64, navg float64) int64 {
	n := len(arr)
	sum := navg * float64(n+1)
	for _, num := range arr {
		sum -= num
	}

	if sum < 0 {
		return -1
	}

	return int64(sum + 0.5)
}

func FizzBuzzCuckooClock(time string) string {
	t := strings.Split(time, ":")
	h, _ := strconv.Atoi(strings.TrimPrefix(t[0], "0"))
	m, _ := strconv.Atoi(strings.TrimPrefix(t[1], "0"))

	if m == 0 {
		if h == 0 {
			return "Cuckoo" + strings.Repeat(" Cuckoo", 11)
		}

		if h > 12 {
			h -= 12
		}

		return "Cuckoo" + strings.Repeat(" Cuckoo", h-1)
	}

	if m == 30 {
		return "Cuckoo"
	}

	ans := ""
	if m%3 == 0 {
		ans += "Fizz"
	}

	if m%5 == 0 {
		if ans != "" {
			ans += " Buzz"
		} else {
			ans += "Buzz"
		}
	}

	if ans != "" {
		return ans
	} else {
		return "tick"
	}

}

func OverTheRoad(address int, n int) int {
	return 2*n + 1 - address
}

func ReverseWords(str string) string {
	strs := strings.Split(str, " ")
	res := make([][]byte, 0, len(strs))
	for _, s := range strs {
		n := len(s)
		tmp := make([]byte, n)
		for i := 0; i < n; i++ {
			tmp[n-1-i] = s[i]
		}
		res = append(res, tmp)
	}

	return string(bytes.Join(res, []byte(" ")))
}

func Add(p int) func(int) int {
	return func(i int) int {
		return i + p
	}
}

func SumCubes(n int) int {
	sum := (1 + n) * n / 2
	return sum * sum
}

func Angle(n int) int {
	return 180 * (n - 2)
}

func PrevMultOfThree(n int) interface{} {
	for ; n > 0; n /= 10 {
		if n%3 == 0 {
			return n
		}
	}

	return nil
}

func Accum(s string) string {
	ans := make([]string, 0, len(s))
	for i, c := range s {
		sc := string(c)
		sc = strings.ToUpper(sc)
		ans = append(ans, sc+strings.Repeat(strings.ToLower(sc), i))
	}

	return strings.Join(ans, "-")
}

func Factorial(n int) int {
	ans := 1
	for i := 1; i <= n; i++ {
		ans *= i
	}

	return ans
}

func CountRedBeads(n int) int {
	if n < 2 {
		return 0
	}

	return 2 * (n - 1)
}

func Solve2(s string) rune {
	hash := make(map[rune]*[2]int)
	for i, c := range s {
		if index, ok := hash[c]; ok {
			index[1] = i
		} else {
			hash[c] = &[2]int{i, 0}
		}
	}

	var ans = 'z' + 1
	var maxVal = -1

	for k, v := range hash {
		val := v[1] - v[0]
		if val < 0 {
			val = 0
		}

		if maxVal < val {
			ans = k
			maxVal = val
		} else if maxVal == val {
			if ans > k {
				ans = k
			}
		}

	}

	return ans
}

func IsTriangle(a, b, c int) bool {
	return a+b > c && a+c > b && b+c > a
}

func Scale(s string, k, n int) string {
	ss := strings.Split(s, "\n")
	l := len(s)
	if l == 0 {
		return ""
	}

	res := make([]string, 0, l*n)

	for _, str := range ss {
		temp := ""
		for _, c := range str {
			tmp := strings.Repeat(string(c), k)
			temp += tmp
		}

		for i := 0; i < n; i++ {
			res = append(res, temp)
		}
	}

	return strings.Join(res, "\n")
}

func Solve3(a, b int) []int {
	if a == 0 || b == 0 {
		return []int{a, b}
	}

	if a >= 2*b {
		return Solve3(a-2*b, b)
	}

	if b >= 2*a {
		return Solve3(a, b-2*a)
	}

	return []int{a, b}

}
