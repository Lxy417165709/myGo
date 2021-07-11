package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

// 目的：模拟协程池(生产者、处理者、消费者)

// 作业
type Job struct {
	jobId       int
	generatorId int
	content     int
}

// 作业处理结果
type Result struct {
	Job
	workerId int
	result   int
}

// 管理者
type Manager struct {
	jobs         chan Job
	results      chan Result
	workerNum    int
	generatorNum int
	takerNum     int
}

func NewManager(jobChanSize int, resultChanSize int, generatorNum int, workerNum int, takerNum int) (jm *Manager) {
	jm = new(Manager)
	jm.jobs = make(chan Job, jobChanSize)
	jm.results = make(chan Result, resultChanSize)
	jm.workerNum = workerNum
	jm.generatorNum = generatorNum
	jm.takerNum = takerNum
	return
}

func (jm *Manager) Simulate() {
	rand.Seed(time.Now().UnixNano())

	// 产生Job
	for i := 0; i < jm.generatorNum; i++ {
		g := Generator{i}
		go g.GenerateJob(jm.jobs)
	}
	// 处理Job
	for i := 0; i < jm.workerNum; i++ {
		w := Worker{i}
		go w.SolveJob(jm.jobs,jm.results)
	}

	// 输出结果
	for i:=0;i<jm.takerNum;i++{
		t := Taker{i}
		go t.PrintJob(jm.results)
	}

}

type Generator struct {
	generatorId int
}

var jobSeq int64
func (g *Generator) GenerateJob(jobs chan Job) {
	// 每个生产者生成10个作业
	for i := 0; i < 10; i++ {

		// 原子操作，防止并发不一致问题
		jobs <- Job{int(atomic.AddInt64(&jobSeq,1)), g.generatorId, rand.Int() % 1000}
		time.Sleep(10 * time.Millisecond)
	}
}

type Worker struct {
	workerId int
}

func (w *Worker) SolveJob(jobs chan Job, results chan Result) {
	for job := range jobs {
		results <- Result{job, w.workerId, 2 * job.content}
		time.Sleep(500 * time.Millisecond)
	}
}

type Taker struct {
	takerId int
}

func (t *Taker) PrintJob(results chan Result) {
	for result := range results {
		fmt.Printf("jobId: %d  generatorId: %d workerId: %d takerId: %d *** %d --> %d\n",
			result.jobId,
			result.generatorId,
			result.workerId,
			t.takerId,
			result.content,
			result.result,
		)
	}
}

func main() {
	manager := NewManager(10, 10, 10,2,5)
	manager.Simulate()
	time.Sleep(200 * time.Second)
}
