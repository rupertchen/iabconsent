package iabconsent

import "time"

type consentType int

const (
	BitField consentType = iota
	SingleRangeWithSingleID
	SingleRangeWithRange
	MultipleRangesWithSingleID
	MultipleRangesWithRange
	MultipleRangesMixed
)

var testTime = time.Unix(1525378200, 8*nsPerDs).UTC()

var consentFixtures = map[consentType]*ParsedConsent{
	// BONMj34ONMj34ABACDENALqAAAAAplY
	BitField: {
		Version:           1,
		Created:           testTime,
		LastUpdated:       testTime,
		CMPID:             1,
		CMPVersion:        2,
		ConsentScreen:     3,
		ConsentLanguage:   "EN",
		VendorListVersion: 11,
		PurposesAllowed: map[int]bool{
			1: true,
			3: true,
			5: true,
		},
		MaxVendorID:     10,
		IsRangeEncoding: false,
		approvedVendorIDs: map[int]bool{
			1:  true,
			2:  true,
			5:  true,
			7:  true,
			9:  true,
			10: true,
		},
	},
	// BONMj34ONMj34ABACDENALqAAAAAqABAD2AAAAAAAAAAAAAAAAAAAAAAAAAA
	SingleRangeWithSingleID: {
		Version:           1,
		Created:           testTime,
		LastUpdated:       testTime,
		CMPID:             1,
		CMPVersion:        2,
		ConsentScreen:     3,
		ConsentLanguage:   "EN",
		VendorListVersion: 11,
		PurposesAllowed: map[int]bool{
			1: true,
			3: true,
			5: true,
		},
		MaxVendorID:     10,
		IsRangeEncoding: true,
		DefaultConsent:  false,
		NumEntries:      1,
		RangeEntries: []*RangeEntry{
			{
				StartVendorID: 123,
				EndVendorID:   123,
			},
		},
	},
	// BONMj34ONMj34ABACDENALqAAAAAqABgD2AdQAAAAAAAAAAAAAAAAAAAAAAAAAA
	SingleRangeWithRange: {
		Version:           1,
		Created:           testTime,
		LastUpdated:       testTime,
		CMPID:             1,
		CMPVersion:        2,
		ConsentScreen:     3,
		ConsentLanguage:   "EN",
		VendorListVersion: 11,
		PurposesAllowed: map[int]bool{
			1: true,
			3: true,
			5: true,
		},
		MaxVendorID:     10,
		IsRangeEncoding: true,
		DefaultConsent:  false,
		NumEntries:      1,
		RangeEntries: []*RangeEntry{
			{
				StartVendorID: 123,
				EndVendorID:   234,
			},
		},
	},
	// BONMj34ONMj34ABACDENALqAAAAAqACAD2AOoAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
	MultipleRangesWithSingleID: {
		Version:           1,
		Created:           testTime,
		LastUpdated:       testTime,
		CMPID:             1,
		CMPVersion:        2,
		ConsentScreen:     3,
		ConsentLanguage:   "EN",
		VendorListVersion: 11,
		PurposesAllowed: map[int]bool{
			1: true,
			3: true,
			5: true,
		},
		MaxVendorID:     10,
		IsRangeEncoding: true,
		DefaultConsent:  false,
		NumEntries:      2,
		RangeEntries: []*RangeEntry{
			{
				StartVendorID: 123,
				EndVendorID:   123,
			},
			{
				StartVendorID: 234,
				EndVendorID:   234,
			},
		},
	},
	// BONMj34ONMj34ABACDENALqAAAAAqACgD2AdUBWQHIAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
	MultipleRangesWithRange: {
		Version:           1,
		Created:           testTime,
		LastUpdated:       testTime,
		CMPID:             1,
		CMPVersion:        2,
		ConsentScreen:     3,
		ConsentLanguage:   "EN",
		VendorListVersion: 11,
		PurposesAllowed: map[int]bool{
			1: true,
			3: true,
			5: true,
		},
		MaxVendorID:     10,
		IsRangeEncoding: true,
		DefaultConsent:  false,
		NumEntries:      2,
		RangeEntries: []*RangeEntry{
			{
				StartVendorID: 123,
				EndVendorID:   234,
			},
			{
				StartVendorID: 345,
				EndVendorID:   456,
			},
		},
	},
	// BONMj34ONMj34ABACDENALqAAAAAqACAD3AVkByAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
	MultipleRangesMixed: {
		Version:           1,
		Created:           testTime,
		LastUpdated:       testTime,
		CMPID:             1,
		CMPVersion:        2,
		ConsentScreen:     3,
		ConsentLanguage:   "EN",
		VendorListVersion: 11,
		PurposesAllowed: map[int]bool{
			1: true,
			3: true,
			5: true,
		},
		MaxVendorID:     10,
		IsRangeEncoding: true,
		DefaultConsent:  false,
		NumEntries:      2,
		RangeEntries: []*RangeEntry{
			{
				StartVendorID: 123,
				EndVendorID:   123,
			},
			{
				StartVendorID: 345,
				EndVendorID:   456,
			},
		},
	},
}
