package sections

import (
	"math"
	"strconv"
	"strings"
)

const TelegramSeparator = "\n"
const RecordSeparator = ","
const RangeSeparator = "-"

type SectionRange struct {
	From int
	To   int
}
type SectionRangePair [2]SectionRange

func NewSectionList(raw string) (result []SectionRangePair) {
	for _, telegram := range strings.Split(raw, TelegramSeparator) {
		records := strings.Split(telegram, RecordSeparator)

		newRange := func(record string) SectionRange {
			from, _ := strconv.ParseInt(strings.Split(record, RangeSeparator)[0], 10, 32)
			to, _ := strconv.ParseInt(strings.Split(record, RangeSeparator)[1], 10, 32)
			return SectionRange{From: int(from), To: int(to)}
		}

		result = append(result, SectionRangePair{
			newRange(records[0]),
			newRange(records[1]),
		})
	}
	return
}

type intervalStrategy = func(l1 int, r1 int, l2 int, r2 int) bool

func FullyIncludes(l1 int, r1 int, l2 int, r2 int) bool {
	return l1 <= l2 && r1 >= r2 || l2 <= l1 && r2 >= r1
}

func Overlaps(l1 int, r1 int, l2 int, r2 int) bool {
	return math.Max(float64(l1), float64(l2)) <= math.Min(float64(r1), float64(r2)) ||
		math.Max(float64(l2), float64(l1)) <= math.Min(float64(r2), float64(r1))
}

func (p *SectionRangePair) Overlaps(strategy intervalStrategy) bool {
	return strategy(p[0].From, p[0].To, p[1].From, p[1].To)
}
