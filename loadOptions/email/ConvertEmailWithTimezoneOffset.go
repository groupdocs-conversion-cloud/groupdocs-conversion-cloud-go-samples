package email

import (
	"fmt"
	"time"

	"github.com/groupdocs-conversion-cloud/groupdocs-conversion-cloud-go-samples/config"
	"github.com/groupdocs-conversion-cloud/groupdocs-conversion-cloud-go/models"
)

func ConvertEmailWithTimezoneOffset() {

	settings := models.ConvertSettings{
		Format:     "pdf",
		FilePath:   "Email/sample.msg",
		OutputPath: "converted",
		LoadOptions: &models.EmailLoadOptions{
			Format:         "msg",
			TimeZoneOffset: (time.Duration(5) * time.Hour).String(),
		},
	}

	result, _, err := config.Client.ConvertApi.ConvertDocument(config.Ctx, settings)

	if err != nil {
		fmt.Printf("ConvertEmailWithTimezoneOffset error: %v\n", err)
		return
	}

	fmt.Printf("Document converted successfully: %v\n", result[0].Url)
}
