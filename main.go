package main

import (
	"database/sql"
	"fmt"
	"runtime"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/smartystreets/go-disruptor"
	"github.com/zhangjunfang/hdata/hdata_core/util"
)

const (
	BufferSize int64 = (1 << 16) * 2048
	BufferMask int64 = BufferSize - 1
)

var ring [BufferSize]*Student

//var ring = [BufferSize]int64{}

type Student struct {
	Id   int64
	Name string
}

func (s *Student) String() string {
	return strconv.Itoa(int(s.Id)) + ":" + s.Name
}

var rows *sql.Rows
var stmt *sql.Stmt

func init() {
	dataSourceName := "root:@tcp(localhost:3306)/zhangboyu?charset=utf8"
	db, err := util.GetConnection("mysql", dataSourceName, 10, 10)
	if err != nil {
		fmt.Println(err)
	}
	rows, err = db.Query(" SELECT u.id ,u.`name` FROM `user` as u ")
	stmt, err = db.Prepare(" INSERT INTO cc(id ,name) VALUES(?,?) ")
}
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	fmt.Println("fsdfsdf", BufferSize)
	controller := disruptor.Configure(BufferSize).WithConsumerGroup(SampleConsumer{}, SampleConsumer{}, SampleConsumer{}).
		BuildShared()
		//Build()

	controller.Start()

	started := time.Now()
	//publish(controller.Writer())
	publishShared(controller.Writer())
	finished := time.Now()
	fmt.Println("耗时：", finished.Sub(started))
	time.Sleep(24 * time.Hour)
	controller.Stop()

}

func publish(writer *disruptor.Writer) {
	sequence := int64(0)
	var id int64
	var name string
	for rows.Next() {
		sequence = writer.Reserve(1) //数据每次递增
		//fmt.Println(sequence)
		rows.Scan(&id, &name)
		ring[sequence&BufferMask] = &Student{
			Id:   id,
			Name: name,
		}
		writer.Commit(sequence-1, sequence) //每生产一个 提交一个
	}
	//writer.Commit(0, sequence) //每生产一批 提交批次
}
func publishShared(writer *disruptor.SharedWriter) {
	sequence := int64(0)
	var id int64
	var name string
	for rows.Next() {
		sequence = writer.Reserve(1) //数据每次递增
		//fmt.Println(sequence)
		rows.Scan(&id, &name)
		ring[sequence&BufferMask] = &Student{
			Id:   id,
			Name: name,
		}
		writer.Commit(sequence-1, sequence) //每生产一个 提交一个
	}
	//writer.Commit(0, sequence) //每生产一批 提交批次
}

type SampleConsumer struct{}

func (this SampleConsumer) Consume(lower, upper int64) {
	for lower <= upper {
		message := ring[lower&BufferMask]
		fmt.Println(lower, "----------", message, "----------", upper)
		stmt.Exec(message.Id, message.Name)
		lower++
	}
}
