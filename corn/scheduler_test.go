/*******
* @Author:qingmeng
* @Description:
* @File:scheduler_test
* @Date:2022/7/19
 */

package corn

import (
	"testing"
	"time"
)

func TestScheduler_insert(t *testing.T) {
	s:=scheduler{node: initList()}
	job:=Job{jobFunc: "hello",atTime: 1*time.Second,lastRun: time.Now()}
	s.AddJob(job)
	get:= s.node.pNext.pNext.job.jobFunc
	if get!="hello"{
		t.Errorf("get %s,want %s",get,"hello")
	}
}

func TestScheduler_Start(t *testing.T) {
	s:=NewScheduler()
	job:=Job{jobFunc: "hello",atTime: 1*time.Second,lastRun: time.Now()}
	job1:=Job{jobFunc: "hello1",atTime: 3*time.Second,lastRun: time.Now()}
	s.AddJob(job)
	s.AddJob(job1)
	s.Start()

}