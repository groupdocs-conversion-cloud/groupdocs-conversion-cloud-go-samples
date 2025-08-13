package wordprocessing

import (
	"fmt"

	"github.com/groupdocs-conversion-cloud/groupdocs-conversion-cloud-go-samples/config"
	"github.com/groupdocs-conversion-cloud/groupdocs-conversion-cloud-go/models"
)

func ConvertWordProcessingByHidingComments() {

	settings := models.ConvertSettings{
		Format:     "pdf",
		FilePath:   "WordProcessing/with_tracked_changes.docx",
		OutputPath: "converted",
		LoadOptions: &models.WordProcessingLoadOptions{
			CommentDisplayMode: models.CommentDisplayModeEnumHidden,
		},
	}

	result, _, err := config.Client.ConvertApi.ConvertDocument(config.Ctx, settings)

	if err != nil {
		fmt.Printf("ConvertWordProcessingByHidingComments error: %v\n", err)
		return
	}

	fmt.Printf("Document converted successfully: %v\n", result[0].Url)
}
