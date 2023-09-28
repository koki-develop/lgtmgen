package repo

type imageRepository struct{}

func newImageRepository() *imageRepository {
	return &imageRepository{}
}
