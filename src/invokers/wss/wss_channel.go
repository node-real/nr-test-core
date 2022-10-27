package wss

type WssChannel struct {
	Wsc chan string
}

func (s *WssChannel) BeforeTest(loop func()) {
	s.NewInstance()
	s.Product(loop)
	s.Garbage()
}

func (s *WssChannel) NewInstance() {
	s.Wsc = make(chan string, 1)
}

func (s *WssChannel) Product(loop func()) {
	go func() {
		loop()
	}()
}

// throw signal to avoid block
func (s *WssChannel) Garbage() {
	//thrown no receiver message
	go func() {
		for {
			select {
			case <-s.Wsc:
			default:
			}
		}
	}()
}

func (s *WssChannel) GetOneSignal() string {
	//thrown no receiver message
	var j string
	for i := range s.Wsc {
		j = i
		break
	}
	return j
}

func (s *WssChannel) GetNSignal(n int) string {
	//thrown no receiver message
	var j string
	for i := range s.Wsc {
		j = i
		n--
		if n <= 0 {
			break
		}
	}
	return j
}
