package router

import (
	"supersign/internal/api"
	v1 "supersign/internal/api/v1"
	"supersign/internal/model"
	"supersign/pkg/log"

	"github.com/gin-gonic/gin"
)

func registerDownloadRouter(store *model.Store, r *gin.RouterGroup) {
	download := v1.Download{
		Base: api.New(store, log.New("Download").L()),
	}
	downloadR := r.Group("/download")

	{
		downloadR.GET("mobileConfig/:uuid", download.MobileConfig)
		downloadR.GET("ipa/:uuid", download.IPA)
		downloadR.GET("icon/:uuid", download.Icon)
	}
}
