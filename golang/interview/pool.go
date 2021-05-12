package interview

import (
	"errors"
)

// 任务
type task struct {
	f    func() // 任务内容
	next *task  // 下一个任务
}

// 任务队列
type queue struct {
	len  int   // 当前任务队列数
	task *task // 链表头,头部弹出
	tail *task // 链表尾,尾部插入
}

type Pool struct {
	// 正在执行的任务队列
	// 容量为cap,可以总限制池同时运行的协程数
	taskWork chan struct{}
	queue    *queue        // 等待任务队列
	done     chan struct{} // 池开关
}

/**
 * 新建协程池
 */
func NewPool(cap int) *Pool {
	return &Pool{
		// 缓冲
		taskWork: make(chan struct{}, cap),
		queue:    &queue{},
		// 无缓冲
		done: make(chan struct{}),
	}
}

/**
 * 待执行任务数
 * 线程不安全
 */
func (q *queue) Len() int {
	return q.len
}

/**
 * 关闭协程池
 */
func (p *Pool) Close() {
	p.done <- struct{}{}
}

/**
 * @brief 尾插法，往任务队列中加入任务
 * @param f 任务函数
 *
 */
func (q *queue) Push(f func()) error {
	if q.len > 1024 {
		return errors.New("queue is full")
	}
	newTask := &task{f: f}
	if q.tail == nil {
		q.tail = newTask
	} else {
		q.tail.next = newTask
	}
	q.len++
	return nil
}

/**
 * @brief 尾弹法，弹出任务
 * @ret func() 任务函数内容
 */
func (q *queue) Pop() func() {
	task := q.task.f
	q.task = q.tail.next
	// 最后一节内容取出
	if q.task == nil {
		q.tail = nil
	}
	q.len--
	return task
}

/**
 * 打开任务开关,循环执行
 */
func (p *Pool) Work() {
	for {
		select {
		// 监控开关
		case <-p.done:
			return
		// 任务中无法插入新的执行任务时，会阻塞
		case p.taskWork <- struct{}{}:
			// 建立协程
			go func() {
				// 执行任务
				p.queue.Pop()()
				// 函数运行完毕取出任务队列占用
				<-p.taskWork
			}()

		default:
			p.taskWork <- struct{}{}
		}
	}
}
