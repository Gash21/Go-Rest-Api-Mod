package model

type Article struct {
	Id      string `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"Desc"`
	Content string `json:"Content"`
}

var Articles []Article

func FindArticle(id string) Article {
	for _, article := range Articles {
		if article.Id == id {
			return article
		}
	}
	return Article{Id: "0"}
}

func CreateNewArticle(article Article) {
	Articles = append(Articles, article)
	//return article
}

func UpdateArticle(art Article) {
	for ii, article := range Articles {
		if article.Id == art.Id {
			Articles[ii].Id = art.Id
			Articles[ii].Title = art.Title
			Articles[ii].Desc = art.Desc
			Articles[ii].Content = art.Content
		}
	}
}

func Populate() {
	Articles = []Article{
		Article{
			Id:      "1",
			Title:   "Gas judul",
			Desc:    "judul untuk deskripsi",
			Content: "Lorem Ipsum",
		},
		Article{
			Id:      "2",
			Title:   "Gas ajah",
			Desc:    "ajah untuk deskripsi",
			Content: "Lorem Ipsum dolores",
		},
	}
}

func init() {
	Populate()
}
