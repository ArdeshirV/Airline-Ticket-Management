package persistance

import (
	"errors"
	"gorm.io/gorm"
	"internal/internal/domain"
	"net/http"
	"strconv"
)

type AirportRepo struct {
	db *gorm.DB
}

func (a *AirportRepo) New(db *gorm.DB) *AirportRepo {
	return &AirportRepo{db: db}
}

func (a *AirportRepo) Save(input *domain.Airport) (*domain.Airport, error) {
	var airport domain.Airport
	db := a.db.Model(&airport)

	checkAirportExist := db.Debug().First(&airport, "ID = ?", input.ID)

	if checkAirportExist.RowsAffected > 0 {
		return &airport, errors.New(strconv.Itoa(http.StatusConflict))
	}

	airport.Name = input.Name
	airport.City = input.City
	airport.Terminal = input.Terminal

	addNewAirport := db.Debug().Create(&airport).Commit()

	if addNewAirport.RowsAffected < 1 {
		return &airport, errors.New(strconv.Itoa(http.StatusForbidden))
	}

	return &airport, nil
}

func (a *AirportRepo) Update(input *domain.Airport) (*domain.Airport, error) {
	var airport domain.Airport
	db := a.db.Model(&airport)

	checkAirportExist := db.Debug().Where(&airport, "ID = ?", input.ID)

	if checkAirportExist.RowsAffected <= 0 {
		return &airport, errors.New(strconv.Itoa(http.StatusNotFound))
	}

	tx := checkAirportExist.Update("ID", input.ID).Update("Name", input.Name).Update("City", input.City).Update("Terminal", input.Terminal)

	if err := tx.Error; err != nil {
		return nil, err
	} else {
		updatedAirport := tx.Commit()
		if updatedAirport.RowsAffected < 1 {
			return &airport, errors.New(strconv.Itoa(http.StatusForbidden))
		}
	}

	return &airport, nil
}

func (a *AirportRepo) Get(id int) (*domain.Airport, error) {
	var airport domain.Airport
	db := a.db.Model(&airport)

	checkAirportExist := db.Debug().Where(&airport, "ID = ?", id)

	if checkAirportExist.RowsAffected <= 0 {
		return &airport, errors.New(strconv.Itoa(http.StatusNotFound))
	}

	tx := checkAirportExist.Find(&airport)

	if err := tx.Error; err != nil {
		return nil, err
	}

	return &airport, nil
}

func (a *AirportRepo) GetAll() (*[]domain.Airport, error) {
	var airports []domain.Airport
	db := a.db.Model(&airports)

	checkAirportExist := db.Debug().Find(&airports)

	if checkAirportExist.RowsAffected <= 0 {
		return &airports, errors.New(strconv.Itoa(http.StatusNotFound))
	}

	tx := db.Debug().Find(&airports)

	if err := tx.Error; err != nil {
		return nil, err
	}

	return &airports, nil
}

func (a *AirportRepo) Delete(id int) error {
	airport, err := a.Get(id)
	if err != nil {
		return err
	}
	db := a.db.Model(&airport)
	deleted := db.Debug().Delete(airport).Commit()
	if deleted.Error != nil {
		return deleted.Error
	}
	return nil
}
