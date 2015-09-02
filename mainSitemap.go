package sitemap

import "time"

const (
	headerXmlns = "http://www.sitemaps.org/schemas/sitemap/0.9"
)

// SitemapContainer is the page
type Container struct {
	Urlset []*SitemapList `xml:"urlset"`
	Xmlns  string         `xml:"xmlns,attr"`
}

// // SitemapList is only a name older tu make it conform to the sitemap structure
// type SitemapList struct {
// 	Map *Sitemap `xml:"url"`
// }

// Sitemap is actualy the real sitemap element
type Sitemap struct {
	Loc        string
	LastMod    time.Time
	Changefreq string
	Priority   float32
}
