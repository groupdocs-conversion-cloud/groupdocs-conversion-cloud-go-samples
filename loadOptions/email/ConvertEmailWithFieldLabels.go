package email

import (
	"fmt"

	"github.com/groupdocs-conversion-cloud/groupdocs-conversion-cloud-go-samples/config"
	"github.com/groupdocs-conversion-cloud/groupdocs-conversion-cloud-go/models"
)

func ConvertEmailWithFieldLabels() {

	settings := models.ConvertSettings{
		Format:     "pdf",
		FilePath:   "Email/sample.msg",
		OutputPath: "converted",
		LoadOptions: &models.EmailLoadOptions{
			Format: "msg",
			FieldLabels: []models.FieldLabel{
				{Field: models.FieldEnumFrom, Label: "Sender"},
				{Field: models.FieldEnumTo, Label: "Receiver"},
			},
		},
	}

	result, _, err := config.Client.ConvertApi.ConvertDocument(config.Ctx, settings)

	if err != nil {
		fmt.Printf("ConvertEmailWithFieldLabels error: %v\n", err)
		return
	}

	fmt.Printf("Document converted successfully: %v\n", result[0].Url)
}
