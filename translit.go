// Copyright (c) 2011 Dmitry Chestnykh
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package translit implements non-standard one-way string transliteration from Cyrillic to Latin.
package translit

import (
	"unicode"
)

// Mapping of Russian Cyrillic characters to ASCII Latin.
var RussianASCII = map[int]string{
	'а': "a",
	'б': "b",
	'в': "v",
	'г': "g",
	'д': "d",
	'е': "e",
	'ё': "yo",
	'ж': "zh",
	'з': "z",
	'и': "i",
	'й': "j",
	'к': "k",
	'л': "l",
	'м': "m",
	'н': "n",
	'о': "o",
	'п': "p",
	'р': "r",
	'с': "s",
	'т': "t",
	'у': "u",
	'ф': "f",
	'х': "h",
	'ц': "c",
	'ч': "ch",
	'ш': "sh",
	'щ': "sch",
	'ъ': "'",
	'ы': "y",
	'ь': "",
	'э': "e",
	'ю': "ju",
	'я': "ja",
}

// Mapping of Cyrillic characters to (Czech/Serbian) Latin.
var CyrillicLatin = map[int] string{
	'а': "a",
	'б': "b",
	'в': "v",
	'г': "g",
	'ґ': "g",
	'ѓ': "ǵ",
	'д': "d",
	'ђ': "đ",
	'е': "e",
	'ё': "yo",
	'ж': "zh",
	'з': "z",
	'ѕ': "ẑ",
	'и': "i",
	'і': "ì",
	'ї': "ï",
	'й': "j",
	'ј': "j",
	'к': "k",
	'л': "l",
	'љ': "lj",
	'м': "m",
	'н': "n",
	'њ': "nj",
	'о': "o",
	'п': "p",
	'р': "r",
	'с': "s",
	'т': "t",
	'ћ': "ć",
	'у': "u",
	'ф': "f",
	'х': "h",
	'ц': "c",
	'ч': "ch",
	'џ': "dž",
	'ш': "sh",
	'щ': "sch",
	'ъ': "'",
	'ы': "y",
	'ь': "",
	'ѣ': "ě",
	'э': "e",
	'є': "ê",
	'ю': "ju",
	'я': "ja",	
}

func Translit(s string, withAdditionalSymbols bool) string {
	var result string
	for _, r := range s {
		if unicode.IsUpper(r) {
			r = unicode.ToLower(r)
		}
		if cyrillic, exist := CyrillicLatin[int(r)]; exist {
			result += cyrillic
		} else {

			unicode.ToLower(r)
			if withAdditionalSymbols && (unicode.IsSpace(r) || unicode.IsMark(r) || unicode.IsPunct(r)) {
				result += string(r)
			}
		}
	}
	return result
}
