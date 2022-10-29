package section8

import (
	"fmt"
	"unicode/utf8"
)

// ---------------------------------------------------------
// EXERCISE: Rune Manipulator
//
//  Please read the comments inside the following code.
//
// EXPECTED OUTPUT
//  Please run the solution.
// ---------------------------------------------------------

func RuneManipulator() {
	words := []string{
		"cool",
		"güzel",
		"jīntiān",
		"今天",
		"read 🤓",
	}

	for _, word := range words {
		fmt.Printf("%q\n", word)

		// Print the byte and rune length of the strings
		// Hint: Use len and utf8.RuneCountInString
		fmt.Printf("\thas %d bytes and %d runes\n", len(word), utf8.RuneCountInString(word))

		// Print the bytes of the strings in hexadecimal
		// Hint: Use % x verb
		fmt.Printf("\tbytes: %x\n", word)

		// Print the runes of the strings in hexadecimal
		// Hint: Use % x verb
		for _, w := range word {
			fmt.Printf("\trunes: %x\n", w)
		}

		// Print the runes of the strings as rune literals
		// Hint: Use for range
		for _, w := range word {
			fmt.Printf("\t%q", w)
		}
		fmt.Println()
		fmt.Println()

		// Print the first rune and its byte size of the strings
		// Hint: Use utf8.DecodeRuneInString
		r, size := utf8.DecodeRuneInString(word)
		fmt.Printf("\tfirst :%q (%d bytes)\n", r, size)

		// Print the last rune of the strings
		// Hint: Use utf8.DecodeLastRuneInString
		r2, size2 := utf8.DecodeLastRuneInString(word)
		fmt.Printf("\tlast :%q (%d bytes)\n", r2, size2)

		// Slice and print the first two runes of the strings
		_, first := utf8.DecodeRuneInString(word)
		_, second := utf8.DecodeRuneInString(word[first:])
		// fmt.Printf("first %d, second %d\n", first, second)
		fmt.Printf("\tfirst 2: %q\n", word[:first+second])

		// Slice and print the last two runes of the strings
		_, last1 := utf8.DecodeLastRuneInString(word)
		_, last2 := utf8.DecodeLastRuneInString(word[:len(word)-last1])
		// fmt.Printf("first %d, second %d\n", first, second)
		fmt.Printf("\tlast 2: %q\n", word[len(word)-last1-last2:])

		// Convert the string to []rune
		rword := []rune(word) // "hello" -> []rune("hello") | []byte("hello")

		// Print the first and last two runes
		fmt.Printf("\tfirst 2 runes: %q\n", string(rword[0:2]))
		fmt.Printf("\tlast 2 runes: %q\n", string(rword[(len(rword)-2):]))

	}
}

