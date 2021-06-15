package controller

import (
	"display-antrian/config"
	"display-antrian/models"
	"encoding/json"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Display(c *gin.Context) {
	c.Header("Access-Control-Allow-Headers", "Content-type")
	c.Header("Access-Control-Allow-Method", "POST, GET, OPTIONS, PUT, DELETE")
	c.Header("Access-Control-Allow-Origin", "*")
	var arrDisplay []models.DisplayAntrian
	var response models.ResponseDisplayAntrian
	dt := time.Now()
	dates := dt.Format("2006.01.02")
	db, errCon := config.ConnectSQL()
	if errCon != nil {
		// logger.Log(pkgName, 4, "failed to connect database, reason: " + errCon.Error())
		log.Panic(errCon)
	}
	var idAntrian int
	loket, errl := db.SQL.Queryx(`SELECT id, nama as loket FROM mst_pelayanan`)
	if errl != nil {
		log.Println(errl)
	}
	defer loket.Close()
	for loket.Next() {
		var Display models.DisplayAntrian
		errScan := loket.Scan(&idAntrian, &Display.Loket)
		if errScan != nil {
			log.Println(errScan)
		}
		e := db.SQL.Get(&Display.Antrian, `select t.no_antrian as antrian from tran_form_isian t 
		left join mst_pelayanan mp on mp.id = t.id_pelayanan
		where status = 'On Progress' and tanggal_kedatangan = $1 and id_pelayanan = $2`, dates, idAntrian)
		if e != nil {
			Display.Antrian = "-"
		}

		arrDisplay = append(arrDisplay, Display)
	}

	// log.Println("MANTAAPPPP ", loket, noAntrian)

	response.Status = 200
	response.Message = "Success"
	response.Data = arrDisplay
	c.Header("Content-Type", "application/json")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(c.Writer).Encode(response)
}

func TextBerjalan(c *gin.Context) {
	c.Header("Access-Control-Allow-Headers", "Content-type")
	c.Header("Access-Control-Allow-Method", "POST, GET, OPTIONS, PUT, DELETE")
	c.Header("Access-Control-Allow-Origin", "*")

	var text []string
	var response models.ResponseTextBerjalan
	db, errCon := config.ConnectSQL()
	if errCon != nil {
		// logger.Log(pkgName, 4, "failed to connect database, reason: " + errCon.Error())
		log.Panic(errCon)
	}

	q, e := db.SQL.Queryx(`select texts from mst_text_berjalan`)
	if e != nil {
		log.Panicln(e)
	}
	defer db.SQL.Close()

	for q.Next() {
		var t string
		err := q.Scan(&t)
		if err != nil {
			log.Panicln(err)
		}
		text = append(text, t)
	}

	response.Status = 200
	response.Message = "Success"
	response.Data = text
	c.Header("Content-Type", "application/json")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(c.Writer).Encode(response)

}
