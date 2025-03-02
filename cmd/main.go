package main

import (
	"fmt"
	"os"

	"golang.org/x/text/language"
)

func main() {
	// コマンドライン引数で渡された各タグ文字列を解析
	for _, tagStr := range os.Args[1:] {
		tag, err := language.Parse(tagStr)
		if err != nil {
			fmt.Printf("%s: error: %v\n", tagStr, err)
		} else if tag == language.Und { // 未定義タグの場合
			fmt.Printf("%s: undefined\n", tagStr)
		} else {
			fmt.Printf("%s: tag %s\n", tagStr, tag)
		}
	}
}
