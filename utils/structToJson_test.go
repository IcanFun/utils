package utils

import (
	"fmt"
	"giac/utils/json"
	"testing"
)

//填入后删除不需要的字段
type TestStruct struct {
	Asset                    string `json:"asset" binding:"required"`
	ICOProgress              string `json:"ICOProgress"`
	Application              string `json:"application"`
	BlockQuery               string `json:"blockQuery"`
	CirculateVolume          string `json:"circulateVolume"`
	ConsensusMechanism       string `json:"consensusMechanism"`
	CrowdfundingPrice        string `json:"crowdfundingPrice"`
	FinancingHistory         string `json:"financingHistory"`
	FullName                 string `json:"fullName"`
	OfficialWebsite          string `json:"officialWebsite"`
	ProjectConsultant        string `json:"projectConsultant"`
	ProjectPosition          string `json:"projectPosition"`
	ProjectProgress          string `json:"projectProgress"`
	ProjectTeam              string `json:"projectTeam"`
	ProjectType              string `json:"projectType"`
	ProjectValuation         string `json:"projectValuation"`
	PublishTime              string `json:"publishTime"`
	PublishVolume            string `json:"publishVolume"`
	Summary                  string `json:"summary" xorm:"MediumText"`
	TechnicalCharacteristics string `json:"technicalCharacteristics"`
	WhitePaper               string `json:"whitePaper"`
	Lang                     string `json:"lang"`
}

//用于写文档时,快速将参数转换成json然后导入小要鸡
//点击三角号运行.
func Test_Struct_To_Json(t *testing.T) {

	mystruct := new(TestStruct)

	res := StructToMap(mystruct)

	b, err := json.Marshal(res)
	if err != nil {
		t.Error(err)
	}
	fmt.Println()
	fmt.Println(string(b))
	fmt.Println()

}
