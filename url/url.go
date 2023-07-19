package url

import (
	"github.com/nizarabdulkholiq/rtm/controller"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func Web(page *fiber.App) {
	//page.Post("/api/whatsauth/request", controller.PostWhatsAuthRequest)  //API from user whatsapp message from iteung gowa
	//page.Get("/ws/whatsauth/qr", websocket.New(controller.WsWhatsAuthQR)) //websocket whatsauth
	// page.Get("/", controller.GetHome)
	// page.Get("/user", controller.Getdatauser)
	// page.Post("/insert", controller.InsertData)
	// page.Get("/userdata/:handphone", controller.GetDataUserbyPhone)
	// page.Delete("delete/:handphone", controller.DeleteDataUser)
	// page.Get("/rapat", controller.Getdatartm)
	// page.Post("/insertrapat", controller.InsertDataRapat)
	// page.Get("/datartm/:agenda", controller.GetDataRtmByAgenda)
	// page.Delete("delete/:lokasi", controller.DeleteDataRtmFromLokasi)
	// page.Post("/job/insert", controller.InsertDataJob)
	// page.Get("/job/get/datajob", controller.GetDataJob)
	// page.Get("/job/get/datajobtitle", controller.GetDataJobtitle)
	// page.Delete("/job/delete/datajob", controller.DeleteDataJob)
	// page.Delete("/job/delete/datajobtitle", controller.DeleteDataJobtitle)
	page.Post("/api/whatsauth/request", controller.PostWhatsAuthRequest)  //API from user whatsapp message from iteung gowa
	page.Get("/ws/whatsauth/qr", websocket.New(controller.WsWhatsAuthQR)) //websocket whatsauth
	page.Get("/", controller.GetHome)
	
	page.Get("/tagihanspp", controller.GetTagihanSPP)
}
