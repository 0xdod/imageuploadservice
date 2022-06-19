package imageuploader

import "context"

type ImageUploader interface {
	Upload(ctx context.Context, name string, data []byte) (string, error)
}
