package simstat

import (
	"fmt"
	"math"
	"sync"
)

type DataSet struct {
	door   sync.Mutex
	data   map[int64]int64
	count  int64
	max    int64
	maxed  bool
	min    int64
	minned bool
	mode   int64
	modeCt int64
}

func NewDataSet() *DataSet {
	return &DataSet{
		data: make(map[int64]int64),
	}
}

func (d *DataSet) Add(i int64) {
	d.door.Lock()
	defer d.door.Unlock()
	d.data[i] += 1
	d.count += 1
	if !d.maxed || d.max < i {
		d.max = i
		d.maxed = true
	}
	if !d.minned || d.min > i {
		d.min = i
		d.minned = true
	}
	if d.modeCt < d.data[i] {
		d.modeCt = d.data[i]
		d.mode = i
	}
}

func (d *DataSet) Min() int64 {
	d.door.Lock()
	defer d.door.Unlock()
	return d.min
}

func (d *DataSet) Max() int64 {
	d.door.Lock()
	defer d.door.Unlock()
	return d.max
}

func (d *DataSet) Mode() (mode int64, appearances int64) {
	d.door.Lock()
	defer d.door.Unlock()
	return d.mode, d.modeCt
}

func (d *DataSet) Count() int64 {
	d.door.Lock()
	defer d.door.Unlock()
	return d.count
}

func (d *DataSet) sum() int64 {
	var sum int64 = 0
	for num, ct := range d.data {
		sum += num * ct
	}
	return sum
}
func (d *DataSet) Sum() int64 {
	d.door.Lock()
	defer d.door.Unlock()
	return d.sum()
}

func (d *DataSet) avg() float64 {
	return float64(d.sum()) / float64(d.count)
}
func (d *DataSet) Avg() float64 {
	d.door.Lock()
	defer d.door.Unlock()
	return d.avg()
}

func (d *DataSet) variance() float64 {
	avg := d.avg()
	var sumf float64 = 0
	for num, ct := range d.data {
		for i := int64(0); i < ct; i++ {
			sumf += (float64(num) - avg) * (float64(num) - avg)
		}
	}
	return sumf / float64(d.count)
}
func (d *DataSet) Variance() float64 {
	d.door.Lock()
	defer d.door.Unlock()
	return d.variance()
}

func (d *DataSet) stdDev() float64 {
	return math.Sqrt(d.variance())
}
func (d *DataSet) StdDev() float64 {
	d.door.Lock()
	defer d.door.Unlock()
	return d.stdDev()
}

func (d *DataSet) String() string {
	d.door.Lock()
	defer d.door.Unlock()
	s := fmt.Sprintf("Total-Records:\t%d\n", d.count)
	s += fmt.Sprintf("Minimum:\t%d\n", d.min)
	s += fmt.Sprintf("Average:\t%f\n", d.avg())
	s += fmt.Sprintf("Maximum:\t%d\n", d.max)
	s += fmt.Sprintf("Std-Dev:\t%f\n", d.stdDev())
	return s
}
