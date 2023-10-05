package lucky

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const htlmQuery string = ".balls .resultBall"
const resultResource string = "results-history-"
const firstResultsYear int = 2004

type HttpCli interface {
	Do(req *http.Request) (*http.Response, error)
}

func NewDefaultClient() *http.Client {
	return &http.Client{
		Timeout: time.Minute,
	}
}

func (o *Options) getFullUrl(year int) string {
	return fmt.Sprintf("%s/%s%d", strings.TrimRight(o.BaseUrl, "/"), resultResource, year)
}

func (o *Options) getHmtl(ctx context.Context, year int) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, o.getFullUrl(year), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", o.UserAgent)
	r, err := o.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if r.StatusCode > http.StatusPermanentRedirect {
		defer r.Body.Close()
		buf := new(bytes.Buffer)
		buf.ReadFrom(r.Body)
		respBytes := buf.String()

		respString := string(respBytes)
		return nil, fmt.Errorf("status code error: %d %s %s", r.StatusCode, r.Status, respString)
	}

	return r, err
}

func (o Options) getNumbersFromHtml(resp *http.Response) ([][]int, error) {
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	balls := doc.Find(htlmQuery)
	all := make([][]int, balls.Length()/7)
	control := 0
	var e error
	balls.Each(func(i int, s *goquery.Selection) {
		index := i + 1
		v, err := strconv.Atoi(s.Text())
		if err != nil {
			e = err
			return
		}

		all[control] = append(all[control], v)
		if index%7 == 0 {
			control++
		}
	})

	if e != nil {
		return nil, e
	}

	return all, nil
}

func (o Options) convertToDraws(d [][]int) []Draw {
	draws := []Draw{}
	for _, nr := range d {
		draws = append(draws, Draw{
			Numbers: nr[:5],
			Starts:  nr[5:7],
		})
	}

	return draws
}

func (o Options) ScrapeFrom(ctx context.Context) (Draws, error) {
	allDraws := Draws{}
	fromYear := o.FromYear
	if o.FromYear < firstResultsYear {
		fromYear = firstResultsYear
	}
	currentYear := time.Now().Year()
	if o.FromYear > currentYear {
		fromYear = currentYear
	}

	years := makeRange(fromYear, currentYear)

	for _, y := range years {
		r, err := o.getHmtl(ctx, y)
		if err != nil {
			return nil, err
		}

		nrs, err := o.getNumbersFromHtml(r)
		if err != nil {
			return nil, err
		}

		allDraws = append(allDraws, o.convertToDraws(nrs)...)
	}

	err := o.saveToFile(allDraws)
	if err != nil {
		return nil, err
	}

	return allDraws, nil
}

func (o Options) saveToFile(d Draws) error {
	if !o.ShouldStore {
		return nil
	}

	file, err := json.Marshal(d)
	if err != nil {
		return err
	}

	err = os.WriteFile(o.StoreFile, file, 0644)
	return err
}

func makeRange(min, max int) []int {
	res := make([]int, max-min+1)
	for i := range res {
		res[i] = min + i
	}
	return res
}
