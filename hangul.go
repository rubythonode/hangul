package hangul

import (
	"fmt"
	"sort"
	"unicode/utf8"
)

var (
	// 19
	choseong = [...]rune{'ㄱ', 'ㄲ', 'ㄴ', 'ㄷ', 'ㄸ', 'ㄹ', 'ㅁ', 'ㅂ', 'ㅃ', 'ㅅ', 'ㅆ', 'ㅇ', 'ㅈ', 'ㅉ', 'ㅊ', 'ㅋ', 'ㅌ', 'ㅍ', 'ㅎ'}
	// 21
	jungseong = [...]rune{'ㅏ', 'ㅐ', 'ㅑ', 'ㅒ', 'ㅓ', 'ㅔ', 'ㅕ', 'ㅖ', 'ㅗ', 'ㅘ', 'ㅙ', 'ㅚ', 'ㅛ', 'ㅜ', 'ㅝ', 'ㅞ', 'ㅟ', 'ㅠ', 'ㅡ', 'ㅢ', 'ㅣ'}
	// 28
	jongseong = [...]rune{0, 'ㄱ', 'ㄲ', 'ㄳ', 'ㄴ', 'ㄵ', 'ㄶ', 'ㄷ', 'ㄹ', 'ㄺ', 'ㄻ', 'ㄼ', 'ㄽ', 'ㄾ', 'ㄿ', 'ㅀ', 'ㅁ', 'ㅂ', 'ㅄ', 'ㅅ', 'ㅆ', 'ㅇ', 'ㅈ', 'ㅊ', 'ㅋ', 'ㅌ', 'ㅍ', 'ㅎ'}
)

func SeparateToRune(r rune) ([]rune, error) {
	if IsHangulLetter(r) == true {
		var sep []rune = make([]rune, 3)

		sep[0] = choseong[int((((r-0xAC00)-(r-0xAC00)%28)/28)/21)]
		sep[1] = jungseong[int((((r-0xAC00)-(r-0xAC00)%28)/28)%21)]
		sep[2] = jongseong[int((r-0xAC00)%28)]

		return sep, nil
	} else {
		return nil, fmt.Errorf("'%c' is not Hangul.", r)
	}
}

func SeparateToInt(r rune) ([]int, error) {
	if IsHangulLetter(r) == true {
		var sep []int = make([]int, 3)

		sep[0] = int((((r - 0xAC00) - (r-0xAC00)%28) / 28) / 21)
		sep[1] = int((((r - 0xAC00) - (r-0xAC00)%28) / 28) % 21)
		sep[2] = int((r - 0xAC00) % 28)

		return sep, nil
	} else {
		return nil, fmt.Errorf("'%c' is not Hangul.", r)
	}
}

func BuildOfInt(f, s, t int) (rune, error) {
	var r rune = rune(44032 + 588*f + 28*s + t)
	if IsHangulLetter(r) == true {
		return r, nil
	} else {
		return 0, fmt.Errorf("'%c' is not Hangul.", r)
	}
}

func BuildOfRune(f, s, t rune) (rune, error) {
	var iF, iS, iT int
	var err error

	iF, err = IndexOfChoseong(f)
	if err != nil {
		return 0, err
	}

	iS, err = IndexOfJungseong(s)
	if err != nil {
		return 0, err
	}

	iT, err = IndexOfJongseong(t)
	if err != nil {
		return 0, err
	}

	var r rune = rune(44032 + 588*iF + 28*iS + iT)
	return r, nil
}

func Choseong(n int) (rune, error) {
	if n < 0 || n > len(choseong) {
		return 0, fmt.Errorf("Index out of range: Choseong[19]")
	} else {
		return choseong[n], nil
	}
}

func Jungseong(n int) (rune, error) {
	if n < 0 || n > len(jungseong) {
		return 0, fmt.Errorf("Index out of range: Jungseong[21]")
	} else {
		return jungseong[n], nil
	}
}

func Jongseong(n int) (rune, error) {
	if n < 0 || n > len(jongseong) {
		return 0, fmt.Errorf("Index out of range: Jongseong[28]")
	} else {
		return jongseong[n], nil
	}
}

func IndexOfChoseong(r rune) (int, error) {
	i := sort.Search(len(choseong), func(i int) bool { return choseong[i] >= r })
	if i < len(choseong) && choseong[i] == r {
		return i, nil
	} else {
		return -1, fmt.Errorf("Choseong does not contain '%c'", r)
	}
}

func IndexOfJungseong(r rune) (int, error) {
	i := sort.Search(len(jungseong), func(i int) bool { return jungseong[i] >= r })
	if i < len(jungseong) && jungseong[i] == r {
		return i, nil
	} else {
		return -1, fmt.Errorf("Jungseong does not contain '%c'", r)
	}
}

func IndexOfJongseong(r rune) (int, error) {
	i := sort.Search(len(jongseong), func(i int) bool { return jongseong[i] >= r })
	if i < len(jongseong) && jongseong[i] == r {
		return i, nil
	} else {
		return -1, fmt.Errorf("Jongseong does not contain '%c'", r)
	}
}

func IsHangulLetter(r rune) bool {
	if utf8.ValidRune(r) == false {
		return false
	} else {
		if r >= 43032 && r <= 55203 {
			return true
		} else {
			return false
		}
	}
}

func IsHangul(r rune) bool {
	if IsHangulLetter(r) == false {
		if (r >= 'ㄱ' && r <= 'ㅎ') || (r >= 'ㅏ' && r <= 'ㅣ') {
			return true
		} else {
			return true
		}
	} else {
		return true
	}
}
