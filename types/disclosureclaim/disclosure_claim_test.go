package disclosureclaim

import (
	"fmt"
	"strconv"
	"testing"
)

func Test_disclosureClaim(t *testing.T) {
	v := 3.1415926535
	t.Log(strconv.FormatFloat(v, 'f', -1, 32)) //float32s2 := strconv.FormatFloat(v, 'E', -1, 64)//float64
	t.Log(strconv.FormatFloat(v, 'g', -1, 32)) //float32s2 := strconv.FormatFloat(v, 'E', -1, 64)//float64
	t.Log(strconv.FormatFloat(v, 'G', -1, 32)) //float32s2 := strconv.FormatFloat(v, 'E', -1, 64)//float64
	t.Log(strconv.FormatFloat(v, 'x', -1, 32)) //float32s2 := strconv.FormatFloat(v, 'E', -1, 64)//float64
	t.Log(strconv.FormatFloat(v, 'X', -1, 32)) //float32s2 := strconv.FormatFloat(v, 'E', -1, 64)//float64

	t.Log(fmt.Sprintf("%f", v))
	t.Log(fmt.Sprint(v))
}
