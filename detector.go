package DocsED

import (
	"bytes"
	"fmt"
	"os"
)

/*
IsDocsEncrypted
return true if file is encrypted , or false
*/
func IsDocsEncrypted(path string) (bool, error) {
	fp, err := os.Open(path)
	if err != nil {
		return false, err
	}

	defer fp.Close()

	stat, err := fp.Stat()
	if err != nil {
		return false, err
	}

	size := stat.Size()

	if size <= 500 {
		return false, fmt.Errorf("the file size is too short , it may be not a docs file")
	}

	head := make([]byte, 8)
	_, err = fp.Read(head)
	if err != nil {
		return false, err
	}

	if !bytes.EqualFold(docsHead, head) {
		return false, nil
	}

	// docx xlsx pptx

	// doc xls ppt

	return false, nil
}
