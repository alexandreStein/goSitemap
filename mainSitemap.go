package sitemap

import (
	"encoding/xml"
	"time"
)

const (
	headerXmlns = "http://www.sitemaps.org/schemas/sitemap/0.9"
)

// Page is the XML page
type Page struct {
	XMLName xml.Name   `xml:"urlset"`
	Urlset  []*Sitemap `xml:"url"`
	Xmlns   string     `xml:"xmlns,attr"`
}

// GetXML return the corresponding sitemap XML
func (c *Page) GetXML() ([]byte, error) {
	c.Xmlns = headerXmlns

	retXML, err := xml.Marshal(c)
	// retXML, err := xml.MarshalIndent(c, "", "	")
	if err != nil {
		return nil, err
	}

	ret := []byte(xml.Header)
	ret = append(ret, retXML...)

	return ret, nil
}

// IndexPage represent the sitemap index element
type IndexPage struct {
	XMLName xml.Name   `xml:"sitemapindex"`
	Urlset  []*Sitemap `xml:"sitemap"`
	Xmlns   string     `xml:"xmlns,attr"`
}

// GetXML return the corresponding sitemap XML
func (c *IndexPage) GetXML() ([]byte, error) {
	c.Xmlns = headerXmlns

	retXML, err := xml.Marshal(c)
	if err != nil {
		return nil, err
	}

	ret := []byte(xml.Header)
	ret = append(ret, retXML...)

	return ret, nil
}

// Sitemap is actualy the real sitemap element
type Sitemap struct {
	Loc        string     `xml:"loc"`
	LastMod    *time.Time `xml:"lastmod,omitempty"`
	Changefreq string     `xml:"changefreq,omitempty"`
	Priority   float32    `xml:"priority,omitempty"`
}
