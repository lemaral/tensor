package addhoctasks

import (
	"fmt"
	"time"
)

// addHocTaskPool is object to store
// task queue
type addHocTaskPool struct {
	queue    []*task
	register chan *task
	running  *task
}

var pool = addHocTaskPool{
	queue:    make([]*task, 0),
	register: make(chan *task),
	running:  nil,
}

func (p *addHocTaskPool) run() {
	ticker := time.NewTicker(10 * time.Second)

	defer func() {
		ticker.Stop()
	}()

	for {
		select {
		case task := <-p.register:
			if p.running == nil {
				go task.run()
				continue
			}

			p.queue = append(p.queue, task)
		case <-ticker.C:
			if len(p.queue) == 0 || p.running != nil {
				continue
			}

			fmt.Println("Running a task.")
			go pool.queue[0].run()
			pool.queue = pool.queue[1:]
		}
	}
}

func StartRunner() {
	pool.run()
}
