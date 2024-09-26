package lkmsService

import (
	"context"
	"fmt"
	"mime/multipart"
	"strings"
	"time"

	"github.com/braciate/braciate-be/internal/api/lkms"
	"github.com/braciate/braciate-be/internal/entity"
	"github.com/braciate/braciate-be/internal/pkg/utils"
)

func (s *LkmsService) CreateLkm(ctx context.Context, req entity.Lkms, logo *multipart.FileHeader) (lkms.LkmsResponse, error) {
	logo.Filename = fmt.Sprintf("%s-%s", strings.Split(logo.Filename, ".")[0], strings.ReplaceAll(time.Now().Format(time.RFC3339), ":", ""))
	fmt.Println(logo.Filename)

	url, err := s.supabase.UploadFile(logo)
	if err != nil {
		s.log.Errorf("error uploading to supabase: %v", err)
		return lkms.LkmsResponse{}, err
	}

	lkmRepo, err := s.lkmsRepository.NewClient(false)
	if err != nil {
		s.log.Errorf("error creating lkm category: %v", err)
		return lkms.LkmsResponse{}, err
	}

	generateID, err := utils.GenerateRandomString(24)
	if err != nil {
		s.log.Errorf("error generating ID: %v", err)
		return lkms.LkmsResponse{}, err
	}

	req.ID = fmt.Sprintf("%d-%s", time.Now().UTC().UnixMilli(), generateID)
	req.LogoFile = url

	newLkm, err := lkmRepo.CreateLkms(ctx, req)
	if err != nil {
		s.log.Errorf("error at service layer: %v", err)
		return lkms.LkmsResponse{}, err
	}

	return lkms.LkmsResponse{
		ID:         newLkm.ID,
		Name:       newLkm.Name,
		CategoryID: newLkm.CategoryID,
		Type:       newLkm.Type,
		LogoLink:   newLkm.LogoFile,
	}, err
}
