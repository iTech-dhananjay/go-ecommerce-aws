package handlers

import (
    "bytes"
    "encoding/json"
    "io/ioutil"
    "net/http"
    "api-gateway/shared"
)

func Login(w http.ResponseWriter, r *http.Request) {
    var reqBody map[string]string
    err := json.NewDecoder(r.Body).Decode(&reqBody)
    if err != nil {
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }

    resp, err := http.Post("http://auth-service:8081/login", "application/json", bytes.NewBuffer(shared.MarshalRequest(reqBody)))
    if err != nil {
        http.Error(w, "Service unavailable", http.StatusServiceUnavailable)
        return
    }
    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)
    w.WriteHeader(resp.StatusCode)
    w.Write(body)
}

func Register(w http.ResponseWriter, r *http.Request) {
    var reqBody map[string]string
    err := json.NewDecoder(r.Body).Decode(&reqBody)
    if err != nil {
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }

    resp, err := http.Post("http://auth-service:8081/register", "application/json", bytes.NewBuffer(shared.MarshalRequest(reqBody)))
    if err != nil {
        http.Error(w, "Service unavailable", http.StatusServiceUnavailable)
        return
    }
    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)
    w.WriteHeader(resp.StatusCode)
    w.Write(body)
}
