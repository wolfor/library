// Text project Text.go
package Text

import (
	"github.com/axgle/mahonia"
)

func CharsetConvertString(text string, srcCharsetCode string, targetCharsetCode string) string {
	srcCoder := mahonia.NewDecoder(srcCharsetCode)
	srcResult := srcCoder.ConvertString(text)
	tagCoder := mahonia.NewDecoder(targetCharsetCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}
