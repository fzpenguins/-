package dao

import (
	"context"
	"gorm.io/gorm"
	"grpc/dal/db"
)

type ImageDao struct {
	*gorm.DB
}

func NewDBClient(ctx context.Context) *gorm.DB {
	Db := db.DB
	return Db.WithContext(ctx)
}

func NewImageDao(ctx context.Context) *ImageDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &ImageDao{NewDBClient(ctx)}
}

func (dao *ImageDao) CreateImage(image *db.Image) (*db.Image, error) {
	err := dao.Model(&db.Image{}).Create(&image).Error
	return image, err
}

func (dao *ImageDao) GetURLByPid(pid int64) (*db.Image, error) {
	var ret *db.Image
	err := dao.Model(&db.Image{}).Where("pid = ?", pid).First(&ret).Error
	return ret, err
}
