package info

import (
	"fmt"

	"github.com/antihax/optional"
	conversion "github.com/groupdocs-conversion-cloud/groupdocs-conversion-cloud-go"
	"github.com/groupdocs-conversion-cloud/groupdocs-conversion-cloud-go-samples/config"
)

func GetDocumentInformation() {

	opts := &conversion.InfoApiGetDocumentMetadataOpts{
		FilePath: optional.NewString("WordProcessing/four-pages.docx"),
	}

	info, _, err := config.Client.InfoApi.GetDocumentMetadata(config.Ctx, opts)

	if err != nil {
		fmt.Printf("GetDocumentInformation error: %v\n", err)
		return
	}

	fmt.Printf("InfoResult.Pages.Count: %v\n", info.PageCount)
}
