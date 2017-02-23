package api

import (
    "domio_public/components/config"
    "net/http"
    "time"
    "log"
    "fmt"
)

type RecordValue struct {
    Value string
}

type Record struct {
    Type            string
    SetIdentifier   string
    ResourceRecords []RecordValue
}

func GetSubscriptionRecords(subscriptionId string, token string) []Record {

    var subRecords []Record

    url := config.Config.ApiUrl

    req, requestErr := http.NewRequest("GET", fmt.Sprintf("%s/subscriptions/%s/records", url, subscriptionId), nil)

    req.Header.Set("Content-Type", "application/json")

    timeout := time.Duration(5 * time.Second)

    client := &http.Client{
        Timeout: timeout,
    }

    req.Header.Add("Authorization", "Bearer " + token)

    resp, requestErr := client.Do(req)

    if requestErr != nil {
        log.Print(requestErr)
        return subRecords
    }
    defer resp.Body.Close()

    decodeError := DecodeJsonRequestBody(resp, &subRecords)

    if (decodeError != nil) {
        log.Print(decodeError)
        return subRecords
    }

    log.Print("*********************************************************")
    log.Print(subRecords)
    log.Print("*********************************************************")

    return subRecords
}