package controllers

import (
	"net/http"

	"github.com/CGRDMZ/rmmbrit-api/errors"
	"github.com/CGRDMZ/rmmbrit-api/services"
	"github.com/gin-gonic/gin"
)

type ShortenerController struct {
	Ss *services.ShortenerService
}

func (sc *ShortenerController) RedirectToOriginalUrl(c *gin.Context) {

	id := c.Param("id")

	urlMap, err := sc.Ss.GetUrlMapInfoAndIncrementVisit(c.Request.Context(), id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if urlMap == nil {
		c.AbortWithError(http.StatusNotFound, errors.NotFoundErr("Url Map", string(id)))
		return
	}

	c.Redirect(http.StatusFound, urlMap.LongUrl)
}

func (sc *ShortenerController) GetUrlMapInfo(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	urlMap, err := sc.Ss.GetUrlMapInfo(c.Request.Context(), id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if urlMap == nil {
		c.AbortWithError(http.StatusNotFound, errors.NotFoundErr("Url Map", string(id)))
		return
	}

	c.HTML(http.StatusOK, "url-info.html", urlMap)

}

type CreateUrlMapRequest struct {
	ShortUrl string `json:"shortUrl" xml:"shortUrl" form:"shortUrl" binding:"-"`
	LongUrl  string `json:"longUrl" xml:"longUrl" form:"longUrl" binding:"required"`
}

func (sc *ShortenerController) AddNewUrlMap(c *gin.Context) {
	var rb CreateUrlMapRequest
	err := c.ShouldBind(&rb)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	urlMap, err := sc.Ss.CreateNewUrlMap(c.Request.Context(), rb.ShortUrl, rb.LongUrl)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, urlMap)
}
