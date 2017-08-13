# 한글 음소 분리기
한글을 초성, 중성, 종성으로 분리하여줍니다. 또한, 초성, 중성, 종성을 조합하여 글자를 만들 수 있습니다.

# Docs
## `SeparateToRune(r rune) ([]rune, error)`
한글 한 글자를 []rune으로 나눈다.
## `SeparateToInt(r rune) ([]int, error)`
한글 한 글자를 []int로 나눈다.
## `BuildOfInt(f, s, t int) (rune, error)`
초성, 중성, 종성을 int로 받아 한 글자로 조합하여 rune으로 반환한다.
## `BuildOfRune(f, s, t rune) (rune, error)`
초성, 중성, 종성을 rune으로 받아 한 글자로 조합하여 rune으로 반환한다.
## `Choseong(n int) (rune, error)`
n번째 초성을 rune으로 반환한다.
## `Jungseong(n int) (rune, error)`
n번째 중성을 rune으로 반환한다.
## `Jongseong(n int) (rune, error)`
n번째 종성을 rune으로 반환한다.
## `IndexOfChoseong(r rune) (int, error)`
초성의 인덱스를 반환한다. (없을 시 -1 반환)
## `IndexOfJungseong(r rune) (int, error)`
중성의 인덱스를 반환한다. (없을 시 -1 반환)
## `IndexOfJongseong(r rune) (int, error)`
종성의 인덱스를 반환한다. (없을 시 -1 반환)
## `IsHangulLetter(r rune) bool`
초성이아닌 한글인지 판단하여 bool으로 반환.
## `IsHangul(r rune) bool`
초성 여부와 상관없이 한글인지 판단하여 bool으로 반환.

# 참고
## 한글 표
| Idx | 초성 | 중성 | 종성 | Idx | 초성 | 중성 | 종성 |
|:---:|:----:|:----:|:----:|:---:|:----:|:----:|:----:|
|  0  |  ㄱ  |  ㅏ  | 없음 |  14 |  ㅊ  |  ㅜ  |  ㄿ  |
|  1  |  ㄲ  |  ㅐ  |  ㄱ  |  15 |  ㅋ  |  ㅝ  |  ㅀ  |
|  2  |  ㄴ  |  ㅑ  |  ㄲ  |  16 |  ㅌ  |  ㅞ  |  ㅁ  |
|  3  |  ㄷ  |  ㅒ  |  ㄳ  |  17 |  ㅍ  |  ㅟ  |  ㅂ  |
|  4  |  ㄸ  |  ㅓ  |  ㄴ  |  18 |  ㅎ  |  ㅡ  |  ㅄ  |
|  5  |  ㄹ  |  ㅔ  |  ㄵ  |  19 |   -  |  ㅢ  |  ㅅ  |
|  6  |  ㅁ  |  ㅕ  |  ㄶ  |  20 |   -  |  ㅣ  |  ㅆ  |
|  7  |  ㅂ  |  ㅖ  |  ㄷ  |  21 |   -  |   -  |  ㅇ  |
|  8  |  ㅃ  |  ㅗ  |  ㄹ  |  22 |   -  |   -  |  ㅈ  |
|  9  |  ㅅ  |  ㅠ  |  ㄺ  |  23 |   -  |   -  |  ㅊ  |
|  10 |  ㅆ  |  ㅘ  |  ㄻ  |  24 |   -  |   -  |  ㅋ  |
|  11 |  ㅇ  |  ㅛ  |  ㄼ  |  25 |   -  |   -  |  ㅌ  |
|  12 |  ㅈ  |  ㅙ  |  ㄽ  |  26 |   -  |   -  |  ㅍ  |
|  13 |  ㅉ  |  ㅘ  |  ㄾ  |  27 |   -  |   -  |  ㅎ  |