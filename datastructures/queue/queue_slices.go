package queue

//import (
//  "fmt"
//)

type QueueS struct {
  QueueS []interface{}
}

func NewQueueS() *QueueS {
  return &QueueS{QueueS: make([]interface{}, 0)}
}

func (self *QueueS) enQueueS(element interface{}) {
  (*self).QueueS = append((*self).QueueS, element)
}

func (self *QueueS) deQueueS() interface{} {
  ret := (*self).QueueS[0]
  (*self).QueueS = (*self).QueueS[1:]
  return ret
}

/*
func main() {
  QueueS := NewQueueS()

  for i := 0; i < 10; i += 1 {
    QueueS.enQueueS(i)
  }

  for i := 0; i < 5; i += 1 {
    fmt.Println(fmt.Sprintf("%v", QueueS.deQueueS()))
  }
  fmt.Println(QueueS)
}
*/
