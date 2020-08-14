// create by d1y<chenhonzhou@gmail.com>

package plist2json

import (
	"errors"
	"os/exec"
	"strings"
)

// the Command only supports osx😂
// but, i don't know if other platforms have the same type of library
const easyCmd = "plutil"

// Easy only supports osx
// format `plist` file to `json`
// https://www.real-world-systems.com/docs/plutil.1.html
func Easy(source string, output string) error {
	if !EasyCheck(source) {
		return errors.New("plist File format error")
	}
	// plutil -convert json data/SoundsSettings.plist -o data.json
	easy := exec.Command(easyCmd, "-convert", "json", source, "-o", output)
	_, err := easy.CombinedOutput()
	if err != nil {
		return err
	}
	return nil
}

// EasyCheck only supports osx
// check `plist` file
// https://scriptingosx.com/2016/11/editing-property-lists
func EasyCheck(path string) bool {
	cache := exec.Command(easyCmd, "-lint", path)
	out, err := cache.CombinedOutput()
	if err != nil {
		return false
	}
	var result = string(out)
	var arr = strings.Split(result, ":")
	if len(arr) > 2 {
		return false
	}
	var t, checkText = arr[1], "OK"
	var ft = strings.TrimSpace(t)
	return ft == checkText
}
