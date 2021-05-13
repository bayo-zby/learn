package interview

type Pool struct {
	queue chan *func()  // 任务队列
	done  chan struct{} // 池开关
}

func NewPool(size int) *Pool {
	// 设置制定容量的通道
	return &Pool{queue: make(chan *func(), size)}
}

func (p *Pool) Work() {
	for {
		select {
		// 任务通道
		case task := <-p.queue:
			go (*task)()
		//开关通道
		case <-p.done:
			return
		}
	}
}

func (p *Pool) Close() {
	p.done <- struct{}{}
}

func (p *Pool) Push(task *func()) {
	p.queue <- task
}
