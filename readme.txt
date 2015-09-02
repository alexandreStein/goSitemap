This package try to implement the sitemap protocol.

It implement the sitemap page and sitemap index.

Documentation can be found at: http://godoc.org/bitbucket.org/astein58/sitemap

Example:

package main

import (
	"fmt"

	"bitbucket.org/astein58/lib_perso/sitemap"
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
