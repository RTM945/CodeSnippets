using Unity.Mathematics;

using static Unity.Mathematics.math;

public static partial class Noise {

	public struct Simplex1D<G> : INoise where G : struct, IGradient {

		public float4 GetNoise4 (float4x3 positions, SmallXXHash4 hash, int frequency) {
			positions *= frequency;
			int4 x0 = (int4)floor(positions.c0), x1 = x0 + 1;

			return default(G).EvaluateCombined(
				Kernel(hash.Eat(x0), x0, positions) + Kernel(hash.Eat(x1), x1, positions)
			);
		}

		static float4 Kernel (SmallXXHash4 hash, float4 lx, float4x3 positions) {
			float4 x = positions.c0 - lx;
			float4 f = 1f - x * x;
			f = f * f * f;
			return f * default(G).Evaluate(hash, x);
		}
	}

	public struct Simplex2D<G> : INoise where G : struct, IGradient {

		public float4 GetNoise4 (float4x3 positions, SmallXXHash4 hash, int frequency) {
			positions *= frequency * (1f / sqrt(3f));
			float4 skew = (positions.c0 + positions.c2) * ((sqrt(3f) - 1f) / 2f);
			float4 sx = positions.c0 + skew, sz = positions.c2 + skew;
			int4
				x0 = (int4)floor(sx), x1 = x0 + 1,
				z0 = (int4)floor(sz), z1 = z0 + 1;

			bool4 xGz = sx - x0 > sz - z0;
			int4 xC = select(x0, x1, xGz), zC = select(z1, z0, xGz);

			SmallXXHash4
				h0 = hash.Eat(x0), h1 = hash.Eat(x1),
				hC = SmallXXHash4.Select(h0, h1, xGz);

			return default(G).EvaluateCombined(
				Kernel(h0.Eat(z0), x0, z0, positions) +
				Kernel(h1.Eat(z1), x1, z1, positions) +
				Kernel(hC.Eat(zC), xC, zC, positions)
			);
		}

		static float4 Kernel (
			SmallXXHash4 hash, float4 lx, float4 lz, float4x3 positions
		) {
			float4 unskew = (lx + lz) * ((3f - sqrt(3f)) / 6f);
			float4 x = positions.c0 - lx + unskew, z = positions.c2 - lz + unskew;
			float4 f = 0.5f - x * x - z * z;
			f = f * f * f * 8f;
			return max(0f, f) * default(G).Evaluate(hash, x, z);
		}
	}

	public struct Simplex3D<G> : INoise where G : struct, IGradient {

		public float4 GetNoise4 (float4x3 positions, SmallXXHash4 hash, int frequency) {
			positions *= frequency * 0.6f;
			float4 skew = (positions.c0 + positions.c1 + positions.c2) * (1f / 3f);
			float4
				sx = positions.c0 + skew,
				sy = positions.c1 + skew,
				sz = positions.c2 + skew;
			int4
				x0 = (int4)floor(sx), x1 = x0 + 1,
				y0 = (int4)floor(sy), y1 = y0 + 1,
				z0 = (int4)floor(sz), z1 = z0 + 1;

			bool4
				xGy = sx - x0 > sy - y0,
				xGz = sx - x0 > sz - z0,
				yGz = sy - y0 > sz - z0;

			bool4
				xA = xGy & xGz,
				xB = xGy | (xGz & yGz),
				yA = !xGy & yGz,
				yB = !xGy | (xGz & yGz),
				zA = (xGy & !xGz) | (!xGy & !yGz),
				zB = !(xGz & yGz);

			int4
				xCA = select(x0, x1, xA),
				xCB = select(x0, x1, xB),
				yCA = select(y0, y1, yA),
				yCB = select(y0, y1, yB),
				zCA = select(z0, z1, zA),
				zCB = select(z0, z1, zB);

			SmallXXHash4
				h0 = hash.Eat(x0), h1 = hash.Eat(x1),
				hA = SmallXXHash4.Select(h0, h1, xA),
				hB = SmallXXHash4.Select(h0, h1, xB);

			return default(G).EvaluateCombined(
				Kernel(h0.Eat(y0).Eat(z0), x0, y0, z0, positions) +
				Kernel(h1.Eat(y1).Eat(z1), x1, y1, z1, positions) +
				Kernel(hA.Eat(yCA).Eat(zCA), xCA, yCA, zCA, positions) +
				Kernel(hB.Eat(yCB).Eat(zCB), xCB, yCB, zCB, positions)
			);
		}

		static float4 Kernel (
			SmallXXHash4 hash, float4 lx, float4 ly, float4 lz, float4x3 positions
		) {
			float4 unskew = (lx + ly + lz) * (1f / 6f);
			float4
				x = positions.c0 - lx + unskew,
				y = positions.c1 - ly + unskew,
				z = positions.c2 - lz + unskew;
			float4 f = 0.5f - x * x - y * y - z * z;
			f = f * f * f * 8f;
			return max(0f, f) * default(G).Evaluate(hash, x, y, z);
		}
	}
}