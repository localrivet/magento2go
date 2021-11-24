package magento2go

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"strings"

	"github.com/localrivet/magento2go/client"
	"github.com/localrivet/magento2go/utils"

	httptransport "github.com/go-openapi/runtime/client"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

func NewCommunityClient(config Config) (*client.MagentoCommunity, error) {
	if config == (Config{}) {
		return nil, errors.New("config cannot be empty")
	}
	if config.AccessToken == "" {
		return nil, errors.New("accessToken cannot be empty")
	}
	if config.Host == "" {
		return nil, errors.New("host cannot be empty")
	}
	if config.Path == "" {
		return nil, errors.New("path cannot be empty")
	}
	if config.Scheme == "" {
		return nil, errors.New("scheme cannot be empty")
	}

	// // create the API client
	transport := httptransport.New(config.Host, config.Path, []string{config.Scheme})

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
		return req.SetHeaderParam("Authorization", fmt.Sprintf("Bearer %s", config.AccessToken))
	})

	// set debug status
	transport.Debug = config.Debug

	// create a new client
	mc := client.New(transport, strfmt.Default)
	var emptyClient *client.MagentoCommunity
	if mc == emptyClient {
		return nil, errors.New("client not created")
	}

	return mc, nil
}
