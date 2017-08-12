package hangul

import (
	"fmt"
	"unicode/utf8"
)

var (
	choseong  = [...]rune{'ㄱ', 'ㄲ', 'ㄴ', 'ㄷ', 'ㄸ', 'ㄹ', 'ㅁ', 'ㅂ', 'ㅃ', 'ㅅ', 'ㅆ', 'ㅇ', 'ㅈ', 'ㅉ', 'ㅊ', 'ㅋ', 'ㅌ', 'ㅍ', 'ㅎ'}
	jungseong = [...]rune{'ㅏ', 'ㅐ', 'ㅑ', 'ㅒ', 'ㅓ', 'ㅔ', 'ㅕ', 'ㅖ', 'ㅗ', 'ㅘ', 'ㅙ', 'ㅚ', 'ㅛ', 'ㅜ', 'ㅝ', 'ㅞ', 'ㅟ', 'ㅠ', 'ㅡ', 'ㅢ', 'ㅣ'}
	jongseong = [...]rune{0, 'ㄱ', 'ㄲ', 'ㄳ', 'ㄴ', 'ㄵ', 'ㄶ', 'ㄷ', 'ㄹ', 'ㄺ', 'ㄻ', 'ㄼ', 'ㄽ', 'ㄾ', 'ㄿ', 'ㅀ', 'ㅁ', 'ㅂ', 'ㅄ', 'ㅅ', 'ㅆ', 'ㅇ', 'ㅈ', 'ㅊ', 'ㅋ', 'ㅌ', 'ㅍ', 'ㅎ'}
)

func SeparateToRune(r rune) ([]rune, error) {
	if IsHangul(r) == true {
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
	if IsHangul(r) == true {
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
	var r rune = rune(0xAC00 + 28*21*f + 28*s + t)
	if IsHangul(r) == true {
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

	var r rune = rune(0xAC00 + 28*21*iF + 28*iS + iT)
	if IsHangul(r) == true {
		return r, nil
	} else {
		return 0, fmt.Errorf("'%c' is not Hangul.", r)
	}
}

func Choseong(n int) rune {
	return choseong[n]
}

func Jungseong(n int) rune {
	return jungseong[n]
}

func Jongseong(n int) rune {
	return jongseong[n]
}

func IndexOfChoseong(r rune) (int, error) {
	for i, v := range choseong {
		if v == r {
			return i, nil
		}
	}
	return 0, fmt.Errorf("Choseong does not contain '%c'", r)
}

func IndexOfJungseong(r rune) (int, error) {
	for i, v := range jungseong {
		if v == r {
			return i, nil
		}
	}
	return 0, fmt.Errorf("Jungseong does not contain '%c", r)
}

func IndexOfJongseong(r rune) (int, error) {
	for i, v := range jongseong {
		if v == r {
			return i, nil
		}
	}
	return 0, fmt.Errorf("Jongseong does not contain '%c'", r)
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
