package domain_test

import (
	"main/src/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseUser(t *testing.T) {
	testCases := []struct {
		desc string
		blob []byte
		user domain.User
	}{
		{
			desc: "default test",
			blob: []byte(`{
				"code": 0,
				"data": {
					"uname": "username",
					"level_info": {
						"current_level": 5,
						"current_min": 10800,
						"current_exp": 19664,
						"next_exp": 28800
					}
				}
			}`),
			user: domain.User{
				Uname: "username",
				Level: domain.Level{
					CurLevel: 5,
					CurExp:   19664,
					NextExp:  28800.0,
				},
			},
		},
		{
			desc: "chinese username",
			blob: []byte(`{
				"code": 0,
				"data": {
					"uname": "中文用户名",
					"level_info": {
						"current_level": 4,
						"current_min": 1080,
						"current_exp": 1964,
						"next_exp": 2880
					}
				}
			}`),
			user: domain.User{
				Uname: "中文用户名",
				Level: domain.Level{
					CurLevel: 4,
					CurExp:   1964,
					NextExp:  2880.0,
				},
			},
		},
		{
			desc: "max level user",
			blob: []byte(`{
				"code": 0,
				"data": {
					"uname": "maxLevelUser",
					"level_info": {
						"current_level": 6,
						"current_min": 28800,
						"current_exp": 31920,
						"next_exp": "--"
					}
				}
			}`),
			user: domain.User{
				Uname: "maxLevelUser",
				Level: domain.Level{
					CurLevel: 6,
					CurExp:   31920,
					NextExp:  "--",
				},
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			user, err := domain.GetUserInfo(tC.blob)
			assert.Equal(t, *user, tC.user)
			assert.Nil(t, err)
		})
	}
}
