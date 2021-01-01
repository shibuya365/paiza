package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type problem struct {
	title      string
	difficulty float64
}

type problems []problem

// sortインタフェース実装
func (p problems) Len() int           { return len(p) }
func (p problems) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p problems) Less(i, j int) bool { return p[i].difficulty > p[j].difficulty }

func main() {
	// 書き込みファイル作成
	f, err := os.Create("index.html")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	// 問題の配列作成
	var pros problems

	// ホームページ読み込み
	res, err := http.Get("https://paiza.jp/challenges/ranks/c/info")
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s\n", res.StatusCode, res.Status)
	}

	// タイトルの取得
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Println(err)
	}

	section := doc.Find(".problem-box__header__title")

	// タイトルを取得
	section.Each(func(i int, line *goquery.Selection) {
		txt := line.Text()
		txt = strings.Replace(txt, "\n", "", -1)
		// コンソールへ新しいニュースのタイトルのみ出力
		// fmt.Println(title)
		var problem problem
		problem.title = txt
		pros = append(pros, problem)
	})

	// 難易度の取得
	section = doc.Find(".problem-box__data dd")

	count := 0
	section.Each(func(i int, line *goquery.Selection) {
		// タイトルを取得
		def := line.Text()
		// コンソールへ新しいニュースのタイトルのみ出力
		if strings.Contains(def, "点") {
			def = strings.Replace(def, "点", "", 1)
			def = strings.Replace(def, "\n", "", -1)

			num, _ := strconv.ParseFloat(def, 64)
			pros[count].difficulty = num
			// defs = append(defs, num)
			count++
		}
	})

	// 昇順にソート
	sort.Sort(pros)

	// HTML表示なら
	// f.WriteString("<table>")

	for i := 0; i < len(pros); i++ {
		// fmt.Println(pros[i].title, pros[i].difficulty)
		line := fmt.Sprintf("□paiza Cランク問題（初級）『%s』平均%.2f点\n", pros[i].title, pros[i].difficulty)
		// HTML表示なら
		// line := fmt.Sprintf("<tr align=\"left\"><th>%s</th><th>平均%.2f点</th></tr>\n", pros[i].title, pros[i].difficulty)
		f.WriteString(line)
	}
	// HTML表示なら
	// f.WriteString("</table>")
}
