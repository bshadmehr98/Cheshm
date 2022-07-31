package main

import (
	"fmt"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/bshadmehr98/cheshm/db"
	"github.com/bshadmehr98/cheshm/models"
	_ "github.com/bshadmehr98/cheshm/routers"

	_ "github.com/bshadmehr98/cheshm/db"
)

func populate() error {
	if v, _ := beego.AppConfig.Bool("populate_data"); !v {
		return nil
	}

	fmt.Println("Populating the database...")

	adminTmp, err := beego.AppConfig.DIY("Admin")
	if err != nil {
		return err
	}

	admin := adminTmp.(map[string]string)
	fmt.Println(admin)

	user := models.User{FirstName: admin["first_name"], LastName: admin["last_name"],
		Email: admin["email"], Type: models.TypeUserAdmin}

	fmt.Println(user)

	if created, id, err := db.DB.ReadOrCreate(&user, "Email"); err == nil {
		if created {
			fmt.Println("New Insert an object. Id:", id)
		} else {
			fmt.Println("Get an object. Id:", id)
		}
	}

	return nil
}

func startApp() {
	err := populate()
	if err != nil {
		panic(err)
	}
}

func main() {
	startApp()

	beego.Run()
}
