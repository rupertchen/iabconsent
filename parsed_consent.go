/*

Package iabconsent provides structs and methods for parsing
Vendor Consent Strings as defined by the IAB Consent String 1.1 Spec.
More info on the spec here:
https://github.com/InteractiveAdvertisingBureau/GDPR-Transparency-and-Consent-Framework/blob/master/Consent%20string%20and%20vendor%20list%20formats%20v1.1%20Final.md#vendor-consent-string-format-.

Copyright (c) 2018 LiveRamp. All rights reserved.

Written by Andy Day, Software Engineer @ LiveRamp
for use in the LiveRamp Pixel Server.

*/
package iabconsent

import (
	"encoding/base64"
	"time"
)

// These constants represent the bit offsets and sizes of the
// fields in the IAB Consent String 1.1 Spec.
const (
	VersionBitOffset        = 0
	VersionBitSize          = 6
	CreatedBitOffset        = 6
	CreatedBitSize          = 36
	UpdatedBitOffset        = 42
	UpdatedBitSize          = 36
	CmpIdOffset             = 78
	CmpIdSize               = 12
	CmpVersionOffset        = 90
	CmpVersionSize          = 12
	ConsentScreenSizeOffset = 102
	ConsentScreenSize       = 6
	ConsentLanguageOffset   = 108
	ConsentLanguageSize     = 12
	VendorListVersionOffset = 120
	VendorListVersionSize   = 12
	PurposesOffset          = 132
	PurposesSize            = 24
	MaxVendorIdOffset       = 156
	MaxVendorIdSize         = 16
	EncodingTypeOffset      = 172
	VendorBitFieldOffset    = 173
	DefaultConsentOffset    = 173
	NumEntriesOffset        = 174
	NumEntriesSize          = 12
	RangeEntryOffset        = 186
	VendorIdSize            = 16
)

// ParsedConsent contains all fields defined in the
// IAB Consent String 1.1 Spec.
type ParsedConsent struct {
	Version           int
	Created           time.Time
	LastUpdated       time.Time
	CMPID             int
	CMPVersion        int
	ConsentScreen     int
	ConsentLanguage   string
	VendorListVersion int
	PurposesAllowed   map[int]bool
	MaxVendorID       int
	IsRange           bool
	approvedVendorIDs map[int]bool
	DefaultConsent    bool
	numEntries        int
	rangeEntries      []*RangeEntry
}

// EveryPurposeAllowed returns true iff every purpose number in ps exists in
// the ParsedConsent, otherwise false.
func (p *ParsedConsent) EveryPurposeAllowed(ps []int) bool {
	for _, rp := range ps {
		if !p.PurposesAllowed[rp] {
			return false
		}
	}
	return true
}

// VendorAllowed returns true if the ParsedConsent contains
// affirmative consent for Vendor of ID |i|.
func (p *ParsedConsent) VendorAllowed(i int) bool {
	if p.IsRange {
		// DefaultConsent indicates the consent for those
		// not covered by any Range Entries. Vendors covered
		// in rangeEntries have the opposite consent of
		// DefaultConsent.
		for _, re := range p.rangeEntries {
			if re.StartVendorID <= i &&
				re.EndVendorID >= i {
				return !p.DefaultConsent
			}
		}
	} else {
		var _, ok = p.approvedVendorIDs[i]
		return ok
	}
	return p.DefaultConsent
}

// RangeEntry contains all fields in the Range Entry
// portion of the Vendor Consent String. This portion
// of the consent string is only populated when the
// EncodingType field is set to 1.
type RangeEntry struct {
	StartVendorID int
	EndVendorID   int
}

