package main

import "fmt"

func main() {
	// ====================================================================
	// Chrome
	// ====================================================================
	chromeBookmarks, err := ParseChromeBookmarks("./testdata/chrome.html")
	if err != nil {
		fmt.Println(err)
	}
	WriteBookmarksToJSONFile(chromeBookmarks, "./testdata/out/chrome.json")
	WriteBookmarksToTSVFile(chromeBookmarks, "./testdata/out/chrome.tsv")

	// ====================================================================
	// Firefox
	// ====================================================================
	firefoxBookmarks, err := ParseFirefoxBookmarks("./testdata/firefox.html")
	if err != nil {
		fmt.Println(err)
	}
	WriteBookmarksToJSONFile(firefoxBookmarks, "./testdata/out/firefox.json")
	WriteBookmarksToTSVFile(firefoxBookmarks, "./testdata/out/firefox.tsv")

	// ====================================================================
	// Safari
	// ====================================================================
	safariBookmarks, err := ParseSafariBookmarks("./testdata/safari.html")
	if err != nil {
		fmt.Println(err)
	}
	WriteBookmarksToJSONFile(safariBookmarks, "./testdata/out/safari.json")
	WriteBookmarksToTSVFile(safariBookmarks, "./testdata/out/safari.tsv")
}
