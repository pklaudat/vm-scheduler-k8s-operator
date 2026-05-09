package azure

import (
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
)

type Client struct {
	credential *azidentity.DefaultAzureCredential
}

func NewClient() (*Client, error) {

	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to create azure credential: %w",
			err,
		)
	}

	return &Client{
		credential: cred,
	}, nil
}