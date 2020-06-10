package notas

import (
	"fmt"
	"strings"
	"time"
)

func parsePageDate(pageText string) (date time.Time, err error) {
	lines = strings.Split(pageText, "\n")

	for _, line := range lines {
		texts := strings.Split(line, " ")
		if len(texts) == 3 {
			dateTxt := texts[2]
			if strings.Contains(dateTxt, "/") {
				return time.Parse("02/01/2006", dateTxt)
			}
		}
	}
	fmt.Println("(parsePageDate) Could not find date in expected form")

	return
}