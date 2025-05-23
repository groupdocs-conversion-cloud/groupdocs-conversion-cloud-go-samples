package presentation

import (
	"fmt"

	"github.com/groupdocs-conversion-cloud/groupdocs-conversion-cloud-go-samples/config"
	"github.com/groupdocs-conversion-cloud/groupdocs-conversion-cloud-go/models"
)

func ConvertPresentationByHidingComments() {

	settings := models.ConvertSettings{
		Format:     "pdf",
		FilePath:   "Presentation/with_notes.pptx",
		OutputPath: "converted",
		LoadOptions: &models.PresentationLoadOptions{
			Format:           "pptx",
			CommentsPosition: models.CommentsPositionEnumNone,
		},
	}

	result, _, err := config.Client.ConvertApi.ConvertDocument(config.Ctx, settings)

	if err != nil {
		fmt.Printf("ConvertPresentationByHidingComments error: %v\n", err)
		return
	}

	fmt.Printf("Document converted successfully: %v\n", result[0].Url)
}
