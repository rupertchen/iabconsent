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
	// TODO
	return time.Time{}
}

func (r *ConsentReader) ReadString(n uint) string {
	// TODO
	return ""
}

func (r *ConsentReader) ReadBoolMap(n uint) map[int]bool {
	// TODO
	return make(map[int]bool)
}

// TODO: export rangeEntry or unexport this func
func (r *ConsentReader) ReadRangeEntries(n uint) []*rangeEntry {
	var ret = make([]*rangeEntry, 0, n)
	for i := uint(0); i < n; i++ {
		var isRange = r.ReadBool()
		var start, end int
		start = r.ReadInt(16)
		if isRange {
			end = r.ReadInt(16)
		} else {
			end = start
		}
		ret = append(ret, &rangeEntry{StartVendorID: start, EndVendorID: end})
	}
	return ret
}

func Parse2(s string) (p *ParsedConsent, err error) {
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
	p.version = r.ReadInt(6)
	p.created = r.ReadTime()
	p.lastUpdated = r.ReadTime()
	p.cmpID = r.ReadInt(12)
	p.cmpVersion = r.ReadInt(12)
	p.consentScreen = r.ReadInt(6)
	p.consentLanguage = r.ReadString(2)
	p.vendorListVersion = r.ReadInt(12)
	p.purposesAllowed = r.ReadBoolMap(24)
	p.maxVendorID = r.ReadInt(16)

	var hasRanges = r.ReadBool()
	if hasRanges {
		p.defaultConsent = r.ReadBool()
		p.numEntries = r.ReadInt(12)
		p.rangeEntries = r.ReadRangeEntries(uint(p.numEntries))
	} else {
		p.approvedVendorIDs = r.ReadBoolMap(uint(p.maxVendorID))
	}

	return p, nil
}
