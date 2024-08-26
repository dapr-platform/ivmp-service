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


Table: o_ai_model
[ 0] id                                             VARCHAR(255)         null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 1] name                                           VARCHAR(255)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 2] created_by                                     VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 3] created_time                                   TIMESTAMP            null: true   primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
[ 4] updated_by                                     VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 5] updated_time                                   TIMESTAMP            null: true   primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
[ 6] config                                         TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 7] description                                    TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: []
[ 8] version                                        VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 9] file_ext                                       VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[10] type                                           VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[11] sub_type                                       VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []


JSON Sample
-------------------------------------
{    "id": "KAdNwidOHNSuWkCGZTjxUgovO",    "name": "jBUhmSsgJOjgwvuYrWZEwYPpe",    "created_by": "gFoGfvsWyMvHUQmlRDrhbXdgc",    "created_time": 78,    "updated_by": "VOhGgbkWmZjxWsnInfrWMKqdX",    "updated_time": 30,    "config": "BNxWsdruIdRCNVxXDxWPDBZkV",    "description": "UdMEPcoHAXcHmTNslGQfqpqBI",    "version": "LNlFDdWFosOTqylvGRpcWujKl",    "file_ext": "OYEfcebBvrTTOxlUhuyCGBuDu",    "type": "wmqosxxwhlyeyHXyuvBOIZrvu",    "sub_type": "EAbhbZfRTlOETVApHohYIXlAq"}



*/

var (
	Ai_model_FIELD_NAME_id = "id"

	Ai_model_FIELD_NAME_name = "name"

	Ai_model_FIELD_NAME_created_by = "created_by"

	Ai_model_FIELD_NAME_created_time = "created_time"

	Ai_model_FIELD_NAME_updated_by = "updated_by"

	Ai_model_FIELD_NAME_updated_time = "updated_time"

	Ai_model_FIELD_NAME_config = "config"

	Ai_model_FIELD_NAME_description = "description"

	Ai_model_FIELD_NAME_version = "version"

	Ai_model_FIELD_NAME_file_ext = "file_ext"

	Ai_model_FIELD_NAME_type = "type"

	Ai_model_FIELD_NAME_sub_type = "sub_type"
)

// Ai_model struct is a row record of the o_ai_model table in the  database
type Ai_model struct {
	ID          string           `json:"id"`           //id
	Name        string           `json:"name"`         //名称
	CreatedBy   string           `json:"created_by"`   //创建者
	CreatedTime common.LocalTime `json:"created_time"` //创建时间
	UpdatedBy   string           `json:"updated_by"`   //更新者
	UpdatedTime common.LocalTime `json:"updated_time"` //更新时间
	Config      string           `json:"config"`       //配置,json格式，每种类型不同的格式
	Description string           `json:"description"`  //描述
	Version     string           `json:"version"`      //版本
	FileExt     string           `json:"file_ext"`     //文件后缀
	Type        string           `json:"type"`         //type
	SubType     string           `json:"sub_type"`     //sub_type

}

var Ai_modelTableInfo = &TableInfo{
	Name: "o_ai_model",
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
			Name:               "config",
			Comment:            `配置,json格式，每种类型不同的格式`,
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
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
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
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "version",
			Comment:            `版本`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "Version",
			GoFieldType:        "string",
			JSONFieldName:      "version",
			ProtobufFieldName:  "version",
			ProtobufType:       "string",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "file_ext",
			Comment:            `文件后缀`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "FileExt",
			GoFieldType:        "string",
			JSONFieldName:      "file_ext",
			ProtobufFieldName:  "file_ext",
			ProtobufType:       "string",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "type",
			Comment:            `type`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "Type",
			GoFieldType:        "string",
			JSONFieldName:      "type",
			ProtobufFieldName:  "type",
			ProtobufType:       "string",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "sub_type",
			Comment:            `sub_type`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "SubType",
			GoFieldType:        "string",
			JSONFieldName:      "sub_type",
			ProtobufFieldName:  "sub_type",
			ProtobufType:       "string",
			ProtobufPos:        12,
		},
	},
}

// TableName sets the insert table name for this struct type
func (a *Ai_model) TableName() string {
	return "o_ai_model"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *Ai_model) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *Ai_model) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *Ai_model) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *Ai_model) TableInfo() *TableInfo {
	return Ai_modelTableInfo
}
