package router

import (
	"log"
	"net/http"
	"os"
	"github.com/gin-gonic/gin"
)

// url path, handler function that will hit that path
type Route struct {
	Name string
	Method string
	Pattern string
	HandlerFunc func(*gin.Context)
}

// declare gin engine
type routes struct {
	router *gin.Engine
}

// variable declare for various routes
type Routes []Route

//serving clientroutes func
func(r routes) UrlShortener(rg *gin.RouterGroup) {
	//groups urls and '/url' acts as parent route
	orderRouterGrouping := rg.Group("/url")
	// grp of urls will initialise for loop and whatever method mathces will be handled accordinly
	for _, route := range urlShortener {
		switch route.Method {
		case http.MethodGet:
			orderRouterGrouping.GET(route.Pattern, route.HandlerFunc)
		case http.MethodPost:
			orderRouterGrouping.POST(route.Pattern, route.HandlerFunc)
		case http.MethodOptions:
			orderRouterGrouping.OPTIONS(route.Pattern, route.HandlerFunc)
		case http.MethodPut:
			orderRouterGrouping.PUT(route.Pattern, route.HandlerFunc)
		case http.MethodDelete:
			orderRouterGrouping.DELETE(route.Pattern, route.HandlerFunc)
		default:
			orderRouterGrouping.GET(route.Pattern, func(c *gin.Context){
				c.JSON(200, gin.H{
					"reuslt": "Specify a valid method for this route",
				})
			})
		}
	}
}

func ClientRoutes(){
	//initialise gin engine
	r := routes{
		router: gin.Default(),
	}
	// grouping of api and send to routes
	api := r.router.Group(os.Getenv("API_VERSION"))
	r.UrlShortener(api)

	//error
	if err := r.router.Run(":" + os.Getenv("PORT")); err != nil {
		log.Println("failed to run server: %V", err)
	}
}