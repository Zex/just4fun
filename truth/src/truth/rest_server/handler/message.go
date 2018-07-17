package truth

import (
  "fmt"
  "log"
  "encoding/json"
  "net/http"

  //cfg "truth/config"
  pb "truth_pb"
)

func UnreadMsg(w http.ResponseWriter, req *http.Request) {
  log.Printf("UnreadMsg (%v, %v)", w, req)
  w.Header().Set("Content-Type", "application/json")
  msgs := []pb.Message{}

  rsp, err := json.Marshal(map[string][]pb.Message{
    "messages": msgs})

  if err != nil {
    log.Printf("json marshal failed: %v", err)
    http.Error(w, "Internal server error", http.StatusInternalServerError)
    return
  }

  fmt.Fprintf(w, "%s", rsp)
}

func SendMsg(w http.ResponseWriter, req *http.Request) {
  log.Printf("SendMsg (%v, %v)", w, req)
  w.Header().Set("Content-Type", "application/json")
  res := pb.OpStatus{pb.StatusCode_Ok, ""}

  rsp, err := json.Marshal(map[string]pb.OpStatus{
    "status": res})

  if err != nil {
    log.Printf("json marshal failed: %v", err)
    http.Error(w, "Internal server error", http.StatusInternalServerError)
    return
  }

  fmt.Fprintf(w, "%s", rsp)
}