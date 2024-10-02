package csv

import (
	"fmt"

	"github.com/groupdocs-conversion-cloud/groupdocs-conversion-cloud-go-samples/config"
	"github.com/groupdocs-conversion-cloud/groupdocs-conversion-cloud-go/models"
)

func ConvertCsvBySpecifyingDelimiter() {

	settings := models.ConvertSettings{
		Format:     "pdf",
		FilePath:   "Spreadsheet/sample.csv",
		OutputPath: "converted",
		LoadOptions: &models.CsvLoadOptions{
			Format:    "csv",
			Separator: ",",
		},
	}

	result, _, err := config.Client.ConvertApi.ConvertDocument(config.Ctx, settings)

	if err != nil {
		fmt.Printf("ConvertCsvBySpecifyingDelimiter error: %v\n", err)
		return
	}

	fmt.Printf("Document converted successfully: %v\n", result[0].Url)
}
