package service

import (
	"grpc/biz/rpc"
	"grpc/dal/db"
	"grpc/dal/db/dao"
	"grpc/dal/milvus"
	"grpc/dal/minio"
	"grpc/proto"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/pkg/errors"
)

func (s *PictureService) Insert(url string) (image *db.Image, err error) {
	var imageByte []byte
	//现将数据上传到minio
	finalStr, err := minio.UploadImage(url)
	if err != nil {
		return nil, errors.WithMessage(err, "service.Insert upload data error")
	}

	//保存到mysql
	//确保上传成功后再写url保存到mysql，保持数据一致

	imageDao := dao.NewImageDao(s.ctx)
	image = &db.Image{
		Url: finalStr,
	}
	image, err = imageDao.CreateImage(image)
	if err != nil {
		return nil, errors.WithMessage(err, "service.Insert create image error")
	}

	if url[0:4] == "/hom" {
		resp, err := os.Open(url)
		if err != nil {
			return nil, errors.WithMessage(err, "service.Insert create image error")
		}
		imageByte, err = ioutil.ReadAll(resp)
		if err != nil {
			return nil, errors.WithMessage(err, "service.Insert create image error")
		}
		defer resp.Close()
	} else {
		resp, err := http.Get(url)
		if err != nil {
			return nil, errors.WithMessage(err, "service.Insert create image error")
		}
		defer resp.Body.Close()

		imageByte, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, errors.WithMessage(err, "service.Insert create image error")
		}
	}

	//rpc调用获取向量数据
	vector, err := rpc.GetImageVector(s.ctx, &proto.ImageRequest{Image: imageByte})
	if err != nil {
		return nil, errors.WithMessage(err, "service.Insert get vector error")
	}

	err = milvus.InsertVector(s.ctx, vector, image.Pid)
	if err != nil {
		return nil, errors.WithMessage(err, "service.Insert insert vector error")
	}

	return image, nil
}
