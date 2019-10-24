package utils

import (
	"fmt"
	"testing"

	"github.com/IcanFun/utils/utils/im"
)

func Test_IM(t *testing.T) {

	err := im.CreateImAccounts("sdfsf1312", "jacky", "", "", "helloabc2")
	if err != nil {
		fmt.Println(err)
		return
	}

}
