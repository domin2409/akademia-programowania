package main

import (
	"io"
	"os"
	"reddit/fetcher"
)

func main() {
	var f fetcher.RedditFetcher // do not change
	var w io.Writer             // do not change

	// update the URL in the Fetch method to the new link
	f = &fetcher.Reddit{URL: "https://www.reddit.com/r/golang.json"}

	//if err != nil {
	//	panic(err)
	//}
	//defer f.

	w, _ = os.Create("output.txt")
	f.Fetch()
	f.Save(w)
}
