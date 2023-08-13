package main

import "math"

// 字符串逐个分析
func myAtoi(s string) int {
	result, sign, i, n := 0, 1, 0, len(s)
	const MinInt32, MaxInt32 = -1 << 31, 1<<31 - 1

	// 1. 如果开始字符是空格, 则一直累增索引
	for ; i < n && s[i] == ' '; i++ {
	}

	// 如果索引大于/等于字符串长度, 则返回0(说明是空字符串)
	if i >= n {
		return 0
	}

	// 2. 当字符不是空格, 且是+/-时候, 符号位赋值;
	switch s[i] {
	case '+':
		i++
	case '-':
		i++
		sign = -1
	}

	// 如果都不是, 继续下面的逻辑
	for ; i < n; i++ {
		// unicode: 0 => 48, 9 => 57
		// 如果不是0-9, 则跳出循环, 直接到return
		if s[i] < 48 || s[i] > 57 {
			break
		}

		// 否则, 按公式 f(x) = f(x) * 10 + x 去计算
		// 比如: 394 =>
		// 1. 3 = 0 * 3 + 3
		// 2. 39 = 3 * 10 + 9
		// 3. 394 = 39 * 10 + 4
		result = result*10 + int(s[i]-'0')

		final := sign * result

		// 判断是否越界
		if final < MinInt32 {
			return MinInt32
		}
		if final > MaxInt32 {
			return MaxInt32
		}
	}

	// 以上逻辑都走通了, 返回最终值
	return sign * result
}

/* *********************** 有限状态机 ************************ */

const (
	START_STATE  = "start"
	SIGN_STATE   = "sign"
	NUMBER_STATE = "number"
	END_STATE    = "end"
)

type State interface {
	Enter(b rune)
	Transfer(b rune) string
}
type StartState struct{}

func (s *StartState) Enter(b rune) {}
func (s *StartState) Transfer(b rune) string {
	switch b {
	case '+', '-':
		return SIGN_STATE
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return NUMBER_STATE
	case ' ':
		return START_STATE
	default:
		return END_STATE
	}
}

type SignState struct {
	sign int
}

func (s *SignState) Enter(b rune) {
	if b == '+' {
		s.sign = 1
	} else {
		s.sign = -1
	}
}
func (s *SignState) Transfer(b rune) string {
	switch b {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return NUMBER_STATE
	default:
		return END_STATE
	}
}

type NumberState struct {
	num int
}

func (s *NumberState) Enter(b rune) {
	s.num *= 10
	s.num += int(b - '0')
}
func (s *NumberState) Transfer(b rune) string {
	if s.num > math.MaxInt32 {
		return END_STATE
	}
	switch b {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return NUMBER_STATE
	default:
		return END_STATE
	}
}

type EndState struct{}

func (s *EndState) Enter(b rune)           {}
func (s *EndState) Transfer(b rune) string { return "end" }

func myAtoiByLimitedState(s string) int {
	mp := map[string]State{
		START_STATE:  &StartState{},
		SIGN_STATE:   &SignState{sign: 1},
		NUMBER_STATE: &NumberState{},
		END_STATE:    &EndState{},
	}
	state := mp[START_STATE]
	for _, b := range s {
		next := state.Transfer(b)
		if next == END_STATE {
			break
		}
		state = mp[next]
		state.Enter(b)
	}
	ans := interface{}(mp[SIGN_STATE]).(*SignState).sign * interface{}(mp[NUMBER_STATE]).(*NumberState).num
	if ans > math.MaxInt32 {
		ans = math.MaxInt32
	} else if ans < math.MinInt32 {
		ans = math.MinInt32
	}
	return ans
}
