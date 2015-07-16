package utils

const URL_ECS_PREFIX = "http://docs.aliyun.com/#/pub/ecs/open-api/"
const URL_OSS_PREFIX = "http://docs.aliyun.com/#/pub/oss/api-reference/"
const DATATYPE_URL_PREFIX = "http://docs.aliyun.com/#/pub/ecs/open-api/datatype&"

var ECSAPI = map[string]bool{
	"instance":      true,
	"disk":          true,
	"snapshot":      true,
	"image":         true,
	"network":       true,
	"securitygroup": true,
	"vpc":           true,
	"vrouter":       true,
	"vswitch":       true,
	"routertable":   true,
	"region":        true,
	"monitor":       true,
	"other":         true,
	"datatype":      true,
	"appendix":      true,
}

var OSSAPI = map[string]bool{
	"service":          true,
	"bucket":           true,
	"object":           true,
	"multipart-upload": true,
	"cors":             true,
}

type Funclist map[string]bool

// Oss mapping relationship from SDK function to Docs.
var ossChart = map[string]string{
	"PutBucket":        "PutBucket",
	"ACL":              "GetBucketAcl",
	"PutBucketWebsite": "PutBucketWebsite",
	"Location":         "GetBucketLoaction",
	"DelBucket":        "DeleteBucket",
	"List":             "GetBucket",
	"Put":              "PutObject",
	"PutCopy":          "CopyObject",
	"Get":              "GetObject",
	"Del":              "DeleteObject",
	"DelMulti":         "DeleteMultipleObjects",
	"Head":             "HeadObject",
	"InitMulti":        "InitiateMultipartUpload",
	"PutPart":          "UploadPart",
	"PutPartCopy":      "UploadPartCopy",
	"Complete":         "Complete MultipartUpload",
	"Abort":            "AbortMultipartUpload",
	"ListParts":        "ListPartsFull",
}

var ossDocs = make(map[string]Funclist)
var ecsDocs = make(map[string]Funclist)
