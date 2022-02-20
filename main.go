package main

import (
	"context"
	"errors"
	"strconv"
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
	parsed_time, parse_err := time.Parse(layout, timeStamp.Time)
	if parse_err != nil {
		unix_val, err := strconv.ParseInt(timeStamp.Time, 10, 64)
		if err != nil {
			return TimestampResponse{}, errors.New("Invalid date format. Function only accepts ISO and Unix date formats.")
		}

		parsed_time = time.Unix(unix_val, 0) // Try to parse as UNIX timestamp...
	}

	return TimestampResponse{UnixTime: parsed_time.Unix(), UTCTime: parsed_time.UTC().String()}, nil
}

func main() {
	lambda.Start(HandleTimestampRequest)
}
