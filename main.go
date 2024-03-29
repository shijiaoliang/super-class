/**
 *******************************************slade********************************************
 * Copyright (c)  slade
 * Created by super-class.
 * User: 605724193@qq.com
 * Date: 2019/08/07
 * Time: 11:18
 ********************************************************************************************
 */

package main

import (
	"time"

	"github.com/gin-gonic/gin"

	"super-class/log"
	"super-class/util"
	"super-class/config"
	"super-class/db"
	"super-class/middleware"
	"super-class/router"
)

func main() {
	defer db.Mysql.Close()

	r := gin.New()

	//=====config=====
	// config init

	//=====mode=====
	//gin.SetMode(gin.ReleaseMode)
	gin.SetMode(config.AppConf.Model)

	//=====log=====
	runtimeDirectory := "runtime"
	if !util.FileExist(runtimeDirectory) {
		util.CreateFile(runtimeDirectory)
	}
	log.ConfigLogger(log.Logger, runtimeDirectory, "app.log", time.Hour*360, time.Hour)

	//=====middleware=====
	r.Use(middleware.Logger(log.Logger))
	r.Use(gin.Recovery())

	//=====route=====
	router.InitRouter(r)

	//=====run=====
	r.Run(config.AppConf.Address)
}
