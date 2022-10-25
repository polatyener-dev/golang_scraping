package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"

	"fyne.io/fyne/v2/widget"
	"github.com/PuerkitoBio/goquery"
)

type Scrapings struct {
	Title string `json:"title"`
}

func main() {

	createJson()

	var jsonData []Scrapings

	jsondatafile, _ := ioutil.ReadFile("output.json")
	json.Unmarshal(jsondatafile, &jsonData)

	ma := app.New()
	mw := ma.NewWindow("List Data")
	mw.Resize(fyne.NewSize(400, 400))

	lst := widget.NewList(
		func() int {
			return len(jsonData)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(lii widget.ListItemID, co fyne.CanvasObject) {
			co.(*widget.Label).SetText(jsonData[lii].Title)
		},
	)

	mw.SetContent(container.NewVScroll(lst))
	mw.ShowAndRun()
}

func createJson() {

	res, err := http.Get("http://yenerpolat.com/blog/")
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("Hata Kodu: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var data []Scrapings
	doc.Find(".article-content .entry-title").Each(
		func(i int, s *goquery.Selection) {
			title := Scrapings{Title: strings.TrimSpace(s.Find("a").Text())}
			data = append(data, title)
		},
	)

	file, _ := json.MarshalIndent(data, "", " ")
	_ = ioutil.WriteFile("output.json", file, 0644)
}
