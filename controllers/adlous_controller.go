package controllers

import (
	"net/http"

	"example.com/go-aldous-backend/external/bnc"
	"github.com/gin-gonic/gin"
)

type AldousController struct {
}

type GetAssetsResponseBody struct {
	Assets []BncAsset `json:"assets"`
}
type BncAsset struct {
	Name      string `json:"name"`
	AssetUUID string `json:"assetUuid"`
}

func (cc AldousController) GetAssets(c *gin.Context) {
	bncResponse, err := bnc.GetAssets()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	assets := []BncAsset{}
	for _, content := range bncResponse.Content {
		asset := BncAsset{}
		asset.AssetUUID = content.ID
		asset.Name = content.Name
		assets = append(assets, asset)
	}
	res := &GetAssetsResponseBody{
		Assets: assets,
	}
	c.JSON(http.StatusOK, res)
}

type PostAssetsRequestBody struct {
}

// TODO: Spec out and implement POST action
func (cc AldousController) PostAssets(c *gin.Context) {
	var reqBody PostAssetsRequestBody
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, reqBody)
}

func AddAldousControllerRoutes(router *gin.Engine) {
	AldousController := AldousController{}
	router.GET("/assets", AldousController.GetAssets)
	router.POST("/bnc/asset", AldousController.PostAssets)
}
