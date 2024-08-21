package fragranitca

import (
	creds "fragnatica/tools"
	"log"
	"net/http/cookiejar"
	"time"

	"github.com/gocolly/colly"
)

func CollyScrape(urls []string) []Images {
	userAgent, err := creds.GetRandomUserAgent("tools/user_agents.txt")
	if err != nil {
		log.Fatal(err)
	}

	c := colly.NewCollector(
		colly.AllowedDomains("www.fragrantica.ru"),
		colly.UserAgent(userAgent),
	)

	cookiejar, _ := cookiejar.New(nil)

	c.SetCookieJar(cookiejar)

	var images []Images

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
		r.Headers.Set("Accept-Language", "en-US,en;q=0.9")
		r.Headers.Set("Connection", "keep-alive")
		r.Headers.Set("Referer", "https://www.google.com/")
		r.Headers.Set("Upgrade-Insecure-Requests", "1")
		r.Headers.Set("DNT", "1")
		c.UserAgent = userAgent
		time.Sleep(2 * time.Second)
		log.Println("Visiting", r.URL)
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Println("Something went wrong:", err)
		if r.StatusCode == 429 {
			log.Println("Sleeping for 1 minute...")
			time.Sleep(3 * time.Minute)
			r.Request.Retry()
		}
	})

	c.OnResponse(func(r *colly.Response) {
		log.Println("Visited", r.Request.URL)
	})

	c.OnHTML("#brands img:not(.hide-for-medium img)", func(e *colly.HTMLElement) {
		imageUrl := Images{}

		imageUrl.Url = e.Attr("src")

		images = append(images, imageUrl)
	})

	c.Visit("https://www.fragrantica.ru")
	for _, url := range urls {
		c.Visit(url)
	}

	return images
}
