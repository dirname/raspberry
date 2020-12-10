package apis

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

type bingParse struct {
	Images []image `json:"images"`
}

type image struct {
	BasePath  string `json:"urlbase"`
	Copyright string `json:"copyright"`
}

func WallpaperIndex(c *gin.Context) {
	resolution := c.Param("resolution")
	url := "https://www.bing.com/HPImageArchive.aspx?format=js&idx=0&n=1&mkt=en-US"
	response, err := http.Get(url)
	if err != nil {
		logrus.Errorf("%s\n", err.Error())
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logrus.Errorf("%s\n", err.Error())
	}
	p := &bingParse{}
	err = json.Unmarshal(body, p)
	if err != nil {
		logrus.Errorf("Failed to parse json: %s\n", err.Error())
	}
	c.JSONP(200, gin.H{"url": fmt.Sprintf("https://bing.com%s_%s.jpg", p.Images[0].BasePath, resolution), "info": p.Images[0].Copyright})
}
