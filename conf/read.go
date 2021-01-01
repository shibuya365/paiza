package conf // 独自の設定ファイルパッケージ

import (
	"bufio"
	"os"
)

// DB設定読み込み関数
func ReadConfDB(web string) ([]string, error) {
	// ニュースを入れるスライス定義
	var news []string

	// 設定ファイルを読み込む
	f, err := os.Open("conf/" + web + ".txt")
	if err != nil {
		return news, err
	}
	// ファイルを最後に閉じる
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		// ここで一行ずつ処理
		news = append(news, scanner.Text())
	}
	if err = scanner.Err(); err != nil {
		return news, err
	}

	return news, nil
}
