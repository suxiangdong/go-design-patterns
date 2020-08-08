package state

import (
	"testing"
)

func TestWorker_Wood(t *testing.T) {
	worker := NewWorker()
	if r := worker.Wood(); r != "星期一综合症\n" {
		t.Fatal("Monday test fail")
	}
	worker.Next()
	if r := worker.Wood(); r != "周日怎么还那么远啊\n" {
		t.Fatal("Tuesday test fail")
	}
	worker.Next()
	if r := worker.Wood(); r != "前不着村后不着店，崩溃\n" {
		t.Fatal("Wednesday test fail")
	}
	worker.Next()
	if r := worker.Wood(); r != "再忍忍\n" {
		t.Fatal("Thursday test fail")
	}
	worker.Next()
	if r := worker.Wood(); r != "最后一天，高兴\n" {
		t.Fatal("Friday test fail")
	}
	worker.Next()
	if r := worker.Wood(); r != "哈哈，周末来了\n" {
		t.Fatal("Saturday test fail")
	}
	worker.Next()
	if r := worker.Wood(); r != "万恶的周一又要来了\n" {
		t.Fatal("Sunday test fail")
	}
}

func TestWorker_SetDay(t *testing.T) {
	worker := NewWorker()
	if r := worker.Wood(); r != "星期一综合症\n" {
		t.Fatal("Monday test fail")
	}
	worker.SetDay(&Sunday{})
	if r := worker.Wood(); r != "万恶的周一又要来了\n" {
		t.Fatal("Sunday test fail")
	}
}
