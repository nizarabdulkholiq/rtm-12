package controller

import (
	"fmt"
	"net/http"
	"rtm/config"

	pkg "github.com/MSyahidAlFajri/packagertm"
	"github.com/gofiber/fiber/v2"
	jobdesk "github.com/harisriyoni3/rtmpackage"
	rtpkg "github.com/rofinafiin/rtm-package"

"github.com/aiteung/musik"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/nizarabdulkholiq/nizar"
	"github.com/nizarabdulkholiq/tagihan1/config"
	"github.com/whatsauth/whatsauth"

)

// var usercol = "data_user"
// var datartm = "data_rtm"

//func WsWhatsAuthQR(c *websocket.Conn) {
//	whatsauth.RunSocket(c, config.PublicKey, config.Usertables[:], config.Ulbimariaconn)
//}
//
//func PostWhatsAuthRequest(c *fiber.Ctx) error {
//	if string(c.Request().Host()) == config.Internalhost {
//		var req whatsauth.WhatsauthRequest
//		err := c.BodyParser(&req)
//		if err != nil {
//			return err
//		}
//		ntfbtn := whatsauth.RunModuleLegacy(req, config.PrivateKey, config.Usertables[:], config.Ulbimariaconn)
//		return c.JSON(ntfbtn)
//	} else {
//		var ws whatsauth.WhatsauthStatus
//		ws.Status = string(c.Request().Host())
//		return c.JSON(ws)
//	}
//
//}

// func GetHome(c *fiber.Ctx) error {
// 	//getip := musik.GetIPaddress()
// 	getip := "Hello guys"
// 	return c.JSON(getip)
// }

// func Getdatauser(c *fiber.Ctx) error {
// 	id := "cc2"
// 	getstats := rtpkg.GetDatauser(id, config.MongoConn, usercol)
// 	fmt.Println(getstats)
// 	return c.JSON(getstats)
// }

// func InsertData(c *fiber.Ctx) error {
// 	database := config.MongoConn
// 	var jumlah rtpkg.User
// 	if err := c.BodyParser(&jumlah); err != nil {
// 		return err
// 	}
// 	Inserted := rtpkg.InsertDataUser(database,
// 		jumlah.Iduser,
// 		jumlah.Nama,
// 		jumlah.Email,
// 		jumlah.Handphone,
// 	)
// 	fmt.Println(Inserted)
// 	return c.JSON(map[string]interface{}{
// 		"status":      http.StatusOK,
// 		"message":     "Data berhasil disimpan.",
// 		"inserted_id": Inserted,
// 	})
// }

// func GetDataUserbyPhone(c *fiber.Ctx) error {
// 	hp := c.Params("handphone")
// 	data := rtpkg.GetDataUserFromPhone(hp, config.MongoConn, "data_user")
// 	fmt.Println(data)
// 	return c.JSON(data)
// }

// func DeleteDataUser(c *fiber.Ctx) error {
// 	hp := c.Params("handphone")
// 	data := rtpkg.DeleteData(hp, config.MongoConn, "data_user")
// 	return c.JSON(data)
// }

// func Getdatartm(c *fiber.Ctx) error {
// 	namarapat := "Rapat Akreditasi"
// 	getstats := pkg.GetDataRtm(namarapat, config.MongoConn, datartm)
// 	fmt.Println(getstats)
// 	return c.JSON(getstats)
// }

// func GetDataRtmByAgenda(c *fiber.Ctx) error {
// 	hp := c.Params("agendarapat")
// 	data := pkg.GetDataRtmFromAgenda(hp, config.MongoConn, "data_rtm")
// 	fmt.Println(data)
// 	return c.JSON(data)
// }

// func DeleteDataRtmFromLokasi(c *fiber.Ctx) error {
// 	hp := c.Params("lokasirapat")
// 	data := pkg.DeleteDataRtm(hp, config.MongoConn, "data_rtm")
// 	return c.JSON(data)
// }

// func InsertDataRapat(c *fiber.Ctx) error {
// 	database := config.MongoConn
// 	var tambah pkg.DataRTM
// 	if err := c.BodyParser(&tambah); err != nil {
// 		return err
// 	}
// 	Inserted := rtpkg.InsertDataUser(database,
// 		tambah.NamaRapat,
// 		tambah.TanggalRapat,
// 		tambah.LokasiRapat,
// 		tambah.AgendaRapat,
// 	)
// 	fmt.Println(Inserted)
// 	return c.JSON(map[string]interface{}{
// 		"status":      http.StatusOK,
// 		"message":     "Data berhasil Tersimpan.",
// 		"inserted_id": Inserted,
// 	})
// }

// // JOB
// func InsertDataJob(c *fiber.Ctx) error {
// 	database := config.MongoConn
// 	var job jobdesk.Job
// 	if err := c.BodyParser(&job); err != nil {
// 		return err
// 	}
// 	Inserted := jobdesk.InsertDataJob(database,
// 		job.Job_title,
// 		job.Deskripsi,
// 		job.Deadline,
// 		job.Priority,
// 	)
// 	fmt.Println(Inserted)
// 	return c.JSON(map[string]interface{}{
// 		"status":      http.StatusOK,
// 		"message":     "Data Job berhasil disimpan.",
// 		"inserted_id": Inserted,
// 	})
// }
// func GetDataJob(c *fiber.Ctx) error {
// 	hp := c.Params("priority")
// 	data := jobdesk.GetDataJob(hp, config.MongoConn, "data_job")
// 	fmt.Println(data)
// 	return c.JSON(data)
// }
// func GetDataJobtitle(c *fiber.Ctx) error {
// 	hp := c.Params("job_title")
// 	data := jobdesk.GetDataJobtitle(hp, config.MongoConn, "data_job")
// 	fmt.Println(data)
// 	return c.JSON(data)
// }
// func DeleteDataJob(c *fiber.Ctx) error {
// 	hp := c.Params("priority")
// 	data := jobdesk.DeleteDataJob(hp, config.MongoConn, "DeleteDataJob")
// 	return c.JSON(data)
// }
// func DeleteDataJobtitle(c *fiber.Ctx) error {
// 	hp := c.Params("job_title")
// 	data := jobdesk.DeleteDataJobtitle(hp, config.MongoConn, "DeleteDataJob")
// 	return c.JSON(data)
// }

var DataTagihanRegistrasi = "tagihanregis"
var DataTagihanSPP = "tagihanspp"


func WsWhatsAuthQR(c *websocket.Conn) {
	whatsauth.RunSocket(c, config.PublicKey, config.Usertables[:], config.Ulbimariaconn)
}

func PostWhatsAuthRequest(c *fiber.Ctx) error {
	if string(c.Request().Host()) == config.Internalhost {
		var req whatsauth.WhatsauthRequest
		err := c.BodyParser(&req)
		if err != nil {
			return err
		}
		ntfbtn := whatsauth.RunModuleLegacy(req, config.PrivateKey, config.Usertables[:], config.Ulbimariaconn)
		return c.JSON(ntfbtn)
	} else {
		var ws whatsauth.WhatsauthStatus
		ws.Status = string(c.Request().Host())
		return c.JSON(ws)
	}

}

func GetHome(c *fiber.Ctx) error {
	getip := musik.GetIPaddress()
	return c.JSON(getip)
}

//	func GetUserDataNomer(c *fiber.Ctx) error {
//		getstats := gaga.GetUserData("081234567890", config.MongoConn, user)
//		return c.JSON(getstats)
//	}

func GetTagihanSPP(c*fiber.Ctx) error {
	getstatus := nizar.GetTagihanSPP("dua")
	return c.JSON(getstatus)
}
