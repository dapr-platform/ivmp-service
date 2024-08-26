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


Table: o_camera
[ 0] id                                             VARCHAR(255)         null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 1] identifier                                     VARCHAR(255)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 2] name                                           VARCHAR(255)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[ 3] created_by                                     VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 4] created_time                                   TIMESTAMP            null: true   primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
[ 5] updated_by                                     VARCHAR(32)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 32      default: []
[ 6] updated_time                                   TIMESTAMP            null: true   primary: false  isArray: false  auto: false  col: TIMESTAMP       len: -1      default: []
[ 7] type                                           INT4                 null: false  primary: false  isArray: false  auto: false  col: INT4            len: -1      default: [0]
[ 8] ai_type                                        INT4                 null: false  primary: false  isArray: false  auto: false  col: INT4            len: -1      default: [0]
[ 9] username                                       VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[10] password                                       VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[11] stream_type                                    VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[12] stream_port                                    INT4                 null: false  primary: false  isArray: false  auto: false  col: INT4            len: -1      default: [554]
[13] stream_path                                    VARCHAR(1024)        null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 1024    default: []
[14] second_stream_path                             VARCHAR(1024)        null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 1024    default: []
[15] ip                                             VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
[16] ai_model                                       VARCHAR(1024)        null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 1024    default: []
[17] ai_status                                      INT4                 null: false  primary: false  isArray: false  auto: false  col: INT4            len: -1      default: [0]
[18] ai_config                                      TEXT                 null: true   primary: false  isArray: false  auto: false  col: TEXT            len: -1      default: [{}]
[19] third_id                                       VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []


JSON Sample
-------------------------------------
{    "id": "JKtrEyjfbBXrOcxNGEJTBPShQ",    "identifier": "BRJITRtmxBggggBeJqqgVgfSH",    "name": "rqvdvVynXreqmNKdiXpHoEqHs",    "created_by": "PvuGYZyxeQqVCBMCamITtYUNx",    "created_time": 2,    "updated_by": "AUpGuTQPGEHwDXyLNPsYrRMpj",    "updated_time": 49,    "type": 22,    "ai_type": 66,    "username": "qXetZFMlwhCNHhlSMREcgCgnF",    "password": "IKdnvqyujbRkQtgUMwvrCpYIS",    "stream_type": "mofVfrSnuHHgVxkJTXxNtUsym",    "stream_port": 7,    "stream_path": "GPRsSrjWwYVwffknRVqVVExYH",    "second_stream_path": "WQjAbbkAMhYOmZwZansFPOFVa",    "ip": "WlruQlMsxZiibxwEHWHsdZqEY",    "ai_model": "htgpSXOGTCRkYXhkZEpsYpkyO",    "ai_status": 88,    "ai_config": "NBsCPYpDWwaJDTUyWgKYLpsgA",    "third_id": "JDwamKepMIdCFsSseRZHijkYw"}



*/

var (
	Camera_FIELD_NAME_id = "id"

	Camera_FIELD_NAME_identifier = "identifier"

	Camera_FIELD_NAME_name = "name"

	Camera_FIELD_NAME_created_by = "created_by"

	Camera_FIELD_NAME_created_time = "created_time"

	Camera_FIELD_NAME_updated_by = "updated_by"

	Camera_FIELD_NAME_updated_time = "updated_time"

	Camera_FIELD_NAME_type = "type"

	Camera_FIELD_NAME_ai_type = "ai_type"

	Camera_FIELD_NAME_username = "username"

	Camera_FIELD_NAME_password = "password"

	Camera_FIELD_NAME_stream_type = "stream_type"

	Camera_FIELD_NAME_stream_port = "stream_port"

	Camera_FIELD_NAME_stream_path = "stream_path"

	Camera_FIELD_NAME_second_stream_path = "second_stream_path"

	Camera_FIELD_NAME_ip = "ip"

	Camera_FIELD_NAME_ai_model = "ai_model"

	Camera_FIELD_NAME_ai_status = "ai_status"

	Camera_FIELD_NAME_ai_config = "ai_config"

	Camera_FIELD_NAME_third_id = "third_id"
)

// Camera struct is a row record of the o_camera table in the  database
type Camera struct {
	ID               string           `json:"id"`                 //id
	Identifier       string           `json:"identifier"`         //标识
	Name             string           `json:"name"`               //名称
	CreatedBy        string           `json:"created_by"`         //创建者
	CreatedTime      common.LocalTime `json:"created_time"`       //创建时间
	UpdatedBy        string           `json:"updated_by"`         //更新者
	UpdatedTime      common.LocalTime `json:"updated_time"`       //更新时间
	Type             int32            `json:"type"`               //类型(0:rtsp,1:virtual)
	AiType           int32            `json:"ai_type"`            //AI类型:(0:none,1:分类...)
	Username         string           `json:"username"`           //用户名
	Password         string           `json:"password"`           //密码
	StreamType       string           `json:"stream_type"`        //流类型(rtsp,rtmp,file)
	StreamPort       int32            `json:"stream_port"`        //流端口
	StreamPath       string           `json:"stream_path"`        //流地址
	SecondStreamPath string           `json:"second_stream_path"` //第二路流地址
	IP               string           `json:"ip"`                 //ip
	AiModel          string           `json:"ai_model"`           //AI模型
	AiStatus         int32            `json:"ai_status"`          //AI状态
	AiConfig         string           `json:"ai_config"`          //AI配置
	ThirdID          string           `json:"third_id"`           //third_id

}

