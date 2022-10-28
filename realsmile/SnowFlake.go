package realsmile

import (
	"fmt"
	"time"
)

// SnowFlake
/**
 *  * Twitter_Snowflake
 *  * SnowFlake的结构如下(每部分用-分开):
 *  * 0 - 0000000000 0000000000 0000000000 0000000000 0 - 00000 - 00000 - 000000000000
 *  * 1位标识，由于long基本类型在Java中是带符号的，最高位是符号位，正数是0，负数是1，所以id一般是正数，最高位是0
 *  * 41位时间戳(毫秒级)，注意，41位时间戳不是存储当前时间的时间戳，而是存储时间戳的差值（当前时间戳 - 开始时间戳)
 *  * 得到的值），这里的的开始时间戳，一般是我们的id生成器开始使用的时间，由我们程序来指定的（如下下面程序IdWorker类的startTime属性）。41位的时间戳，可以使用69年，年T = (1L << 41) / (1000L * 60 * 60 * 24 * 365) = 69<br>
 *  * 10位的数据机器位，可以部署在1024个节点，包括5位datacenterId和5位workerId
 *  * 12位序列，毫秒内的计数，12位的计数顺序号支持每个节点每毫秒(同一机器，同一时间戳)产生4096个ID序号
 *  * 加起来刚好64位，为一个Long型。
 *  * SnowFlake的优点是，整体上按照时间自增排序，并且整个分布式系统内不会产生ID碰撞(由数据中心ID和机器ID作区分)，并且效率较高，经测试，SnowFlake每秒能够产生26万ID左右。
 *  *
 *
 * @author Real 57916127@qq.com
 * @date   2022-10-08
 **/
type SnowFlake struct {
	workerId      int64
	datacenterId  int64
	lastSequence  int64
	lastTimestamp int64
}

var (
	Snow SnowFlake
)

func init() {
	Snow.Init(31, 31)
}
func (sf *SnowFlake) Init(datacenterId, workerId int64) {
	//机器ID  2进制5位  32位减掉1位 31个
	var workerIdMax int64 = 31
	if workerId <= workerIdMax && workerId >= 0 {
		var datacenterIdMax int64 = 31
		if datacenterId <= datacenterIdMax && datacenterId >= 0 {
			sf.workerId = workerId
			sf.datacenterId = datacenterId
		} else {
			panic(fmt.Sprintf("datacenter Id can't be greater than %d or less than 0", datacenterIdMax))
		}
	} else {
		panic(fmt.Sprintf("worker Id can't be greater than %d or less than 0", workerId))
	}
	sf.lastTimestamp = -1
	sf.lastSequence = 0
}
func (sf *SnowFlake) NextId() int64 {
	timestamp := time.Now().UnixMilli()
	if timestamp < sf.lastTimestamp {
		panic(fmt.Sprintf("Clock moved backwards.  Refusing to generate id for %v milliseconds", sf.lastTimestamp-timestamp))
	} else {
		if sf.lastTimestamp == timestamp {
			var sequenceMax int64 = 4095
			sf.lastSequence = (sf.lastSequence + 1) & sequenceMax
			if sf.lastSequence == 0 {
				timestamp = sf.nextMillis()
			}
		} else {
			sf.lastSequence = 0
		}
		sf.lastTimestamp = timestamp
		/**
		 * 开始时间截 2022年1月1日
		 */
		var startTimeStamp int64 = 1640966400000
		/**
		 * 先左移动12位，空出12位序列
		 *
		 **/
		var workerIdShift int64 = 12
		/**
		 * 左移17位，空出序列和5位工作id
		 *
		 **/
		var datacenterIdShift int64 = 17
		/**
		 * 左移22位，空出12位序列、5为工作id和5位数据中心id
		 *
		 **/
		var timestampLeftShift int64 = 22
		return ((timestamp - startTimeStamp) << timestampLeftShift) | (sf.datacenterId << datacenterIdShift) | (sf.workerId << workerIdShift) | sf.lastSequence
	}
}
func (sf *SnowFlake) nextMillis() int64 {
	timestamp := time.Now().UnixMilli()
	for timestamp <= sf.lastTimestamp {
		timestamp = time.Now().UnixMilli()
	}
	return timestamp
}
