package services

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/microcosm-cc/bluemonday"
	"github.com/pkg/errors"
)

// htmlから正規表現を用いて今日のcontribution数を抜き出す関数
func extractContributionCount(html string) int {
	// bluemondayの設定
	p := bluemonday.UGCPolicy()
	p.AllowElements("rect")
	p.AllowAttrs("data-count", "data-date").OnElements("rect")

	t := time.Now().Format("2006-01-02")
	sanitaizedHtml := p.Sanitize(html)

	pattern := fmt.Sprintf("(?i)<rect data-count=\"(\\d+)\" data-date=\"%s\"></rect>", t)
	re := regexp.MustCompile(pattern)
	strContributionCount := re.FindStringSubmatch(sanitaizedHtml)[1]
	contributionCount, _ := strconv.Atoi(strContributionCount)


	return contributionCount
}

// rectタグに近いDOMを手に入れる関数
func getTargetCloseHtml(doc *goquery.Document) (string, error) {
	html, err := doc.Find("div.js-calendar-graph svg.js-calendar-graph-svg g").Html()
	if err != nil {
		return "", errors.Wrapf(err, "dom get failed")
	}
	return html, nil
}

func GetTodaysPublicContributions(userName string) (int, error) {
	targetUrl := "https://github.com/" + userName

	res, err := http.Get(targetUrl)
	if err != nil {
		return 0, errors.Wrapf(err, "cant get URL")
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return 0, errors.Wrapf(err, "status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return 0, errors.Wrapf(err, "cant fetch document")
	}

	html, _ := getTargetCloseHtml(doc)
	contributionCount := extractContributionCount(html)

	return contributionCount, nil
}
