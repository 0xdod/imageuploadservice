package imageuploader

import "context"

type ImageUploader interface {
	Upload(ctx context.Context, data []byte, name string) (string, error)
}
