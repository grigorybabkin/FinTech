package storage

import _ "github.com/lib/pq"

type url struct {
	fullURL, shortURL string
}

var urlsTest = []url{
	url{"https://www.ozon.ru/category/elektronika-15500/", "hIkp_O92Q3"},
	url{"https://www.ozon.ru/highlight/zakrytaya-rasprodazha-20906/", "k8usjOpl_9"},
	url{"https://www.ozon.ru/category/zhenskaya-odezhda-7501/", "yuIFG_620A"},
	url{"https://www.ozon.ru/travel/?perehod=ozon_menu_header", "yhbF_23lV8"},
}

const (
	host     = "127.0.0.1"
	port     = 5432
	user     = "root"
	password = "postgres"
	dbname   = "root"
)
