package service

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nytimes/gizmo/examples/nyt"
	"github.com/nytimes/gizmo/server"
	"github.com/nytimes/gizmo/web"
	"golang.org/x/net/context"
)

func (s *RPCService) GetMostPopular(ctx context.Context, r *MostPopularRequest) (*MostPopularResponse, error) {
	var (
		err error
		res []*nyt.MostPopularResult
	)
	defer server.MonitorRPCRequest()(ctx, "GetMostPopular", err)

	res, err = s.client.GetMostPopular(r.ResourceType, r.Section, uint(r.TimePeriodDays))
	if err != nil {
		return nil, err
	}
	return &MostPopularResponse{res}, nil
}

func (s *RPCService) GetMostPopularJSON(r *http.Request) (int, interface{}, error) {
	res, err := s.GetMostPopular(
		context.Background(),
		&MostPopularRequest{
			mux.Vars(r)["resourceType"],
			mux.Vars(r)["section"],
			uint32(web.GetUInt64Var(r, "timeframe")),
		})
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	return http.StatusOK, res.Results, nil
}
