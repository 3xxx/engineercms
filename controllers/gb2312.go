package controllers //gogb2312

// just for convert gb2312 to utf-8

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

const (
	UTF8_B1 = 0x80
	UTF8_B2 = 0x80
	UTF8_B3 = 0xe0
)

func readable(b byte) byte {
	if b <= 9 {
		return b + byte('0')
	}
	return b - 10 + byte('a')
}

func tohex(buf []byte) string {
	if len(buf) == 0 {
		return "[]"
	}

	var (
		i   int
		b   byte
		hex []byte = make([]byte, len(buf)*3+2)
	)

	hex[0] = byte('[')
	i = 1
	for _, b = range buf {
		hex[i] = readable((b & 0xf0) >> 4)
		i++
		hex[i] += readable(b & 0x0f)
		i++
		hex[i] = byte(' ')
		i++
	}
	// overwrite the space
	hex[i-1] = ']'
	hex[i] = ' '

	return string(hex)
}

func nearbygbk(buf []byte, tlen int, plen int) string {
	if tlen <= plen {
		return tohex(buf[0:tlen])
	}
	return tohex(buf[tlen-plen : tlen])
}

func nearbygbks(buf string, tlen int, plen int) string {
	if tlen <= plen {
		return tohex([]byte(buf[0:tlen]))
	}
	return tohex([]byte(buf[tlen-plen : tlen]))
}

// get a slice of utf-8 string, from tlen-plen, to tlen
func nearbyutf8(buf []byte, tlen int, plen int) string {
	if tlen <= plen {
		return tohex(buf) + string(buf)
	}

	start := tlen - plen
	if buf[start] < byte(0x7f) || buf[start] > byte(0xe0) {
		return tohex(buf[start:]) + string(buf[start:])
	}
	start = start - 1
	if buf[start] < byte(0x7f) || buf[start] > byte(0xe0) {
		return tohex(buf[start:]) + string(buf[start:])
	}
	if start < 1 {
		return ""
	}
	start = start - 1
	if buf[start] < byte(0x7f) || buf[start] > byte(0xe0) {
		return tohex(buf[start:]) + string(buf[start:])
	}
	return ""
}

// param: input: input bytes array
// return: output: output bytes array
//         err: error if there are errors when convert
//         ic: input has been converted
//         oc: output has been converted
func ConvertGB2312(input []byte) (output []byte, err error, ic int, oc int) {
	ilen := len(input)
	output = make([]byte, (ilen/2)*3+3)
	olen := 0
	i := 0
	for i < ilen-1 {
		if input[i] <= 0x7f {
			output[olen] = input[i]
			olen++
			i++
		} else {
			gb := int(input[i])<<8 | int(input[i+1])
			u8, ok := gb2312toutf8[gb]
			if !ok {
				err = fmt.Errorf("gb2312 has no character %x, at %d.\nnearby input: %s\nnearby output: %s",
					gb, ilen, nearbygbk(input[0:i], i, 20),
					nearbyutf8(output[0:olen], olen, 30))
				ic = i
				oc = olen
				return
			}
			if u8 >= 0x10000 {
				output[olen] = byte(u8 >> 16)
				olen++
				output[olen] = byte((u8 >> 8) & 0xff)
				olen++
				output[olen] = byte(u8 & 0xff)
				olen++
			} else {
				output[olen] = byte(u8 >> 8)
				olen++
				output[olen] = byte(u8 & 0xff)
				olen++
			}
			i = i + 2
		}
	}

	// the last character
	if i == ilen-1 {
		if byte(input[ilen-1]) <= 0x7f {
			output[olen] = input[ilen-1]
			olen++
			i++
		}
	}

	ilen = i
	output = output[0:olen]
	return output, nil, ilen, olen
}

func ConvertGB2312String(input string) (soutput string, err error, ic int, oc int) {
	ilen := len(input)
	output := make([]byte, (ilen/2)*3+3)
	olen := 0
	i := 0
	for i < ilen-1 {
		bi := byte(input[i])
		if bi <= 0x7f {
			output[olen] = bi
			olen++
			i++
		} else {
			bii := byte(input[i+1])
			gb := int(bi)<<8 | int(bii)
			u8, ok := gb2312toutf8[gb]
			if !ok {
				err = fmt.Errorf("gb2312 has no character %x, at %d\nnearby input: %s\nnearby output: %s",
					gb, ilen, nearbygbks(input[0:i], i, 20),
					nearbyutf8(output[0:olen], olen, 30))
				ic = i
				oc = olen
				return
			}
			if u8 >= 0x10000 {
				output[olen] = byte(u8 >> 16)
				olen++
				output[olen] = byte((u8 >> 8) & 0xff)
				olen++
				output[olen] = byte(u8 & 0xff)
				olen++
			} else {
				output[olen] = byte(u8 >> 8)
				olen++
				output[olen] = byte(u8 & 0xff)
				olen++
			}
			i = i + 2
		}
	}
	if i == ilen-1 {
		if byte(input[ilen-1]) <= 0x7f {
			output[olen] = input[ilen-1]
			olen++
			i++
		}
	}

	output = output[0:olen]
	soutput = string(output)
	return soutput, nil, ilen, olen
}

