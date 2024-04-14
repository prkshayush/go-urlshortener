package router

import (
	"net/http"
	"urlshortener/constant"
	"urlshortener/controller"
)

// variable and post method
var urlShortener = Routes{
	Route{"URL Shortening service", http.MethodPost, constant.UrlShortenerPath, controller.ShortenUrl},
	// after connnecting with db to send long url
	// declare for constant and controller
	Route{"Redirecting to original URL", http.MethodGet, constant.RedirectUrlPath, controller.RedirectUrl},
}