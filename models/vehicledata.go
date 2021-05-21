package models

type Vehicledata struct {
	ID        int
	Dtime     string
	Imei      string
	Latitude  string
	Longitude string
	Speed     string
	Status    string
}

func GetAllVehicleData() []Vehicledata {
	var vehicleData []Vehicledata
	db, err := dbConn()
	if err != nil {
		panic("fail to connect to database .")
	}
	defer db.Close()

	results, err := db.Query("SELECT * FROM gps_data")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	for results.Next() {
		var tag Vehicledata
		err = results.Scan(&tag.ID, &tag.Dtime, &tag.Imei, &tag.Latitude, &tag.Longitude, &tag.Speed, &tag.Status)
		//err = results.Scan(&tag.ID, &tag.Imei, &tag.Dtime, &tag.Latitude, &tag.Longitude, &tag.Speed, &tag.Status)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		//fmt.Println(tag.Imei + "  " + tag.Longitude + "  " + tag.Latitude)
		vehicleData = append(vehicleData, tag)
		// and then print out the tag's Name attribute
		// log.Printf(tag.imei)
		// log.Printf(tag.latitude)
		// log.Printf(tag.longitude)
	}
	return vehicleData
}
