package fetcher

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type response struct {
	Data struct {
		Children []struct {
			Data struct {
				Title string `json:"title"`
				URL   string `json:"url"`
			} `json:"data"`
		} `json:"children"`
	} `json:"data"`
}

type RedditFetcher interface {
	Fetch() error
	Save(io.Writer) error
}

type Reddit struct {
	URL string
	r   response
}

func (r *Reddit) Fetch() error {
	resp, err := http.Get(r.URL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&r.r)
	if err != nil {
		return err
	}

	return nil
}

func (r *Reddit) Save(w io.Writer) error {
	for _, child := range r.r.Data.Children {
		_, err := fmt.Fprintf(w, "%s\n%s\n\n", child.Data.Title, child.Data.URL)
		if err != nil {
			return err
		}
	}
	return nil
}
