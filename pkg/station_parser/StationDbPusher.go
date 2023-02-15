package StationParser

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/essentialkaos/translit"
	"github.com/jackc/pgx/v5"
	"log"
	"nokogiriwatir/radio-main/pkg/database"
	"os"
	"strings"
)

type Stations struct {
	Sheet []struct {
		Slug        string `sql:"slug"`
		Name        string `json,sql:"name"`
		Site        string `json:"site,omitempty"`
		Email       string `json:"email,omitempty"`
		Region      string `json:"region"`
		City        string `json:"city"`
		Address     string `json:"address,omitempty"`
		Facebook    string `json:"facebook,omitempty"`
		Twitter     string `json:"twitter,omitempty"`
		Ok          string `json:"одноклассники,omitempty"`
		Vk          string `json:"vk,omitempty"`
		Wiki        string `json:"wiki,omitempty"`
		Genre       string `json:"genre,omitempty"`
		PhoneNumber string `json:"phoneNumber,omitempty"`
		Stream      string `json:"stream"`
	} `json:"Sheet1"`
}

func StationsMigrate(filePath string, config *database.DbConfig) error {
	stations := parseStations(filePath)

	stations = addSlugToStations(stations)

	err := WriteStationInDB(stations, config)

	if err != nil {
		log.Fatalln("Error occurred while writing in database" + err.Error())
	}

	return nil
}

func addSlugToStations(stations *Stations) *Stations {
	for in, el := range stations.Sheet {
		stations.Sheet[in].Slug = strings.ToLower(
			strings.Replace(
				translit.EncodeToISO9B(el.Name), " ", "_", -1),
		)
	}

	return stations
}

func parseStations(filePath string) *Stations {
	file, err := os.Open(filePath)

	fileInfo, _ := file.Stat()

	stationsRaw := make([]byte, fileInfo.Size())

	if err != nil {
		log.Fatalf("%s, Can't open file %s \n", err, filePath)
	}

	_, er := file.Read(stationsRaw)

	if er != nil {
		log.Fatalln("Can't read json with stations")
	}

	var stations *Stations

	err = json.Unmarshal(stationsRaw, &stations)

	if err != nil {
		log.Fatalln(err)
	}

	return stations
}

func WriteStationInDB(stations *Stations, config *database.DbConfig) error {
	var err error
	conn := DbConnection(config)

	defer conn.Close(context.Background())

	var raw string
	for in, _ := range stations.Sheet {
		_ = conn.QueryRow(context.Background(), `INSERT INTO stations (id, slug, name, site, email, region, city, address, facebook, twitter, ok, vk, wiki, genre, phoneNumber, stream) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16)`,
			in, stations.Sheet[in].Slug, stations.Sheet[in].Name, stations.Sheet[in].Site, stations.Sheet[in].Email, stations.Sheet[in].Region,
			stations.Sheet[in].City, stations.Sheet[in].Address, stations.Sheet[in].Facebook, stations.Sheet[in].Twitter, stations.Sheet[in].Ok, stations.Sheet[in].Vk,
			stations.Sheet[in].Wiki, stations.Sheet[in].Genre, stations.Sheet[in].PhoneNumber, stations.Sheet[in].Stream).Scan(&raw)
	}

	if err != nil {
		return err
	}

	return nil
}

func DbConnection(config *database.DbConfig) *pgx.Conn {
	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s", config.User, config.Password, config.Address, config.Name)

	conn, err := pgx.Connect(context.Background(), connStr)

	if err != nil {
		log.Fatalln("Unable to connect to database\n" + err.Error())
	}

	return conn
}
