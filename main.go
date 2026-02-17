package main

import (
	"fmt"
	"net/http"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	var target string

	if path == "/" {
		target = "https://github.com/AnotherPillow/sdv-wiki"
	} else if strings.HasPrefix(path, "/m:") {
		// /m:Index -> stardewvalleywiki.com/Modding:Index
		target = "https://stardewvalleywiki.com/Modding:" + path[3:]
	} else if strings.HasPrefix(path, "/m/") {
		// /m/Category:Recommendations -> stardewmodding.wiki.gg/wiki/Category:Recommendations
		target = "https://stardewmodding.wiki.gg/wiki" + path[2:]
	} else {
		// /Abigail -> stardewvalleywiki.com/Abigail

		target = "https://stardewvalleywiki.com/" + path[1:]
	}

	http.Redirect(w, r, target, http.StatusMovedPermanently)
}

func main() {
	port := getPort()
	http.HandleFunc("/", handler)
	fmt.Printf("Listening on :%s", port)
	http.ListenAndServe(":"+port, nil)
}
