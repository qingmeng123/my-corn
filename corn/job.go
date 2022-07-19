/*******
* @Author:qingmeng
* @Description:
* @File:job
* @Date:2022/7/18
 */

package corn

import (
	"errors"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type timeUnit int

const (
	seconds timeUnit = iota + 1
	minutes
	hours
	days
	weeks
)

var (
	ErrTimeFormat           = errors.New("时间格式错误")
	ErrNotAFunction         = errors.New("只有功能可以安排到作业队列中")
	ErrParamsNotAdapted     = errors.New("参数的数量不适应")
	ErrPeriodNotSpecified   = errors.New("未指定的工作时间")
	ErrParameterCannotBeNil = errors.New("nil 参数不能与反射一起使用")
)

// Job 他job的封装好喜欢，借鉴借鉴源码吧
type Job struct {
	interval uint64	//运行之间的暂停间隔单位
	jobFunc string	//需要执行的类名
	uint timeUnit	//时间单位
	atTime time.Duration	//时间间隔
	lastRun time.Time 	//上次执行的时间
	nextRun time.Time	//下次执行的时间
	err      error                    // 和job有关的error
	funcs    map[string]interface{}   // 函数任务存储的映射
	params 	interface{} // 方法参数
}

func NewJob(interval uint64) *Job {
	return &Job{
		interval: interval,
		uint: seconds,	//默认为秒
		lastRun:  time.Unix(0, 0),
		nextRun:  time.Unix(0, 0),
		funcs:    make(map[string]interface{}),
	}
}

// Second 将单位设置为秒,并计算出atTime
func (j *Job) Second() *Job {
	j.uint=seconds
	j.atTime= time.Duration(int(j.uint) * int(j.interval))*time.Second

	return j
}

//添加具体时间
func (j *Job) At(t string) *Job {
	hour, min, sec, err := formatTime(t)
	if err != nil {
		j.err = ErrTimeFormat
		return j
	}
	// 计算出atTime
	j.atTime = time.Duration(hour)*time.Hour + time.Duration(min)*time.Minute + time.Duration(sec)*time.Second
	return j
}

//执行任务
//func (j *Job) Do() {
//	j.lastRun=time.Now()
//	j.nextRun=time.Now().Add(j.atTime)
//	j.jobFunc="hello"
//}

// Do 指定每次作业运行时应调用的 jobFunc
func (j *Job) Do(jobFun interface{}, params ...interface{}) error {
	if j.err != nil {
		return j.err
	}

	//反射获取方法
	typ := reflect.TypeOf(jobFun)
	if typ.Kind() != reflect.Func {
		return ErrNotAFunction
	}
	fname := getFunctionName(jobFun)
	j.funcs[fname] = jobFun
	j.params = params
	j.jobFunc = fname

	j.lastRun=time.Now()
	j.nextRun=time.Now().Add(j.atTime)
	return nil
}

// 对于给定的函数 fn，获取函数的名称
func getFunctionName(fn interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
}


//标准化时间
func formatTime(t string) (hour, min, sec int, err error) {
	ts := strings.Split(t, ":")
	if len(ts) < 2 || len(ts) > 3 {
		return 0, 0, 0, ErrTimeFormat
	}

	if hour, err = strconv.Atoi(ts[0]); err != nil {
		return 0, 0, 0, err
	}
	if min, err = strconv.Atoi(ts[1]); err != nil {
		return 0, 0, 0, err
	}
	if len(ts) == 3 {
		if sec, err = strconv.Atoi(ts[2]); err != nil {
			return 0, 0, 0, err
		}
	}

	if hour < 0 || hour > 23 || min < 0 || min > 59 || sec < 0 || sec > 59 {
		return 0, 0, 0, ErrTimeFormat
	}

	return hour, min, sec, nil
}

