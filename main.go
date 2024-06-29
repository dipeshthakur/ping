// made by Sitikantha Mishra

package main

import ("net";"github.com/gin-gonic/gin";"time";"fmt")

type dest struct {
    URL     string  `json:"url"`
    Port  string  `json:"port"`
}

func main() {
        r := gin.Default()
        r.GET("/ping", func(c *gin.Context) {
                c.JSON(200, gin.H{
                        "message": "pong",
                })
        })
        r.POST("/portcheck", func(c *gin.Context) {
                var newdest dest
                if err := c.BindJSON(&newdest); err != nil {
                 return
                }
                url := newdest.URL
                port := newdest.Port
                full_url := url + ":" +port
                //conn, err := net.Dial("tcp",full_url)
                conn, err := net.DialTimeout("tcp",full_url,20  * time.Second)
                _ = conn
                if err != nil {
                c.JSON(200, gin.H{
                        "message": fmt.Sprint("failed for"," ",url,":",port),
                })
                }else{
                c.JSON(200, gin.H{
                        "message": fmt.Sprint("Successed for"," ",url,":",port),
                })
                }
        })
        r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
