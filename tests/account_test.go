package tests_test

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type AccountTestSuite struct {
	suite.Suite
}

func TestAccountTestSuite(t *testing.T) {
	suite.Run(t, new(AccountTestSuite))

}
