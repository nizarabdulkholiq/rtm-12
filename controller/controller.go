package controller

import (
	"fmt"
	"net/http"
	"rtm/config"

	pkg "github.com/MSyahidAlFajri/packagertm"
	"github.com/gofiber/fiber/v2"
	rtpkg "github.com/rofinafiin/rtm-package"
)

var usercol = "data_user"

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

func GetHome(c *fiber.Ctx) error {
	//getip := musik.GetIPaddress()
	getip := "Hello guys"
	return c.JSON(getip)
}

func Getdatauser(c *fiber.Ctx) error {
	id := "cc2"
	getstats := rtpkg.GetDatauser(id, config.MongoConn, usercol)
	fmt.Println(getstats)
	return c.JSON(getstats)
}

func InsertData(c *fiber.Ctx) error {
	database := config.MongoConn
	var jumlah rtpkg.User
	if err := c.BodyParser(&jumlah); err != nil {
		return err
	}
	Inserted := rtpkg.InsertDataUser(database,
		jumlah.Iduser,
		jumlah.Nama,
		jumlah.Email,
		jumlah.Handphone,
	)
	fmt.Println(Inserted)
	return c.JSON(map[string]interface{}{
		"status":      http.StatusOK,
		"message":     "Data berhasil disimpan.",
		"inserted_id": Inserted,
	})
}

func GetDataUserbyPhone(c *fiber.Ctx) error {
	hp := c.Params("handphone")
	data := rtpkg.GetDataUserFromPhone(hp, config.MongoConn, "data_user")
	fmt.Println(data)
	return c.JSON(data)
}

func DeleteDataUser(c *fiber.Ctx) error {
	hp := c.Params("handphone")
	data := rtpkg.DeleteData(hp, config.MongoConn, "data_user")
	return c.JSON(data)
}

func Getdatartm(c *fiber.Ctx) error {
	namarapat := "Rapat Akreditasi"
	getstats := pkg.GetDataRtm(namarapat, config.MongoConn, usercol)
	fmt.Println(getstats)
	return c.JSON(getstats)
}

func GetDataRtmByAgenda(c *fiber.Ctx) error {
	hp := c.Params("agendarapat")
	data := pkg.GetDataRtmFromAgenda(hp, config.MongoConn, "data_rtm")
	fmt.Println(data)
	return c.JSON(data)
}

func DeleteDataRtmFromLokasi(c *fiber.Ctx) error {
	hp := c.Params("lokasirapat")
	data := pkg.DeleteDataRtm(hp, config.MongoConn, "data_rtm")
	return c.JSON(data)
}

func InsertDataRapat(c *fiber.Ctx) error {
	database := config.MongoConn
	var tambah pkg.DataRTM
	if err := c.BodyParser(&tambah); err != nil {
		return err
	}
	Inserted := rtpkg.InsertDataUser(database,
		tambah.NamaRapat,
		tambah.TanggalRapat,
		tambah.LokasiRapat,
		tambah.AgendaRapat,
	)
	fmt.Println(Inserted)
	return c.JSON(map[string]interface{}{
		"status":      http.StatusOK,
		"message":     "Data berhasil Tersimpan.",
		"inserted_id": Inserted,
	})
}
