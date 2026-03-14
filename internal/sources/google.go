package sources

type GoogleTakeoutMetadata struct {
	Title        string `json:"title"`
	Description  string `json:"description"`
	ImageViews   string `json:"imageViews"`
	CreationTime struct {
		Timestamp string `json:"timestamp"`
		Formatted string `json:"formatted"`
	} `json:"creationTime"`
	PhotoTakenTime struct {
		Timestamp string `json:"timestamp"`
		Formatted string `json:"formatted"`
	} `json:"photoTakenTime"`
	GeoData struct {
		Latitude      float64 `json:"latitude"`
		Longitude     float64 `json:"longitude"`
		Altitude      float64 `json:"altitude"`
		LatitudeSpan  float64 `json:"latitudeSpan"`
		LongitudeSpan float64 `json:"longitudeSpan"`
	} `json:"geoData"`
	GeoDataExif struct {
		Latitude      float64 `json:"latitude"`
		Longitude     float64 `json:"longitude"`
		Altitude      float64 `json:"altitude"`
		LatitudeSpan  float64 `json:"latitudeSpan"`
		LongitudeSpan float64 `json:"longitudeSpan"`
	} `json:"geoDataExif"`
	URL                string `json:"url"`
	GooglePhotosOrigin struct {
		MobileUpload struct {
			DeviceFolder struct {
				LocalFolderName string `json:"localFolderName"`
			} `json:"deviceFolder"`
			DeviceType string `json:"deviceType"`
		} `json:"mobileUpload"`
	} `json:"googlePhotosOrigin"`
	AppSource struct {
		AndroidPackageName string `json:"androidPackageName"`
	} `json:"appSource"`
}
