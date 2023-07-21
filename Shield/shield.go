package Shield

type resMeta struct {
	err error
}

type execMeta struct {
	response chan resMeta
	f        func() error
}

type Shield struct {
	execChan chan execMeta
}

func (s *Shield) handle() {
	for {
		exec := <-s.execChan
		err := exec.f()
		exec.response <- resMeta{
			err: err,
		}
	}
}

func (s *Shield) Protect(f func() error) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()
	conn := make(chan resMeta)
	s.execChan <- execMeta{
		f:        f,
		response: conn,
	}
	res := <-conn
	close(conn)
	err = res.err
	return
}

func (s *Shield) Close() {
	close(s.execChan)
}

func NewShield() *Shield {
	s := &Shield{
		execChan: make(chan execMeta),
	}
	go s.handle()
	return s
}