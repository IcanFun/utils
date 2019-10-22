package utils

import (
	"fmt"
	"giac/utils/im"
	"testing"
)

func Test_IM(t *testing.T) {

	err := im.CreateImAccounts("sdfsf1312", "jacky", "", "", "helloabc2")
	if err != nil {
		fmt.Println(err)
		return
	}

}
