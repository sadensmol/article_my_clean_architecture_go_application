package tests_test

import (
	"context"
	"fmt"
	jsonpath "github.com/steinfletcher/apitest-jsonpath"
	"net/http"
	"net/http/cookiejar"
	"time"

	contractv1 "github.com/sadensmol/article_my_clean_architecture_go_application/api/proto/v1"
	"github.com/steinfletcher/apitest"
	"google.golang.org/grpc"
	"google.golang.org/grpc/benchmark"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	testAccountID int64 = 1
)

func (s *AccountTestSuite) TestGetAccountById() {
	clientConn := benchmark.NewClientConn("localhost:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	client := contractv1.NewAccountServiceClient(clientConn)

	response, err := client.GetById(context.Background(), &contractv1.GetAccountRequest{
		Id: testAccountID,
	})

	s.Require().NoError(err)

	s.Require().NotNil(response)
	s.Require().Equal(testAccountID, response.GetId())
	s.Require().Equal("test_account", response.GetName())
	s.Require().Equal(contractv1.AccountStatus_ACCOUNT_STATUS_OPEN, response.GetStatus())
	s.Require().Equal(contractv1.AccessLevel_ACCOUNT_ACCESS_LEVEL_FULL_ACCESS, response.AccessLevel)
	s.Require().NotNil(response.GetOpenedDate())
	s.Require().Nil(response.GetClosedDate())
}
func (s *AccountTestSuite) TestGetAccountByIdHttp() {
	cookieJar, _ := cookiejar.New(nil)
	cli := &http.Client{
		Timeout: time.Second * 1,
		Jar:     cookieJar,
	}

	apitest.New(). // configuration
			EnableNetworking(cli).
			Get(fmt.Sprintf("http://localhost:8090/api/v1/account/%d", testAccountID)).
			Expect(s.T()).
			Assert(jsonpath.Chain().
				Equal("id", "1").
				Equal("name", "test_account").
				Equal("status", "ACCOUNT_STATUS_OPEN").
				Equal("accessLevel", "ACCOUNT_ACCESS_LEVEL_FULL_ACCESS").
				Present("openedDate").
				NotPresent("closedDate").
				End()).
			Status(http.StatusOK).
			End()
}
