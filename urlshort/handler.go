package urlshort

import (
	"log"
	"net/http"

	"gopkg.in/yaml.v2"
)

type UrlShort struct {
	Path string
	Url  string
}

func MapHandler(urlsshort map[string]string, h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		longurl, present := urlsshort[r.URL.Path]
		if present {
			http.Redirect(w, r, longurl, 301)
			return
		}
		h.ServeHTTP(w, r)
		return
	}
}

func YAMLHandler(yml string, h http.Handler) http.HandlerFunc {
	return MapHandler(YAMLParse(yml), h)
}

func YAMLParse(yml string) map[string]string {
	urls := []UrlShort{}
	urlsshort := map[string]string{}
	err := yaml.Unmarshal([]byte(yml), &urls)
	if err != nil {
		log.Fatalln(err)
	}
	for _, url := range urls {
		urlsshort[url.Path] = url.Url
	}
	return urlsshort
}
