package main

import (
    "log"
    "github.com/gin-gonic/gin"
    "os"
)

func main() {
    log.Println("start server...")
    r := gin.Default()
    r.GET("/v1/servers", get_servers)
    log.Fatal(r.Run())
}

func get_servers(context *gin.Context) {
    servers := make(gin.H)

     groups, _ := os.ReadDir("servers")
     for i := range groups {
         if !groups[i].IsDir() {
             continue
         }
         group := groups[i].Name()
         local := []string{}

         names, _ := os.ReadDir("servers/" + group)
         for j := range names {
             if names[j].IsDir() {
                local = append(local, names[j].Name())
             }
         }
         servers[group] = local
     }

    context.JSON(200, servers)
}