/*
Unicode和UTF-8之间的转换关系表
UCS-4编码	UTF-8字节流
U+00000000 – U+0000007F	0xxxxxxx
U+00000080 – U+000007FF	110xxxxx 10xxxxxx
U+00000800 – U+0000FFFF	1110xxxx 10xxxxxx 10xxxxxx
U+00010000 – U+001FFFFF	11110xxx 10xxxxxx 10xxxxxx 10xxxxxx
U+00200000 – U+03FFFFFF	111110xx 10xxxxxx 10xxxxxx 10xxxxxx 10xxxxxx
U+04000000 – U+7FFFFFFF	1111110x 10xxxxxx 10xxxxxx 10xxxxxx 10xxxxxx 10xxxxxx
*/
var (
	utf8FirstByte = []byte{0, 0, byte(0xc0), byte(0xe0), byte(0xf0), byte(0xf8), byte(0xfc)}
	utf8FirstMask = []byte{0, 0, byte(0xe0), byte(0xf0), byte(0xf8), byte(0xfc), byte(0xfe)}
)

// check if the first N byte of buf is utf-8
func isutf8(buf []byte, blen int) int {
	if blen <= 2 {
		return 0
	}
	// invalid utf-8 code
	if buf[0] == byte(0xc0) || buf[0] == byte(0xc1) ||
		buf[0] == byte(0xfe) || buf[0] == byte(0xff) {
		return 0
	}

	var (
		i      int = 3
		maxlen int = 6
	)
	if blen < 6 {
		maxlen = blen
	}
	for ; i < maxlen; i++ {
		fb := utf8FirstByte[i]
		mk := utf8FirstMask[i]
		if buf[0]&mk == fb {
			res := true
			for j := 1; j <= i-1; j++ {
				b := buf[j]
				if b == byte(0xc0) || b == byte(0xc1) ||
					b == byte(0xfe) || b == byte(0xff) {
					return 0
				}
				if b&byte(0xc0) != 0x80 {
					res = false
					break
				}
			}
			if res {
				return i
			}
		}
	}
	return 0
}
func isutf8s(buf string, blen int) int {
	if blen <= 2 {
		return 0
	}
	// invalid utf-8 code
	if buf[0] == byte(0xc0) || buf[0] == byte(0xc1) ||
		buf[0] == byte(0xfe) || buf[0] == byte(0xff) {
		return 0
	}

	var (
		i      int = 3
		maxlen int = 6
	)
	if blen < 6 {
		maxlen = blen
	}
	for ; i < maxlen; i++ {
		fb := utf8FirstByte[i]
		mk := utf8FirstMask[i]
		if buf[0]&mk == fb {
			res := true
			for j := 1; j <= i-1; j++ {
				b := buf[j]
				if b == byte(0xc0) || b == byte(0xc1) ||
					b == byte(0xfe) || b == byte(0xff) {
					return 0
				}
				if b&byte(0xc0) != 0x80 {
					res = false
					break
				}
			}
			if res {
				return i
			}
		}
	}
	return 0
}

func ConvertHybirdString(input string) (soutput string, err error, ic int, oc int) {
	ilen := len(input)
	output := make([]byte, (ilen/2)*3+3)
	olen := 0
	i := 0
	for i < ilen-1 {
		bi := byte(input[i])
		if bi <= 0x7f {
			output[olen] = bi
			olen++
			i++
		} else {
			utf8 := isutf8s(input[i:], ilen-i)
			if utf8 == 0 {
				bii := byte(input[i+1])
				gb := int(bi)<<8 | int(bii)
				u8, ok := gb2312toutf8[gb]
				if !ok {
					err = fmt.Errorf("ConvertHybirdString: gb2312 has no character %x, at %d\nnearby input: %s\nnearby output: %s",
						gb, ilen, nearbygbks(input[0:i], i, 20),
						nearbyutf8(output[0:olen], olen, 30))
					ic = i
					oc = olen
					return
				}
				output[olen] = byte(u8 >> 16)
				olen++
				output[olen] = byte((u8 >> 8) & 0xff)
				olen++
				output[olen] = byte(u8 & 0xff)
				olen++
				i = i + 2
			} else {
				for ul := 0; ul < utf8; ul++ {
					output[olen] = input[i]
					olen++
					i++
				}
			}
		}
	}
	if i == ilen-1 {
		if byte(input[ilen-1]) <= 0x7f {
			output[olen] = input[ilen-1]
			olen++
			i++
		}
	}

	output = output[0:olen]
	soutput = string(output)
	return soutput, nil, ilen, olen
}

const (
	UTF8_B7FF     = 0xc0
	UTF8_1FFFFF   = 0xf0
	UTF8_3FFFFFF  = 0xf8
	UTF8_7FFFFFFF = 0xfc
)

