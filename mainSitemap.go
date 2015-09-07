package sitemap

import (
	"encoding/xml"
	"time"
)

const (
	headerXmlns = "http://www.sitemaps.org/schemas/sitemap/0.9"
)

// Those values represent the posible change frequences
const (
	Always  ChangeFreq = "always"
	Hourly  ChangeFreq = "hourly"
	Daily   ChangeFreq = "daily"
	Weekly  ChangeFreq = "weekly"
	Monthly ChangeFreq = "monthly"
	Yearly  ChangeFreq = "yearly"
	Never   ChangeFreq = "never"
)

// ChangeFreq is a type to only permit coerant choice
type ChangeFreq string

// Page is the XML page
type Page struct {
	// Set the name of the element
	XMLName xml.Name   `xml:"urlset"`
	Urlset  []*Sitemap `xml:"url"`
	Xmlns   string     `xml:"xmlns,attr"`
}

// GetXML return the corresponding sitemap XML
func (c *Page) GetXML() ([]byte, error) {
	// Add the sitemap attribut
	c.Xmlns = headerXmlns

	// Marshal the object to bytes
	retXML, err := xml.Marshal(c)
	// retXML, err := xml.MarshalIndent(c, "", "	")
	if err != nil {
		return nil, err
	}

	// Add the XML header at the first place and concat the rest of the XML export
	ret := []byte(xml.Header)
	ret = append(ret, retXML...)

	return ret, nil
}

// IndexPage represent the sitemap index element
type IndexPage struct {
	// Set the name of the element
	XMLName xml.Name   `xml:"sitemapindex"`
	Urlset  []*Sitemap `xml:"sitemap"`
	Xmlns   string     `xml:"xmlns,attr"`
}

// GetXML return the corresponding sitemap XML
func (c *IndexPage) GetXML() ([]byte, error) {
	// Add the sitemap attribut
	c.Xmlns = headerXmlns

	// Marshal the object to bytes
	retXML, err := xml.Marshal(c)
	// retXML, err := xml.MarshalIndent(c, "", "	")
	if err != nil {
		return nil, err
	}

	// Add the XML header at the first place and concat the rest of the XML export
	ret := []byte(xml.Header)
	ret = append(ret, retXML...)

	return ret, nil
}

// Sitemap is actualy the real sitemap element
type Sitemap struct {
	Loc        string     `xml:"loc"`
	LastMod    *time.Time `xml:"lastmod,omitempty"`
	Changefreq ChangeFreq `xml:"changefreq,omitempty"`
	Priority   float32    `xml:"priority,omitempty"`
}
