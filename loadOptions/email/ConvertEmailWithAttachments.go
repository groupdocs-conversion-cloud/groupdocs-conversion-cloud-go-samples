package email

import (
	"fmt"

	"github.com/groupdocs-conversion-cloud/groupdocs-conversion-cloud-go-samples/config"
	"github.com/groupdocs-conversion-cloud/groupdocs-conversion-cloud-go/models"
)

func ConvertEmailWithAttachments() {

	settings := models.ConvertSettings{
		Format:     "pdf",
		FilePath:   "Email/embedded-image-and-attachment.eml",
		OutputPath: "converted",
		LoadOptions: &models.EmailLoadOptions{
			Format:             "eml",
			ConvertAttachments: true,
		},
	}

	result, _, err := config.Client.ConvertApi.ConvertDocument(config.Ctx, settings)

	if err != nil {
		fmt.Printf("ConvertEmailWithAttachments error: %v\n", err)
		return
	}

	fmt.Printf("Document converted successfully: %v\n", result[0].Url)
}
