package beacons

import (
	"github.com/samber/lo"
	"golang.org/x/exp/slices"
	"math"
)

type ReadingList struct {
	Readings []*Reading `parser:"@@*"`
}

type RowCoverageMap = map[int]bool
type CoveredInterval = [2]int

func (l *ReadingList) CoveredPositions(targetRow int) (int, []int) {
	var intervals = []CoveredInterval{}
	minX, maxX := math.MaxInt, math.MinInt

	for _, reading := range l.Readings {
		distanceToBeacon := reading.SensorBeaconDistance()
		distanceToTargetRow := int(math.Abs(float64(reading.SensorAt.Y - targetRow)))

		if distanceToTargetRow >= distanceToBeacon {
			continue
		}

		sensorsCoverageOnLine := distanceToBeacon - distanceToTargetRow
		left, right := reading.SensorAt.X-sensorsCoverageOnLine, reading.SensorAt.X+sensorsCoverageOnLine

		if left < minX {
			minX = left
		}
		if right > maxX {
			maxX = right
		}
		intervals = append(intervals, [2]int{left, right})
	}

	slices.SortFunc(intervals, func(a CoveredInterval, b CoveredInterval) bool {
		return a[0] < b[0]
	})

	totalGaps, gapsStarts := 0, []int{}
	atX := intervals[0][0]

	for idx := range intervals {
		start, end := intervals[idx][0], intervals[idx][1]
		if start > atX+1 {
			gapsStarts = append(gapsStarts, atX+1)
			totalGaps += start - (atX + 1)
		}
		atX = lo.Max([]int{atX, end})
	}

	return maxX - minX - totalGaps, gapsStarts
}

type Reading struct {
	SensorAt      Coordinate `parser:"Preamble @@ Separator"`
	ClosestBeacon Coordinate `parser:"@@? Newline?"`
}

func (r *Reading) SensorBeaconDistance() int {
	dx, dy := math.Abs(float64(r.SensorAt.X-r.ClosestBeacon.X)), math.Abs(float64(r.SensorAt.Y-r.ClosestBeacon.Y))
	return int(dx + dy)
}
