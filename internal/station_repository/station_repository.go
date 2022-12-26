package station_repository

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Station struct {
	ID          string `sql="id"`
	Slug        string `sql="sql"`
	Name        string `sql="name"`
	Site        string `sql="site"`
	Email       string `sql="email"`
	Region      string `sql="region"`
	City        string `sql="city"`
	Address     string `sql="address"`
	Facebook    string `sql="facebook"`
	Twitter     string `sql="twitter"`
	OK          string `sql="ok"`
	VK          string `sql="vk"`
	Wiki        string `sql="wiki"`
	Genre       string `sql="genre"`
	PhoneNumber string `sql="phonenumber"`
	Stream      string `sql="stream"`
}

const _stationDB = "stations"

func GetStation(id string, cnct *sql.DB) (Station, error) {
	query := fmt.Sprintf(`Select * from %s where id = %s `, _stationDB, id)
	data, err := cnct.QueryContext(context.Background(), query)

	if err != nil {
		return Station{}, err
	}

	var station Station

	for data.Next() {
		er := data.Scan(&station.ID, &station.Slug, &station.Name, &station.Site, &station.Email, &station.Region, &station.City, &station.Address, &station.Facebook, &station.Twitter, &station.OK, &station.VK, &station.Wiki, &station.Genre, &station.PhoneNumber, &station.Stream)
		if er != nil {
			return Station{}, er
		}
	}
	return station, nil
}
