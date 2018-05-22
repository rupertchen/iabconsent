package iabconsent

import (
	"encoding/base64"
	"fmt"
	"time"

	"github.com/rupertchen/go-bits"
)

type ConsentReader struct {
	*bits.Reader
}

func NewConsentReader(src []byte) *ConsentReader {
	return &ConsentReader{bits.NewReader(bits.NewBitmap(src))}
}

func (r *ConsentReader) ReadInt(n uint) int {
	return int(r.ReadBits(n))
}

func (r *ConsentReader) ReadTime() time.Time {
	var ds = int64(r.ReadBits(36))
	return time.Unix(ds/dsPerS, (ds%dsPerS)*nsPerDs).UTC()
}

func (r *ConsentReader) ReadString(n uint) string {
	var buf = make([]byte, 0, n)
	for i := uint(0); i < n; i++ {
		buf = append(buf, byte(r.ReadBits(6))+'A')
	}
	return string(buf)
}

func (r *ConsentReader) ReadPurposes(n uint) map[int]bool {
	var m = make(map[int]bool)
	for i := uint(0); i < n; i++ {
		if r.ReadBool() {
			m[int(i)+1] = true
		}
	}
	return m
}

func (r *ConsentReader) ReadRangeEntries(n uint) []*RangeEntry {
	var ret = make([]*RangeEntry, 0, n)
	for i := uint(0); i < n; i++ {
		var isRange = r.ReadBool()
		var start, end int
		start = r.ReadInt(16)
		if isRange {
			end = r.ReadInt(16)
		} else {
			end = start
		}
		ret = append(ret, &RangeEntry{StartVendorID: start, EndVendorID: end})
	}
	return ret
}

// Parse takes a base64 Raw URL Encoded string which represents
// a Vendor Consent String and returns a ParsedConsent with
// it's fields populated with the values stored in the string.
//
// Example Usage:
//
//   var pc, err = iabconsent.Parse("BONJ5bvONJ5bvAMAPyFRAL7AAAAMhuqKklS-gAAAAAAAAAAAAAAAAAAAAAAAAAA")
func Parse(s string) (p *ParsedConsent, err error) {
	// This func leverages named returns to return partially parsed content when there is an error

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()

	b, err := base64.RawURLEncoding.DecodeString(s)
	if err != nil {
		return nil, err
	}

	var r = NewConsentReader(b)

	// This block of code directly describes the format of the payload.
	p = &ParsedConsent{}
	// TODO: Is setting consentString, still interesting?
	p.Version = r.ReadInt(6)
	p.Created = r.ReadTime()
	p.LastUpdated = r.ReadTime()
	p.CMPID = r.ReadInt(12)
	p.CMPVersion = r.ReadInt(12)
	p.ConsentScreen = r.ReadInt(6)
	p.ConsentLanguage = r.ReadString(2)
	p.VendorListVersion = r.ReadInt(12)
	p.PurposesAllowed = r.ReadPurposes(24)
	p.MaxVendorID = r.ReadInt(16)

	p.IsRangeEncoding = r.ReadBool()
	if p.IsRangeEncoding {
		p.DefaultConsent = r.ReadBool()
		p.NumEntries = r.ReadInt(12)
		p.RangeEntries = r.ReadRangeEntries(uint(p.NumEntries))
	} else {
		p.approvedVendorIDs = r.ReadPurposes(uint(p.MaxVendorID))
	}

	return p, nil
}
