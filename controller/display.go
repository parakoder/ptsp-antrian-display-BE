package controller

import (
	"display-antrian/config"
	"display-antrian/models"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)


func inTimeSpan(start, end, check time.Time) bool {
	return check.After(start) && check.Before(end)
}
var idJam int

func getJamKedatanganID() int {
	// tx := m.Conn.MustBegin()
	// var jam2 bool
	fmt.Println("TES")
	dt := time.Now()
	layoutJam := "15:04"
	dates := dt.Format("15:04")
	datesParse, _ := time.Parse(layoutJam, dates)

	// ======================== jam ke 1 ========================
	start1 := "08:00"
	startParse1, _ := time.Parse(layoutJam, start1)

	end1 := "09:00"
	endParse1, _ := time.Parse(layoutJam, end1)

	jam1 := inTimeSpan(startParse1, endParse1, datesParse)

	// ======================== jam ke 2 ========================
	start2 := "09:00"
	startParse2, _ := time.Parse(layoutJam, start2)

	end2 := "10:00"
	endParse2, _ := time.Parse(layoutJam, end2)

	jam2 := inTimeSpan(startParse2, endParse2, datesParse)

	// ======================== jam ke 3 ========================
	start3 := "10:00"
	startParse3, _ := time.Parse(layoutJam, start3)

	end3 := "11:00"
	endParse3, _ := time.Parse(layoutJam, end3)

	jam3 := inTimeSpan(startParse3, endParse3, datesParse)

	// ======================== jam ke 4 ========================
	start4 := "11:00"
	startParse4, _ := time.Parse(layoutJam, start4)

	end4 := "12:00"
	endParse4, _ := time.Parse(layoutJam, end4)

	jam4 := inTimeSpan(startParse4, endParse4, datesParse)

	// ======================== jam ke 5 ========================
	start5 := "13:00"
	startParse5, _ := time.Parse(layoutJam, start5)

	end5 := "14:00"
	endParse5, _ := time.Parse(layoutJam, end5)

	jam5 := inTimeSpan(startParse5, endParse5, datesParse)

	// ======================== jam ke 6 ========================
	start6 := "14:00"
	startParse6, _ := time.Parse(layoutJam, start6)

	end6 := "15:00"
	endParse6, _ := time.Parse(layoutJam, end6)
	jam6 := inTimeSpan(startParse6, endParse6, datesParse)
	start7 := "21:00"
	startParse7, _ := time.Parse(layoutJam, start7)

	end7 := "23:00"
	endParse7, _ := time.Parse(layoutJam, end7)

	jam7 := inTimeSpan(startParse7, endParse7, datesParse)

	if jam1 == true {
		idJam = 1
	} else if jam2 == true {
		idJam = 2
	} else if jam3 == true {
		idJam = 3
	} else if jam4 == true {
		idJam = 4
	} else if jam5 == true {
		idJam = 5
	} else if jam6 == true {
		idJam = 6
	} else if jam7 == true {
		idJam = 7
	}

	log.Println("INI DIA ID JAM NYA ", idJam)

	return idJam
}

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
	idJam := getJamKedatanganID()
	var idAntrian int
	loket, errl := db.SQL.Queryx(`SELECT id, nama as loket FROM mst_pelayanan ORDER BY id ASC`)
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
		where status = 'On Progress' and tanggal_kedatangan = $1 and id_pelayanan = $2 and jam_kedatangan =$3`, dates, idAntrian, idJam)
		if e != nil {
			Display.Antrian = "-"
		}

		arrDisplay = append(arrDisplay, Display)
	}
	defer db.SQL.Close()
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
