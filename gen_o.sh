dp-cli gen --connstr "postgresql://things:things2024@localhost:5432/thingsdb?sslmode=disable" \
--tables=o_camera,o_ai_model,o_scene --model_naming "{{ toUpperCamelCase ( replace . \"o_\" \"\") }}"  \
--file_naming "{{ toLowerCamelCase ( replace . \"o_\" \"\") }}" \
--module ivmp-service

