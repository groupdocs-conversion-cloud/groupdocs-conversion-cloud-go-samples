package presentation

import (
	"fmt"

	"github.com/groupdocs-conversion-cloud/groupdocs-conversion-cloud-go-samples/config"
	"github.com/groupdocs-conversion-cloud/groupdocs-conversion-cloud-go/models"
)

func ConvertPresentationWithHiddenSlidesIncluded() {

	settings := models.ConvertSettings{
		Format:     "pdf",
		FilePath:   "Presentation/with_hidden_page.pptx",
		OutputPath: "converted",
		LoadOptions: &models.PresentationLoadOptions{
			Format:           "pptx",
			ShowHiddenSlides: true,
		},
	}

	result, _, err := config.Client.ConvertApi.ConvertDocument(config.Ctx, settings)

	if err != nil {
		fmt.Printf("ConvertPresentationWithHiddenSlidesIncluded error: %v\n", err)
		return
	}

	fmt.Printf("Document converted successfully: %v\n", result[0].Url)
}
