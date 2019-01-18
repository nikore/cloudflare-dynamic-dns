package iputils_test

import (
	"github.com/nikore/cloudflare-dynamic-dns/pkg/iputils"
	. "gopkg.in/check.v1"
	"testing"
)

func TestPublicIp(t *testing.T) {
	TestingT(t)
}

type PublicIpTestSuite struct{}

var _ = Suite(&PublicIpTestSuite{})

func (s *PublicIpTestSuite) TestGetPublicIp(c *C) {
	ip, err := iputils.GetPublicIp()

	c.Assert(err, IsNil)
	c.Assert(ip, Matches, `(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}`)
}
