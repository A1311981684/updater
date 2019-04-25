package generate

import (
	"github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
)

func TestCreateUpdate(t *testing.T) {
	convey.Convey("Test create update file", t, func(c convey.C) {
		abs := "F:\\VirtualBox\\Ubuntu\\GoPath\\GoPath\\src\\updator\\testFiles\\"
		var uc = UpdateContent{
			Scripts:[]string{abs+"post_1.sh", abs+"pre_2.sh", abs+"pre_1.sh"},
			Name:"app",
			Version: VersionControl{
				To:"1.5.0",
				From:[]string{"1.1","1.2"},
				UpdateLog:[]string{"1. Fix nothing", "2. Update nothing", "3. Replace nothing"},
			},
			Paths:[]string{abs + "app"},
		}
		f, err := os.Create(abs + "zzz.update")
		if err != nil {
			panic(err)
		}
		defer func() {
			err := f.Close()
			if err != nil {
				panic(err)
			}
		}()
		err = CreateUpdate(uc, f)
		c.So(err , convey.ShouldBeNil)
	})
}
