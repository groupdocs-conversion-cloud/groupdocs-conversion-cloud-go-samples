package txt

import (
	"fmt"

	"github.com/groupdocs-conversion-cloud/groupdocs-conversion-cloud-go-samples/config"
	"github.com/groupdocs-conversion-cloud/groupdocs-conversion-cloud-go/models"
)

func ConvertTxtByControllingLeadingSpacesBehavior() {

	settings := models.ConvertSettings{
		Format:     "pdf",
		FilePath:   "Text/sample.txt",
		OutputPath: "converted",
		LoadOptions: &models.TxtLoadOptions{
			LeadingSpacesOptions:           models.LeadingSpacesOptionsEnumConvertToIndent,
			DetectNumberingWithWhitespaces: true,
		},
	}

	result, _, err := config.Client.ConvertApi.ConvertDocument(config.Ctx, settings)

	if err != nil {
		fmt.Printf("ConvertTxtByControllingLeadingSpacesBehavior error: %v\n", err)
		return
	}

	fmt.Printf("Document converted successfully: %v\n", result[0].Url)
}
