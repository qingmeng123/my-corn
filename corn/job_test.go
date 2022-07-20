/*******
* @Author:qingmeng
* @Description:
* @File:job_test
* @Date:2022/7/19
 */

package corn

import (
	"fmt"
	"testing"
)

func TestJob_Do(t *testing.T) {
	scheduler := NewScheduler()
	scheduler.Every(2).Second().Do(func() { fmt.Println("hello") })
	fmt.Println("job", scheduler.node.pNext.pNext.job)
}
