package beacons

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProcessReports(t *testing.T) {
	fakeData := &ReadingList{
		Readings: []*Reading{
			{
				SensorAt: Coordinate{X: 2, Y: 18}, ClosestBeacon: Coordinate{X: -2, Y: 15}}, {
				SensorAt: Coordinate{X: 9, Y: 16}, ClosestBeacon: Coordinate{X: 10, Y: 16}}, {
				SensorAt: Coordinate{X: 13, Y: 2}, ClosestBeacon: Coordinate{X: 15, Y: 3}}, {
				SensorAt: Coordinate{X: 12, Y: 14}, ClosestBeacon: Coordinate{X: 10, Y: 16}}, {
				SensorAt: Coordinate{X: 10, Y: 20}, ClosestBeacon: Coordinate{X: 10, Y: 16}}, {
				SensorAt: Coordinate{X: 14, Y: 17}, ClosestBeacon: Coordinate{X: 10, Y: 16}}, {
				SensorAt: Coordinate{X: 8, Y: 7}, ClosestBeacon: Coordinate{X: 2, Y: 10}}, {
				SensorAt: Coordinate{X: 2}, ClosestBeacon: Coordinate{X: 2, Y: 10}}, {
				SensorAt: Coordinate{Y: 11}, ClosestBeacon: Coordinate{X: 2, Y: 10}}, {
				SensorAt: Coordinate{X: 20, Y: 14}, ClosestBeacon: Coordinate{X: 25, Y: 17}}, {
				SensorAt: Coordinate{X: 17, Y: 20}, ClosestBeacon: Coordinate{X: 21, Y: 22}}, {
				SensorAt: Coordinate{X: 16, Y: 7}, ClosestBeacon: Coordinate{X: 15, Y: 3}}, {
				SensorAt: Coordinate{X: 14, Y: 3}, ClosestBeacon: Coordinate{X: 15, Y: 3}}, {
				SensorAt: Coordinate{X: 20, Y: 1}, ClosestBeacon: Coordinate{X: 15, Y: 3}},
		},
	}
	t.Run("should report 26 covered positions in line y=10", func(t *testing.T) {
		coveredPositions, _ := fakeData.CoveredPositions(10)
		assert.Equal(t, 26, coveredPositions)
	})
}
