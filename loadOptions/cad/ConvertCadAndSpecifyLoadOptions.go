package cad

import (
	"fmt"

	"github.com/groupdocs-conversion-cloud/groupdocs-conversion-cloud-go-samples/config"
	"github.com/groupdocs-conversion-cloud/groupdocs-conversion-cloud-go/models"
)

func ConvertCadAndSpecifyLoadOptions() {

	settings := models.ConvertSettings{
		Format:     "pdf",
		FilePath:   "Cad/Sample.dwg",
		OutputPath: "converted",
		LoadOptions: &models.CadLoadOptions{
			Format:          "dwg",
			BackgroundColor: "Gray",
		},
	}

	result, _, err := config.Client.ConvertApi.ConvertDocument(config.Ctx, settings)

	if err != nil {
		fmt.Printf("ConvertCadAndSpecifyLoadOptions error: %v\n", err)
		return
	}

	fmt.Printf("Document converted successfully: %v\n", result[0].Url)
}
