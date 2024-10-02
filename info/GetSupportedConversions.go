package info

import (
	"fmt"

	"github.com/antihax/optional"
	conversion "github.com/groupdocs-conversion-cloud/groupdocs-conversion-cloud-go"
	"github.com/groupdocs-conversion-cloud/groupdocs-conversion-cloud-go-samples/config"
)

func GetSupportedConversions() {

	opts := &conversion.InfoApiGetSupportedConversionTypesOpts{
		Format: optional.NewString("docx"),
	}

	response, _, err := config.Client.InfoApi.GetSupportedConversionTypes(config.Ctx, opts)

	if err != nil {
		fmt.Printf("GetSupportedConversions error: %v\n", err)
		return
	}

	fmt.Printf("response.length: %v\n", len(response))
}
