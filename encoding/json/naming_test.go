package json

import (
	"testing"
)

type conversion struct {
	original  string
	converted string
}

var examplePythonConversions = []conversion{
	{"EndpointURL", "endpoint_url"},
	{"NewUUIDValue", "new_uuid_value"},
	{"APlusB", "a_plus_b"},
	{"TBSCertList", "tbs_cert_list"},
	{"H264Version", "h264_version"},
}

func Test_PythonConversions(t *testing.T) {
	for _, c := range examplePythonConversions {
		if changeFieldNamingStyle(c.original, PythonNamingStyle) != c.converted {
			t.Error(c.original, " did not convert to ", c.converted)
		}
	}
}

var exampleJavascriptConversions = []conversion{
	{"EndpointURL", "endpointURL"},
	{"NewUUIDValue", "newUUIDValue"},
	{"APlusB", "aPlusB"},
	{"TBSCertList", "tbsCertList"},
	{"H264Version", "h264Version"},
}

func Test_JavascriptConversions(t *testing.T) {
	for _, c := range exampleJavascriptConversions {
		if changeFieldNamingStyle(c.original, JavascriptNamingStyle) != c.converted {
			t.Error(c.original, " did not convert to ", c.converted)
		}
	}
}

// var exampleCSharpConversions = []conversion{
// 	{"EndpointURL", "endpointUrl"},
// 	{"NewUUIDValue", "newUuidValue"},
// 	{"APlusB", "aPlusB"},
// 	{"TBSCertList", "tbsCertList"},
// 	{"H264Version", "h264Version"},
// }

// func Test_CSharpConversions(t *testing.T) {
// 	for _, c := range exampleCSharpConversions {
// 		if changeFieldNamingStyle(c.original, CSharpNamingStyle) != c.converted {
// 			t.Error(c.original, " did not convert to ", c.converted)
// 		}
// 	}
// }
