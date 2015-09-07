# sitemap
Is a package that generate XML sitemap from the given page objects.

It implement the sitemap page and sitemap index.

Documentation can be found at: http://godoc.org/github.com/alexandreStein/sitemap

Example:

```golang
package main

import (
	"fmt"
	"time"

	"github.com/alexandreStein/sitemap"
)

func main() {
	// Set the last modification to now
	now := time.Now()

	// Initialize a standard sitemap
	sMap := sitemap.Page{
		Urlset: []*sitemap.Sitemap{
			&sitemap.Sitemap{
				Loc:        "https://www.test.com",
				LastMod:    &now,
				Changefreq: sitemap.Daily,
				Priority:   0.5,
			},
		},
	}
	// Build the XML output
	x, err := sMap.GetXML()
	fmt.Println(string(x), err)

	// Initialize an index sitemap
	sMapi := sitemap.IndexPage{
		Urlset: []*sitemap.Sitemap{
			&sitemap.Sitemap{
				Loc: "https://www.test.com/sitemap.xml",
			},
		},
	}
	// Same as sitemap page build the XML using GetXML
	x, err = sMapi.GetXML()
	fmt.Println(string(x), err)
}

```
