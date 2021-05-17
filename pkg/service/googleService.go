package service

import "github.com/myownhatred/botAPI/pkg/repository"

type Picture interface {
}

type Video interface {
}

type Translate interface {
}

type GoogleService struct {
	Picture
	Video
	Translate
}

func NewGoogleService(rep *repository.GoogleRepository) *GoogleService {
	return &GoogleService{}
}
