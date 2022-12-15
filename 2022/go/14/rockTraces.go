package day14

type rockTraces []*rockTrace

func (traces rockTraces) isOn(x, y int) bool {
	for _, trace := range traces {
		if trace.isOn(x, y) {
			return true
		}
	}
	return false
}

func (traces rockTraces) isBelow(y int) bool {
	for _, trace := range traces {
		if trace.isBelow(y) == false {
			return false
		}
	}

	return true
}
