package lkmsService

import (
	"context"
	"mime/multipart"

	"github.com/braciate/braciate-be/internal/api/lkms"
	lkmsRepository "github.com/braciate/braciate-be/internal/api/lkms/repository"
	"github.com/braciate/braciate-be/internal/entity"
	"github.com/braciate/braciate-be/internal/pkg/supabase"
	"github.com/sirupsen/logrus"
)

type LkmsService struct {
	lkmsRepository lkmsRepository.RepositoryItf
	log            *logrus.Logger
	supabase       supabase.SupabaseInterface
}

type LkmsServiceItf interface {
	CreateLkm(ctx context.Context, req entity.Lkms, logo *multipart.FileHeader) (lkms.LkmsResponse, error)
	GetLkmsByCategoryIDAndType(ctx context.Context, id string, lkmType string) ([]lkms.LkmsResponse, error)
}

func New(log *logrus.Logger, repo lkmsRepository.RepositoryItf, supabase supabase.SupabaseInterface) LkmsServiceItf {
	return &LkmsService{
		lkmsRepository: repo,
		log:            log,
		supabase:       supabase,
	}
}
