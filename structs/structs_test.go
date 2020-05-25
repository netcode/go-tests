package structs

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	expected := 40.0
	if got != expected {
		t.Errorf("Got %.2f expected %.2f", got, expected)
	}
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		shape    Shape
		expected float64
	}{
		{Rectangle{5.0, 7.0}, 35.0},
		{Circle{15.00}, 706.8583470577034},
		{Triangle{12, 6}, 36.0},
	}

	for _, areaTest := range areaTests {
		got := GetShapeArea(areaTest.shape)
		if got != areaTest.expected {
			t.Errorf("Shape %#v: Got %g expected %g", areaTest.shape, got, areaTest.expected)
		}
	}

}
