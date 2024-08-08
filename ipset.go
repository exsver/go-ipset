package ipset

import "encoding/xml"

type Ipsets struct {
	XMLName xml.Name `xml:"ipsets"`
	Ipset   []Ipset  `xml:"ipset"`
}

type Ipset struct {
	Name     string       `xml:"name,attr"`
	Type     string       `xml:"type"`
	Revision int          `xml:"revision"`
	Header   IpsetHeader  `xml:"header"`
	Members  IpsetMembers `xml:"members"`
}

type IpsetHeader struct {
	Family     string `xml:"family"`
	HashSize   int    `xml:"hashsize"`
	MaxElem    int    `xml:"maxelem"`
	BucketSize int    `xml:"bucketsize"`
	InitVal    string `xml:"initval"`
	MemSize    int    `xml:"memsize"`
	References int    `xml:"references"`
	NumEntries int    `xml:"numentries"`
}

type IpsetMembers struct {
	Members []IpsetMember `xml:"member"`
}

type IpsetMember struct {
	Elem    string `xml:"elem"`
	Timeout int    `xml:"timeout"`
	Packets int    `xml:"packets"`
	Bytes   int    `xml:"bytes"`
	Comment string `xml:"comment"`
}
