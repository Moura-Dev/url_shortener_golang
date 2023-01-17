package repository

import (
	"url_shortner/db"
	"url_shortner/models"
)

func GetAllGolies() ([]models.Goly, error) {
	golies := []models.Goly{}
	err := db.Conn.Select(&golies, "SELECT * FROM goly")

	if err != nil {
		return []models.Goly{}, err
	}
	return golies, nil
}

func GetGoly(id uint64) (models.Goly, error) {
	golyModel := models.Goly{}
	err := db.Conn.Get(&golyModel, "SELECT * FROM goly WHERE id = $1", id)
	if err != nil {
		return models.Goly{}, err
	}
	return golyModel, nil
}
func GetGolyUrl(uuid string) (models.Goly, error) {
	golyModel := models.Goly{}
	err := db.Conn.Get(&golyModel, "SELECT * FROM goly WHERE goly = $1", uuid)
	if err != nil {
		return models.Goly{}, err
	}
	return golyModel, nil
}

func CreateGoly(goly models.Goly) (models.Goly, error) {
	_, err := db.Conn.NamedExec("INSERT INTO goly (redirect_url, goly, random) VALUES (:redirect_url, :goly, :random)", goly)
	if err != nil {
		return goly, err
	}
	return goly, nil
}

func UpdateGoly(goly models.Goly) (models.Goly, error) {
	_, err := db.Conn.NamedExec("UPDATE goly SET redirect_url = :redirect_url, goly = :goly, random = :random WHERE id = :id", goly)
	if err != nil {
		return goly, err
	}
	return goly, nil
}

func DeleteGoly(id uint64) error {
	_, err := db.Conn.Exec("DELETE FROM goly WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
