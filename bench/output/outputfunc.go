package output

import (
	"io"
	"time"

	"github.com/nicholasjackson/bench/bench/results"
)

type OutputFunc func(interval time.Duration, r results.ResultSet, w io.Writer)
