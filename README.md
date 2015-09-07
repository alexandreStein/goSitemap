# sitemap
Is a package that generate XML sitemap from the given page objects.

It implement the sitemap page and sitemap index.

Documentation can be found at: http://godoc.org/github.com/alexandreStein/sitemap

Example:

package main

import (
	"fmt"

	"github.com/alexandreStein/sitemap"
)

func main() {
	sMap := sitemap.Page{
		Urlset: []*sitemap.Sitemap{
			&sitemap.Sitemap{
				Loc:        "https://www.test.com",
				Changefreq: "daily",
				Priority:   0.5,
			},
		},
	}
	x, err := sMap.GetXML()
	fmt.Println(string(x), err)

	sMapi := sitemap.IndexPage{
		Urlset: []*sitemap.Sitemap{
			&sitemap.Sitemap{
				Loc: "https://www.test.com",
			},
		},
	}

	x, err = sMapi.GetXML()
	fmt.Println(string(x), err)
}
