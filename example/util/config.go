package util

import sls "github.com/galaxydi/go-loghub"

// Project define Project for test
var Project = &sls.LogProject{
	Name:            "loghub-test",
	Endpoint:        "cn-hangzhou.log.aliyuncs.com",
	AccessKeyID:     "xxx",
	AccessKeySecret: "xxx",
}
