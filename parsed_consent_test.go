package iabconsent

import (
	"sort"

	"github.com/go-check/check"
)

type ParsedConsentSuite struct{}

func (p *ParsedConsentSuite) TestParseConsentStrings(c *check.C) {
	var cases = []struct {
		Type          consentType
		EncodedString string
	}{
		{
			Type:          BitField,
			EncodedString: "BONMj34ONMj34ABACDENALqAAAAAplY",
		},
		{
			Type:          SingleRangeWithSingleID,
			EncodedString: "BONMj34ONMj34ABACDENALqAAAAAqABAD2AAAAAAAAAAAAAAAAAAAAAAAAAA",
		},
		{
			Type:          SingleRangeWithRange,
			EncodedString: "BONMj34ONMj34ABACDENALqAAAAAqABgD2AdQAAAAAAAAAAAAAAAAAAAAAAAAAA",
		},
		{
			Type:          MultipleRangesWithSingleID,
			EncodedString: "BONMj34ONMj34ABACDENALqAAAAAqACAD2AOoAAAAAAAAAAAAAAAAAAAAAAAAAAAAA",
		},
		{
			Type:          MultipleRangesWithRange,
			EncodedString: "BONMj34ONMj34ABACDENALqAAAAAqACgD2AdUBWQHIAAAAAAAAAAAAAAAAAAAAAAAAAAAAA",
		},
		{
			Type:          MultipleRangesMixed,
			EncodedString: "BONMj34ONMj34ABACDENALqAAAAAqACAD3AVkByAAAAAAAAAAAAAAAAAAAAAAAAAAAAA",
		},
	}

	for _, tc := range cases {
		c.Log(tc)
		pc, err := Parse(tc.EncodedString)
		c.Check(err, check.IsNil)

		normalizeParsedConsent(pc)
		normalizeParsedConsent(consentFixtures[tc.Type])

		c.Assert(pc, check.DeepEquals, consentFixtures[tc.Type])
	}
}

func normalizeParsedConsent(p *ParsedConsent) {
	sort.Slice(p.rangeEntries, func(i, j int) bool {
		return p.rangeEntries[i].StartVendorID < p.rangeEntries[j].StartVendorID
	})
}

var _ = check.Suite(&ParsedConsentSuite{})