/*
Unicode和UTF-8之间的转换关系表
UCS-4编码	UTF-8字节流
U+00000000 – U+0000007F	0xxxxxxx
U+00000080 – U+000007FF	110xxxxx 10xxxxxx
U+00000800 – U+0000FFFF	1110xxxx 10xxxxxx 10xxxxxx
U+00010000 – U+001FFFFF	11110xxx 10xxxxxx 10xxxxxx 10xxxxxx
U+00200000 – U+03FFFFFF	111110xx 10xxxxxx 10xxxxxx 10xxxxxx 10xxxxxx
U+04000000 – U+7FFFFFFF	1111110x 10xxxxxx 10xxxxxx 10xxxxxx 10xxxxxx 10xxxxxx
*/
func unicode2utf8(u int) int {
	if u < 0x7f {
		return u
	}
	var u8 int
	if u >= 0x80 && u <= 0x7ff {
		b1 := (u & 0x3f) | UTF8_B1
		b2 := ((u >> 6) & 0x1f) | UTF8_B7FF
		u8 = b1 | (b2 << 8)
	} else if u >= 0x800 && u <= 0xffff {
		b1 := (u & 0x3f) | UTF8_B1
		b2 := ((u >> 6) & 0x3f) | UTF8_B2
		b3 := (u >> 12) | UTF8_B3
		u8 = b1 | (b2 << 8) | (b3 << 16)
	} else if u >= 0x10000 && u <= 0x1FFFFF {
		b1 := (u & 0x3f) | UTF8_B1
		b2 := ((u >> 6) & 0x3f) | UTF8_B1
		b3 := ((u >> 12) & 0x3f) | UTF8_B1
		b4 := ((u >> 18) & 0x7) | UTF8_1FFFFF
		u8 = b1 | (b2 << 8) | (b3 << 16) | (b4 << 24)
	} else if u >= 0x200000 && u <= 0x3FFFFFF {
		return 0
	} else if u >= 04000000 && u <= 0x7FFFFFFF {
		return 0
	}

	return u8
}

func readcp936(fp string) {
	f, err := os.Open(fp)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	sf, err := os.Create("./gbk2utf8.tmp")
	if err != nil {
		panic(err.Error())
	}
	sf2, err2 := os.Create("./gbk2unicode.tmp")
	if err2 != nil {
		panic(err2.Error())
	}

	sf.WriteString("package gogb2312\n\nvar gb2312toutf8=map[int]int{")
	sf2.WriteString("package gogb2312\n\nvar gb2312tounicode=map[int]int{")

	rd := bufio.NewReader(f)
	i := 0
	line, _, err := rd.ReadLine()
	for ; err == nil; line, _, err = rd.ReadLine() {
		gb, unicode, u8, e := parseline(line, i)
		if e != nil {
			fmt.Print(e)
			i++
			continue
		}
		i++
		serr := savecode(sf, gb, u8)
		if serr != nil {
			fmt.Println("save code failed", serr)
		}
		serr = savecode(sf2, gb, unicode)
		if serr != nil {
			fmt.Println("save code failed", serr)
		}
	}
	if err != nil && err != io.EOF {
		fmt.Println(err)
	}
	fmt.Printf("Lines of %s: %d\n", fp, i)
	sf.WriteString("}\n")
	sf.Close()
	sf2.WriteString("}\n")
	sf2.Close()
}

func savecode(rd *os.File, gb, u8 int) error {
	s := fmt.Sprintf("0x%x:0x%x,\n", gb, u8)
	_, err := rd.WriteString(s)
	return err
}

var re_space = regexp.MustCompile(`\s+`)

func parseline(line []byte, i int) (int, int, int, error) {
	comma := bytes.Index(line, []byte{'#'})
	if comma >= 0 {
		line = line[0:comma]
	}
	if l := len(line); l < 4 {
		return -1, -1, -1, fmt.Errorf("line %d length invalid: %d %q\n", i, l, line)
	}
	rl := bytes.TrimSpace(line)
	rl = re_space.ReplaceAll(rl, []byte{' '})
	chs := bytes.Split(rl, []byte{' '})
	if len(chs) != 2 {
		return -1, -1, -1,
			fmt.Errorf("line %d has %d numbers: %s\n", i, len(chs), rl)
	}
	ret, err := strconv.ParseInt(string(bytes.ToLower(chs[0])), 0, 32)
	if err != nil {
		return -1, -1, -1,
			fmt.Errorf("convert %q to int failed at line %d: %s\n", chs[0], i, err)
	}
	gb2312 := int(ret)
	if gb2312 <= 0x7f {
		return gb2312, gb2312, gb2312, fmt.Errorf("No need convert for ascii 0x%x\n", gb2312)
	}
	ret, err = strconv.ParseInt(string(bytes.ToLower(chs[1])), 0, 32)
	if err != nil {
		return -1, -1, -1,
			fmt.Errorf("convert %q to int failed at line %d: %s\n", chs[1], i, err)
	}
	unicode := int(ret)
	utf8 := unicode2utf8(unicode)

	return gb2312, unicode, utf8, nil
}
