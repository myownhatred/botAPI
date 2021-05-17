package repository

type Picture interface {
}

type Video interface {
}

type Translate interface {
}

type GoogleRepository struct {
	Picture
	Video
	Translate
}

func NewGoogleRepository() *GoogleRepository {
	return &GoogleRepository{}
}