var CameraTableInfo = &TableInfo{
	Name: "o_camera",
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
			Name:               "identifier",
			Comment:            `标识`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "Identifier",
			GoFieldType:        "string",
			JSONFieldName:      "identifier",
			ProtobufFieldName:  "identifier",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
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
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
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
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
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
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
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
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
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
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "type",
			Comment:            `类型(0:rtsp,1:virtual)`,
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
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "ai_type",
			Comment:            `AI类型:(0:none,1:分类...)`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "AiType",
			GoFieldType:        "int32",
			JSONFieldName:      "ai_type",
			ProtobufFieldName:  "ai_type",
			ProtobufType:       "int32",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "username",
			Comment:            `用户名`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "Username",
			GoFieldType:        "string",
			JSONFieldName:      "username",
			ProtobufFieldName:  "username",
			ProtobufType:       "string",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "password",
			Comment:            `密码`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "Password",
			GoFieldType:        "string",
			JSONFieldName:      "password",
			ProtobufFieldName:  "password",
			ProtobufType:       "string",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "stream_type",
			Comment:            `流类型(rtsp,rtmp,file)`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "StreamType",
			GoFieldType:        "string",
			JSONFieldName:      "stream_type",
			ProtobufFieldName:  "stream_type",
			ProtobufType:       "string",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "stream_port",
			Comment:            `流端口`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "StreamPort",
			GoFieldType:        "int32",
			JSONFieldName:      "stream_port",
			ProtobufFieldName:  "stream_port",
			ProtobufType:       "int32",
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "stream_path",
			Comment:            `流地址`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(1024)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       1024,
			GoFieldName:        "StreamPath",
			GoFieldType:        "string",
			JSONFieldName:      "stream_path",
			ProtobufFieldName:  "stream_path",
			ProtobufType:       "string",
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
			Name:               "second_stream_path",
			Comment:            `第二路流地址`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(1024)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       1024,
			GoFieldName:        "SecondStreamPath",
			GoFieldType:        "string",
			JSONFieldName:      "second_stream_path",
			ProtobufFieldName:  "second_stream_path",
			ProtobufType:       "string",
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
			Name:               "ip",
			Comment:            `ip`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "IP",
			GoFieldType:        "string",
			JSONFieldName:      "ip",
			ProtobufFieldName:  "ip",
			ProtobufType:       "string",
			ProtobufPos:        16,
		},

		&ColumnInfo{
			Index:              16,
			Name:               "ai_model",
			Comment:            `AI模型`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(1024)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       1024,
			GoFieldName:        "AiModel",
			GoFieldType:        "string",
			JSONFieldName:      "ai_model",
			ProtobufFieldName:  "ai_model",
			ProtobufType:       "string",
			ProtobufPos:        17,
		},

		&ColumnInfo{
			Index:              17,
			Name:               "ai_status",
			Comment:            `AI状态`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "INT4",
			DatabaseTypePretty: "INT4",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT4",
			ColumnLength:       -1,
			GoFieldName:        "AiStatus",
			GoFieldType:        "int32",
			JSONFieldName:      "ai_status",
			ProtobufFieldName:  "ai_status",
			ProtobufType:       "int32",
			ProtobufPos:        18,
		},

		&ColumnInfo{
			Index:              18,
			Name:               "ai_config",
			Comment:            `AI配置`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "TEXT",
			DatabaseTypePretty: "TEXT",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "TEXT",
			ColumnLength:       -1,
			GoFieldName:        "AiConfig",
			GoFieldType:        "string",
			JSONFieldName:      "ai_config",
			ProtobufFieldName:  "ai_config",
			ProtobufType:       "string",
			ProtobufPos:        19,
		},

		&ColumnInfo{
			Index:              19,
			Name:               "third_id",
			Comment:            `third_id`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       255,
			GoFieldName:        "ThirdID",
			GoFieldType:        "string",
			JSONFieldName:      "third_id",
			ProtobufFieldName:  "third_id",
			ProtobufType:       "string",
			ProtobufPos:        20,
		},
	},
}

// TableName sets the insert table name for this struct type
func (c *Camera) TableName() string {
	return "o_camera"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *Camera) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *Camera) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *Camera) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *Camera) TableInfo() *TableInfo {
	return CameraTableInfo
}
