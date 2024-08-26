package model

import (
	"database/sql"
	"github.com/dapr-platform/common"
	"time"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = common.LocalTime{}
)

/*
DB Table Details
-------------------------------------


Table: o_scene
[ 0] id                                             VARCHAR(255)         null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 1] name                                           VARCHAR(255)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 2] created_by                                     VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 3] created_time                                   TIMESTAMP            null: true   primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
[ 4] updated_by                                     VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 5] updated_time                                   TIMESTAMP            null: true   primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
[ 6] type                                           INT4                 null: false  primary: false  isArray: false  auto: false  col: INT4            len: -1      default: [0]
[ 7] config                                         TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 8] description                                    TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []


JSON Sample
-------------------------------------
{    "id": "XBWVvlglnuSTeuNAbNcunOLXa",    "name": "DSAUxPHlbZhhntTTsGKfkDZUu",    "created_by": "FdDrvoirHmUXoALsMsWxIXoJT",    "created_time": 30,    "updated_by": "dOkLDiNgFGSfmUfkXeANGXJkm",    "updated_time": 13,    "type": 34,    "config": "NLPcoGAqvZwbEKZlCKjmkTPhg",    "description": "qRYgcQyvjUPtorbOOttKYlFXh"}



*/

var (
	Scene_FIELD_NAME_id = "id"

	Scene_FIELD_NAME_name = "name"

	Scene_FIELD_NAME_created_by = "created_by"

	Scene_FIELD_NAME_created_time = "created_time"

	Scene_FIELD_NAME_updated_by = "updated_by"

	Scene_FIELD_NAME_updated_time = "updated_time"

	Scene_FIELD_NAME_type = "type"

	Scene_FIELD_NAME_config = "config"

	Scene_FIELD_NAME_description = "description"
)

// Scene struct is a row record of the o_scene table in the  database
type Scene struct {
	ID          string           `json:"id"`           //id
	Name        string           `json:"name"`         //名称
	CreatedBy   string           `json:"created_by"`   //创建者
	CreatedTime common.LocalTime `json:"created_time"` //创建时间
	UpdatedBy   string           `json:"updated_by"`   //更新者
	UpdatedTime common.LocalTime `json:"updated_time"` //更新时间
	Type        int32            `json:"type"`         //类型
	Config      string           `json:"config"`       //配置
	Description string           `json:"description"`  //描述

}

var SceneTableInfo = &TableInfo{
	Name: "o_scene",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            `id`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "ID",
			GoFieldType:        "string",
			JSONFieldName:      "id",
			ProtobufFieldName:  "id",
			ProtobufType:       "string",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "name",
			Comment:            `名称`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "Name",
			GoFieldType:        "string",
			JSONFieldName:      "name",
			ProtobufFieldName:  "name",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "created_by",
			Comment:            `创建者`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(32)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       32,
			GoFieldName:        "CreatedBy",
			GoFieldType:        "string",
			JSONFieldName:      "created_by",
			ProtobufFieldName:  "created_by",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "created_time",
			Comment:            `创建时间`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TIMESTAMP",
			DatabaseTypePretty: "TIMESTAMP",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TIMESTAMP",
			ColumnLength:       -1,
			GoFieldName:        "CreatedTime",
			GoFieldType:        "common.LocalTime",
			JSONFieldName:      "created_time",
			ProtobufFieldName:  "created_time",
			ProtobufType:       "uint64",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "updated_by",
			Comment:            `更新者`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(32)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       32,
			GoFieldName:        "UpdatedBy",
			GoFieldType:        "string",
			JSONFieldName:      "updated_by",
			ProtobufFieldName:  "updated_by",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "updated_time",
			Comment:            `更新时间`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TIMESTAMP",
			DatabaseTypePretty: "TIMESTAMP",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TIMESTAMP",
			ColumnLength:       -1,
			GoFieldName:        "UpdatedTime",
			GoFieldType:        "common.LocalTime",
			JSONFieldName:      "updated_time",
			ProtobufFieldName:  "updated_time",
			ProtobufType:       "uint64",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "type",
			Comment:            `类型`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "Type",
			GoFieldType:        "int32",
			JSONFieldName:      "type",
			ProtobufFieldName:  "type",
			ProtobufType:       "int32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "config",
			Comment:            `配置`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TEXT",
			DatabaseTypePretty: "TEXT",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TEXT",
			ColumnLength:       -1,
			GoFieldName:        "Config",
			GoFieldType:        "string",
			JSONFieldName:      "config",
			ProtobufFieldName:  "config",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "description",
			Comment:            `描述`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TEXT",
			DatabaseTypePretty: "TEXT",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TEXT",
			ColumnLength:       -1,
			GoFieldName:        "Description",
			GoFieldType:        "string",
			JSONFieldName:      "description",
			ProtobufFieldName:  "description",
			ProtobufType:       "string",
			ProtobufPos:        9,
		},
	},
}

// TableName sets the insert table name for this struct type
func (s *Scene) TableName() string {
	return "o_scene"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (s *Scene) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (s *Scene) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (s *Scene) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (s *Scene) TableInfo() *TableInfo {
	return SceneTableInfo
}
