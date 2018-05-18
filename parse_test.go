package iabconsent_test

import (
	"time"

	"github.com/LiveRamp/iabconsent"
	"github.com/go-check/check"
)

type parseSuite struct{}

var _ = check.Suite(&parseSuite{})

func (s *parseSuite) TestConsentReader_ReadInt(c *check.C) {
	var tests = []struct {
		expected int
		n        uint
	}{
		{1, 1},
		{0, 1},
		{5, 3},
		{2, 3},
	}

	var r = iabconsent.NewConsentReader([]byte{0xaa})
	for _, t := range tests {
		c.Check(r.ReadInt(t.n), check.Equals, t.expected)
	}
	c.Check(r.HasUnread(), check.Equals, false)
}

func (s *parseSuite) TestConsentReader_ReadTime(c *check.C) {
	// 2018-05-18 17:48:31.5 +0000 UTC
	// 1526665711.5 s
	// 15266657115 deci-seconds
	// 0x38df6b35b deci-seconds (hex)
	var r = iabconsent.NewConsentReader([]byte{0x38, 0xdf, 0x6b, 0x35, 0xB0})
	c.Check(r.ReadTime(), check.DeepEquals, time.Unix(1526665711, int64(500*time.Millisecond)).UTC())
	c.Check(r.NumUnread(), check.Equals, 4)
}

func (s *parseSuite) TestConsentReader_ReadBoolMap(c *check.C) {
	var tests = []struct {
		expected map[int]bool
		n        uint
	}{
		{map[int]bool{
			2: true,
		}, 2},
		{map[int]bool{
			2: true,
			3: true,
			5: true,
		}, 6},
	}

	var r = iabconsent.NewConsentReader([]byte{0x5a})
	for _, t := range tests {
		c.Check(r.ReadPurposes(t.n), check.DeepEquals, t.expected)
	}
	c.Check(r.HasUnread(), check.Equals, false)
}
