package met

type Objects struct {
	Total     int   `json:"total"`
	ObjectIDs []int `json:"objectIDs"`
}

// Object represents the structure of the response for an art object from the MET API.
type Object struct {
	ObjectID              int           `json:"objectID"`
	IsHighlight           bool          `json:"isHighlight"`
	AccessionNumber       string        `json:"accessionNumber"`
	AccessionYear         string        `json:"accessionYear"`
	IsPublicDomain        bool          `json:"isPublicDomain"`
	PrimaryImage          string        `json:"primaryImage"`
	PrimaryImageSmall     string        `json:"primaryImageSmall"`
	AdditionalImages      []string      `json:"additionalImages"`
	Constituents          []Constituent `json:"constituents"`
	Department            string        `json:"department"`
	ObjectName            string        `json:"objectName"`
	Title                 string        `json:"title"`
	Culture               string        `json:"culture"`
	Period                string        `json:"period"`
	Dynasty               string        `json:"dynasty"`
	Reign                 string        `json:"reign"`
	Portfolio             string        `json:"portfolio"`
	ArtistRole            string        `json:"artistRole"`
	ArtistPrefix          string        `json:"artistPrefix"`
	ArtistDisplayName     string        `json:"artistDisplayName"`
	ArtistDisplayBio      string        `json:"artistDisplayBio"`
	ArtistSuffix          string        `json:"artistSuffix"`
	ArtistAlphaSort       string        `json:"artistAlphaSort"`
	ArtistNationality     string        `json:"artistNationality"`
	ArtistBeginDate       string        `json:"artistBeginDate"`
	ArtistEndDate         string        `json:"artistEndDate"`
	ArtistGender          string        `json:"artistGender"`
	ArtistWikidataURL     string        `json:"artistWikidata_URL"`
	ArtistULANURL         string        `json:"artistULAN_URL"`
	ObjectDate            string        `json:"objectDate"`
	ObjectBeginDate       int           `json:"objectBeginDate"`
	ObjectEndDate         int           `json:"objectEndDate"`
	Medium                string        `json:"medium"`
	Dimensions            string        `json:"dimensions"`
	Measurements          []Measurement `json:"measurements"`
	CreditLine            string        `json:"creditLine"`
	GeographyType         string        `json:"geographyType"`
	City                  string        `json:"city"`
	State                 string        `json:"state"`
	County                string        `json:"county"`
	Country               string        `json:"country"`
	Region                string        `json:"region"`
	Subregion             string        `json:"subregion"`
	Locale                string        `json:"locale"`
	Locus                 string        `json:"locus"`
	Excavation            string        `json:"excavation"`
	River                 string        `json:"river"`
	Classification        string        `json:"classification"`
	RightsAndReproduction string        `json:"rightsAndReproduction"`
	LinkResource          string        `json:"linkResource"`
	MetadataDate          string        `json:"metadataDate"`
	Repository            string        `json:"repository"`
	ObjectURL             string        `json:"objectURL"`
	Tags                  []Tag         `json:"tags"`
	ObjectWikidataURL     string        `json:"objectWikidata_URL"`
	IsTimelineWork        bool          `json:"isTimelineWork"`
	GalleryNumber         string        `json:"GalleryNumber"`
}

// Constituent struct models the information about an artist or contributor.
type Constituent struct {
	ConstituentID          int    `json:"constituentID"`
	Role                   string `json:"role"`
	Name                   string `json:"name"`
	ConstituentULANURL     string `json:"constituentULAN_URL"`
	ConstituentWikidataURL string `json:"constituentWikidata_URL"`
	Gender                 string `json:"gender"`
}

// Measurement struct models the dimensions of an art object.
type Measurement struct {
	ElementName         string             `json:"elementName"`
	ElementDescription  string             `json:"elementDescription"`
	ElementMeasurements map[string]float64 `json:"elementMeasurements"`
}

// Tag struct models the tagging system used for categorizing the art object.
type Tag struct {
	Term        string `json:"term"`
	AATURL      string `json:"AAT_URL"`
	WikidataURL string `json:"Wikidata_URL"`
}
