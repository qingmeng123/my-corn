/*******
* @Author:qingmeng
* @Description:
* @File:scheduler
* @Date:2022/7/19
 */

package corn

import (
	"fmt"
	"reflect"
	"time"
)

type scheduler struct {
	node *Node	//头结点
	state bool	//true为启动态
}

var (
	defaultScheduler = NewScheduler()
)

func NewScheduler()*scheduler  {
	return &scheduler{
		node: initList(),
		state: true,
	}
}

//将工作的job放入管道
func (s *scheduler) putWorkJob(jobChan chan<- Job)  {
	for{
		job := getJobByTime(s.node)

		//未暂停时才会把工作的job放进去
		if s.state{
			jobChan<-*job
		}
	}
}

//放入job到scheduler
func (s *scheduler) AddJob(val Job)  {
	val.lastRun=time.Now()
	val.nextRun=time.Now().Add(val.atTime)
	insertList(&(s.node),&val)
}

//启动scheduler
func (s *scheduler) Start() {
	jobChan:=make(chan Job,1000)
	go s.putWorkJob(jobChan)
	for {
		job:=<-jobChan
		if job.jobFunc!=""{
			_, err := call(job.funcs,job.jobFunc)
			if err!=nil{
				fmt.Println("get method err:",err)
		}
		}
	}
}

//调用反射出的方法
func call(m map[string]interface{}, name string, params ...interface{}) ([]reflect.Value, error) {
	f := reflect.ValueOf(m[name])
	if len(params) != f.Type().NumIn() {
		fmt.Println(len(params))
		fmt.Println(f.Type().NumIn())
		return nil, ErrParamsNotAdapted
	}
	in := make([]reflect.Value, len(params))
	for k, v := range params {
		in[k] = reflect.ValueOf(v)
	}
	return f.Call(in), nil
}


// 调度周期
func (s *scheduler) Every(interval uint64) *Job {
	s.state=true
	job := NewJob(interval)
	insertList(&(s.node),job)
	return job
}

//暂停
func (s *scheduler) Pause()  {
	s.state=false
}

//取消暂停
func (s *scheduler) On() {
	s.state=true
}

//重置
func (s *scheduler) Clear() {
	cleanList(s.node)
}

//删除一个定时任务
func (s *scheduler) Remove(val string) {
	loc:=locateList(s.node,val)
	delList(&s.node,loc)
}
