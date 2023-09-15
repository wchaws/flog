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
	JSONLogFormat         = `{"no": "%d", "host":"%s", "user-identifier":"%s", "time":"%s", "method": "%s", "request": "%s", "protocol":"%s", "status":%d, "bytes":%d, "referer": "%s"}`
	JSONLogNK             = `{"no": "%d", "level":"%s","time":"%s","logger":"gateway","message":"%s","app_name":"gateway","code":0,"file":"log.go:75","module_name":"gopkg.mihoyo.com/plat/kit/http/middleware","function":"Log.func1","hostname":"%s","env":"testing","request_id":"%s","method":"%s","url":"%s","status_code":%d,"content_length":%d,"user_agent":"%s","elapsed":%d,"remote_ip":"%s","client_ip":"%s"}`
	WAFLog                = `{"timestamp":%d,"formatVersion":1,"webaclId":"arn:aws:wafv2:us-east-1:012345678:regional/webacl/testwaf/02bb031c-f4fe-4318-a00f-88f99723bd18","terminatingRuleId":"Default_Action","terminatingRuleType":"REGULAR","action":"ALLOW","terminatingRuleMatchDetails":[],"httpSourceName":"APPSYNC","httpSourceId":"arn:aws:appsync:us-east-1:0123456789:apis/b7xkryfoandybmswajvnoqycoq","ruleGroupList":[],"rateBasedRuleList":[],"nonTerminatingMatchingRules":[],"requestHeadersInserted":null,"responseCodeSent":null,"httpRequest":{"clientIp":"%s","country":"%s","headers":[{"name":"accept","value":"*/*"},{"name":"accept-encoding","value":"gzip, deflate, br"},{"name":"accept-language","value":"zh-CN,zh;q=0.9"},{"name":"authorization","value":"eyJraWQiOiIwamFQMmhuem5pT05LQ1pCbkMwV3VKUVlLU0JuRzl1bnpIRFpNdnRGd2owPSIsImFsZyI6IlJTMjU2In0.eyJzdWIiOiIyNTVlMmI4Yy0xZTAyLTRlZWMtODViOC1mYjMzMmIxOGVhMzYiLCJpc3MiOiJodHRwczpcL1wvY29nbml0by1pZHAudXMtZWFzdC0xLmFtYXpvbmF3cy5jb21cL3VzLWVhc3QtMV83b0YwaVUwSzYiLCJjb2duaXRvOnVzZXJuYW1lIjoiMjU1ZTJiOGMtMWUwMi00ZWVjLTg1YjgtZmIzMzJiMThlYTM2Iiwib3JpZ2luX2p0aSI6IjkwM2Q5YjQxLTg5YzktNDQxNC1iNjdlLWNiY2M0MmQ1MGY4NSIsImF1ZCI6IjdjYmcwcGlvNzF1NjI0aXIzYmkxdmtkYzZoIiwiZXZlbnRfaWQiOiJkZTBhNTU0NC00ZTVmLTQ2MGEtOTUzMS0yMTBhMGViODg3NjQiLCJ0b2tlbl91c2UiOiJpZCIsImF1dGhfdGltZSI6MTY3MjkyODA2OCwiZXhwIjoxNjcyOTMxNjY4LCJpYXQiOjE2NzI5MjgwNjgsImp0aSI6IjM2ZDY2ZTA5LTAxZTItNDg1My04NmM3LTdjMDRhMGM0YTRlZCIsImVtYWlsIjoidGFvbWF0QGFtYXpvbi5jb20ifQ.HKbbd0-accUwrh5tUBR73Rv3Mn3KvYfkrdNRqjnaPX7Y582FnmpSeMU2SwTVUravKX29xeT2MYzYA1w9mS-OObAu2vqEeRtjcayfzWFeOAiNYR9NAMQ4mCE8QezxowpH5sRJWBqLQ2gi_CRUgQwumAVxv21vz2SYzDzfZMaQHwR3bL_3RMFtcj2EcjdiGJHN8srFLwg8ZrkI5icX96b4K6Ghhf8GXTL4tHjT7so1XsqedcylrfwkH2yJA2nm2nVgyC1LyK5JGu_HWnhjOwj2FzqgciIRhfGFAJtO5J4MLHIMwaI46--9nwY6nurC7XpYUsJP4hnlH2KvUAA6u6narA"},{"name":"cloudfront-forwarded-proto","value":"%s"},{"name":"cloudfront-is-desktop-viewer","value":"true"},{"name":"cloudfront-is-mobile-viewer","value":"false"},{"name":"cloudfront-is-smarttv-viewer","value":"false"},{"name":"cloudfront-is-tablet-viewer","value":"false"},{"name":"cloudfront-viewer-asn","value":"16509"},{"name":"cloudfront-viewer-country","value":"%s"},{"name":"content-length","value":"480"},{"name":"content-type","value":"application/json"},{"name":"host","value":"ed53tvkwujhu5cokt4ikulyoju.appsync-api.us-east-1.amazonaws.com"},{"name":"origin","value":"https://d3lt12oxze37kc.cloudfront.net"},{"name":"sec-ch-ua","value":"\"Not?A_Brand\";v=\"8\", \"Chromium\";v=\"108\", \"Google Chrome\";v=\"108\""},{"name":"sec-ch-ua-mobile","value":"?0"},{"name":"sec-ch-ua-platform","value":"\"macOS\""},{"name":"sec-fetch-dest","value":"empty"},{"name":"sec-fetch-mode","value":"cors"},{"name":"sec-fetch-site","value":"cross-site"},{"name":"user-agent","value":"%s"},{"name":"via","value":"2.0 42cd5a36cf9ad881ea4b618b3995860a.cloudfront.net (CloudFront)"},{"name":"x-amz-cf-id","value":"AX8s3To-T6kcYQF0bID-a1oTfOqeaRhyjIYlcLmuCFH2CTUjYEv8EQ=="},{"name":"x-amz-user-agent","value":"aws-amplify/3.0.7"},{"name":"x-amzn-trace-id","value":"Root=1-63b6dedf-15c57f0718bea4e91c64207e"},{"name":"x-forwarded-for","value":"%s, %s"},{"name":"x-forwarded-port","value":"%d"},{"name":"x-forwarded-proto","value":"%s"}],"uri":"%s","args":"","httpVersion":"%s","httpMethod":"%s","requestId":"%s"}}`
	WAFALBLog             = `{ "timestamp": %d, "formatVersion": 1, "webaclId": "arn:aws:wafv2:us-east-1:012345678912:regional/webacl/test-waf/8221a98c-d0d6-4e95-9f8b-ea2f4a5488c4", "terminatingRuleId": "Default_Action", "terminatingRuleType": "REGULAR", "action": "ALLOW", "terminatingRuleMatchDetails": [], "httpSourceName": "ALB", "httpSourceId": "012345678912-app/alb/05d7c48b12495e0a", "ruleGroupList": [ { "ruleGroupId": "AWS#AWSManagedRulesAmazonIpReputationList", "terminatingRule": null, "nonTerminatingMatchingRules": [], "excludedRules": null, "customerConfig": null }, { "ruleGroupId": "AWS#AWSManagedRulesCommonRuleSet", "terminatingRule": null, "nonTerminatingMatchingRules": [], "excludedRules": null, "customerConfig": null }, { "ruleGroupId": "AWS#AWSManagedRulesKnownBadInputsRuleSet", "terminatingRule": null, "nonTerminatingMatchingRules": [], "excludedRules": null, "customerConfig": null }, { "ruleGroupId": "AWS#AWSManagedRulesSQLiRuleSet", "terminatingRule": { "ruleId": "SQLi_BODY", "action": "BLOCK", "ruleMatchDetails": null }, "nonTerminatingMatchingRules": [], "excludedRules": null, "customerConfig": null } ], "rateBasedRuleList": [], "nonTerminatingMatchingRules": [], "requestHeadersInserted": null, "responseCodeSent": null, "httpRequest": { "clientIp": "%s", "country": "%s", "headers": [ { "name": "Host", "value": "example.host.io" }, { "name": "User-Agent", "value": "%s" } ], "uri": "%s", "args": "", "httpVersion": "%s", "httpMethod": "%s", "requestId": "%s" }, "webaclName": "test-waf", "host": "test-alb.us-east-1.elb.amazonaws.com", "userAgent": "%s" }`
	WAFALBSQLIngestionLog = `{ "timestamp":%d, "formatVersion": 1, "webaclId": "arn:aws:wafv2:us-east-1:012345678912:regional/webacl/test-waf/8221a98c-d0d6-4e95-9f8b-ea2f4a5488c4", "terminatingRuleId": "AWS-AWSManagedRulesSQLiRuleSet", "terminatingRuleType": "MANAGED_RULE_GROUP", "action": "BLOCK", "terminatingRuleMatchDetails": [ { "conditionType": "SQL_INJECTION", "location": "BODY", "matchedData": [ "--------------------------310d893689fd5869\r\nContent-Disposition", "THIS", "IS", "THE", "SQL_INJECTION", "!", "--------------------------310d893689fd5869--\r" ], "sensitivityLevel": "LOW" } ], "httpSourceName": "ALB", "httpSourceId": "012345678912-app/alb/05d7c48b12495e0a", "ruleGroupList": [ { "ruleGroupId": "AWS#AWSManagedRulesAmazonIpReputationList", "terminatingRule": null, "nonTerminatingMatchingRules": [], "excludedRules": null, "customerConfig": null }, { "ruleGroupId": "AWS#AWSManagedRulesCommonRuleSet", "terminatingRule": null, "nonTerminatingMatchingRules": [], "excludedRules": null, "customerConfig": null }, { "ruleGroupId": "AWS#AWSManagedRulesKnownBadInputsRuleSet", "terminatingRule": null, "nonTerminatingMatchingRules": [], "excludedRules": null, "customerConfig": null }, { "ruleGroupId": "AWS#AWSManagedRulesSQLiRuleSet", "terminatingRule": { "ruleId": "SQLi_BODY", "action": "BLOCK", "ruleMatchDetails": null }, "nonTerminatingMatchingRules": [], "excludedRules": null, "customerConfig": null } ], "rateBasedRuleList": [], "nonTerminatingMatchingRules": [], "requestHeadersInserted": null, "responseCodeSent": null, "httpRequest": { "clientIp": "%s", "country": "%s", "headers": [ { "name": "Host", "value": "example.host.io" }, { "name": "Accept", "value": "*/*" }, { "name": "User-Agent", "value": "%s" }, { "name": "Content-Length", "value": "147" }, { "name": "Content-Type", "value": "multipart/form-data; boundary=------------------------310d893689fd5869" } ], "uri": "%s", "args": "", "httpVersion": "%s", "httpMethod": "%s", "requestId": "%s" }, "labels": [ { "name": "awswaf:managed:aws:sql-database:SQLi_Body" } ], "webaclName": "test-waf", "host": "test-alb.us-east-1.elb.amazonaws.com", "userAgent": "%s" }`
	WAFNonSQLIngestionLog = `{ "timestamp":%d, "formatVersion": 1, "webaclId": "arn:aws:wafv2:us-east-1:012345678912:regional/webacl/test-waf/8221a98c-d0d6-4e95-9f8b-ea2f4a5488c4", "terminatingRuleId": "AWS-AWSManagedRulesSQLiRuleSet", "terminatingRuleType": "MANAGED_RULE_GROUP", "action": "BLOCK", "terminatingRuleMatchDetails": [ { "conditionType": "SQL_INJECTION", "location": "BODY", "matchedData": [ "--------------------------310d893689fd5869\r\nContent-Disposition", "THIS", "IS", "THE", "WRONGLY", "BLOCKED", "PHOTO", "!", "--------------------------310d893689fd5869--\r" ], "sensitivityLevel": "LOW" } ], "httpSourceName": "ALB", "httpSourceId": "012345678912-app/alb/05d7c48b12495e0a", "ruleGroupList": [ { "ruleGroupId": "AWS#AWSManagedRulesAmazonIpReputationList", "terminatingRule": null, "nonTerminatingMatchingRules": [], "excludedRules": null, "customerConfig": null }, { "ruleGroupId": "AWS#AWSManagedRulesCommonRuleSet", "terminatingRule": null, "nonTerminatingMatchingRules": [], "excludedRules": null, "customerConfig": null }, { "ruleGroupId": "AWS#AWSManagedRulesKnownBadInputsRuleSet", "terminatingRule": null, "nonTerminatingMatchingRules": [], "excludedRules": null, "customerConfig": null }, { "ruleGroupId": "AWS#AWSManagedRulesSQLiRuleSet", "terminatingRule": { "ruleId": "SQLi_BODY", "action": "BLOCK", "ruleMatchDetails": null }, "nonTerminatingMatchingRules": [], "excludedRules": null, "customerConfig": null } ], "rateBasedRuleList": [], "nonTerminatingMatchingRules": [], "requestHeadersInserted": null, "responseCodeSent": null, "httpRequest": { "clientIp": "%s", "country": "%s", "headers": [ { "name": "Host", "value": "waf-testalb-2062338277.us-east-2.elb.amazonaws.com" }, { "name": "User-Agent", "value": "curl/7.87.0" }, { "name": "accept", "value": "text/plain" }, { "name": "Content-Type", "value": "multipart/form-data" }, { "name": "Authorization", "value": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJodHRwOi8vc2NoZW1hcy54bWxzb2FwLm9yZy93cy8yMDA1LzA1L2lkZW50aXR5L2NsYWltcy9uYW1laWRlbnRpZmllciI6IjE3NCIsImh0dHA6Ly9zY2hlbWFzLnhtbHNvYXAub3JnL3dzLzIwMDUvMDUvaWRlbnRpdHkvY2xhaW1zL25hbWUiOiI4NzY4NzgzNzMxIiwiQXNwTmV0LklkZW50aXR5LlNlY3VyaXR5U3RhbXAiOiJIUFk1WFMzWU80TFRRUUFENVRCNFVBNFJYU0lJSFA2UCIsImh0dHA6Ly9zY2hlbWFzLm1pY3Jvc29mdC5jb20vd3MvMjAwOC8wNi9pZGVudGl0eS9jbGFpbXMvcm9sZSI6IkN1c3RvbWVyIiwiaHR0cDovL3d3dy5hc3BuZXRib2lsZXJwbGF0ZS5jb20vaWRlbnRpdHkvY2xhaW1zL3RlbmFudElkIjoiMSIsInN1YiI6IjE3NCIsImp0aSI6IjMzOTA0NDM4LTEzZjktNDIwZi1iYTI0LWIwNTU2NTI3ZTU1YyIsImlhdCI6MTYwMTQxNjA4MywibmJmIjoxNjAxNDE2MDgzLCJleHAiOjE2MDE1MDI0ODMsImlzcyI6IlVuaWZ5IiwiYXVkIjoiVW5pZnkifQ.Bm-1DC0qcqk6JGKYBaM9_qLW6onC27K5hvbFXAc1h6o" }, { "name": "Content-Length", "value": "14" } ], "uri": "%s", "args": "", "httpVersion": "%s", "httpMethod": "POST", "requestId": "%s" }, "labels": [ { "name": "awswaf:managed:aws:sql-database:SQLi_Body" } ], "webaclName": "test-waf", "host": "test-alb.us-east-1.elb.amazonaws.com", "userAgent": "%s" }`
	CloudFrontRealTimeLog = `%.3f	%s	%.3f	%d	%d	%s	https	%s.cloudfront.net	%s	%d	%s	%s	%s.cloudfront.net	%.3f	HTTP/2.0	IPv4	Mozilla/5.0%%20(Macintosh;%%20Intel%%20Mac%%20OS%%20X%%2010_15_7)%%20AppleWebKit/537.36%%20(KHTML,%%20like%%20Gecko)%%20Chrome/108.0.0.0%%20Safari/537.36	-	-	v=v1.2.0	Hit	-	TLSv1.3	TLS_AES_128_GCM_SHA256	Hit	-	-	application/json	-	-	-	17929	Hit	JP	gzip,%%20deflate,%%20br	*/*	*	host:d2omsvnrx2yu4i.cloudfront.net%%0Apragma:no-cache%%0Acache-control:no-cache%%0Asec-ch-ua:%%22Not?A_Brand%%22;v%%22%%22,%%20%%22Chromiu%%22;v%%2210%%22,%%2%%22Google%%20Chrome%%22;v%%2210%%22%%0Asec-ch-ua-mobile:?0%%0Auser-agent:Mozilla/5.0%%20(Macintosh;%%20Intel%%20Mac%%20OS%%20X%%2010_15_7)%%20AppleWebKit/537.36%%20(KHTML,%%20like%%20Gecko)%%20Chrome/108.0.0.0%%20Safari/537.36%%0Asec-ch-ua-platform%%22macO%%22%%0Aaccept:*/*%%0Asec-fetch-site:same-origin%%0Asec-fetch-mode:cors%%0Asec-fetch-dest:empty%%0Aaccept-encoding:gzip,%%20deflate,%%20br%%0Aaccept-language:en,zh-CN;q=0.9,zh;q=0.8%%0A	host%%0Apragma%%0Acache-control%%0Asec-ch-ua%%0Asec-ch-ua-mobile%%0Auser-agent%%0Asec-ch-ua-platform%%0Aaccept%%0Asec-fetch-site%%0Asec-fetch-mode%%0Asec-fetch-dest%%0Aaccept-encoding%%0Aaccept-language%%0A	13	E32QJQ394BWBEI	d2omsvnrx2yu4i.cloudfront.net	-	-	16509`
	LogNK                 = `{"datetime":"%s", "body":"%s"}`
	NestedJSON            = `{"timestamp":"%s","correlationId":"%d","processInfo":{"startTime":"%s","hostname":"ltvtix0apidev01","domainId":"%s","groupId":"group-2","groupName":"grp_dev_bba","serviceId":"instance-1","serviceName":"ins_dev_bba","version":"7.7.20210130"},"transactionSummary":{"path":"%s","protocol":"https","protocolSrc":"%d","status":"exception","serviceContexts":[{"service":"NSC_APP-117127_DCTM_Get Documentum Token","monitor":true,"client":"Pass Through","org":null,"app":null,"method":"getTokenUsingPOST","status":"exception","duration":%d}]}}`
)

