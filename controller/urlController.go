package controller

import (
	"fmt"
	"net/http"
	"time"
	"urlshortener/constant"
	"urlshortener/database"
	"urlshortener/helper"
	"urlshortener/types"

	"github.com/gin-gonic/gin"
)

func ShortenUrl(c *gin.Context) {
	var shortUrlBody types.ShortUrlBody
	// after dbcalls
	//it will bind the shorturlbody to json
    err :=	c.BindJSON(shortUrlBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": true, "message": constant.BindError})
		return 
	}
	// now we need a code generator for random code generator so we will make helper functions
	// inside genrandomstring we specify the length of the stirng

	code := helper.GenRandomString(6)

	// query to check if same url doesnot exist in connection.go
	// after dbcalls 
	record, _ := database.Mgr.GetUrlFromCode(code, constant.UrlCollection)
	//if error then datbase is not up and running
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": err.Error()})
	// 	return
	// }
	//checking only if record exist then only error
	if record.UrlCode != "" {
		//means record is present
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": "Code is already in use"})
	}
	var url types.UrlDB
	url.CreatedAt = time.Now().Unix()
	url.ExpireAt = time.Now().Unix()
	url.UrlCode = code
	url.LongUrl = shortUrlBody.LonguRL
	// we need to have a base url for shorturl
	// see const
	url.ShortUrl = constant.BaseUrl + code

	// now we will inseert
	resp, err := database.Mgr.Insert(url, constant.UrlCollection)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": true, "message": "Some error see urlcontroller"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"error": false, "data": resp, "shortURL": url.ShortUrl})
}

func RedirectUrl (c *gin.Context){
	code := c.Param("code")

	//check for if the code exists or not
	record, _ := database.Mgr.GetUrlFromCode(code, constant.UrlCollection)
	if record.UrlCode == "" {
		//means record is present
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": "Provide a valid url"})
	}
	fmt.Println(record.LongUrl)

	//redirect to original url
	c.Redirect(http.StatusPermanentRedirect, record.LongUrl)
}