package common

import (
	"fmt"

	"github.com/groupdocs-conversion-cloud/groupdocs-conversion-cloud-go-samples/config"
	"github.com/groupdocs-conversion-cloud/groupdocs-conversion-cloud-go/models"
)

func ConvertUsingCustomFont() {

	settings := models.ConvertSettings{
		Format:     "pdf",
		FilePath:   "Presentation/uses-custom-font.pptx",
		OutputPath: "converted",
		FontsPath:  "font/ttf",
	}

	result, _, err := config.Client.ConvertApi.ConvertDocument(config.Ctx, settings)

	if err != nil {
		fmt.Printf("ConvertUsingCustomFont error: %v\n", err)
		return
	}

	fmt.Printf("Document converted successfully: %v\n", result[0].Url)
}
