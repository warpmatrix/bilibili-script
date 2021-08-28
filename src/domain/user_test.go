package domain

import (
	"main/src/client"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseUser(t *testing.T) {
	testCases := []struct {
		desc string
		blob []byte
		user User
	}{{
		desc: "default test",
		blob: []byte(`{
			"code": 0,
			"data": {
				"uname": "username"
			}
		}`),
		user: User{
			Name: "username",
		},
	}, {
		desc: "chinese username",
		blob: []byte(`{
			"code": 0,
			"data": {
				"uname": "中文用户名"
			}
		}`),
		user: User{
			Name: "中文用户名",
		},
	}, {
		desc: "unmax level user",
		blob: []byte(`{
			"code": 0,
			"data": {
				"level_info": {
					"current_level": 5,
					"current_min": 10800,
					"current_exp": 19664,
					"next_exp": 28800
				}
			}
		}`),
		user: User{
			Level: level{
				CurLevel: 5,
				CurExp:   19664,
				NextExp:  28800.0,
			},
		},
	}, {
		desc: "max level user",
		blob: []byte(`{
			"code": 0,
			"data": {
				"level_info": {
					"current_level": 6,
					"current_min": 28800,
					"current_exp": 31920,
					"next_exp": "--"
				}
			}
		}`),
		user: User{
			Level: level{
				CurLevel: 6,
				CurExp:   31920,
				NextExp:  "--",
			},
		},
	}, {
		desc: "vip",
		blob: []byte(`{
			"code": 0,
			"data": {
				"vip": {
					"type": 2,
					"status": 1,
					"label": {
						"text": "年度大会员"
					}
				}
			}
		}`),
		user: User{
			Vip: vip{
				Typ:  2,
				St:   1,
				text: "年度大会员",
				Remain: map[string]interface{}{
					"label": map[string]interface{}{"text": "年度大会员"},
				},
			},
		},
	}}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			data, err := client.ParseBlob(tC.blob)
			assert.Nil(t, err)
			ptr, err := parseUserInfo(data.(map[string]interface{}))
			assert.Equal(t, *ptr, tC.user)
			assert.Nil(t, err)
		})
	}
}