func RuneManipulatorDA() {
	words := []string{
		"cool",
		"güzel",
		"jīntiān",
		"今天",
		"read 🤓",
	}

	for _, s := range words {
		fmt.Printf("%q\n", s)

		// Print the byte and rune length of the strings
		// Hint: Use len and utf8.RuneCountInString
		fmt.Printf("\thas %d bytes and %d runes\n",
			len(s), utf8.RuneCountInString(s))
		/*
			Bài tập này giúp hiểu cách in độ dài của strings theo byte / rune.
			- Theo byte thì dùng len()
			- Theo rune thì dùng utf8.RuneCountInString() <- tham số truyền vào là string
			- Trả về: %d
		*/

		// Print the bytes of the strings in hexadecimal
		// Hint: Use % x verb
		fmt.Printf("\tbytes   : % x\n", s)

		/*
			Bài tập này giúp biết cách in string theo hệ số 16 (hexadecimals)
			[% x] để giữa các bytes có khoảng cách cho dễ phân biệt.
		*/

		// Print the runes of the strings in hexadecimal
		// Hint: Use % x verb
		fmt.Print("\trunes   :")
		for _, r := range s {
			fmt.Printf("% x", r)
		}
		fmt.Println()

		/*
			Bài tập này giúp biết cách in string theo dạng rune.
			(nếu in thông thường thì theo dạng bytes)
			- Go hỗ trợ hàm range để loop từng kí tự của string và trả về dạng rune.
			- Để in ra dạng rune thì dùng %x
		*/

		// Print the runes of the strings as rune literals
		// Hint: Use for range
		fmt.Print("\trunes   :")
		for _, r := range s {
			fmt.Printf("%q", r)
		}
		fmt.Println()

		/*
			Run Literals đại khái hiểu là in ra đúng dạng mà mắt thường nhìn thấy string
			- dùng %q để in ra rune
		*/

		// Print the first rune and its byte size of the strings
		// Hint: Use utf8.DecodeRuneInString
		r, size := utf8.DecodeRuneInString(s)
		fmt.Printf("\tfirst   : %q (%d bytes)\n", r, size)

		/*
			Chú ý là: firs rune # first byte
			-> chúng ta không thể dựa vào len() để tính index
			mà phải decode string để trả về size (số bytes của ký tự đầu tiên của chuỗi)
			- dùng hàm utf8.DecodeRuneInString(s string)
			-> trả về first rune và size (số bytes của rune đó)
			- để in ra rune: %q
			- để in ra size: %d
			- \t để thụt lề sang phải
		*/

		// Print the last rune of the strings
		// Hint: Use utf8.DecodeLastRuneInString
		r, size = utf8.DecodeLastRuneInString(s)
		fmt.Printf("\tlast    : %q (%d bytes)\n", r, size)

		/*
			Tương tụ cách in first rune, last rune cũng phải dùng hàm để
			decode string. Go hỗ trợ sẵn hàm DecodeLastRuneInString(s string)
			-> hàm này trả về last rune và size của nó.
		*/

		// Slice and print the first two runes of the strings
		_, first := utf8.DecodeRuneInString(s)
		_, second := utf8.DecodeRuneInString(s[first:])
		fmt.Printf("\tfirst 2 : %q\n", s[:first+second])

		/*
			Chú ý: first two runes tức là in ra 2 runes đầu tiên.
			Go có hàm hỗ trợ in ra first run -> chỉ cần chúng ta chỉ cần
			cắt string để tạo 1 string mới và tiếp tục in phần tử đầu tiên
			của string mới thì sẽ có được size của run thứ 2 trong string gốc.
			- str -> first rune, size
			- str[:size] -> chính là slice string để tạo 1 string mới có fist rune
			chính second rune của slice gốc

		*/

		// Slice and print the last two runes of the strings
		_, last1 := utf8.DecodeLastRuneInString(s)
		_, last2 := utf8.DecodeLastRuneInString(s[:len(s)-last1])
		fmt.Printf("\tlast 2  : %q\n", s[len(s)-last2-last1:])

		/*
			Tương tự cách tư duy của first two rune, chúng ta sẽ dùng hàm DecodeLastRuneInString()
			để lấy size của last rune.
			Tiếp theo kết hợp len() để tính ra vị trí last rune thứ 2.
			Tiếp theo tạo 1 string mới str[:len(s)-size], lại dùng DecodeLastRuneInString()
			để lấy size của last rune thứ 2.
			Cuối cùng chúng ta kết hợp len(), last1, last2 để tạo 1 string mới str[len(s)-last1-last2:]
			chứa last two rune, và in nó ra thôi.
		*/

		// Convert the string to []rune
		// Print the first and last two runes
		rs := []rune(s)
		fmt.Printf("\tfirst 2 : %q\n", string(rs[:2]))
		fmt.Printf("\tlast 2  : %q\n", string(rs[len(rs)-2:]))

		/*
			Để convert string tu rune ta dùng công thức:
			Ví dụ: str := "hello" -> r := []rune(str) <- dấu ngoặc đơn, k phải kép
			- Khi chuyển thành dạng []rune thì các phần tử đều là dạng rune # string thì
			các phần tử là dạng byte. Do đó khi in ra là khác nhau.
			- Muốn in rune ra string thì cần phải chuyển rune về string bằng cách ép kiểu string
			Ví dụ: fmt.Println(r[:2]) -> [104, 101]
			Nhưng fmt.Println(string(r[0:2])) -> "he"
		*/
	}
}
