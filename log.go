package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

var no int64 = 0

const (
	// ApacheCommonLog : {host} {user-identifier} {auth-user-id} [{datetime}] "{method} {request} {protocol}" {response-code} {bytes}
	ApacheCommonLog = "%s - %s [%s] \"%s %s %s\" %d %d"
	// ApacheCombinedLog : {host} {user-identifier} {auth-user-id} [{datetime}] "{method} {request} {protocol}" {response-code} {bytes} "{referrer}" "{agent}"
	ApacheCombinedLog = "%s - %s [%s] \"%s %s %s\" %d %d \"%s\" \"%s\""
	// ApacheErrorLog : [{timestamp}] [{module}:{severity}] [pid {pid}:tid {thread-id}] [client %{client}:{port}] %{message}
	ApacheErrorLog = "[%s] [%s:%s] [pid %d:tid %d] [client %s:%d] %s"
	// RFC3164Log : <priority>{timestamp} {hostname} {application}[{pid}]: {message}
	RFC3164Log = "<%d>%s %s %s[%d]: %s"
	// RFC5424Log : <priority>{version} {iso-timestamp} {hostname} {application} {pid} {message-id} {structured-data} {message}
	RFC5424Log = "<%d>%d %s %s %s %d ID%d %s %s"
	// CommonLogFormat : {host} {user-identifier} {auth-user-id} [{datetime}] "{method} {request} {protocol}" {response-code} {bytes}
	CommonLogFormat = "%s - %s [%s] \"%s %s %s\" %d %d"
	// JSONLogFormat : {"host": "{host}", "user-identifier": "{user-identifier}", "datetime": "{datetime}", "method": "{method}", "request": "{request}", "protocol": "{protocol}", "status", {status}, "bytes": {bytes}, "referer": "{referer}"}
	JSONLogFormat = `{"no": "%d", "host":"%s", "user-identifier":"%s", "time":"%s", "method": "%s", "request": "%s", "protocol":"%s", "status":%d, "bytes":%d, "referer": "%s"}`
	JSONLogNK     = `{"no": "%d", "level":"%s","time":"%s","logger":"gateway","message":"%s","app_name":"gateway","code":0,"file":"log.go:75","module_name":"gopkg.mihoyo.com/plat/kit/http/middleware","function":"Log.func1","hostname":"%s","env":"testing","request_id":"%s","method":"%s","url":"%s","status_code":%d,"content_length":%d,"user_agent":"%s","elapsed":%d,"remote_ip":"%s","client_ip":"%s"}`
	WAFLog        = `{"timestamp":%d,"formatVersion":1,"webaclId":"arn:aws:wafv2:us-east-1:012345678:regional/webacl/testwaf/02bb031c-f4fe-4318-a00f-88f99723bd18","terminatingRuleId":"Default_Action","terminatingRuleType":"REGULAR","action":"ALLOW","terminatingRuleMatchDetails":[],"httpSourceName":"APPSYNC","httpSourceId":"arn:aws:appsync:us-east-1:0123456789:apis/b7xkryfoandybmswajvnoqycoq","ruleGroupList":[],"rateBasedRuleList":[],"nonTerminatingMatchingRules":[],"requestHeadersInserted":null,"responseCodeSent":null,"httpRequest":{"clientIp":"%s","country":"%s","headers":[{"name":"accept","value":"*/*"},{"name":"accept-encoding","value":"gzip, deflate, br"},{"name":"accept-language","value":"zh-CN,zh;q=0.9"},{"name":"authorization","value":"eyJraWQiOiIwamFQMmhuem5pT05LQ1pCbkMwV3VKUVlLU0JuRzl1bnpIRFpNdnRGd2owPSIsImFsZyI6IlJTMjU2In0.eyJzdWIiOiIyNTVlMmI4Yy0xZTAyLTRlZWMtODViOC1mYjMzMmIxOGVhMzYiLCJpc3MiOiJodHRwczpcL1wvY29nbml0by1pZHAudXMtZWFzdC0xLmFtYXpvbmF3cy5jb21cL3VzLWVhc3QtMV83b0YwaVUwSzYiLCJjb2duaXRvOnVzZXJuYW1lIjoiMjU1ZTJiOGMtMWUwMi00ZWVjLTg1YjgtZmIzMzJiMThlYTM2Iiwib3JpZ2luX2p0aSI6IjkwM2Q5YjQxLTg5YzktNDQxNC1iNjdlLWNiY2M0MmQ1MGY4NSIsImF1ZCI6IjdjYmcwcGlvNzF1NjI0aXIzYmkxdmtkYzZoIiwiZXZlbnRfaWQiOiJkZTBhNTU0NC00ZTVmLTQ2MGEtOTUzMS0yMTBhMGViODg3NjQiLCJ0b2tlbl91c2UiOiJpZCIsImF1dGhfdGltZSI6MTY3MjkyODA2OCwiZXhwIjoxNjcyOTMxNjY4LCJpYXQiOjE2NzI5MjgwNjgsImp0aSI6IjM2ZDY2ZTA5LTAxZTItNDg1My04NmM3LTdjMDRhMGM0YTRlZCIsImVtYWlsIjoidGFvbWF0QGFtYXpvbi5jb20ifQ.HKbbd0-accUwrh5tUBR73Rv3Mn3KvYfkrdNRqjnaPX7Y582FnmpSeMU2SwTVUravKX29xeT2MYzYA1w9mS-OObAu2vqEeRtjcayfzWFeOAiNYR9NAMQ4mCE8QezxowpH5sRJWBqLQ2gi_CRUgQwumAVxv21vz2SYzDzfZMaQHwR3bL_3RMFtcj2EcjdiGJHN8srFLwg8ZrkI5icX96b4K6Ghhf8GXTL4tHjT7so1XsqedcylrfwkH2yJA2nm2nVgyC1LyK5JGu_HWnhjOwj2FzqgciIRhfGFAJtO5J4MLHIMwaI46--9nwY6nurC7XpYUsJP4hnlH2KvUAA6u6narA"},{"name":"cloudfront-forwarded-proto","value":"%s"},{"name":"cloudfront-is-desktop-viewer","value":"true"},{"name":"cloudfront-is-mobile-viewer","value":"false"},{"name":"cloudfront-is-smarttv-viewer","value":"false"},{"name":"cloudfront-is-tablet-viewer","value":"false"},{"name":"cloudfront-viewer-asn","value":"16509"},{"name":"cloudfront-viewer-country","value":"%s"},{"name":"content-length","value":"480"},{"name":"content-type","value":"application/json"},{"name":"host","value":"ed53tvkwujhu5cokt4ikulyoju.appsync-api.us-east-1.amazonaws.com"},{"name":"origin","value":"https://d3lt12oxze37kc.cloudfront.net"},{"name":"sec-ch-ua","value":"\"Not?A_Brand\";v=\"8\", \"Chromium\";v=\"108\", \"Google Chrome\";v=\"108\""},{"name":"sec-ch-ua-mobile","value":"?0"},{"name":"sec-ch-ua-platform","value":"\"macOS\""},{"name":"sec-fetch-dest","value":"empty"},{"name":"sec-fetch-mode","value":"cors"},{"name":"sec-fetch-site","value":"cross-site"},{"name":"user-agent","value":"%s"},{"name":"via","value":"2.0 42cd5a36cf9ad881ea4b618b3995860a.cloudfront.net (CloudFront)"},{"name":"x-amz-cf-id","value":"AX8s3To-T6kcYQF0bID-a1oTfOqeaRhyjIYlcLmuCFH2CTUjYEv8EQ=="},{"name":"x-amz-user-agent","value":"aws-amplify/3.0.7"},{"name":"x-amzn-trace-id","value":"Root=1-63b6dedf-15c57f0718bea4e91c64207e"},{"name":"x-forwarded-for","value":"%s, %s"},{"name":"x-forwarded-port","value":"%d"},{"name":"x-forwarded-proto","value":"%s"}],"uri":"%s","args":"","httpVersion":"%s","httpMethod":"%s","requestId":"%s"}}`
	LogNK         = `{"datetime":"%s", "body":"%s"}`
)

