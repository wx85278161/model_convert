package model_convert

import (
	"fmt"
	"testing"
)

func TestGenerateMDAdd(t *testing.T) {
	type JiugonggeSmallAwardPool struct {
		Id         int     `gorm:"column:id;default:" json:"id" form:"id"`
		GameId     int     `gorm:"column:game_id;default:" json:"game_id" form:"game_id"`
		BlockId    int     `gorm:"column:block_id;default:" json:"block_id" form:"block_id"`
		PropId     int     `gorm:"column:prop_id;default:" json:"prop_id" form:"prop_id"`
		PropNum    int     `gorm:"column:prop_num;default:" json:"prop_num" form:"prop_num"`
		ExpireIn   int     `gorm:"column:expire_in;default:" json:"expire_in" form:"expire_in"`
		Percentage float64 `gorm:"column:percentage;default:" json:"percentage" form:"percentage"`
		TotalNum   int     `gorm:"column:total_num;default:" json:"total_num" form:"total_num"`
	}
	rs := generateMDAdd(JiugonggeSmallAwardPool{}, map[string]interface{}{
		"${model_chinese_name}": "九宫格小奖池",
		"${md_order}":           1,
	})
	fmt.Println(rs)
}

func TestGenerateMDUpdate(t *testing.T) {
	type JiugonggeSmallAwardPool struct {
		Id         int     `gorm:"column:id;default:" json:"id" form:"id"`
		GameId     int     `gorm:"column:game_id;default:" json:"game_id" form:"game_id"`
		BlockId    int     `gorm:"column:block_id;default:" json:"block_id" form:"block_id"`
		PropId     int     `gorm:"column:prop_id;default:" json:"prop_id" form:"prop_id"`
		PropNum    int     `gorm:"column:prop_num;default:" json:"prop_num" form:"prop_num"`
		ExpireIn   int     `gorm:"column:expire_in;default:" json:"expire_in" form:"expire_in"`
		Percentage float64 `gorm:"column:percentage;default:" json:"percentage" form:"percentage"`
		TotalNum   int     `gorm:"column:total_num;default:" json:"total_num" form:"total_num"`
	}
	rs := generateMDUpdate(JiugonggeSmallAwardPool{}, map[string]interface{}{
		"${model_chinese_name}": "九宫格小奖池",
		"${md_order}":           1,
	})
	fmt.Println(rs)
}

func TestGenerateMDList(t *testing.T) {
	type DzpTaskUnion struct {
		Id            int    `gorm:"column:id;default:" json:"id" form:"id"`
		TaskUnionId   int    `gorm:"column:task_union_id;default:" json:"task_union_id" form:"task_union_id"`
		Description   string `gorm:"column:description;default:" json:"description" form:"description"`
		OsDescription string `gorm:"column:os_description;default:" json:"os_description" form:"os_description"`

		State     int    `gorm:"column:state;default:" json:"state" form:"state"`
		TaskTimes int    `gorm:"column:task_times;default:" json:"task_times" form:"task_times"`
		RefreshAt string `gorm:"column:refresh_at;default:" json:"refresh_at" form:"refresh_at"`
	}

	rs := generateMDList(DzpTaskUnion{}, map[string]interface{}{
		"${model_chinese_name}": "任务池",
		"${md_order}":           2,
	})
	fmt.Println(rs)
}

func TestGenerateMDelete(t *testing.T) {
	type JiugonggeSmallAwardPool struct {
		Id         int     `gorm:"column:id;default:" json:"id" form:"id"`
		GameId     int     `gorm:"column:game_id;default:" json:"game_id" form:"game_id"`
		BlockId    int     `gorm:"column:block_id;default:" json:"block_id" form:"block_id"`
		PropId     int     `gorm:"column:prop_id;default:" json:"prop_id" form:"prop_id"`
		PropNum    int     `gorm:"column:prop_num;default:" json:"prop_num" form:"prop_num"`
		ExpireIn   int     `gorm:"column:expire_in;default:" json:"expire_in" form:"expire_in"`
		Percentage float64 `gorm:"column:percentage;default:" json:"percentage" form:"percentage"`
		TotalNum   int     `gorm:"column:total_num;default:" json:"total_num" form:"total_num"`
	}
	rs := generateMDDelete(JiugonggeSmallAwardPool{}, map[string]interface{}{
		"${model_chinese_name}": "九宫格小奖池",
		"${md_order}":           1,
	})
	fmt.Println(rs)
}

func TestGenerateMGet(t *testing.T) {
	type JiugonggeSmallAwardPool struct {
		Id         int     `gorm:"column:id;default:" json:"id" form:"id"`
		GameId     int     `gorm:"column:game_id;default:" json:"game_id" form:"game_id"`
		BlockId    int     `gorm:"column:block_id;default:" json:"block_id" form:"block_id"`
		PropId     int     `gorm:"column:prop_id;default:" json:"prop_id" form:"prop_id"`
		PropNum    int     `gorm:"column:prop_num;default:" json:"prop_num" form:"prop_num"`
		ExpireIn   int     `gorm:"column:expire_in;default:" json:"expire_in" form:"expire_in"`
		Percentage float64 `gorm:"column:percentage;default:" json:"percentage" form:"percentage"`
		TotalNum   int     `gorm:"column:total_num;default:" json:"total_num" form:"total_num"`
	}
	rs := generateMDGet(JiugonggeSmallAwardPool{}, map[string]interface{}{
		"${model_chinese_name}": "九宫格小奖池",
		"${md_order}":           1,
	})
	fmt.Println(rs)
}

func TestGenerateMD(t *testing.T) {
	type DzpTask struct {
		Id            int    `gorm:"column:id;default:" json:"id" form:"id"`
		GameId        int    `gorm:"column:game_id;default:" json:"game_id" form:"game_id"`
		GameAreaId    int    `gorm:"column:game_area_id;default:" json:"game_area_id" form:"game_area_id"`
		TaskUnionId   int    `gorm:"column:task_union_id;default:" json:"task_union_id" form:"task_union_id"`
		Description   string `gorm:"column:description;default:" json:"description" form:"description"`
		OsDescription string `gorm:"column:os_description;default:" json:"os_description" form:"os_description"`

		State     int    `gorm:"column:state;default:" json:"state" form:"state"`
		TaskTimes int    `gorm:"column:task_times;default:" json:"task_times" form:"task_times"`
		RefreshAt string `gorm:"column:refresh_at;default:" json:"refresh_at" form:"refresh_at"`

		FlexibleKey string `gorm:"column:flexible_key;default:" json:"flexible_key" form:"flexible_key"`
	}
	rs := GenerateMarkDown(DzpTask{}, map[string]interface{}{
		"${model_chinese_name}": "九宫格任务",
		"${md_order}":           1,
	})
	fmt.Println(rs)
}