package httpforwarded_test

import (
	"testing"

	. "gopkg.in/check.v1"
)

type TestSuite struct{}

var _ = Suite(&TestSuite{})

func Test(t *testing.T) { TestingT(t) }
