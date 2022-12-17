package beacons

type Coordinate struct {
	X int `parser:"XEquals @Int"`
	Y int `parser:"YEquals @Int"`
}
