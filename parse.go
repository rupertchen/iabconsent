package iabconsent

import (
	"github.com/rupertchen/go-bits"
	"time"
	"encoding/base64"
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

func Parse2(s string) (*ParsedConsent, error) {

	var b, err = base64.RawURLEncoding.DecodeString(s)
	if err != nil {
		return nil, err
	}

	var r = NewConsentReader(b)

	// This block of code directly describes the format of the payload.
	var p = &ParsedConsent{}
	// Is setting consentString, still interesting?
	p.version = r.ReadInt(6)
	p.created = r.ReadTime()
	p.lastUpdated = r.ReadTime()
	p.cmpID = r.ReadInt(12)
	p.cmpVersion = r.ReadInt(12)
	p.consentScreen = r.ReadInt(6)
	p.consentLanguage = r.ReadString(2)
	p.vendorListVersion = r.ReadInt(12)
	p.purposesAllowed = r.ReadBoolMap(24)

	return p, nil
}
