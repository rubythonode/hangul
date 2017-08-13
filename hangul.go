package hangul

import (
	"fmt"
	"unicode/utf8"
)

const (
	CHOSEONG_SIZE  = 19
	JUNGSEONG_SIZE = 21
	JONGSEONG_SIZE = 28
)

func SeparateToRune(r rune) ([]rune, error) {
	if IsHangulLetter(r) == true {
		var sep []rune = make([]rune, 3)
		var err error

		sep[0], err = Choseong(int((((r - 0xAC00) - (r-0xAC00)%JONGSEONG_SIZE) / JONGSEONG_SIZE) / JUNGSEONG_SIZE))
		if err != nil {
			return nil, err
		}
		sep[1], err = Jungseong(int((((r - 0xAC00) - (r-0xAC00)%JONGSEONG_SIZE) / JONGSEONG_SIZE) % JUNGSEONG_SIZE))
		if err != nil {
			return nil, err
		}
		sep[2], err = Jongseong(int((r - 0xAC00) % JONGSEONG_SIZE))
		if err != nil {
			return nil, err
		}

		return sep, nil
	} else {
		return nil, fmt.Errorf("'%c' is not Hangul.", r)
	}
}

func SeparateToInt(r rune) ([]int, error) {
	if IsHangulLetter(r) == true {
		var sep []int = make([]int, 3)

		sep[0] = int((((r - 0xAC00) - (r-0xAC00)%JONGSEONG_SIZE) / JONGSEONG_SIZE) / JUNGSEONG_SIZE)
		sep[1] = int((((r - 0xAC00) - (r-0xAC00)%JONGSEONG_SIZE) / JONGSEONG_SIZE) % JUNGSEONG_SIZE)
		sep[2] = int((r - 0xAC00) % JONGSEONG_SIZE)

		return sep, nil
	} else {
		return nil, fmt.Errorf("'%c' is not Hangul.", r)
	}
}

func BuildOfInt(f, s, t int) (rune, error) {
	var r rune = rune(0xAC00 + 588*f + JONGSEONG_SIZE*s + t)
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

	var r rune = rune(0xAC00 + 588*iF + JONGSEONG_SIZE*iS + iT)
	return r, nil
}

func Choseong(n int) (rune, error) {
	if n < 0 || n > CHOSEONG_SIZE {
		return 0, fmt.Errorf("Index out of range: Choseong[19]")
	} else {
		return rune(0x1100 + n), nil
	}
}

func Jungseong(n int) (rune, error) {
	if n < 0 || n > JUNGSEONG_SIZE {
		return 0, fmt.Errorf("Index out of range: Jungseong[21]")
	} else {
		return rune(0x1161 + n), nil
	}
}

func Jongseong(n int) (rune, error) {
	if n < 0 || n > JONGSEONG_SIZE {
		return 0, fmt.Errorf("Index out of range: Jongseong[28]")
	} else {
		return rune(0x11A7 + n), nil
	}
}

func IndexOfChoseong(r rune) (int, error) {
	if r >= rune(0x1100) && r <= rune(0x1113) {
		return int(r) - 0x1100, nil
	} else {
		return -1, fmt.Errorf("Choseong does not contain '%c'", r)
	}
}

func IndexOfJungseong(r rune) (int, error) {
	if r >= rune(0x1161) && r <= rune(0x1176) {
		return int(r) - 0x1161, nil
	} else {
		return -1, fmt.Errorf("Jungseong does not contain '%c'", r)
	}
}

func IndexOfJongseong(r rune) (int, error) {
	if r >= rune(0x11A7) && r <= rune(0x11C3) {
		return int(r) - 0x11A7, nil
	} else {
		return -1, fmt.Errorf("Jongseong does not contain '%c'", r)
	}
}

func IsHangulLetter(r rune) bool {
	if utf8.ValidRune(r) == false {
		return false
	} else {
		if r >= 0xAC00 && r <= 0xD7A3 {
			return true
		} else {
			return false
		}
	}
}

func IsHangul(r rune) bool {
	if IsHangulLetter(r) == false {
		if (r >= 'ㄱ' && r <= 'ㅎ') || (r >= 'ㅏ' && r <= 'ㅣ') || (r >= rune(0x1100) && r <= rune(0x1113)) || (r >= rune(0x1161) && r <= rune(0x1176)) || (r >= rune(0x11A7) && r <= rune(0x11C3)) {
			return true
		} else {
			return false
		}
	} else {
		return true
	}
}
