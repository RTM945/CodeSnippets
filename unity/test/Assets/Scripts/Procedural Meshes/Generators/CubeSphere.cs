using Unity.Mathematics;
using UnityEngine;

using static Unity.Mathematics.math;

namespace ProceduralMeshes.Generators {

	public struct CubeSphere : IMeshGenerator {

		struct Side {
			public int id;
			public float3 uvOrigin, uVector, vVector;
		}

		public Bounds Bounds => new Bounds(Vector3.zero, new Vector3(2f, 2f, 2f));

		public int VertexCount => 6 * 4 * Resolution * Resolution;

		public int IndexCount => 6 * 6 * Resolution * Resolution;

		public int JobLength => 6 * Resolution;

		public int Resolution { get; set; }

		public void Execute<S> (int i, S streams) where S : struct, IMeshStreams {
			int u = i / 6;
			Side side = GetSide(i - 6 * u);
			int vi = 4 * Resolution * (Resolution * side.id + u);
			int ti = 2 * Resolution * (Resolution * side.id + u);

			float3 uA = side.uvOrigin + side.uVector * u / Resolution;
			float3 uB = side.uvOrigin + side.uVector * (u + 1) / Resolution;
			float3 pA = CubeToSphere(uA), pB = CubeToSphere(uB);

			var vertex = new Vertex();
			vertex.tangent = float4(normalize(pB - pA), -1f);

			for (int v = 1; v <= Resolution; v++, vi += 4, ti += 2) {
				float3 pC = CubeToSphere(uA + side.vVector * v / Resolution);
				float3 pD = CubeToSphere(uB + side.vVector * v / Resolution);

				vertex.position = pA;
				vertex.normal = normalize(cross(pC - pA, vertex.tangent.xyz));
				vertex.texCoord0 = 0f;
				streams.SetVertex(vi + 0, vertex);

				vertex.position = pB;
				vertex.normal = normalize(cross(pD - pB, vertex.tangent.xyz));
				vertex.texCoord0 = float2(1f, 0f);
				streams.SetVertex(vi + 1, vertex);

				vertex.position = pC;
				vertex.tangent.xyz = normalize(pD - pC);
				vertex.normal = normalize(cross(pC - pA, vertex.tangent.xyz));
				vertex.texCoord0 = float2(0f, 1f);
				streams.SetVertex(vi + 2, vertex);

				vertex.position = pD;
				vertex.normal = pD;
				vertex.normal = normalize(cross(pD - pB, vertex.tangent.xyz));
				vertex.texCoord0 = 1f;
				streams.SetVertex(vi + 3, vertex);

				streams.SetTriangle(ti + 0, vi + int3(0, 2, 1));
				streams.SetTriangle(ti + 1, vi + int3(1, 2, 3));

				pA = pC;
				pB = pD;
			}
		}

		static float3 CubeToSphere (float3 p) => p * sqrt(
			1f - ((p * p).yxx + (p * p).zzy) / 2f + (p * p).yxx * (p * p).zzy / 3f
		);

		static Side GetSide (int id) => id switch
		{
			0 => new Side {
				id = id,
				uvOrigin = -1f,
				uVector = 2f * right(),
				vVector = 2f * up()
			},
			1 => new Side {
				id = id,
				uvOrigin = float3(1f, -1f, -1f),
				uVector = 2f * forward(),
				vVector = 2f * up()
			},
			2 => new Side {
				id = id,
				uvOrigin = -1f,
				uVector = 2f * forward(),
				vVector = 2f * right()
			},
			3 => new Side {
				id = id,
				uvOrigin = float3(-1f, -1f, 1f),
				uVector = 2f * up(),
				vVector = 2f * right()
			},
			4 => new Side {
				id = id,
				uvOrigin = -1f,
				uVector = 2f * up(),
				vVector = 2f * forward()
			},
			_ => new Side {
				id = id,
				uvOrigin = float3(-1f, 1f, -1f),
				uVector = 2f * right(),
				vVector = 2f * forward()
			}
		};
	}
}