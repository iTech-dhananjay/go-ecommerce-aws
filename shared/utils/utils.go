package utils

import (
    "encoding/json"
    "log"
)

func MarshalRequest(data interface{}) []byte {
    reqBody, err := json.Marshal(data)
    if err != nil {
        log.Fatalf("Error marshalling request: %v", err)
    }
    return reqBody
}
