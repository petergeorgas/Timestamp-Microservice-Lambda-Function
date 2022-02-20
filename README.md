## Go Timestamp Microservice Lambda 
This repository houses the code for implementing a simple timestamp microservice to be run as an AWS Lambda Function. 

### Usage
To use, simply send a request with a JSON body of the format: 
``` js
{
  "time":"YYYY-MM-DD"
}
```
or 
``` js
{
  "time":"<unix timestamp>"
}
```
and you will receive back both the Unix timestamp and UTC representation of the timestamp you sent: 
``` js
{
  "unix": 1645383157,
  "utc": "2022-02-20 18:52:37 +0000 UTC"
}
```
**Note**: malformed strings should be accounted for. Currently the service only deals with ISO and Unix-timestamp formatted time strings.
