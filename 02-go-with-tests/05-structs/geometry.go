package main

func Perimeter(rectangle Rectangle) float64 {
	return (rectangle.Width + rectangle.Height) * 2.0
}

func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}
