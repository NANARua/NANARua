package structure

import "testing"

//我们可以使用保留字 struct 来定义自己的类型。一个通过 struct 定义出来的类型是一些已命名的域的集合，这些域用来保存数据。
//定义我们自己的类型 Rectangle，它可以封装长方形的信息
type Rectangle struct {
	Width  float64
	Height float64
}

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0
	if got != want {
		t.Errorf("got %2.f want %2.f", got, want)
	}
}

func TestArea(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Area(rectangle)
	want := 100.0
	if got != want {
		t.Errorf("got %2.f want %2.f", got, want)
	}
}
