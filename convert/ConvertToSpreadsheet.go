package convert

import (
	"fmt"

	"github.com/groupdocs-conversion-cloud/groupdocs-conversion-cloud-go-samples/config"
	"github.com/groupdocs-conversion-cloud/groupdocs-conversion-cloud-go/models"
)

func ConvertToSpreadsheet() {

	settings := models.ConvertSettings{
		Format:     "xlsx",
		FilePath:   "WordProcessing/four-pages.docx",
		OutputPath: "converted",
		ConvertOptions: &models.SpreadsheetConvertOptions{
			FromPage:   2,
			PagesCount: 1,
			Zoom:       150,
		},
	}

	result, _, err := config.Client.ConvertApi.ConvertDocument(config.Ctx, settings)

	if err != nil {
		fmt.Printf("ConvertToSpreadsheet error: %v\n", err)
		return
	}

	fmt.Printf("Document converted successfully: %v\n", result[0].Url)
}
