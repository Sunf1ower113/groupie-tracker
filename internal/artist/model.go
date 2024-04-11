package artist

type Artists struct {
	Id    int    `json:"id"`
	Image string `json:"image"`
	Name  string `json:"name"`
}
type Relation struct {
	Location string
	Date     string
}
type Artist struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	locations    []string
	concertDates []string
	Rels         []Relation
}

type Location struct {
	Id        int      `json:"id"`
	Locations []string `json:"locations"`
}

type Date struct {
	Id    int      `json:"id"`
	Dates []string `json:"dates"`
}
