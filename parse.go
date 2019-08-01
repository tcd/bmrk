package main

import (
	"io/ioutil"
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

// ParseChromeBookmarks reads a chrome bookmarks html file and reutrns
// an array of `Bookmark` structs parsed from the file, or an error.
func ParseChromeBookmarks(path string) ([]Bookmark, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return []Bookmark{}, err

	}
	return parseChromeBookmarks(string(bytes)), nil
}

func parseChromeBookmarks(s string) []Bookmark {
	var bookmarks []Bookmark
	doc := strings.NewReader(s)

	z := html.NewTokenizer(doc)

	for {
		t1 := z.Next()

		switch {
		case t1 == html.ErrorToken:
			return bookmarks
		case t1 == html.StartTagToken:
			T1 := z.Token()

			if T1.Data == "a" {
				var bm Bookmark

				for _, a := range T1.Attr {
					if a.Key == "href" {
						bm.Link = a.Val
					}
					if a.Key == "add_date" {
						bm.AddDate = a.Val
					}
				}
				t2 := z.Next()
				if t2 == html.TextToken {
					T2 := z.Text()
					t3 := z.Next()
					if t3 == html.EndTagToken {
						bm.Description = string(T2)
						bookmarks = append(bookmarks, bm)
					}
				}
			}
		}
	}
}

// ParseFirefoxBookmarks reads a firefox bookmarks html file and reutrns
// an array of `Bookmark` structs parsed from the file, or an error.
func ParseFirefoxBookmarks(path string) ([]Bookmark, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return []Bookmark{}, err

	}
	return parseFirefoxBookmarks(string(bytes)), nil
}

func parseFirefoxBookmarks(s string) []Bookmark {
	var bookmarks []Bookmark

	re := regexp.MustCompile(`\n\s*`)

	doc := strings.NewReader(s)

	z := html.NewTokenizer(doc)

	for {
		t1 := z.Next()

		switch {
		case t1 == html.ErrorToken:
			return bookmarks
		case t1 == html.StartTagToken:
			T1 := z.Token()

			if T1.Data == "a" {
				var bm Bookmark

				for _, a := range T1.Attr {
					if a.Key == "href" {
						bm.Link = a.Val
					}
					if a.Key == "add_date" {
						bm.AddDate = a.Val
					}
					if a.Key == "tags" {
						bm.Tags = a.Val
					}
				}

				t2 := z.Next()
				if t2 == html.TextToken {
					T2 := z.Text()
					t3 := z.Next()
					if t3 == html.EndTagToken {
						bm.Description = re.ReplaceAllString(string(T2), " ")
						bookmarks = append(bookmarks, bm)
					}
				}
			}
		}
	}
}

// ParseSafariBookmarks reads a chrome bookmarks html file and reutrns
// an array of `Bookmark` structs parsed from the file, or an error.
func ParseSafariBookmarks(path string) ([]Bookmark, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return []Bookmark{}, err

	}
	return parseChromeBookmarks(string(bytes)), nil
}

func parseSafariBookmarks(s string) []Bookmark {
	var bookmarks []Bookmark
	doc := strings.NewReader(s)

	z := html.NewTokenizer(doc)

	for {
		t1 := z.Next()

		switch {
		case t1 == html.ErrorToken:
			return bookmarks
		case t1 == html.StartTagToken:
			T1 := z.Token()

			if T1.Data == "a" {
				var bm Bookmark

				for _, a := range T1.Attr {
					if a.Key == "href" {
						bm.Link = a.Val
					}
				}

				t2 := z.Next()
				if t2 == html.TextToken {
					T2 := z.Text()
					t3 := z.Next()
					if t3 == html.EndTagToken {
						bm.Description = string(T2)
						bookmarks = append(bookmarks, bm)
					}
				}
			}
		}
	}
}
