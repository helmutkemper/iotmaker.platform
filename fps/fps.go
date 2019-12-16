package fps

var fps = 60

func Set(value int) {
	fps = value
}

func Get() int {
	return fps
}
