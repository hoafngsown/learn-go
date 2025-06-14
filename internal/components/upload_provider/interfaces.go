package uploadprovider

import (
	"context"
	"learn-go/v2/internal/common"
)

type UploadProvider interface {
	SaveFileUploaded(ctx context.Context, data []byte, dst string) (*common.Image, error)
}
