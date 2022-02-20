package main

import (
	"context"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
)

type TimestampEvent struct {
	Time string `json:"time"`
}

type TimestampResponse struct {
	UnixTime int64  `json:"unix"`
	UTCTime  string `json:"utc"`
}

func HandleTimestampRequest(ctx context.Context, timeStamp TimestampEvent) (TimestampResponse, error) {
	const layout = "2006-01-02"
	parsed_time, _ := time.Parse(layout, timeStamp.Time)
	return TimestampResponse{UnixTime: parsed_time.Unix(), UTCTime: parsed_time.UTC().String()}, nil
}

func main() {
	lambda.Start(HandleTimestampRequest)
}
