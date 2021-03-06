package controllers

import (
	"net/http"
	"net/url"
	"github.com/CGRDMZ/rmmbrit-api/sherrors"
	"github.com/CGRDMZ/rmmbrit-api/services"
	"github.com/gin-gonic/gin"
)

type ShortenerController struct {
	Ss *services.ShortenerService
}

func (sc *ShortenerController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func (sc *ShortenerController) RedirectToOriginalUrl(c *gin.Context) {

	id := c.Param("id")

	// decode the short url
	id, err := url.PathUnescape(id)
	if err != nil {
		c.Error(err)
		return
	}

	urlMap, err := sc.Ss.GetUrlMapInfoAndIncrementVisit(c.Request.Context(), id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if urlMap == nil {
		c.AbortWithError(http.StatusNotFound, sherrors.NotFoundErr("Url Map", string(id)))
		return
	}

	c.Redirect(http.StatusFound, urlMap.LongUrl)
}

func (sc *ShortenerController) GetAllUrlMapInfo(c *gin.Context) {
	urlMaps, err := sc.Ss.GetAllUrlMaps(c.Request.Context())
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	// encode the url so that it doesn't contain any special characters that reserved for url
	for _, urlMap := range urlMaps {
		urlMap.ShortUrl = url.PathEscape(urlMap.ShortUrl)
	}

	c.HTML(http.StatusOK, "url-list.html", urlMaps)
}

func (sc *ShortenerController) GetUrlMapInfo(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// decode the short url
	id, err := url.PathUnescape(id)
	if err != nil {
		c.Error(err)
		return
	}

	urlMap, err := sc.Ss.GetUrlMapInfo(c.Request.Context(), id)
	if err != nil {
		c.Error(err)
		return
	}
	if urlMap == nil {
		c.Error(sherrors.NotFoundErr("Url Map", string(id)))
		return
	}

	// encode the url so that it doesn't contain any special characters that reserved for url
	urlMap.ShortUrl = url.PathEscape(urlMap.ShortUrl)

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
		c.Error(err)
		return
	}

	// encode short url so that it doesn't contain any special characters that reserved for url
	urlMap.ShortUrl = url.PathEscape(urlMap.ShortUrl)

	c.Redirect(http.StatusFound, "/info/"+urlMap.ShortUrl)
}

