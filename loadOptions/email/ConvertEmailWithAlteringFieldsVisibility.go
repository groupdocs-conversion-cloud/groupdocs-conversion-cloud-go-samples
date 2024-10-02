package email

import (
	"fmt"

	"github.com/groupdocs-conversion-cloud/groupdocs-conversion-cloud-go-samples/config"
	"github.com/groupdocs-conversion-cloud/groupdocs-conversion-cloud-go/models"
)

func ConvertEmailWithAlteringFieldsVisibility() {

	settings := models.ConvertSettings{
		Format:     "pdf",
		FilePath:   "Email/sample.msg",
		OutputPath: "converted",
		LoadOptions: &models.EmailLoadOptions{
			Format:                  "msg",
			DisplayHeader:           false,
			DisplayFromEmailAddress: false,
			DisplayToEmailAddress:   false,
			DisplayCcEmailAddress:   false,
			DisplayBccEmailAddress:  false,
		},
	}

	result, _, err := config.Client.ConvertApi.ConvertDocument(config.Ctx, settings)

	if err != nil {
		fmt.Printf("ConvertEmailWithAlteringFieldsVisibility error: %v\n", err)
		return
	}

	fmt.Printf("Document converted successfully: %v\n", result[0].Url)
}
