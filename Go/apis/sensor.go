package apis

import "github.com/gin-gonic/gin"

func IndexAPI(c *gin.Context) {
	c.JSONP(200, gin.H{"test": "ok"})
}
