package structure

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

func Area(width, height float64) float64 {
	return width * height
}
