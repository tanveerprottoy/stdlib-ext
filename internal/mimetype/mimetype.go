package mimetype

import "strings"

var MimeTypes = [2]string{
	"image/",
	"application/",
}

func IsMatchingMIMEType(mimeType string) bool {
	for _, m := range MimeTypes {
		if strings.HasPrefix(mimeType, m) {
			return true
		}
	}
	return false
}
