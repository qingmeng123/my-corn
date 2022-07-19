/*******
* @Author:qingmeng
* @Description:
* @File:circleList_test
* @Date:2022/7/19
 */

package corn

import (
	"fmt"
	"testing"
	"time"
)

var jobs =[]struct{
	name string
	uint timeUnit
	out string
}{
	{"one",seconds,"one"},
	{"three",minutes,"three"},
	{"two",seconds,"two"},

}


func TestInsertList(t *testing.T) {
	list:=initList()
	p:=list.pNext
	for _, test := range jobs {
		job:=Job{jobFunc: test.name,uint: test.uint}
		insertList(&list,&job)
		get:=job.jobFunc
		if get!=test.out{
			t.Errorf("get %s,want %s",get,test.out)
		}
		p=p.pNext
	}

}

func TestGetJobByTime(t *testing.T) {
	list:=initList()
	job1:=Job{jobFunc: "hello1",atTime:time.Second,nextRun: time.Now().Add(time.Second)}
	job2:=Job{jobFunc: "hello2",atTime:5*time.Second,nextRun: time.Now().Add(2*time.Second)}
	insertList(&list,&job1)
	insertList(&list,&job2)
	for{
		fmt.Println(getJobByTime(list).jobFunc)

	}

}
