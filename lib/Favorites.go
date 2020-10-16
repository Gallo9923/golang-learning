package lib

// Stores favorites
var favorites []string

// Initialization logic for the package
func init() {
	favorites = make([]string, 3)
	favorites[0] = "github.com/Gallo9923"
	favorites[1] = "Chris"
	favorites[2] = "Gallo"
}

// Add add a favorite into the in-memory collection
func Add(favorite string) {
	favorites = append(favorites, favorite)
}

// GetAll Returns all favorites
func GetAll() []string {
	return favorites
}
