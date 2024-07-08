package service

import (
	"grpc/biz/rpc"
	"grpc/dal/db"
	"grpc/dal/db/dao"
	"grpc/dal/milvus"
	"grpc/proto"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/pkg/errors"
)

func (s *PictureService) SearchByImage(url string) (images []*db.Image, err error) {
	var ids []int64
	var data []byte

	if url[0:4] == "/hom" {
		resp, err := os.Open(url)
		if err != nil {
			return nil, errors.WithMessage(err, "service.Insert read image error")
		}
		data, err = ioutil.ReadAll(resp)
		if err != nil {
			return nil, errors.WithMessage(err, "service.Insert read image error")
		}
		defer resp.Close()
	} else {
		resp, err := http.Get(url)
		if err != nil {
			return nil, errors.WithMessage(err, "GetUrlByID failed")
		}
		data, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, errors.WithMessage(err, "GetUrlByID failed")
		}
	}

	vector, err := rpc.GetImageVector(s.ctx, &proto.ImageRequest{Image: data})
	if err != nil {
		return nil, errors.WithMessage(err, "GetTextVector failed")
	}
	ids, err = milvus.Search(s.ctx, vector)
	if err != nil {
		return nil, errors.WithMessage(err, "Search failed")
	}
	imageDao := dao.NewImageDao(s.ctx)
	images = make([]*db.Image, len(ids))
	for index, id := range ids {
		pic, err := imageDao.GetURLByPid(id)
		if err != nil {
			return nil, errors.WithMessage(err, "GetUrlByID failed")
		}
		temp := &db.Image{
			Pid: pic.Pid,
			Url: pic.Url,
		}
		images[index] = temp
	}
	return images, nil
}
