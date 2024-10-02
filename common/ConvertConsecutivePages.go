package common

import (
	"fmt"

	"github.com/groupdocs-conversion-cloud/groupdocs-conversion-cloud-go-samples/config"
	"github.com/groupdocs-conversion-cloud/groupdocs-conversion-cloud-go/models"
)

func ConvertConsecutivePages() {

	settings := models.ConvertSettings{
		Format:     "pdf",
		FilePath:   "WordProcessing/four-pages.docx",
		OutputPath: "converted",
		ConvertOptions: &models.PdfConvertOptions{
			FromPage:   2,
			PagesCount: 2,
		},
	}

	result, _, err := config.Client.ConvertApi.ConvertDocument(config.Ctx, settings)

	if err != nil {
		fmt.Printf("ConvertConsecutivePages error: %v\n", err)
		return
	}

	fmt.Printf("Document converted successfully: %v\n", result[0].Url)
}
