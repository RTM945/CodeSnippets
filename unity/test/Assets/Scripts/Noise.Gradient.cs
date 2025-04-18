using Unity.Mathematics;

using static Unity.Mathematics.math;

public static partial class Noise {

	public interface IGradient {
		float4 Evaluate (SmallXXHash4 hash, float4 x);

		float4 Evaluate (SmallXXHash4 hash, float4 x, float4 y);

		float4 Evaluate (SmallXXHash4 hash, float4 x, float4 y, float4 z);

		float4 EvaluateCombined (float4 value);
	}

	public struct Turbulence<G> : IGradient where G : struct, IGradient {

		public float4 Evaluate (SmallXXHash4 hash, float4 x) =>
			default(G).Evaluate(hash, x);

		public float4 Evaluate (SmallXXHash4 hash, float4 x, float4 y) =>
			default(G).Evaluate(hash, x, y);

		public float4 Evaluate (SmallXXHash4 hash, float4 x, float4 y, float4 z) =>
			default(G).Evaluate(hash, x, y, z);

		public float4 EvaluateCombined (float4 value) =>
			abs(default(G).EvaluateCombined(value));
	}

	public struct Value : IGradient {

		public float4 Evaluate (SmallXXHash4 hash, float4 x) => hash.Floats01A * 2f - 1f;

		public float4 Evaluate (SmallXXHash4 hash, float4 x, float4 y) =>
			hash.Floats01A * 2f - 1f;

		public float4 Evaluate (SmallXXHash4 hash, float4 x, float4 y, float4 z) =>
			hash.Floats01A * 2f - 1f;

		public float4 EvaluateCombined (float4 value) => value;
	}

	public struct Perlin : IGradient {

		public float4 Evaluate (SmallXXHash4 hash, float4 x) =>
			BaseGradients.Line(hash, x);

		public float4 Evaluate (SmallXXHash4 hash, float4 x, float4 y) =>
			BaseGradients.Square(hash, x, y) * (2f / 0.53528f);

		public float4 Evaluate (SmallXXHash4 hash, float4 x, float4 y, float4 z) =>
			BaseGradients.Octahedron(hash, x, y, z) * (1f / 0.56290f);

		public float4 EvaluateCombined (float4 value) => value;
	}

	public struct Simplex : IGradient {

		public float4 Evaluate (SmallXXHash4 hash, float4 x) =>
			BaseGradients.Line(hash, x) * (32f / 27f);

		public float4 Evaluate (SmallXXHash4 hash, float4 x, float4 y) =>
			BaseGradients.Circle(hash, x, y) * (5.832f / sqrt(2f));

		public float4 Evaluate (SmallXXHash4 hash, float4 x, float4 y, float4 z) =>
			BaseGradients.Sphere(hash, x, y, z) * (1024f / (125f * sqrt(3f)));

		public float4 EvaluateCombined (float4 value) => value;
	}

	public static class BaseGradients {

		public static float4 Line (SmallXXHash4 hash, float4 x) =>
			(1f + hash.Floats01A) * select(-x, x, ((uint4)hash & 1 << 8) == 0);

		public static float4 Square (SmallXXHash4 hash, float4 x, float4 y) {
			float4x2 v = SquareVectors(hash);
			return v.c0 * x + v.c1 * y;
		}

		public static float4 Circle (SmallXXHash4 hash, float4 x, float4 y) {
			float4x2 v = SquareVectors(hash);
			return (v.c0 * x + v.c1 * y) * rsqrt(v.c0 * v.c0 + v.c1 * v.c1);
		}

		public static float4 Octahedron (
			SmallXXHash4 hash, float4 x, float4 y, float4 z
		) {
			float4x3 v = OctahedronVectors(hash);
			return v.c0 * x + v.c1 * y + v.c2 * z;
		}

		public static float4 Sphere (SmallXXHash4 hash, float4 x, float4 y, float4 z) {
			float4x3 v = OctahedronVectors(hash);
			return
				(v.c0 * x + v.c1 * y + v.c2 * z) *
				rsqrt(v.c0 * v.c0 + v.c1 * v.c1 + v.c2 * v.c2);
		}

		static float4x2 SquareVectors (SmallXXHash4 hash) {
			float4x2 v;
			v.c0 = hash.Floats01A * 2f - 1f;
			v.c1 = 0.5f - abs(v.c0);
			v.c0 -= floor(v.c0 + 0.5f);
			return v;
		}

		static float4x3 OctahedronVectors (SmallXXHash4 hash) {
			float4x3 g;
			g.c0 = hash.Floats01A * 2f - 1f;
			g.c1 = hash.Floats01D * 2f - 1f;
			g.c2 = 1f - abs(g.c0) - abs(g.c1);
			float4 offset = max(-g.c2, 0f);
			g.c0 += select(-offset, offset, g.c0 < 0f);
			g.c1 += select(-offset, offset, g.c1 < 0f);
			return g;
		}
	}
}