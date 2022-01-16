package sample

import (
	"github.com/google/uuid"
	"math/rand"
	"pcbook/pb"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomKeyboardLayout() pb.Keyboard_Layout {
	switch rand.Intn(3) {
	case 1:
		return pb.Keyboard_QWERTY
	case 2:
		return pb.Keyboard_QWERTZ
	default:
		return pb.Keyboard_AZERTY
	}
}

func randomBool() bool {
	return rand.Intn(2) == 1
}

func randomCPUBrand() string {
	return randomStringFromSet("Intel", "AMD")
}

func randomGPUBrand() string {
	return randomStringFromSet("NVIDIA", "AMD")
}

func randomLaptopBrand() string {
	return randomStringFromSet("Apple", "Lenovo")
}

func randomLaptopName(brand string) string {
	switch brand {
	case "Apple":
		return randomStringFromSet("macbook air", "macbook pro")
	case "Lenovo":
		return randomStringFromSet("Thinkpad x1", "Thinkpad p1")
	default:
		return "Other Laptops"

	}
}

func randomGPUName(brand string) string {
	if brand == "NVIDIA" {
		return randomStringFromSet("RTX 2060", "RTX 2070", "GTX 1660", "GTX 1070")

	}
	return randomStringFromSet("RX 590", "RX 580", "RX 5700", "RX Vega")
}
func randomCPUName(brand string) string {
	if brand == "Intel" {
		return randomStringFromSet(
			"Xeon E",
			"Core i9",
			"Core i7",
			"Core i5",
			"Core i3")
	}
	return randomStringFromSet(
		"Ryzen 7",
		"Ryzen 5",
		"Ryzen 3")
}

func randomStringFromSet(a ...string) string {
	n := len(a)
	if n == 0 {
		return ""
	}
	return a[rand.Intn(n)]

}

func randomInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func randomFloat64(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func radomFloat32(min, max float32) float32 {
	return min + rand.Float32()*(max-min)
}

func randomScreenPanel() pb.Screen_Panel {
	if rand.Intn(2) == 1 {
		return pb.Screen_IPS
	}
	return pb.Screen_OLED
}

func randomScreenResolution() *pb.Screen_Resolustion {
	height := uint32(randomInt(1080, 4320))
	width := height * 16 / 9
	resolution := &pb.Screen_Resolustion{
		Height: height,
		Width:  width,
	}
	return resolution
}

func randomID() string {
	return uuid.New().String()
}
