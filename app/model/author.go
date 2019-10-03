package model

import (
	"fmt"
	"log"
	"os"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

type Author struct {
	Id       int    `json:"Id" sql:",pk"`
	Name     string `json:"Name" sql:",notnull"`
	Username string `json:"Username" sql:",notnull"`
	Email    string `json:"Email" sql:",notnull, unique"`
	Phone    string `json:"Phone" sql:",notnull"`
}

var Authors []Author
var db *pg.DB

func FindAuthor(id int) Author {
	opt, err := pg.ParseURL("postgres://evaizee:secret@localhost:5432/tutorial?sslmode=disable")
	if err != nil {
		log.Println(err)
	}
	db = pg.Connect(opt)

	user := &Author{Id: id}
	db.Begin()
	defer db.Close()
	err = db.Select(user)
	if err != nil {
		fmt.Println(err)
	}

	db.Close()
	return Author{Id: user.Id, Name: user.Name, Username: user.Username, Email: user.Email, Phone: user.Phone}
}

func CreateNewAuthor(author Author) {
	Authors = append(Authors, author)
	//return article
}

func UpdateAuthor(author Author) {
	for ii, article := range Authors {
		if article.Id == author.Id {
			Authors[ii].Id = author.Id
			Authors[ii].Name = author.Name
			Authors[ii].Username = author.Username
			Authors[ii].Email = author.Email
			Authors[ii].Phone = author.Phone
		}
	}
}

func populateAuthor() {
	auth := new(Author)
	exists, err := db.Model(auth).Where("id = ?", 1).Exists()

	if err != nil {
		panic(err)
	}

	if !exists {
		author2 := Author{
			Id:       3,
			Name:     "Lorem Ipsum",
			Username: "lorem2233",
			Email:    "lorem@ips.biz.gas",
			Phone:    "31845009",
		}

		author3 := Author{
			Id:       2,
			Name:     "Joahn Doe",
			Username: "johandou",
			Email:    "lorjohanem@ips.biz",
			Phone:    "318450008",
		}

		author1 := Author{
			Id:       4,
			Name:     "Lorem Ipsum",
			Username: "lorem2233",
			Email:    "lorem@ips.biz.co",
			Phone:    "31845009",
		}

		err := db.Insert(author1, author2, author3)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("populate author done")
	} else {
		fmt.Println("populate author not running")
	}
}

func init() {
	opt, err := pg.ParseURL("postgres://evaizee:secret@localhost:5432/tutorial?sslmode=disable")
	if err != nil {
		log.Println(err)
	}
	db = pg.Connect(opt)

	db.Begin()

	err1 := db.CreateTable(&Author{}, &orm.CreateTableOptions{
		Temp:          false, // create temp table
		FKConstraints: true,
		IfNotExists:   true,
	})
	if err1 != nil {
		log.Println(err1)
	}

	file, err := os.OpenFile("authorInfo.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0664)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	log.SetOutput(file)
	log.Println("author ready")
	populateAuthor()
	db.Close()
	log.Printf("connection closed")
}
