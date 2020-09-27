package response

import "github.com/gogf/gf/net/ghttp"

type JsonResponse struct {
	Code    int         `json:"code"` // 错误码(0:成功，1:失败，>1:错误码)
	Message string      `json:"msg"`  // 提示信息
	Data    interface{} `json:"data"` // 返回数据（由业务接口定义）
}

// 标准化响应JSON结果
func Json(r *ghttp.Request, code int, message string, data ...interface{}) {
	responseData := interface{}(nil)
	if len(data) > 0 {
		responseData = data[0]
	}
	r.Response.WriteJson(JsonResponse{
		Code:    code,
		Message: message,
		Data:    responseData,
	})
}

// 返回JSON数据并退出当前HTTP执行函数（HttpHandler）
func JsonExit(r *ghttp.Request, err int, message string, data ...interface{}) {
	Json(r, err, message, data)
	r.Exit()
}
