package utils

type RenderStruct struct {
	S int `json:"s"`
	M string `json:"m"`
	D interface{} `json:"d"`
}

func GetSuccessRender(list interface{}) (*RenderStruct) {
	render := new(RenderStruct)
	render.S = 200
	render.M = "success"
	if list == nil {
		list = []string{}
	}
	render.D = list
	return render
}

func GetErrorRender(message string, code int) (*RenderStruct) {
	render := new(RenderStruct)
	render.S = code
	render.M = message
	render.D = []string{}
	return render
}

func GetClientErrRender() (*RenderStruct) {
	render := new(RenderStruct)
	render.S = 400
	render.M = "input error"
	render.D = []string{}
	return render
}

func GetServerErrRender() (*RenderStruct) {
	render := new(RenderStruct)
	render.S = 500
	render.M = "server error"
	render.D = []string{}
	return render
}
