package tests_test

import (
	"context"
	"net/http"
	"net/http/cookiejar"
	"time"

	contractv1 "github.com/sadensmol/article_my_clean_architecture_go_application/api/proto/v1"
	"github.com/steinfletcher/apitest"
	"github.com/steinfletcher/apitest-jsonpath"
	"google.golang.org/grpc"
	"google.golang.org/grpc/benchmark"
	"google.golang.org/grpc/credentials/insecure"
)

func (s *AccountTestSuite) TestCreateAccount() {
	clientConn := benchmark.NewClientConn("localhost:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	client := contractv1.NewAccountServiceClient(clientConn)
	response, err := client.Create(context.Background(), &contractv1.CreateAccountRequest{
		Name:        "test",
		AccessLevel: 2,
	})

	s.Require().NoError(err)

	s.Require().NotNil(response)
	s.Require().NotNil(response.GetId())
}
func (s *AccountTestSuite) TestCreateAccountHttp() {
	cookieJar, _ := cookiejar.New(nil)
	cli := &http.Client{
		Timeout: time.Second * 1,
		Jar:     cookieJar,
	}

	apitest.New(). // configuration
			EnableNetworking(cli).
			Post("http://localhost:8090/api/v1/account").
			Expect(s.T()).
			Assert(jsonpath.Present("id")).
			Status(http.StatusOK).
			End()
}
