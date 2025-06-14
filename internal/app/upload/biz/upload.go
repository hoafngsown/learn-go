package uploadbiz

import (
	"bytes"
	"context"
	"fmt"
	"image"

	// Import image format decoders - REQUIRED for PNG, JPEG, GIF support
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	uploadmodel "learn-go/v2/internal/app/upload/model"
	"learn-go/v2/internal/common"
	uploadprovider "learn-go/v2/internal/components/upload_provider"
	"path/filepath"
	"time"
)

type UploadStore interface {
	Upload(ctx context.Context, data []byte, folder, fileName string) (*common.Image, error)
}

var _ UploadStore = (*uploadBiz)(nil)

type uploadBiz struct {
	provider uploadprovider.UploadProvider
}

func NewUploadBiz(provider uploadprovider.UploadProvider) *uploadBiz {
	return &uploadBiz{provider: provider}
}

func (biz *uploadBiz) Upload(ctx context.Context, data []byte, folder, fileName string) (*common.Image, error) {
	fileBytes := bytes.NewBuffer(data)

	w, h, err := getImageDimension(fileBytes)

	if err != nil {
		return nil, common.ErrorInvalidRequest(uploadmodel.ErrorFileIsNotImage)
	}

	fileExt := filepath.Ext(fileName)                                // .jpg, .png, .jpeg
	fileName = fmt.Sprintf("%d%s", time.Now().Nanosecond(), fileExt) // 1718390400000000000.jpg

	img, err := biz.provider.SaveFileUploaded(ctx, data, fmt.Sprintf("%s/%s", folder, fileName))

	if err != nil {
		return nil, common.ErrorCannotSaveFile(err)
	}

	img.Width = w
	img.Height = h
	img.Extension = fileExt

	return &common.Image{
		Id:        img.Id,
		Url:       img.Url,
		Width:     w,
		Height:    h,
		Extension: fileExt,
		CloudName: "s3",
	}, nil
}

func getImageDimension(file *bytes.Buffer) (int, int, error) {
	bufferCopy := bytes.NewBuffer(file.Bytes())
	img, _, err := image.DecodeConfig(bufferCopy)

	if err != nil {
		return 0, 0, err
	}

	return img.Width, img.Height, nil
}
