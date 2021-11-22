package magento2go

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"magento2go/client"
	"magento2go/utils"
	"strings"

	httptransport "github.com/go-openapi/runtime/client"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

func NewClient(config *struct {
	accessToken string
	host        string
	path        string
	scheme      string
	debug       bool
}) (*client.MagentoCommunity, error) {
	if config == nil {
		return nil, errors.New("config cannot be empty")
	}

	if config.accessToken == "" {
		return nil, errors.New("accessToken cannot be empty")
	}
	if config.host == "" {
		return nil, errors.New("host cannot be empty")
	}
	if config.path == "" {
		return nil, errors.New("path cannot be empty")
	}
	if config.scheme == "" {
		return nil, errors.New("scheme cannot be empty")
	}

	// // create the API client
	transport := httptransport.New(config.host, config.path, []string{config.scheme})

	// intercept the application/json consumer and fix the overloaded "value":string|[]string issue
	transport.Consumers[runtime.JSONMime] = runtime.ConsumerFunc(func(reader io.Reader, v interface{}) error {
		buf, _ := ioutil.ReadAll(reader)

		// we need to fix the json response before marshalling it
		// into a struct of the expected type
		output := utils.FlattenValueArrays(string(buf)) // flatten all "value":[]string into "value":string
		output = utils.FixErrorParameters(output)       // fix error parameters

		newReader := strings.NewReader(output)
		dec := json.NewDecoder(newReader)
		dec.UseNumber()
		err := dec.Decode(v)

		return err
	})

	// set the bearer authorization header
	transport.DefaultAuthentication = runtime.ClientAuthInfoWriterFunc(func(req runtime.ClientRequest, _ strfmt.Registry) error {
		return req.SetHeaderParam("Authorization", fmt.Sprintf("Bearer %s", config.accessToken))
	})

	// set debug status
	transport.Debug = config.debug

	// create a new client
	c := client.New(transport, strfmt.Default)
	var emptyClient *client.MagentoCommunity
	if c == emptyClient {
		return nil, errors.New("client not created")
	}

	return c, nil
}
