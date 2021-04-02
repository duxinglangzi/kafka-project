package main

import (
	"com/conor/kafka/consumer"
	"com/conor/kafka/producer"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var (
	// HOSTS  = []string{"39.104.21.XXX:9092","39.104.21.XXX:9091"}
	HOSTS = []string{"39.104.21.126:9092"}
)

const (
	TOPIC = "filebeats-topic"
	KEY   = "test-kafka-key"
)

func main() {

	// 生产者连接kafka
	producer.SetKafkaSetting(HOSTS)

	// 消费者连接kafka
	consumer.InitMQ(HOSTS)

	// 启动一个监听，假装一个请求数据接口
	mux := http.NewServeMux()
	mux.HandleFunc("/sendMessage", sendMessage)
	// 设置监听的端口
	err := http.ListenAndServe(":8091", mux)
	if err != nil {
		fmt.Println("ListenAndServe: ", err)
	}

	fmt.Println(" Started success ........")

}

func sendMessage(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // 解析url传递的参数，对于POST则解析响应包的主体（request body）
	readAll, _ := ioutil.ReadAll(r.Body)
	mapJSON := make(map[string]interface{})
	json.Unmarshal(readAll, &mapJSON)
	// 获取 url 连接中的参数
	for k, v := range r.Form {
		mapJSON[k] = strings.Join(v, "")
	}
	marshal, _ := json.Marshal(mapJSON)
	params := string(marshal)
	pushKafkaMessage(TOPIC, KEY, params)
	w.WriteHeader(200)
	io.WriteString(w, `{"status":"ok"}`)
}

func pushKafkaMessage(topic, key string, params string) {
	kaf, err := producer.GetKafkaProducer()
	if err != nil {
		log.Printf("[kafka_producer_message] [ %v ]  请求连接异常， %s \n", key, err.Error())
		return
	}
	_, offset, _, err := kaf.SendByHashPartition(topic, key, params)
	if err != nil {
		log.Printf("[kafka_push_message] [ %v ]  发生请求异常， %s \n", key, err.Error())
		return
	}
	log.Printf("[kafka_push_message] [ %v ] 发送请求成功，当前offset: %v \n", topic, offset)
}
