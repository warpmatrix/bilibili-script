package domain_test

import (
	"main/src/domain"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCookie(t *testing.T) {
	cookie := "SESSDATA=37dfc8ba%2C1640068106%2C60f11%2A61\n" +
		"bili_jct=51685a18c19d95388285d238d3c81ad0\n" +
		"DedeUserID=25032900"
	os.Setenv("COOKIE", cookie)
	cookie_kv := domain.GetCookie()
	assert.Equal(t, "bili_jct=51685a18c19d95388285d238d3c81ad0;SESSDATA=37dfc8ba%2C1640068106%2C60f11%2A61;DedeUserID=25032900;", cookie_kv)
}
