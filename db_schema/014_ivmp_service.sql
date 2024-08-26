-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
create table o_camera(
    id varchar(255) not null,
    identifier varchar(255) not null,
    name varchar(255) not null,
    created_by      VARCHAR(32),
    created_time    TIMESTAMP,
    updated_by      VARCHAR(32),
    updated_time    TIMESTAMP,
    type int4 not null default 0,
    ai_type int4 not null default 0,
    username varchar(255) null,
    password varchar(255) null,
    stream_type varchar(255) null,
    stream_port int4 not null default 554,
    stream_path varchar(1024) null,
    second_stream_path varchar(1024) null,
    ip varchar(255) null,
    ai_model varchar(1024) null,
    ai_status int4 not null default 0,
    ai_config text default '{}',
    third_id varchar(255) null,
    primary key (id)
);
comment on table o_camera is '摄像头';
comment on column o_camera.id is 'id';
comment on column o_camera.identifier is '标识';
comment on column o_camera.name is '名称';
comment on column o_camera.created_by is '创建者';
comment on column o_camera.created_time is '创建时间';
comment on column o_camera.updated_by is '更新者';
comment on column o_camera.updated_time is '更新时间';
comment on column o_camera.type is '类型(0:rtsp,1:virtual)';
comment on column o_camera.ai_type is 'AI类型:(0:none,1:分类...)';
comment on column o_camera.username is '用户名';
comment on column o_camera.password is '密码';
comment on column o_camera.stream_type is '流类型(rtsp,rtmp,file)';
comment on column o_camera.stream_port is '流端口';
comment on column o_camera.stream_path is '流地址';
comment on column o_camera.second_stream_path is '第二路流地址';
comment on column o_camera.ip is 'ip';
comment on column o_camera.ai_model is 'AI模型';
comment on column o_camera.ai_status is 'AI状态';
comment on column o_camera.ai_config is 'AI配置';
create index idx_o_camera_name on o_camera (name);

create table o_scene(
    id varchar(255) not null,
    name varchar(255) not null,
    created_by      VARCHAR(32),
    created_time    TIMESTAMP,
    updated_by      VARCHAR(32),
    updated_time    TIMESTAMP,
    type int4 not null default 0,
    config text null,
    description text null,
    primary key (id)
);
comment on table o_scene is '场景';
comment on column o_scene.id is 'id';
comment on column o_scene.name is '名称';
comment on column o_scene.created_by is '创建者';
comment on column o_scene.created_time is '创建时间';
comment on column o_scene.updated_by is '更新者';
comment on column o_scene.updated_time is '更新时间';
comment on column o_scene.type is '类型';
comment on column o_scene.config is '配置';
comment on column o_scene.description is '描述';
create index idx_o_scene_name on o_scene (name);

create table o_ai_model(
    id varchar(255) not null,
    name varchar(255) not null,
    created_by      VARCHAR(32),
    created_time    TIMESTAMP,
    updated_by      VARCHAR(32),
    updated_time    TIMESTAMP,
    type varchar(255) not null default '',
    sub_type varchar(255) not null default '',
    config text null,
    description text null,
    version varchar(255) null,
    file_ext varchar(255) null,
    primary key (id)
);
comment on table o_ai_model is 'AI模型';
comment on column o_ai_model.id is 'id';
comment on column o_ai_model.name is '名称';
comment on column o_ai_model.created_by is '创建者';
comment on column o_ai_model.created_time is '创建时间';
comment on column o_ai_model.updated_by is '更新者';
comment on column o_ai_model.updated_time is '更新时间';
comment on column o_ai_model.type is '类型';
comment on column o_ai_model.sub_type is '子类型';
comment on column o_ai_model.config is '配置,json格式，每种类型不同的格式';
comment on column o_ai_model.description is '描述';
comment on column o_ai_model.version is '版本';
comment on column o_ai_model.file_ext is '文件后缀';
create index idx_o_ai_model_name on o_ai_model (name);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE IF EXISTS o_camera cascade;
DROP TABLE IF EXISTS o_scene cascade;
DROP TABLE IF EXISTS o_ai_model cascade;
-- +goose StatementEnd
