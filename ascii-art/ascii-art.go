package asciiart

import (
	"fmt"
	"strings"
)

func GetAscii(input, style string) (string, error) {
	bannerFile, err := GetBannerFile(style)
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}
	lines := make([]string, 0)
	words := strings.Split(input, "\n")

	for _, word := range words {
		if word == "" {
			lines = append(lines, "")
			continue
		}
		w, err := GetWord(word, bannerFile)
		if err != nil {
			return "", err
		}
		lines = append(lines, w...)
	}

	str := ""
	for i, l := range lines {
		if i != len(lines)-1 {
			str += l + "\n"
			continue
		}
		str += l
	}

	return str, nil
}
