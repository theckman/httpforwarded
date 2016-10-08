package httpforwarded_test

import (
	"github.com/theckman/go-httpforwarded"
	. "gopkg.in/check.v1"
)

func (*TestSuite) TestParse(c *C) {
	testParseSingleParam(c)
	testParseMultiParam(c)
	testParseMultiLine(c)
	testParseMultiParamValue(c)
	testParseAllTheThings(c)
}

func testParseSingleParam(c *C) {
	var err error
	var params map[string][]string

	values := []string{"for=192.0.2.1"}

	params, err = httpforwarded.Parse(values)
	c.Assert(err, IsNil)
	c.Check(len(params), Equals, 1)

	forVal, ok := params["for"]
	c.Assert(ok, Equals, true)
	c.Assert(len(forVal), Equals, 1)
	c.Check(forVal[0], Equals, "192.0.2.1")
}

func testParseMultiParam(c *C) {
	var err error
	var params map[string][]string

	values := []string{"for=192.0.2.1; proto=http"}

	params, err = httpforwarded.Parse(values)
	c.Assert(err, IsNil)
	c.Check(len(params), Equals, 2)

	forVal, ok := params["for"]
	c.Assert(ok, Equals, true)
	c.Assert(len(forVal), Equals, 1)
	c.Check(forVal[0], Equals, "192.0.2.1")

	proto, ok := params["proto"]
	c.Assert(ok, Equals, true)
	c.Assert(len(proto), Equals, 1)
	c.Check(proto[0], Equals, "http")
}

func testParseMultiLine(c *C) {
	var err error
	var params map[string][]string

	values := []string{
		"for=192.0.2.1",
		"proto=http; by=192.0.2.200",
	}

	params, err = httpforwarded.Parse(values)
	c.Assert(err, IsNil)
	c.Check(len(params), Equals, 3)

	forVal, ok := params["for"]
	c.Assert(ok, Equals, true)
	c.Assert(len(forVal), Equals, 1)
	c.Check(forVal[0], Equals, "192.0.2.1")

	proto, ok := params["proto"]
	c.Assert(ok, Equals, true)
	c.Assert(len(proto), Equals, 1)
	c.Check(proto[0], Equals, "http")

	by, ok := params["by"]
	c.Assert(ok, Equals, true)
	c.Assert(len(by), Equals, 1)
	c.Check(by[0], Equals, "192.0.2.200")
}

func testParseMultiParamValue(c *C) {
	var err error
	var params map[string][]string

	values := []string{"for=192.0.2.1, for=192.0.2.4; proto=http"}

	params, err = httpforwarded.Parse(values)
	c.Assert(err, IsNil)
	c.Check(len(params), Equals, 2)

	forVal, ok := params["for"]
	c.Assert(ok, Equals, true)
	c.Assert(len(forVal), Equals, 2)
	c.Check(forVal[0], Equals, "192.0.2.1")
	c.Check(forVal[1], Equals, "192.0.2.4")

	proto, ok := params["proto"]
	c.Assert(ok, Equals, true)
	c.Assert(len(proto), Equals, 1)
	c.Check(proto[0], Equals, "http")
}

func testParseAllTheThings(c *C) {
	var err error
	var params map[string][]string

	values := []string{
		"for=192.0.2.1; proto=http",
		"by=192.0.2.200",
		"for=192.0.2.4, for=192.0.2.10; by=192.0.2.202",
	}

	params, err = httpforwarded.Parse(values)
	c.Assert(err, IsNil)
	c.Check(len(params), Equals, 3)

	forVal, ok := params["for"]
	c.Assert(ok, Equals, true)
	c.Assert(len(forVal), Equals, 3)
	c.Check(forVal[0], Equals, "192.0.2.1")
	c.Check(forVal[1], Equals, "192.0.2.4")
	c.Check(forVal[2], Equals, "192.0.2.10")

	proto, ok := params["proto"]
	c.Assert(ok, Equals, true)
	c.Assert(len(proto), Equals, 1)
	c.Check(proto[0], Equals, "http")

	by, ok := params["by"]
	c.Assert(ok, Equals, true)
	c.Assert(len(by), Equals, 2)
	c.Check(by[0], Equals, "192.0.2.200")
	c.Check(by[1], Equals, "192.0.2.202")
}
