package golang_united_school_homework

import (
	"errors"
)

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	if len(b.shapes) < b.shapesCapacity {
		b.shapes = append(b.shapes, shape)
	}
	return errors.New("The box is full")

}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if i > len(b.shapes) {
		return nil, errors.New("Requested index is out of range")
	}
	for index := 0; index < len(b.shapes); index++ {
		if index == i {
			return b.shapes[index], nil
		}
	}
	return nil, errors.New("There is no such index")
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	if i > len(b.shapes) {
		return nil, errors.New("Requested index is out of range")
	}

	for index := 0; index < len(b.shapes); index++ {
		if index == i {
			memo := b.shapes[index]
			b.shapes = append(b.shapes[:index], b.shapes[index+1:]...)
			return memo, nil
		}
	}
	return nil, errors.New("There is no such index")
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	if i > len(b.shapes) {
		return nil, errors.New("Requested index is out of range")
	}
	for index := 0; index < len(b.shapes); index++ {
		if index == i {
			memo := b.shapes[index]
			b.shapes[index] = shape
			return memo, nil
		}
	}
	return nil, errors.New("There is no such index")
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	var sum float64
	for index := 0; index < len(b.shapes); index++ {
		sum += b.shapes[index].CalcPerimeter()
	}
	return sum

}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	var sum float64
	for index := 0; index < len(b.shapes); index++ {
		sum += b.shapes[index].CalcArea()
	}
	return sum
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	var countCircle int
	for index := 0; index < len(b.shapes); index++ {
		switch b.shapes[index].(type) {
		case Circle:
			b.shapes = append(b.shapes[:index], b.shapes[index+1])
			countCircle += 1
		}
	}
	if countCircle == 0 {
		return errors.New("There is no circles in box")
	}
	return nil
}
