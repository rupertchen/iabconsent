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

var testTime = time.Unix(1525378200, 8*nsPerDs)

var consentFixtures = map[consentType]*ParsedConsent{
	// BONMj34ONMj34ABACDENALqAAAAAplY
	BitField: {
		consentString:     "0000010011100011010011001000111101111110000011100011010011001000111101111110000000000000010000000000100000110001000011010000000010111010100000000000000000000000000000001010011001010110",
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
		MaxVendorID: 10,
		IsRange:     false,
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
		consentString:     "000001001110001101001100100011110111111000001110001101001100100011110111111000000000000001000000000010000011000100001101000000001011101010000000000000000000000000000000101010000000000001000000000011110110000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
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
		MaxVendorID:       10,
		IsRange:           true,
		approvedVendorIDs: map[int]bool{},
		DefaultConsent:    false,
		numEntries:        1,
		rangeEntries: []*RangeEntry{
			{
				StartVendorID: 123,
				EndVendorID:   123,
			},
		},
	},
	// BONMj34ONMj34ABACDENALqAAAAAqABgD2AdQAAAAAAAAAAAAAAAAAAAAAAAAAA
	SingleRangeWithRange: {
		consentString:     "0000010011100011010011001000111101111110000011100011010011001000111101111110000000000000010000000000100000110001000011010000000010111010100000000000000000000000000000001010100000000000011000000000111101100000000111010100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
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
		MaxVendorID:       10,
		IsRange:           true,
		approvedVendorIDs: map[int]bool{},
		DefaultConsent:    false,
		numEntries:        1,
		rangeEntries: []*RangeEntry{
			{
				StartVendorID: 123,
				EndVendorID:   234,
			},
		},
	},
	// BONMj34ONMj34ABACDENALqAAAAAqACAD2AOoAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
	MultipleRangesWithSingleID: {
		consentString:     "00000100111000110100110010001111011111100000111000110100110010001111011111100000000000000100000000001000001100010000110100000000101110101000000000000000000000000000000010101000000000001000000000001111011000000000111010100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
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
		MaxVendorID:       10,
		IsRange:           true,
		approvedVendorIDs: map[int]bool{},
		DefaultConsent:    false,
		numEntries:        2,
		rangeEntries: []*RangeEntry{
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
		consentString:     "0000010011100011010011001000111101111110000011100011010011001000111101111110000000000000010000000000100000110001000011010000000010111010100000000000000000000000000000001010100000000000101000000000111101100000000111010101000000010101100100000001110010000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
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
		MaxVendorID:       10,
		IsRange:           true,
		approvedVendorIDs: map[int]bool{},
		DefaultConsent:    false,
		numEntries:        2,
		rangeEntries: []*RangeEntry{
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
		consentString:     "000001001110001101001100100011110111111000001110001101001100100011110111111000000000000001000000000010000011000100001101000000001011101010000000000000000000000000000000101010000000000010000000000011110111000000010101100100000001110010000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
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
		MaxVendorID:       10,
		IsRange:           true,
		approvedVendorIDs: map[int]bool{},
		DefaultConsent:    false,
		numEntries:        2,
		rangeEntries: []*RangeEntry{
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
