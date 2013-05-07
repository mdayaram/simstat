package simstat

import (
	"fmt"
	"math"
	"sync"
)

type DataSet struct {
	door   sync.Mutex
	data   map[int]int
	count  int
	max    int
	min    int
	mode   int
	modeCt int
}

func NewDataSet() *DataSet {
	return &DataSet{
		data: make(map[int]int),
	}
}

func (d *DataSet) Add(i int) {
	d.door.Lock()
	defer d.door.Unlock()
	d.data[i] += 1
	d.count += 1
	if d.max < i {
		d.max = i
	}
	if d.min > i {
		d.min = i
	}
	if d.modeCt < d.data[i] {
		d.modeCt = d.data[i]
		d.mode = i
	}
}

func (d *DataSet) Min() int {
	d.door.Lock()
	defer d.door.Unlock()
	return d.min
}

func (d *DataSet) Max() int {
	d.door.Lock()
	defer d.door.Unlock()
	return d.max
}

func (d *DataSet) Mode() (mode int, appearances int) {
	d.door.Lock()
	defer d.door.Unlock()
	return d.mode, d.modeCt
}

func (d *DataSet) Count() int {
	d.door.Lock()
	defer d.door.Unlock()
	return d.count
}

func (d *DataSet) sum() int {
	sum := 0
	for num, ct := range d.data {
		sum += num * ct
	}
	return sum
}
func (d *DataSet) Sum() int {
	d.door.Lock()
	defer d.door.Unlock()
	return d.sum()
}

func (d *DataSet) avg() float64 {
	sum := 0
	for num, ct := range d.data {
		sum += num * ct
	}
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
		for i := 0; i < ct; i++ {
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
