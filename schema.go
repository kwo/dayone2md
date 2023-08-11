package dayone2md

import "time"

type Journal struct {
	Metadata *Metadata `json:"metadata"`
	Entries  []Entry   `json:"entries"`
}

type Metadata struct {
	Version string `json:"version"`
}

type Entry struct {
	CreationDate        time.Time `json:"creationDate"`
	ModifiedDate        time.Time `json:"modifiedDate"`
	Date                time.Time `json:"-"` // CreationDate with timezone applied
	TimeZone            string    `json:"timeZone"`
	Duration            int       `json:"duration"`
	Pinned              bool      `json:"isPinned"`
	Starred             bool      `json:"starred"`
	AllDay              bool      `json:"isAllDay"`
	UUID                string    `json:"uuid"`
	Title               string    `json:"-"` // calculated
	Text                string    `json:"text"`
	Tags                []string  `json:"tags,omitempty"`
	Photos              []Photo   `json:"photos,omitempty"`
	Location            *Location `json:"location,omitempty"`
	Weather             *Weather  `json:"weather,omitempty"`
	EditingTime         float64   `json:"editingTime,omitempty"`
	CreationDevice      string    `json:"creationDevice,omitempty"`
	CreationDeviceType  string    `json:"creationDeviceType,omitempty"`
	CreationDeviceModel string    `json:"creationDeviceModel,omitempty"`
	CreationOSName      string    `json:"creationOSName,omitempty"`
	CreationOSVersion   string    `json:"creationOSVersion,omitempty"`
}

type Photo struct {
	Date                 string    `json:"date"`
	Filename             string    `json:"filename"`
	FileSize             int       `json:"fileSize"`
	Height               int       `json:"height"`
	Identifier           string    `json:"identifier"`
	IsSketch             bool      `json:"isSketch"`
	MD5                  string    `json:"md5"`
	Type                 string    `json:"type"`
	Width                int       `json:"width"`
	AppleCloudIdentifier string    `json:"appleCloudIdentifier,omitempty"`
	CameraMake           string    `json:"cameraMake,omitempty"`
	CameraModel          string    `json:"cameraModel,omitempty"`
	CreationDevice       string    `json:"creationDevice,omitempty"`
	Duration             int       `json:"duration,omitempty"`
	ExposureBiasValue    int       `json:"exposureBiasValue,omitempty"`
	Favorite             bool      `json:"favorite,omitempty"`
	FocalLength          string    `json:"focalLength,omitempty"`
	FStop                string    `json:"fnumber,omitempty"`
	LensMake             string    `json:"lensMake,omitempty"`
	LensModel            string    `json:"lensModel,omitempty"`
	Location             *Location `json:"location,omitempty"`
	OrderInEntry         int       `json:"orderInEntry,omitempty"`
}

type Location struct {
	Label     string  `json:"userLabel"`
	Address   string  `json:"placeName"`
	City      string  `json:"localityName"`
	State     string  `json:"administrativeArea"`
	Country   string  `json:"country"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Altitude  float64 `json:"altitude"`
}

type Weather struct {
	Conditions         string    `json:"conditionsDescription"`
	MoonPhase          float64   `json:"moonPhase"`
	MoonPhaseCode      string    `json:"moonPhaseCode"`
	PressureMB         float64   `json:"pressureMB"`
	RelativeHumidity   int       `json:"relativeHumidity"`
	SunriseDate        time.Time `json:"sunriseDate"`
	SunsetDate         time.Time `json:"sunsetDate"`
	TemperatureCelsius float64   `json:"temperatureCelsius"`
	VisibilityKM       float64   `json:"visibilityKM"`
	WeatherCode        string    `json:"weatherCode"`
	WeatherServiceName string    `json:"weatherServiceName"`
	WindBearing        int       `json:"windBearing"`
	WindSpeedKPH       float64   `json:"windSpeedKPH"`
}
