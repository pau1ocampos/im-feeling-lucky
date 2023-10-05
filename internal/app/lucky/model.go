package lucky

type Draw struct {
	Numbers []int `json:"numbers"`
	Starts  []int `json:"stars"`
}

type Draws []Draw
