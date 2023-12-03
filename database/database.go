package database

import (
	"fmt"

	"go-graphql/constants"
	"go-graphql/graph/model"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


var DB *Database

type Database struct {
	DB *gorm.DB
}

func PostgresDB() (*Database, error) {
	db, err := gorm.Open(postgres.Open(fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		//database host
		constants.Host,
		//database port
		constants.Port,
		//database user
		constants.User,
		//database pass
		constants.Password,
		//database name
		constants.Dbname,
		//ssl
		"disable")), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}
	return &Database{
		DB: db,
	}, nil
}

func (app *Database) CreateJobListing(input model.CreateJobListingInput) *model.JobListing {
	var (
		dataid string
	)
	app.DB.Debug().Raw(`insert into public.info(title, description,company,url)
	values(?,?,?,?) returning _id`, input.Title, input.Description, input.Company, input.URL).Scan(&dataid)
	return &model.JobListing{
		ID:          dataid,
		Title:       input.Title,
		Description: input.Description,
		Company:     input.Company,
		URL:         input.URL,
	}
}

func (app *Database) UpdateJobListing(id string, input model.UpdateJobListingInput) *model.JobListing {
	var data *model.JobListing
	app.DB.Debug().Table("public.info").Where("_id= " + id + "").Updates(&struct {
		Title       *string `json:"title"`
		Description *string `json:"description"`
		Url         *string `json:"url"`
		Company     *string `json:"company"`
	}{
		Title:       input.Title,
		Description: input.Description,
		Url:         input.URL,
		Company:     input.Company,
	})

	app.DB.Debug().Raw(`select *, _id as id from public.info where _id= ?`, id).Scan(&data)
	return data
}

func (app *Database) JobListing() []*model.JobListing {
	var data []*model.JobListing
	app.DB.Debug().Raw(`select *, _id as id from public.info`).Scan(&data)
	return data
}

func (app *Database) JobListingByID(id string) *model.JobListing {
	var data *model.JobListing
	app.DB.Debug().Raw(`select *, _id as id from public.info where _id=?`, id).Scan(&data)
	return data
}

func (app *Database) DeleteJobID(id string) *model.DeleteJobResponse {
	app.DB.Debug().Exec(`delete from public.info where _id=?`, id)
	return &model.DeleteJobResponse{
		DeleteJobID: &id,
	}
}
