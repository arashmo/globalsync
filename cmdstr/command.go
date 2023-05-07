package cmdstr

import (
	"fmt"

	"github.com/arashmo/globalsync/servers"
//	"github.com/arashmo/globalsync/sshsync"
	"github.com/gin-gonic/gin"
)

func copier(c *gin.Context)  {

src := servers.SearchSrcData
fmt.Println(src)

}
