package ascii

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func Push(s string) map[int]string {
	alph := make(map[int]string)
	i := 1
	j := 0
	count := 0
	simbol := ""
	for i < len(s) {
		if s[i] == '\n' {
			j++
		}
		if j == 8 {
			alph[' '+count] = simbol
			count++
			simbol = ""
			j = 0
			i += 2
		}
		simbol += string(s[i])
		i++
	}
	alph[' '+count] = simbol
	return alph
}

func Readfile(s string) (string, error) {
	content, err := ioutil.ReadFile("pkg/ascii/files/" + s)
	if err != nil {
		return "", err
	}

	if !checkBanner(content) {
		return "", errors.New("not correct file")
	}

	return string(content), nil
}

func checkBanner(content []byte) bool {
	h := sha256.New()
	h.Write(content)
	Standardhash := []byte{195, 236, 117, 132, 251, 126, 207, 189, 115, 158, 107, 63, 111, 99, 253, 31, 229, 87, 210, 174, 62, 36, 248, 112, 115, 13, 156, 248, 178, 85, 158, 148}
	Shadowhash := []byte{120, 204, 214, 22, 104, 14, 185, 6, 143, 225, 70, 93, 177, 200, 82, 206, 175, 253, 140, 15, 49, 142, 58, 160, 65, 78, 22, 53, 80, 142, 133, 191}
	Thinkertoy := []byte{71, 44, 250, 176, 86, 116, 6, 209, 3, 19, 171, 34, 32, 170, 94, 100, 183, 251, 241, 121, 109, 150, 193, 28, 119, 82, 70, 115, 45, 206, 126, 122}
	if string(h.Sum(nil)) == string(Standardhash) || string(h.Sum(nil)) == string(Thinkertoy) || string(h.Sum(nil)) == string(Shadowhash) {
		return true
	}
	return false
}

func print(q string, count int) string {
	a := strings.Split(q, "\n")
	s := ""
	for i := count; i < count+1; i++ {
		s += a[i]
	}
	return s
}

func printAll(s string, alph map[int]string) string {
	s1 := ""
	for j := 0; j < 8; j++ {
		for i := 0; i < len(s); i++ {
			if s[i] == 13 {
				break
			}
			if s[i] < 32 || s[i] > 126 {
				fmt.Println(s[i])
				return "err1"
			}
			s1 += print(alph[int(s[i])], j)
		}
		s1 += "\n"
	}
	return s1
}

func checker(s []string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] != "" {
			return false
		}
	}
	return true
}

func Final(s string, alph map[int]string) string {
	final := ""
	if s == "" {
		return ""
	}
	if s == "\\n" {
		return "\n"
	}
	s1 := strings.Split(s, "\r\n")
	if checker(s1) {
		s1 = s1[:len(s1)-1]
	}
	for i := 0; i < len(s1); i++ {
		if s1[i] != "" {
			if printAll(s1[i], alph) == "err1" {
				return "err1"
			} else {
				final += printAll(s1[i], alph)
			}
		} else {
			final += "\n"
		}
	}
	return final
}

func WriteFile(file string, result []byte) {
	if len(file) < 9 || file[:9] != "--output=" {
		fmt.Println("Not correct input")
		return
	}
	filename := file[9:len(file)-4] + ".txt"
	ioutil.WriteFile(filename, []byte(result), 0664)
}

func Ascii(s string, filename string) (string, int) {
	content, err := Readfile(filename)
	if err != nil {
		log.Println(err.Error())
		return http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError
	}
	alph := Push(string(content))
	result := Final(s, alph)
	if result == "err1" {
		return http.StatusText(http.StatusBadRequest), http.StatusBadRequest
	}
	return result, http.StatusOK
}