// NewApacheCommonLog creates a log string with apache common log format
func NewApacheCommonLog(t time.Time) string {
	return fmt.Sprintf(
		ApacheCommonLog,
		gofakeit.IPv4Address(),
		RandAuthUserID(),
		t.Format(Apache),
		gofakeit.HTTPMethod(),
		RandResourceURI(),
		RandHTTPVersion(),
		gofakeit.HTTPStatusCode(),
		gofakeit.Number(0, 30000),
	)
}

// NewApacheCombinedLog creates a log string with apache combined log format
func NewApacheCombinedLog(t time.Time) string {
	return fmt.Sprintf(
		ApacheCombinedLog,
		gofakeit.IPv4Address(),
		RandAuthUserID(),
		t.Format(Apache),
		gofakeit.HTTPMethod(),
		RandResourceURI(),
		RandHTTPVersion(),
		gofakeit.HTTPStatusCode(),
		gofakeit.Number(30, 100000),
		gofakeit.URL(),
		gofakeit.UserAgent(),
	)
}

// NewApacheErrorLog creates a log string with apache error log format
func NewApacheErrorLog(t time.Time) string {
	return fmt.Sprintf(
		ApacheErrorLog,
		t.Format(ApacheError),
		gofakeit.Word(),
		gofakeit.LogLevel("apache"),
		gofakeit.Number(1, 10000),
		gofakeit.Number(1, 10000),
		gofakeit.IPv4Address(),
		gofakeit.Number(1, 65535),
		gofakeit.HackerPhrase(),
	)
}

