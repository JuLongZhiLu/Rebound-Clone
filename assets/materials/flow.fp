varying mediump vec2 var_texcoord0;
varying mediump vec4 var_position;

uniform lowp sampler2D texture_sampler;
uniform lowp vec4 tint;

// 从 Lua 传来的 vec4
// x = 累积时间, y = X轴速度, z = Y轴速度
uniform mediump vec4 time_and_speed; 

void main()
{
	// 1. 获取原始 UV
	vec2 uv = var_texcoord0;

	// 2. 根据时间计算偏移量
	// 偏移 = 时间 * 速度
	float time = time_and_speed.x;
	vec2 speed = time_and_speed.yz;

	uv += time * speed;

	// 3. 采样纹理
	gl_FragColor = texture2D(texture_sampler, uv) * tint;
}