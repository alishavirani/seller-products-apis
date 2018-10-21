package structs

type Config struct {
	MongoDB struct {
		Host        string `json:"host"`
		Port        string `json:"port"`
		Db          string `json:"db"`
		Collections struct {
			Products string `json:"products"`
		} `json:"collections"`
	} `json:"MongoDB"`
	Server struct {
		Port string `json:"port"`
	} `json:"server"`
}