func NewNestedJSON(t time.Time) string {
	return fmt.Sprintf(
		NestedJSON,
		t.Format(RFC5424),
		gofakeit.Uint64(),
		t.Format(RFC5424),
		gofakeit.UUID(),
		gofakeit.URL(),
		gofakeit.Number(0, 1000),
		gofakeit.Uint16(),
	)
}

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

func NewWAFALBLog(t time.Time) string {
	//no++
	return fmt.Sprintf(
		WAFALBLog,
		t.UnixMilli(),
		gofakeit.IPv4Address(),
		gofakeit.CountryAbr(),
		gofakeit.UserAgent(),
		RandResourceURI(),
		RandHTTPVersion(),
		gofakeit.HTTPMethod(),
		gofakeit.UUID(),
		gofakeit.UserAgent(),
	)
}

// `"timestamp":%s, "clientIp": "%s","country":"%","user-agent":"%s","uri": "%s","args": "%s","httpVersion":"%s","httpMethod":"%s","requestId":"%s"`
func NewWAFSQLIngestionLog(t time.Time) string {
	//no++
	return fmt.Sprintf(
		WAFALBSQLIngestionLog,
		t.UnixMilli(),
		gofakeit.IPv4Address(),
		gofakeit.CountryAbr(),
		gofakeit.UserAgent(),
		RandResourceURI(),
		RandHTTPVersion(),
		gofakeit.HTTPMethod(),
		gofakeit.UUID(),
		gofakeit.UserAgent(),
	)
}

