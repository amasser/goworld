#version 330

in vec4 color0;
in vec3 normal0;
in vec3 position0;

layout(location=0) out vec4 out_diffuse;
layout(location=1) out vec4 out_normal;
layout(location=2) out vec4 out_position;

void main() {
    out_diffuse = color0;

    vec4 pack_normal = vec4((normal0 + 1.0) / 2.0, 1);
    out_normal = pack_normal;

    out_position = vec4(position0,1);
}
