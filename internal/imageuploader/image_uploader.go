package imagestorage

type ImageUploader interface {
	Upload(data []byte, name string) (string, error)
}
