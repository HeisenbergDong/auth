package gin

import (
	"auth/global"
	"fmt"
	"github.com/fvbock/endless"
)

func Run() {
	go func() {
		global.ADDRESS = fmt.Sprintf(":%d", global.CONFIG.System.Addr) // 初始化服务地址和端口
		s := endless.NewServer(global.ADDRESS, global.ROUTER)
		s.ListenAndServe()
	}()
}
