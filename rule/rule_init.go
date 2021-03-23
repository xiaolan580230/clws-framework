package rule

import (
	"clws-framework/api/api_example"
	"clws-framework/core/clRouter"
)

func InitRule() {

	// 普通用户登录
	clRouter.AddRule(clRouter.RouterRule{
		Ac:       "api_example",
		Param:    []clRouter.RouterParam{
			clRouter.NewParam("param1", "", clRouter.ParamTypeSafe, true),
			clRouter.NewParam("param2", "", clRouter.ParamTypeSafe, true),
		},
		Callback: api_example.ApiExample,
		Login:    false,
	})

}