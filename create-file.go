package createfile

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func Createfile(file string, timest int64, paymentID string) {
	// if _, err := os.Stat(file); err == nil {
	// 	os.Remove(file)
	// }
	out := fmt.Sprintf("{\"%s\", \"%d\"}\n", paymentID, timest)
	fi, err := os.OpenFile(file, os.O_APPEND|os.O_WRONLY, 0644)
	check(err)
	defer fi.Close()
	if _, err = fi.WriteString(out); err != nil {
		panic(err)
	}
	fi.Sync()
	return
}