// `"timestamp":%s, "clientIp": "%s","country":"%","user-agent":"%s","uri": "%s","args": "%s","httpVersion":"%s","httpMethod":"%s","requestId":"%s"`
func NewWAFNonSQLIngestionLog(t time.Time) string {
	//no++
	return fmt.Sprintf(
		WAFNonSQLIngestionLog,
		t.UnixMilli(),
		gofakeit.IPv4Address(),
		gofakeit.CountryAbr(),
		gofakeit.UserAgent(),
		RandResourceURI(),
		RandHTTPVersion(),
		gofakeit.UserAgent(),
	)
}

func NewCloudFrontRealTimeLog(t time.Time) string {
	return fmt.Sprintf(
		CloudFrontRealTimeLog,
		float64(t.UnixMilli())/1000,  // timestamp
		gofakeit.IPv4Address(),       // c-ip
		gofakeit.Float32Range(0, 1),  // time-to-first-byte
		gofakeit.HTTPStatusCode(),    // sc-status
		gofakeit.IntRange(100, 5000), // sc-bytes
		gofakeit.HTTPMethod(),        // cs-method
		gofakeit.LetterN(10),         // cs-host
		strings.TrimLeft(
			strings.TrimLeft(
				gofakeit.URL(),
				"https://",
			),
			"http://",
		), // cs-uri-stem
		gofakeit.IntRange(100, 200), // cs-bytes
		gofakeit.CountryAbr(),       // x-edge-location
		gofakeit.UUID(),             // x-edge-request-id
		gofakeit.LetterN(10),        // x-host-header
		gofakeit.Float32Range(0, 1), // time-taken
	)
}
