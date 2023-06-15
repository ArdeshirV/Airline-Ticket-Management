package persistance

import (
	"errors"
	"gorm.io/gorm"
	"internal/internal/domain"
	"net/http"
	"strconv"
)

type AirlineRepo struct {
	db *gorm.DB
}

func (r *AirlineRepo) New(db *gorm.DB) *AirlineRepo {
	return &AirlineRepo{db: db}
}

func (r *AirlineRepo) Save(input *domain.Airline) (*domain.Airline, error) {
	var airline domain.Airline
	db := r.db.Model(&airline)

	checkAirlineExist := db.Debug().First(&airline, "ID = ?", input.ID)

	if checkAirlineExist.RowsAffected > 0 {
		return &airline, errors.New(strconv.Itoa(http.StatusConflict))
	}

	airline.Name = input.Name
	airline.Logo = input.Logo

	addNewAirline := db.Debug().Create(&airline).Commit()

	if addNewAirline.RowsAffected < 1 {
		return &airline, errors.New(strconv.Itoa(http.StatusForbidden))
	}

	return &airline, nil
}

func (r *AirlineRepo) Update(input *domain.Airline) (*domain.Airline, error) {
	var airline domain.Airline
	db := r.db.Model(&airline)

	checkAirlineExist := db.Debug().Where(&airline, "ID = ?", input.ID)

	if checkAirlineExist.RowsAffected <= 0 {
		return &airline, errors.New(strconv.Itoa(http.StatusNotFound))
	}

	tx := checkAirlineExist.Update("ID", input.ID).Update("Name", input.Name).Update("Logo", input.Logo)

	if err := tx.Error; err != nil {
		return nil, err
	} else {
		updatedAirline := tx.Commit()
		if updatedAirline.RowsAffected < 1 {
			return &airline, errors.New(strconv.Itoa(http.StatusForbidden))
		}
	}

	return &airline, nil
}

func (r *AirlineRepo) Get(id int) (*domain.Airline, error) {
	var airline domain.Airline
	db := r.db.Model(&airline)

	checkAirlineExist := db.Debug().Where(&airline, "ID = ?", id)

	if checkAirlineExist.RowsAffected <= 0 {
		return &airline, errors.New(strconv.Itoa(http.StatusNotFound))
	}

	tx := db.Debug().Where("ID = ?", id).Find(&airline)

	if err := tx.Error; err != nil {
		return nil, err
	} else {
		updatedAirline := tx.Commit()
		if updatedAirline.RowsAffected < 1 {
			return &airline, errors.New(strconv.Itoa(http.StatusForbidden))
		}
	}

	return &airline, nil
}

func (r *AirlineRepo) GetAll() (*[]domain.Airline, error) {
	var airlines []domain.Airline
	db := r.db.Model(&airlines)

	checkAirlineExist := db.Debug().Find(&airlines)

	if checkAirlineExist.RowsAffected <= 0 {
		return &airlines, errors.New(strconv.Itoa(http.StatusNotFound))
	}

	tx := db.Debug().Find(&airlines)

	if err := tx.Error; err != nil {
		return nil, err
	} else {
		updatedAirline := tx.Commit()
		if updatedAirline.RowsAffected < 1 {
			return &airlines, errors.New(strconv.Itoa(http.StatusForbidden))
		}
	}

	return &airlines, nil
}

func (r *AirlineRepo) Delete(id int) error {
	airline, err := r.Get(id)
	if err != nil {
		return err
	}
	db := r.db.Model(&airline)
	deleted := db.Debug().Delete(airline).Commit()
	if deleted.Error != nil {
		return deleted.Error
	}
	return nil
}
