/*******
* @Author:qingmeng
* @Description:
* @File:circleList
* @Date:2022/7/18
 */

package corn

import (
	"fmt"
	"time"
)

type Node struct{
	job   *Job
	pNext *Node
}

func initList() *Node{
	pHead:=new(Node)
	pHead.pNext=pHead
	return pHead   //返回头指针
}

//清空循环链表
func cleanList(list *Node){
	if isempty(list) {
		return
	}
	phead:=list.pNext   //头结点
	p:=list.pNext.pNext //第一个结点
	q:=p
	for  p!=list.pNext  {
		q=p.pNext
		p=nil
		p=q
	}
	phead.pNext=phead
}
//插入结点并大致排序
func insertList(list **Node,val *Job) {
	j:=0
	length:=listLength(*list)
	p,q:=(*list).pNext,(*list).pNext    //头结点(头结点不含值)
	for  j<length&&p.pNext.job.uint<val.uint  {	//按时间单位uint排序
		p=p.pNext
		j++
	}
	pnew:=new(Node)
	pnew.job =val
	//插到比val的uint大的结点的前一个
	pnew.pNext=p.pNext
	p.pNext=pnew
	if pnew.pNext==q {
		*list=pnew
	}
}

func delList(list **Node,loc int)   {
	j:=1
	p,q:=(*list).pNext,(*list).pNext    //头结点
	//查找loc-1结点
	for  j<loc  {
		p=p.pNext
		j++
	}
	cur:=p.pNext
	p.pNext=cur.pNext
	if p.pNext==q {
		*list=p
	}
	cur=nil
}

//通过方法名返回需要查找的结点位置
func locateList(list *Node,val string)int{
	q:=list.pNext.pNext //第一个结点
	loc:=0
	for q!=list.pNext {
		loc ++
		if q.job.jobFunc ==val {
			break
		}
		q=q.pNext
	}
	return loc
}

func traverse(list *Node) {
	if isempty(list) {
		fmt.Println("空链表")
		return
	}
	fmt.Println("链表内容如下：")
	p:=list.pNext.pNext //第一个结点
	for  p!=list.pNext {
		fmt.Printf("%v",p.job)
		p=p.pNext
	}
	fmt.Println()
}

//无限循环寻找一个当前时间的结点,并修改job值
func getJobByTime(list *Node) *Job {
	p:=list.pNext.pNext //第一个结点
	for  !isempty(p) {	//头结点中的job会有空指针异常
		if p.job==nil{
			p=p.pNext
		}
		//!!!!淦，这个地方让我卡了好久。比较时间用equal一直比不出来,还以为是传递问题,结果时间最小单位的问题，好久没用的嗯是容易错。
		if time.Now().Second()==p.job.nextRun.Second(){
			p.job.lastRun=time.Now()
			p.job.nextRun=time.Now().Add(p.job.atTime)
			return p.job
		}
		p=p.pNext
	}

	return new(Job)
}

func isempty(list *Node)bool {
	if list.pNext==list {
		return true
	}else {
		return false
	}
}
func listLength(list *Node) int {
	if isempty(list) {
		return 0
	}
	var len =0
	p:=list.pNext.pNext  //第一个结点
	for p!=list.pNext {
		len++
		p=p.pNext
	}
	return len
}