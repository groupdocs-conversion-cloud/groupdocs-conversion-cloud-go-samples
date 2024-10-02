package convert

import (
	"fmt"

	"github.com/groupdocs-conversion-cloud/groupdocs-conversion-cloud-go-samples/config"
	"github.com/groupdocs-conversion-cloud/groupdocs-conversion-cloud-go/models"
)

func ConvertToHtml() {

	settings := models.ConvertSettings{
		Format:     "html",
		FilePath:   "WordProcessing/four-pages.docx",
		OutputPath: "converted",
		ConvertOptions: &models.WebConvertOptions{
			FromPage:               1,
			PagesCount:             1,
			FixedLayout:            true,
			FixedLayoutShowBorders: true,
		},
	}

	result, _, err := config.Client.ConvertApi.ConvertDocument(config.Ctx, settings)

	if err != nil {
		fmt.Printf("ConvertToHtml error: %v\n", err)
		return
	}

	fmt.Printf("Document converted successfully: %v\n", result[0].Url)
}
