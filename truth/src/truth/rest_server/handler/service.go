package truth

import (
  "fmt"
  "github.com/golang/glog"
  "encoding/json"
  "net/http"

  cfg "truth/config"
  pb "truth_pb"
)

func Version(w http.ResponseWriter, req *http.Request) {
  glog.Infof("Version (%v, %v)", w, req)
  w.Header().Set("Content-Type", "application/json")

  rsp, err := json.Marshal(map[string]string{
    "version": cfg.VERSION })

  if err != nil {
    glog.Infof("json marshal failed: %v", err)
    http.Error(w, "Internal server error", http.StatusInternalServerError)
    return
  }

  fmt.Fprintf(w, "%s", rsp)
}

func ListServices(w http.ResponseWriter, req *http.Request) {
  glog.Infof("ListServices (%v, %v)", w, req)
  w.Header().Set("Content-Type", "application/json")
  services := []pb.Service{}

  rsp, err := json.Marshal(map[string][]pb.Service{
    "services": services })

  if err != nil {
    glog.Infof("json marshal failed: %v", err)
    http.Error(w, "Internal server error", http.StatusInternalServerError)
    return
  }

  fmt.Fprintf(w, "%s", rsp)
}
