package api

import (
  "net/http"
  cql "github.com/gocql/gocql"
  "github.com/golang/glog"
  "phi/common"
)

type UserProfileApi struct {
  cass_sess *cql.Session
}

func NewUserProfileApi() *UserProfileApi {
  return &UserProfileApi{
    cass_sess: CassSession("phi_userprofile"),
  }
}

//  select user_id,sum(checkin_count) FROM user_checkin_count group by user_id
func (p *UserProfileApi) MetricsUserCheckin(w http.ResponseWriter, r *http.Request) {
  glog.Infof("%s %v", phi.CurrentScope(), r.URL)
}

func (p *UserProfileApi) MetricsVenueCheckin(w http.ResponseWriter, r *http.Request) {
  glog.Infof("%s %v", phi.CurrentScope(), r.URL)
}

func (p *UserProfileApi) MetricsCityCheckin(w http.ResponseWriter, r *http.Request) {
  glog.Infof("%s %v", phi.CurrentScope(), r.URL)
}

func (p *UserProfileApi) MetricsPois(w http.ResponseWriter, r *http.Request) {
  glog.Infof("%s %v", phi.CurrentScope(), r.URL)
}

func (p *UserProfileApi) MetricsCities(w http.ResponseWriter, r *http.Request) {
  glog.Infof("%s %v", phi.CurrentScope(), r.URL)
}

func (p *UserProfileApi) SetHandler(mux *http.ServeMux) {
  mux.HandleFunc("/v1/metrics/profile/user_checkin", p.MetricsUserCheckin)
  mux.HandleFunc("/v1/metrics/profile/venue_checkin", p.MetricsVenueCheckin)
  mux.HandleFunc("/v1/metrics/profile/city_checkin", p.MetricsCityCheckin)
  mux.HandleFunc("/v1/metrics/profile/pois", p.MetricsPois)
  mux.HandleFunc("/v1/metrics/profile/cities", p.MetricsCities)
}