// Parse takes a base64 Raw URL Encoded string which represents
// a Vendor Consent String and returns a ParsedConsent with
// it's fields populated with the values stored in the string.
// Example Usage:
//	var pc, err = iabconsent.Parse("BONJ5bvONJ5bvAMAPyFRAL7AAAAMhuqKklS-gAAAAAAAAAAAAAAAAAAAAAAAAAA")
func Parse(s string) (*ParsedConsent, error) {
	var b []byte
	var err error

	b, err = base64.RawURLEncoding.DecodeString(s)
	if err != nil {
		return nil, err
	}

	var bs = ParseBytes(b)
	var version, cmpID, cmpVersion, consentScreen, vendorListVersion, maxVendorID, numEntries int
	var created, updated time.Time
	var isRangeEntries, defaultConsent, isIDRange bool
	var consentLanguage string
	var purposesAllowed = make(map[int]bool)
	var approvedVendorIDs = make(map[int]bool)

	version, err = bs.ParseInt(VersionBitOffset, VersionBitSize)
	if err != nil {
		return nil, err
	}
	created, err = bs.ParseTime(CreatedBitOffset, CreatedBitSize)
	if err != nil {
		return nil, err
	}
	updated, err = bs.ParseTime(UpdatedBitOffset, UpdatedBitSize)
	if err != nil {
		return nil, err
	}
	cmpID, err = bs.ParseInt(CmpIdOffset, CmpIdSize)
	if err != nil {
		return nil, err
	}
	cmpVersion, err = bs.ParseInt(CmpVersionOffset, CmpVersionSize)
	if err != nil {
		return nil, err
	}
	consentScreen, err = bs.ParseInt(ConsentScreenSizeOffset, ConsentScreenSize)
	if err != nil {
		return nil, err
	}
	consentLanguage, err = bs.ParseString(ConsentLanguageOffset, ConsentLanguageSize)
	if err != nil {
		return nil, err
	}
	vendorListVersion, err = bs.ParseInt(VendorListVersionOffset, VendorListVersionSize)
	if err != nil {
		return nil, err
	}
	purposesAllowed, err = bs.ParseBitList(PurposesOffset, PurposesSize)
	if err != nil {
		return nil, err
	}
	maxVendorID, err = bs.ParseInt(MaxVendorIdOffset, MaxVendorIdSize)
	if err != nil {
		return nil, err
	}
	isRangeEntries, err = bs.ParseBool(EncodingTypeOffset)
	if err != nil {
		return nil, err
	}

	var rangeEntries []*RangeEntry

	if isRangeEntries {
		defaultConsent, err = bs.ParseBool(DefaultConsentOffset)
		if err != nil {
			return nil, err
		}
		numEntries, err = bs.ParseInt(NumEntriesOffset, NumEntriesSize)
		if err != nil {
			return nil, err
		}

		// Track how many range entry bits we've parsed since it's variable.
		var parsedBits = 0

		for i := 0; i < numEntries; i++ {
			var startVendorID, endVendorID int

			isIDRange, err = bs.ParseBool(RangeEntryOffset + parsedBits)
			parsedBits++

			if isIDRange {
				startVendorID, err = bs.ParseInt(RangeEntryOffset+parsedBits, VendorIdSize)
				if err != nil {
					return nil, err
				}
				parsedBits += VendorIdSize
				endVendorID, err = bs.ParseInt(RangeEntryOffset+parsedBits, VendorIdSize)
				if err != nil {
					return nil, err
				}
				parsedBits += VendorIdSize
			} else {
				startVendorID, err = bs.ParseInt(RangeEntryOffset+parsedBits, VendorIdSize)
				if err != nil {
					return nil, err
				}
				endVendorID = startVendorID
				parsedBits += VendorIdSize
			}

			rangeEntries = append(rangeEntries, &RangeEntry{
				StartVendorID: startVendorID,
				EndVendorID:   endVendorID,
			})
		}
	} else {
		approvedVendorIDs, err = bs.ParseBitList(VendorBitFieldOffset, maxVendorID)
		if err != nil {
			return nil, err
		}
	}

	return &ParsedConsent{
		consentString:     bs.value,
		Version:           version,
		Created:           created,
		LastUpdated:       updated,
		CMPID:             cmpID,
		CMPVersion:        cmpVersion,
		ConsentScreen:     consentScreen,
		ConsentLanguage:   consentLanguage,
		VendorListVersion: vendorListVersion,
		PurposesAllowed:   purposesAllowed,
		MaxVendorID:       maxVendorID,
		IsRange:           isRangeEntries,
		approvedVendorIDs: approvedVendorIDs,
		DefaultConsent:    defaultConsent,
		numEntries:        numEntries,
		rangeEntries:      rangeEntries,
	}, nil
}
