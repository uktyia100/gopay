package douyin

import (
	"encoding/json"
	"testing"

	"github.com/go-pay/gopay"
)

func TestClient_CensorImage(t *testing.T) {
	ast, err := client.GetAccessToken(ctx)
	if err != nil {
		t.Log(err)
	}
	bm := make(gopay.BodyMap)
	bm.Set("access_token", ast.AccessToken)
	bm.Set("image", "http://xxx.jpg")
	resp, err := client.CensorImage(ctx, bm)
	if err != nil {
		t.Logf("fail：%v", err)
		return
	}
	t.Logf("resp:%+v, \nerr:%s", resp, err)
}

func TestClient_TextAnti(t *testing.T) {
	ast, err := client.GetAccessToken(ctx)
	if err != nil {
		t.Log(err)
	}
	bm := make(gopay.BodyMap)
	bm.SetBodyMap("headers", func(bm gopay.BodyMap) {
		bm.Set("X-Token", ast.AccessToken)
	})
	t.Logf("ast:%s", ast.AccessToken)
	type Task struct {
		Content string `json:"content"`
	}
	var tasks []*Task
	task := &Task{
		Content: "习近平",
	}
	task2 := &Task{
		Content: "哈哈",
	}
	task3 := &Task{
		Content: "操你妈",
	}
	tasks = append(tasks, task, task2, task3)
	bm.Set("tasks", tasks)
	resp, err := client.TextAnti(ctx, bm)
	if err != nil {
		t.Logf("fail：%v", err)
		return
	}

	rst, _ := json.Marshal(resp)
	t.Logf("resp:%+v, \nerr:%s\n%+v", string(rst), err, resp)
}
