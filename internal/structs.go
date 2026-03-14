package main

type Metadata struct {
	Title              string             `json:"title"`
	Description        string             `json:"description"`
	ImageViews         string             `json:"imageViews"`
	CreationTime       TimeInfo           `json:"creationTime"`
	PhotoTakenTime     TimeInfo           `json:"photoTakenTime"`
	GeoData            GeoData            `json:"geoData"`
	GeoDataExif        GeoData            `json:"geoDataExif"`
	URL                string             `json:"url"`
	GooglePhotosOrigin GooglePhotosOrigin `json:"googlePhotosOrigin"`
	AppSource          AppSource          `json:"appSource"`
}

type TimeInfo struct {
	Timestamp string `json:"timestamp"`
	Formatted string `json:"formatted"`
}

type GeoData struct {
	Latitude      float64 `json:"latitude"`
	Longitude     float64 `json:"longitude"`
	Altitude      float64 `json:"altitude"`
	LatitudeSpan  float64 `json:"latitudeSpan"`
	LongitudeSpan float64 `json:"longitudeSpan"`
}

type GooglePhotosOrigin struct {
	MobileUpload MobileUpload `json:"mobileUpload"`
}

type MobileUpload struct {
	DeviceFolder DeviceFolder `json:"deviceFolder"`
	DeviceType   string       `json:"deviceType"`
}

type DeviceFolder struct {
	LocalFolderName string `json:"localFolderName"`
}

type AppSource struct {
	AndroidPackageName string `json:"androidPackageName"`
}