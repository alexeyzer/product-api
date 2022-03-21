package service

import (
	"bytes"
	"context"
	"github.com/alexeyzer/product-api/internal/client"
	"github.com/alexeyzer/product-api/internal/pkg/datastruct"
	"github.com/alexeyzer/product-api/internal/pkg/repository"
	"github.com/google/uuid"
)

type MediaService interface {
	CreateMedia(ctx context.Context, productID int64, image []byte, contentTypeS3 string, contentType datastruct.ContentType) (*datastruct.Media, error)
	GetMedia(ctx context.Context, ID int64) (*datastruct.Media, error)
	DeleteMedia(ctx context.Context, ID int64) error
	ListMedias(ctx context.Context, productID int64) ([]*datastruct.Media, error)
}

type mediaService struct {
	dao repository.DAO
	s3  client.S3Client
}

func (s *mediaService) DeleteMedia(ctx context.Context, ID int64) error {
	err := s.dao.MediaQuery().Delete(ctx, ID)
	if err != nil {
		return err
	}
	return nil
}

func (s *mediaService) GetMedia(ctx context.Context, ID int64) (*datastruct.Media, error) {
	media, err := s.dao.MediaQuery().Get(ctx, ID)
	if err != nil {
		return nil, err
	}
	return media, nil
}

func (s *mediaService) ListMedias(ctx context.Context, productID int64) ([]*datastruct.Media, error) {
	medias, err := s.dao.MediaQuery().List(ctx, productID)
	if err != nil {
		return nil, err
	}

	return medias, nil
}

func (s *mediaService) CreateMedia(ctx context.Context, productID int64, image []byte, contentTypeS3 string, contentType datastruct.ContentType) (*datastruct.Media, error) {

	req := datastruct.Media{
		ContentType: contentType,
		ProductID:   productID,
	}
	if len(image) > 0 && contentType != "" {
		res, err := s.s3.UploadFileD(ctx, uuid.New().String(), bytes.NewReader(image), contentTypeS3)
		if err != nil {
			return nil, err
		}
		req.Url = res.Location
	}

	res, err := s.dao.MediaQuery().Create(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func NewMediaService(dao repository.DAO, s3 client.S3Client) MediaService {
	return &mediaService{dao: dao, s3: s3}
}
