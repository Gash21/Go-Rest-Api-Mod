package model

import (
	"log"
	"os"
)

type Author struct {
	Id       int    `json:"Id"`
	Name     string `json:"Name"`
	Username string `json:"Username"`
	Email    string `json:"Email"`
	Phone    string `json:"Phone"`
}

var Authors []Author

func FindAuthor(id int) Author {
	for _, auth := range Authors {
		if auth.Id == id {
			return auth
		}
	}
	return Author{Id: 0}
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
	Authors = []Author{
		Author{
			Id:       1,
			Name:     "Lorem Ipsum",
			Username: "lorem2233",
			Email:    "lorem@ips.biz",
			Phone:    "31845009",
		},
		Author{
			Id:       1,
			Name:     "Joahn Doe",
			Username: "johandou",
			Email:    "lorjohanem@ips.biz",
			Phone:    "318450008",
		},
	}
	log.Println("populate author done")
}

func init() {
	file, err := os.OpenFile("authorInfo.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0664)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	log.SetOutput(file)
	log.Println("author ready")
	populateAuthor()
}
