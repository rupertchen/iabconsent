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
	"time"
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
	IsRangeEncoding   bool
	approvedVendorIDs map[int]bool
	DefaultConsent    bool
	NumEntries        int
	RangeEntries      []*RangeEntry
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

// VendorAllowed returns true if the ParsedConsent contains affirmative consent
// for VendorID v.
func (p *ParsedConsent) VendorAllowed(v int) bool {
	if p.IsRangeEncoding {
		for _, re := range p.RangeEntries {
			if re.StartVendorID <= v && v <= re.EndVendorID {
				return !p.DefaultConsent
			}
		}
		return p.DefaultConsent
	}

	return p.approvedVendorIDs[v]
}

// RangeEntry contains all fields in the RangeEncoding Entry
// portion of the Vendor Consent String. This portion
// of the consent string is only populated when the
// EncodingType field is set to 1.
type RangeEntry struct {
	StartVendorID int
	EndVendorID   int
}