// NewRFC3164Log creates a log string with syslog (RFC3164) format
func NewRFC3164Log(t time.Time) string {
	return fmt.Sprintf(
		RFC3164Log,
		gofakeit.Number(0, 191),
		t.Format(RFC3164),
		strings.ToLower(gofakeit.Username()),
		gofakeit.Word(),
		gofakeit.Number(1, 10000),
		gofakeit.HackerPhrase(),
	)
}

// NewRFC5424Log creates a log string with syslog (RFC5424) format
func NewRFC5424Log(t time.Time) string {
	return fmt.Sprintf(
		RFC5424Log,
		gofakeit.Number(0, 191),
		gofakeit.Number(1, 3),
		t.Format(RFC5424),
		gofakeit.DomainName(),
		gofakeit.Word(),
		gofakeit.Number(1, 10000),
		gofakeit.Number(1, 1000),
		"-", // TODO: structured data
		gofakeit.HackerPhrase(),
	)
}

// NewCommonLogFormat creates a log string with common log format
func NewCommonLogFormat(t time.Time) string {
	return fmt.Sprintf(
		CommonLogFormat,
		gofakeit.IPv4Address(),
		RandAuthUserID(),
		t.Format(CommonLog),
		gofakeit.HTTPMethod(),
		RandResourceURI(),
		RandHTTPVersion(),
		gofakeit.HTTPStatusCode(),
		gofakeit.Number(0, 30000),
	)
}

// NewJSONLogFormat creates a log string with json log format
func NewJSONLogFormat(t time.Time) string {
	no++
	return fmt.Sprintf(
		JSONLogFormat,
		no,
		gofakeit.IPv4Address(),
		RandAuthUserID(),
		t.Format(CommonLog),
		gofakeit.HTTPMethod(),
		RandResourceURI(),
		RandHTTPVersion(),
		gofakeit.HTTPStatusCode(),
		gofakeit.Number(0, 30000),
		gofakeit.URL(),
	)
}

func NewJSONLogNK(t time.Time, n uint) string {
	no++
	return fmt.Sprintf(
		JSONLogNK,
		no,
		gofakeit.LogLevel("apache"),
		t.Format(RFC5424),
		gofakeit.LetterN(n*340),
		gofakeit.IPv4Address(),
		RandAuthUserID(),
		gofakeit.HTTPMethod(),

		gofakeit.URL(),

		gofakeit.HTTPStatusCode(),
		gofakeit.Number(0, 30000),
		gofakeit.UserAgent(),
		gofakeit.Number(0, 1000),
		gofakeit.IPv4Address(),
		gofakeit.IPv4Address(),
	)
}

// `"timestamp":%s, "clientIp": "%s","country":"%","cloudfront-forwarded-proto":"%s","cloudfront-viewer-country":"%s","user-agent":"%s","x-forwarded-for":"ip%s,ip%s","x-forwarded-port":"%d","x-forwarded-proto":"%s","uri": "%s","args": "%s","httpVersion":"%s","httpMethod":"%s","requestId":"%s"`
func WAFFormat(t time.Time) string {
	//no++
	return fmt.Sprintf(
		WAFLog,
		t.UnixMilli(),
		gofakeit.IPv4Address(),
		gofakeit.CountryAbr(),
		RandHTTPVersion(),
		gofakeit.CountryAbr(),
		gofakeit.UserAgent(),
		gofakeit.IPv4Address(),
		gofakeit.IPv4Address(),
		gofakeit.Number(1, 65535),
		RandHTTPVersion(),
		RandResourceURI(),
		RandHTTPVersion(),
		gofakeit.HTTPMethod(),
		gofakeit.UUID(),
	)
}
