package common

import (
	"fmt"

	"github.com/groupdocs-conversion-cloud/groupdocs-conversion-cloud-go-samples/config"
	"github.com/groupdocs-conversion-cloud/groupdocs-conversion-cloud-go/models"
)

func ConvertSpecificPages() {

	settings := models.ConvertSettings{
		Format:     "pdf",
		FilePath:   "WordProcessing/four-pages.docx",
		OutputPath: "converted",
		ConvertOptions: &models.PdfConvertOptions{
			Pages: []int32{1, 3},
		},
	}

	result, _, err := config.Client.ConvertApi.ConvertDocument(config.Ctx, settings)

	if err != nil {
		fmt.Printf("ConvertSpecificPages error: %v\n", err)
		return
	}

	fmt.Printf("Document converted successfully: %v\n", result[0].Url)
}
