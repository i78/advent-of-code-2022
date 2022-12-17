package beacons

import "github.com/alecthomas/participle/v2/lexer"

var (
	SensorReadingListLexer = lexer.MustSimple([]lexer.SimpleRule{
		{`Preamble`, `Sensor at `},
		{`XEquals`, `x=`},
		{`YEquals`, `, y=`},
		{`Separator`, `: closest beacon is at `},
		{`Int`, `[-]?\d+`},
		{"Newline", `\n`},
	})
)
