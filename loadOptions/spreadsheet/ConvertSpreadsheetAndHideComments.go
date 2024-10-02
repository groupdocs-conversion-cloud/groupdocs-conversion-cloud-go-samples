package spreadsheet

import (
	"fmt"

	"github.com/groupdocs-conversion-cloud/groupdocs-conversion-cloud-go-samples/config"
	"github.com/groupdocs-conversion-cloud/groupdocs-conversion-cloud-go/models"
)

func ConvertSpreadsheetAndHideComments() {

	settings := models.ConvertSettings{
		Format:     "pdf",
		FilePath:   "Spreadsheet/sample.xlsx",
		OutputPath: "converted",
		LoadOptions: &models.SpreadsheetLoadOptions{
			HideComments:    true,
			OnePagePerSheet: true,
		},
	}

	result, _, err := config.Client.ConvertApi.ConvertDocument(config.Ctx, settings)

	if err != nil {
		fmt.Printf("ConvertSpreadsheetAndHideComments error: %v\n", err)
		return
	}

	fmt.Printf("Document converted successfully: %v\n", result[0].Url)
}
