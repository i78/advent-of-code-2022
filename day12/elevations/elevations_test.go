package elevations

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDistanceToTarget(t *testing.T) {
	elevations, start, destination := NewElevationsMap(testData)

	t.Run("should return 0 when invoked on target itself", func(t *testing.T) {
		dist, found := elevations.DistanceBfsWithCriteria(start, elevations.NodeCoordinateEqualsCriteria(start), elevations.UpwardsElevationConstraint())
		assert.Equal(t, 0, dist)
		assert.True(t, found)
	})

	t.Run("should return 1 when invoked next to target fulfilling constraints", func(t *testing.T) {
		foundCriteria := elevations.NodeCoordinateEqualsCriteria(destination)
		elevationConstraint := elevations.UpwardsElevationConstraint()
		dist, found := elevations.DistanceBfsWithCriteria(destination.Step(West), foundCriteria, elevationConstraint)
		assert.Equal(t, 1, dist)
		assert.True(t, found)
	})

	t.Run("should return 2 when invoked 2 steps away from target fulfilling constraints", func(t *testing.T) {
		foundCriteria := elevations.NodeCoordinateEqualsCriteria(destination)
		elevationConstraint := elevations.UpwardsElevationConstraint()
		dist, found := elevations.DistanceBfsWithCriteria(destination.Step(West).Step(North), foundCriteria, elevationConstraint)
		assert.Equal(t, 2, dist)
		assert.True(t, found)
	})

	t.Run("should return 3 when invoked 3 steps away from target fulfilling constraints", func(t *testing.T) {
		foundCriteria := elevations.NodeCoordinateEqualsCriteria(destination)
		elevationConstraint := elevations.UpwardsElevationConstraint()
		dist, found := elevations.DistanceBfsWithCriteria(destination.Step(West).Step(North).Step(East), foundCriteria, elevationConstraint)
		assert.Equal(t, 3, dist)
		assert.True(t, found)
	})

	t.Run("should report n alternatives for finding route from start to estination", func(t *testing.T) {
		foundCriteria := elevations.NodeCoordinateEqualsCriteria(destination)
		elevationConstraint := elevations.UpwardsElevationConstraint()
		// dist, found := elevations.Distance(destination, destination, 64)
		dist, found := elevations.DistanceBfsWithCriteria(start, foundCriteria, elevationConstraint)
		assert.Equal(t, 31, dist)
		assert.True(t, found)
	})

	t.Run("should report 29 steps back from E to first A when atCriteriaUsed", func(t *testing.T) {
		foundCriteria := elevations.AtEqualsCriteria('a')
		elevationConstraint := elevations.DownwardsElevationConstraint()
		dist, found := elevations.DistanceBfsWithCriteria(destination, foundCriteria, elevationConstraint)
		assert.Equal(t, 29, dist)
		assert.True(t, found)
	})

}
