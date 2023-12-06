package ui

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gdamore/tcell/v2"
)

func MakeTcellColor(hexString string) (tcell.Color, error) {

	// replace 0x or 0X with empty String
	hexString = strings.Replace(hexString, "0x", "", -1)
	hexString = strings.Replace(hexString, "0X", "", -1)
	hexString = strings.Replace(hexString, "#", "", -1)

	hexInt, err := strconv.ParseInt(hexString, 16, 64)
	if err != nil {
		fmt.Println(err)
		return tcell.ColorDefault, err
	}

	return tcell.NewHexColor(int32(hexInt)), nil
}
