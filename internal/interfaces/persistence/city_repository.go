package persistance

import (
	"errors"
	"gorm.io/gorm"
	"internal/internal/domain"
	"net/http"
	"strconv"
)

type CityRepo struct {
	db *gorm.DB
}

func (a *CityRepo) New(db *gorm.DB) *CityRepo {
	return &CityRepo{db: db}
}

func (a *CityRepo) Save(input *domain.City) (*domain.City, error) {
	var city domain.City
	db := a.db.Model(&city)

	checkCityExist := db.Debug().First(&city, "ID = ?", input.ID)

	if checkCityExist.RowsAffected > 0 {
		return &city, errors.New(strconv.Itoa(http.StatusConflict))
	}

	city.Name = input.Name

	addNewCity := db.Debug().Create(&city).Commit()

	if addNewCity.RowsAffected < 1 {
		return &city, errors.New(strconv.Itoa(http.StatusForbidden))
	}

	return &city, nil
}

func (a *CityRepo) Update(input *domain.City) (*domain.City, error) {
	var city domain.City
	db := a.db.Model(&city)

	checkCityExist := db.Debug().Where(&city, "ID = ?", input.ID)

	if checkCityExist.RowsAffected <= 0 {
		return &city, errors.New(strconv.Itoa(http.StatusNotFound))
	}

	tx := checkCityExist.Update("ID", input.ID).Update("Name", input.Name)

	if err := tx.Error; err != nil {
		return nil, err
	} else {
		updatedCity := tx.Commit()
		if updatedCity.RowsAffected < 1 {
			return &city, errors.New(strconv.Itoa(http.StatusForbidden))
		}
	}

	return &city, nil
}

func (a *CityRepo) Get(id int) (*domain.City, error) {
	var city domain.City
	db := a.db.Model(&city)

	checkCityExist := db.Debug().Where(&city, "ID = ?", id)

	if checkCityExist.RowsAffected <= 0 {
		return &city, errors.New(strconv.Itoa(http.StatusNotFound))
	}

	tx := checkCityExist.Find(&city)

	if err := tx.Error; err != nil {
		return nil, err
	}

	return &city, nil
}

func (a *CityRepo) GetAll() (*[]domain.City, error) {
	var cities []domain.City
	db := a.db.Model(&cities)

	checkCityExist := db.Debug().Find(&cities)

	if checkCityExist.RowsAffected <= 0 {
		return &cities, errors.New(strconv.Itoa(http.StatusNotFound))
	}

	tx := db.Debug().Find(&cities)

	if err := tx.Error; err != nil {
		return nil, err
	}

	return &cities, nil
}

func (a *CityRepo) Delete(id int) error {
	city, err := a.Get(id)
	if err != nil {
		return err
	}
	db := a.db.Model(&city)
	deleted := db.Debug().Delete(city).Commit()
	if deleted.Error != nil {
		return deleted.Error
	}
	return nil
}
