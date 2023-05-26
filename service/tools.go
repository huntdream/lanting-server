package service

import (
	"github.com/huntdream/lanting-server/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/huntdream/lanting-server/app"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

// GetUploadToken Get qiniu upload token
func GetUploadToken(c *gin.Context) {
	bucket := app.Config.Storage.Bucket

	secretKey := app.Config.Storage.SecretKey
	accessKey := app.Config.Storage.AccessKey

	putPolicy := storage.PutPolicy{
		Scope:      bucket,
		ReturnBody: `{"key":"$(key)","hash":"$(etag)","size":$(fsize),"name":"$(fname)","type": $(mimeType),"width": $(imageInfo.width),"height": $(imageInfo.height)}`,
	}

	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)

	c.JSON(http.StatusOK, gin.H{
		"token": upToken,
	})
}

// GetMediaFromArticle get media
func GetMediaFromArticle(nodes []model.ArticleNode, media []model.ArticleMedia) []model.ArticleMedia {
	for _, child := range nodes {
		if len(child.Children) != 0 {
			media = GetMediaFromArticle(child.Children, media)
		}

		if child.Type == "audio" || child.Type == "video" || child.Type == "image" {
			media = append(media, model.ArticleMedia{
				Src:  child.Src,
				Type: child.Type,
			})
		}
	}

	return media
}
