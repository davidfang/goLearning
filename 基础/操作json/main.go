package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// 假数据
	b := []byte(`{"views": [],
	"dicts": [
	  {
		"defKey": "Gender",
		"defName": "性别",
		"intro": "",
		"items": [
		  {
			"defKey": "1",
			"defName": "男",
			"intro": "",
			"parentKey": "",
			"enabled": true,
			"attr1": "",
			"attr2": "",
			"attr3": "",
			"sort": "1",
			"id": "3622D417-DA1A-408F-BEE1-11D328D534A0"
		  },
		  {
			"defKey": "2",
			"defName": "女",
			"intro": "",
			"parentKey": "",
			"enabled": true,
			"attr1": "",
			"attr2": "",
			"attr3": "",
			"sort": "2",
			"id": "380A0790-64A7-481E-831C-32F7BEE1502B"
		  },
		  {
			"defKey": "0",
			"defName": "未知",
			"intro": "",
			"parentKey": "",
			"enabled": true,
			"attr1": "",
			"attr2": "",
			"attr3": "",
			"sort": "3",
			"id": "FA239F4D-1276-40D9-B230-F66BD35C3C27"
		  }
		],
		"id": "BF9E20E0-80D3-486D-BD58-5FADCF3E4A1D"
	  },
	  {
		"defKey": "Political",
		"defName": "政治面貌",
		"intro": "",
		"items": [
		  {
			"defKey": "90",
			"defName": "未知",
			"intro": "",
			"parentKey": "",
			"enabled": true,
			"attr1": "",
			"attr2": "",
			"attr3": "",
			"id": "DA65D752-AF04-4A11-81D8-14A38692A64A",
			"sort": "0"
		  },
		  {
			"defKey": "10",
			"defName": "共青团员",
			"intro": "",
			"parentKey": "",
			"enabled": true,
			"attr1": "",
			"attr2": "",
			"attr3": "",
			"id": "61F3145A-7599-4CCB-B298-D5EE788107BE",
			"sort": "1"
		  },
		  {
			"defKey": "20",
			"defName": "中共党员",
			"intro": "",
			"parentKey": "",
			"enabled": true,
			"attr1": "",
			"attr2": "",
			"attr3": "",
			"id": "ED16D25A-AB2F-4FA0-BB48-2B9031FA28C4",
			"sort": "2"
		  },
		  {
			"defKey": "30",
			"defName": "民主党派",
			"intro": "",
			"parentKey": "",
			"enabled": true,
			"attr1": "",
			"attr2": "",
			"attr3": "",
			"id": "0FB7974A-AE11-438F-86E0-B163316F9272",
			"sort": "3"
		  },
		  {
			"defKey": "40",
			"defName": "群众",
			"intro": "",
			"parentKey": "",
			"enabled": true,
			"attr1": "",
			"attr2": "",
			"attr3": "",
			"id": "7D247234-7E97-45B1-8C56-4A17A370854A",
			"sort": "4"
		  }
		],
		"id": "06EED564-BBA9-4747-8D73-AF809A330CB8"
	  },
	  {
		"defKey": "Marital",
		"defName": "婚姻状况",
		"intro": "婚姻状况的码表",
		"items": [
		  {
			"defKey": "0",
			"defName": "未说明",
			"intro": "",
			"parentKey": "",
			"enabled": true,
			"attr1": "",
			"attr2": "",
			"attr3": "",
			"sort": "0",
			"id": "A7928FE2-349A-4702-9682-2EF7205E077B"
		  },
		  {
			"defKey": "1",
			"defName": "未婚",
			"intro": "",
			"parentKey": "",
			"enabled": true,
			"attr1": "",
			"attr2": "",
			"attr3": "",
			"sort": "1",
			"id": "20EE81BC-74EE-47DA-A56F-9663B23F44BD"
		  },
		  {
			"defKey": "2",
			"defName": "已婚",
			"intro": "",
			"parentKey": "",
			"enabled": true,
			"attr1": "",
			"attr2": "",
			"attr3": "",
			"sort": "2",
			"id": "4DCA10A8-417E-4A8D-BDF6-0A278C060ADC"
		  },
		  {
			"defKey": "3",
			"defName": "丧偶",
			"intro": "",
			"parentKey": "",
			"enabled": true,
			"attr1": "",
			"attr2": "",
			"attr3": "",
			"sort": "3",
			"id": "826062A7-057C-4892-B338-06459F5B808D"
		  },
		  {
			"defKey": "4",
			"defName": "离婚",
			"intro": "",
			"parentKey": "",
			"enabled": true,
			"attr1": "",
			"attr2": "",
			"attr3": "",
			"sort": "4",
			"id": "B23200B4-5E59-4F5E-A779-D981A040FA32"
		  }
		],
		"id": "EA1587B7-3954-437A-BFE0-FCB0453BCABA"
	  },
	  {
		"defKey": "StudentStatus",
		"defName": "学生状态",
		"intro": "",
		"items": [
		  {
			"defKey": "1",
			"defName": "正常",
			"intro": "",
			"parentKey": "",
			"enabled": true,
			"attr1": "",
			"attr2": "",
			"attr3": "",
			"sort": "1",
			"id": "E9CA1CC9-8851-4F6B-86BA-B9CF0E44EB73"
		  },
		  {
			"defKey": "2",
			"defName": "毕业",
			"intro": "",
			"parentKey": "",
			"enabled": true,
			"attr1": "",
			"attr2": "",
			"attr3": "",
			"sort": "2",
			"id": "DEC51D7C-99DF-430C-817D-0499862D3CCC"
		  },
		  {
			"defKey": "3",
			"defName": "肄业",
			"intro": "",
			"parentKey": "",
			"enabled": true,
			"attr1": "",
			"attr2": "",
			"attr3": "",
			"sort": "3",
			"id": "8853D6B6-75D3-4479-9006-FC731CD85B20"
		  },
		  {
			"defKey": "4",
			"defName": "休学",
			"intro": "",
			"parentKey": "",
			"enabled": true,
			"attr1": "",
			"attr2": "",
			"attr3": "",
			"sort": "4",
			"id": "C74BA8CF-1DC6-4C79-BAAC-F11EB9C6AF01"
		  }
		],
		"id": "4642BC5F-02EE-4E17-A60C-CF22F86A0282"
	  },
	  {
		"defKey": "GBNation",
		"defName": "民族",
		"intro": "",
		"items": [
		  {
			"defKey": "0",
			"defName": "未知",
			"sort": "0",
			"parentKey": "",
			"intro": "",
			"enabled": true,
			"attr1": "",
			"attr2": "",
			"attr3": "",
			"id": "0DAE5DA4-9CD2-4C28-BB88-5C5C1AB25B24"
		  },
		  {
			"defKey": "1",
			"defName": "汉族",
			"intro": "",
			"parentKey": "",
			"enabled": true,
			"attr1": "",
			"attr2": "",
			"attr3": "",
			"id": "9224DF53-F7C0-447D-B8ED-0A39F799EE19",
			"sort": "1"
		  },
		  {
			"defKey": "2",
			"defName": "蒙古族",
			"intro": "",
			"parentKey": "",
			"enabled": true,
			"attr1": "",
			"attr2": "",
			"attr3": "",
			"id": "D57818E3-9206-45BB-AE79-27C64A4AB80F",
			"sort": "2"
		  },
		  {
			"defKey": "3",
			"defName": "回族",
			"intro": "",
			"parentKey": "",
			"enabled": true,
			"attr1": "",
			"attr2": "",
			"attr3": "",
			"id": "0A1A3CA9-6D68-4E15-8BD0-9A2FF428D804",
			"sort": "3"
		  },
		  {
			"defKey": "4",
			"defName": "藏族",
			"intro": "",
			"parentKey": "",
			"enabled": true,
			"attr1": "",
			"attr2": "",
			"attr3": "",
			"id": "7CC6B6BE-47EA-460E-ACFA-C377468DEA11",
			"sort": "4"
		  },
		  {
			"defKey": "5",
			"defName": "维吾尔族",
			"intro": "",
			"parentKey": "",
			"enabled": true,
			"attr1": "",
			"attr2": "",
			"attr3": "",
			"id": "A666D51F-D249-4FAC-B1F3-78C371836CB3",
			"sort": "5"
		  },
		  {
			"defKey": "6",
			"defName": "苗族",
			"intro": "",
			"parentKey": "",
			"enabled": true,
			"attr1": "",
			"attr2": "",
			"attr3": "",
			"id": "A0C9E1BA-D87B-4695-ADFA-287FDA32BB5A",
			"sort": "6"
		  },
		  {
			"defKey": "7",
			"defName": "彝族",
			"intro": "",
			"parentKey": "",
			"enabled": true,
			"attr1": "",
			"attr2": "",
			"attr3": "",
			"id": "149B5B9E-C1D1-4790-8CCF-0ED5F4B25172",
			"sort": "7"
		  },
		  {
			"defKey": "8",
			"defName": "壮族",
			"intro": "",
			"parentKey": "",
			"enabled": true,
			"attr1": "",
			"attr2": "",
			"attr3": "",
			"id": "F9A3E65C-BF4D-4C6B-ADB7-8C9CF0487360",
			"sort": "8"
		  },
		  {
			"defKey": "9",
			"defName": "布依族",
			"intro": "",
			"parentKey": "",
			"enabled": true,
			"attr1": "",
			"attr2": "",
			"attr3": "",
			"id": "14F17DE4-E96A-460B-98A6-F84EC8CF9885",
			"sort": "9"
		  },
		  {
			"defKey": "10",
			"defName": "朝鲜族",
			"intro": "",
			"parentKey": "",
			"enabled": true,
			"attr1": "",
			"attr2": "",
			"attr3": "",
			"id": "8A81AB18-B1BF-4797-A6E5-DEDB2C6566B0",
			"sort": "10"
		  },
		  {
			"defKey": "11",
			"defName": "满族",
			"intro": "",
			"parentKey": "",
			"enabled": true,
			"attr1": "",
			"attr2": "",
			"attr3": "",
			"id": "7D460947-FBD2-4E4D-8366-3B38DCAF09D1",
			"sort": "11"
		  }
		],
		"id": "115EDEFC-0323-410E-81AB-CCAB8879837A"
	  },
	  {
		"defKey": "GradeLevel",
		"defName": "受教育程度",
		"sort": "",
		"intro": "",
		"items": [
		  {
			"defKey": "0",
			"defName": "未知",
			"sort": "0",
			"parentKey": "",
			"intro": "",
			"enabled": true,
			"attr1": "",
			"attr2": "",
			"attr3": "",
			"id": "4E88BB19-11B0-4D86-935B-58B67BBD2A03"
		  },
		  {
			"defKey": "1",
			"defName": "高中及以下",
			"sort": "1",
			"parentKey": "",
			"intro": "",
			"enabled": true,
			"attr1": "",
			"attr2": "",
			"attr3": "",
			"id": "794F5ED5-820F-45CD-9802-937B85CC8D3D"
		  },
		  {
			"defKey": "2",
			"defName": "大专",
			"sort": "2",
			"parentKey": "",
			"intro": "",
			"enabled": true,
			"attr1": "",
			"attr2": "",
			"attr3": "",
			"id": "153E9980-C3E9-45B3-9DF6-B60C9B54467D"
		  },
		  {
			"defKey": "3",
			"defName": "本科",
			"sort": "3",
			"parentKey": "",
			"intro": "",
			"enabled": true,
			"attr1": "",
			"attr2": "",
			"attr3": "",
			"id": "F1C4480A-E113-4863-AAB6-698BF396CA10"
		  },
		  {
			"defKey": "4",
			"defName": "硕士",
			"sort": "4",
			"parentKey": "",
			"intro": "",
			"enabled": true,
			"attr1": "",
			"attr2": "",
			"attr3": "",
			"id": "C0027F4F-7E44-47B8-9B5A-469B6FE4C777"
		  },
		  {
			"defKey": "5",
			"defName": "博士",
			"sort": "5",
			"parentKey": "",
			"intro": "",
			"enabled": true,
			"attr1": "",
			"attr2": "",
			"attr3": "",
			"id": "3C9DA6FD-101E-4E8B-8EB8-6B2FDA0E5A4C"
		  },
		  {
			"defKey": "6",
			"defName": "博士后",
			"sort": "6",
			"parentKey": "",
			"intro": "",
			"enabled": true,
			"attr1": "",
			"attr2": "",
			"attr3": "",
			"id": "99D3DD33-CF5D-44A0-B083-0C018838B08B"
		  }
		],
		"id": "9E7C9788-B805-4C7D-8531-FD1D9DC79B05"
	  },
	  {
		"defKey": "UserStatus",
		"defName": "用户状态",
		"sort": "",
		"intro": "",
		"id": "3618143F-B3EF-476D-B63A-24F87D7FF900",
		"items": [
		  {
			"defKey": "1",
			"defName": "正常",
			"sort": "1",
			"parentKey": "",
			"intro": "",
			"enabled": true,
			"attr1": "",
			"attr2": "",
			"attr3": "",
			"id": "9334B14B-F730-4FD8-B5B2-09BA75EC897C"
		  },
		  {
			"defKey": "2",
			"defName": "注销",
			"sort": "2",
			"parentKey": "",
			"intro": "",
			"enabled": true,
			"attr1": "",
			"attr2": "",
			"attr3": "",
			"id": "6E0EC689-F060-46A1-BA8D-85CC80302658"
		  },
		  {
			"defKey": "3",
			"defName": "拉黑",
			"sort": "3",
			"parentKey": "",
			"intro": "",
			"enabled": true,
			"attr1": "",
			"attr2": "",
			"attr3": "",
			"id": "348F49C8-FA50-4311-AF9B-EF48762140CB"
		  }
		]
	  },
	  {
		"defKey": "UserTag",
		"defName": "用户标签",
		"sort": "",
		"intro": "多选项值测试",
		"id": "A72AAFC7-5876-473F-B19A-41E5B2CD5A41",
		"items": [
		  {
			"defKey": "阳光",
			"defName": "阳光",
			"sort": "1",
			"parentKey": "",
			"intro": "",
			"enabled": true,
			"attr1": "",
			"attr2": "",
			"attr3": "",
			"id": "E83CA222-C923-401C-A617-A277E41F3ED1"
		  },
		  {
			"defKey": "帅气",
			"defName": "帅气",
			"sort": "2",
			"parentKey": "",
			"intro": "",
			"enabled": true,
			"attr1": "",
			"attr2": "",
			"attr3": "",
			"id": "5B6364BB-107C-4C30-8BD4-387C06132BE9"
		  },
		  {
			"defKey": "年轻",
			"defName": "年轻",
			"sort": "3",
			"parentKey": "",
			"intro": "",
			"enabled": true,
			"attr1": "",
			"attr2": "",
			"attr3": "",
			"id": "DEC99ED3-BBDA-472F-99ED-DDA6FFAD27A3"
		  },
		  {
			"defKey": "活力",
			"defName": "活力",
			"sort": "4",
			"parentKey": "",
			"intro": "",
			"enabled": true,
			"attr1": "",
			"attr2": "",
			"attr3": "",
			"id": "8F083C89-0CC5-419B-9BC1-9E6C00D11AC5"
		  },
		  {
			"defKey": "智慧",
			"defName": "智慧",
			"sort": "5",
			"parentKey": "",
			"intro": "",
			"enabled": true,
			"attr1": "",
			"attr2": "",
			"attr3": "",
			"id": "8629D88D-F729-407C-B6DD-75676A85FE55"
		  }
		]
	  }
	],
	"viewGroups": [
	  {
		"defKey": "business",
		"defName": "业务",
		"refEntities": [
		  "E5C6B8CF-1673-4C09-A9B7-228F43BDAABF",
		  "1F7F5EC0-45F9-4170-B27E-25331E770952"
		],
		"refViews": [],
		"refDiagrams": [],
		"refDicts": [],
		"id": "DD12B255-7DBB-44B2-826E-14CD12759C1A"
	  }
	]}`)
	// 声明接口
	var i interface{}
	err := json.Unmarshal(b, &i)
	if err != nil {
		fmt.Println(err)
	}
	// 自动转到map
	// fmt.Println(i)
	// 可以判断类型
	m := i.(map[string]interface{})
	fmt.Println(m["viewGroups"])
}
