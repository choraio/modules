package server

import (
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	validatorv1 "github.com/choraio/mods/validator/api/v1"
	v1 "github.com/choraio/mods/validator/types/v1"
)

type queryPolicy struct {
	*baseSuite
	res *v1.QueryPolicyResponse
	err error
}

func TestQueryPolicy(t *testing.T) {
	gocuke.NewRunner(t, &queryPolicy{}).
		Path("./features/query_policy.feature").
		Run()
}

func (s *queryPolicy) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
}

func (s *queryPolicy) Policy(a gocuke.DocString) {
	var policy validatorv1.Policy
	err := jsonpb.UnmarshalString(a.Content, &policy)
	require.NoError(s.t, err)

	err = s.srv.ss.PolicyTable().Save(s.ctx, &validatorv1.Policy{
		SignedBlocksWindow: policy.SignedBlocksWindow,
		MinSignedPerWindow: policy.MinSignedPerWindow,
	})
	require.NoError(s.t, err)
}

func (s *queryPolicy) QueryPolicy() {
	var req v1.QueryPolicyRequest
	s.res, s.err = s.srv.Policy(s.ctx, &req)
}

func (s *queryPolicy) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *queryPolicy) ExpectTheError(a gocuke.DocString) {
	require.EqualError(s.t, s.err, a.Content)
}

func (s *queryPolicy) ExpectResponse(a gocuke.DocString) {
	var expected v1.QueryPolicyResponse
	err := jsonpb.UnmarshalString(a.Content, &expected)
	require.NoError(s.t, err)

	require.Equal(s.t, &expected, s.res)
}
