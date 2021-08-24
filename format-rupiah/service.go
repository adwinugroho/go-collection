package formatrupiah

import (
	"strings"

	"github.com/dustin/go-humanize"
)

func FormatRupiah(amount float64) string {
	var humanizeValue = humanize.CommafWithDigits(amount, 0)
	var stringValue = strings.ReplaceAll(humanizeValue, ",", ".")
	return stringValue
}
