package articles

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
)

func getArticle(page int) (*Articles, error) {
	resp, err := http.Get(fmt.Sprintf("https://jsonmock.hackerrank.com/api/articles?page=%d", page))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	articles := Articles{}
	err = json.Unmarshal(body, &articles)
	if err != nil {
		return nil, err
	}

	return &articles, nil
}

func TopArticles(limit int32) []string {
	if limit == 0 {
		return []string{}
	}

	article1, _ := getArticle(1)
	total_pages := article1.TotalPages
	artTitles := make([]ArticleTitle, 0)
	total := article1.Total

	if len(article1.Data) == 0 {
		return []string{}
	}
	for _, title := range article1.Data {
		if title.Title == "" {
			title.Title = title.StoryTitle
		}
		artTitles = append(artTitles, title)
	}
	total = total - len(artTitles)

	// if page more than 1
	if total_pages > 1 {
		titleCh := make(chan ArticleTitle, total)

		// get title info
		for i := 2; i <= total_pages; i++ {
			go func(page int) {
				article2, _ := getArticle(page)
				for _, title := range article2.Data {
					if title.Title == "" {
						title.Title = title.StoryTitle
					}
					titleCh <- title
				}
			}(i)
		}

		for i := 0; i < total; i++ {
			artTitles = append(artTitles, <-titleCh)
		}

		close(titleCh)
	}

	// Sort the slice
	sort.Slice(artTitles, func(i, j int) bool {
		if artTitles[i].Title == "" {
			return false
		}
		if artTitles[i].NumComments == artTitles[j].NumComments {
			return artTitles[i].Title > artTitles[j].Title
		}
		return artTitles[i].NumComments > artTitles[j].NumComments
	})

	// output
	output := make([]string, 0)
	for i := 0; i < int(limit); i++ {
		output = append(output, artTitles[i].Title)
	}

	return output
}

type ArticleTitle struct {
	Title       string `json:"title"`
	StoryTitle  string `json:"story_title"`
	NumComments int    `json:"num_comments"`
}

type Articles struct {
	Page       int            `json:"page"`
	Total      int            `json:"total"`
	TotalPages int            `json:"total_pages"`
	Data       []ArticleTitle `json:"data"`
}
