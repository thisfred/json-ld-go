package jsonld

import (
	"encoding/json"
	. "gopkg.in/check.v1"
	"io/ioutil"
	"testing"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type JsonLdSuite struct{}

type JsonLdTest struct {
	Id      string   `json:"@id"`
	Type    []string `json:"@type"`
	Name    string
	Purpose string
	Input   string
	Context string
	Expect  string
}

type Manifest struct {
	Context     string `json:"@context"`
	Id          string `json:"@id"`
	Type        string `json:"@type"`
	Name        string
	Description string
	BaseIri     string
	Sequence    []JsonLdTest
}

var _ = Suite(&JsonLdSuite{})

func RunTest(c *C, test *JsonLdTest) {
	_, err := ioutil.ReadFile("./json-ld.org/test-suite/tests/" + test.Context)
	c.Assert(err, IsNil)
	_, err = ioutil.ReadFile("./json-ld.org/test-suite/tests/" + test.Input)
	c.Assert(err, IsNil)
	_, err = ioutil.ReadFile("./json-ld.org/test-suite/tests/" + test.Expect)
	c.Assert(err, IsNil)
}

func (s *JsonLdSuite) TestManifest(c *C) {
	file, err := ioutil.ReadFile("./json-ld.org/test-suite/tests/compact-manifest.jsonld")
	c.Assert(err, IsNil)
	var manifest Manifest
	err = json.Unmarshal(file, &manifest)
	c.Assert(err, IsNil)
	test := manifest.Sequence[0]
	RunTest(c, &test)
}
