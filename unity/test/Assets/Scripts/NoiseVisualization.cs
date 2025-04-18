using Unity.Collections;
using Unity.Jobs;
using Unity.Mathematics;
using UnityEngine;

using static Noise;

public class NoiseVisualization : Visualization {

	static ScheduleDelegate[,] noiseJobs = {
		{
			Job<Lattice1D<LatticeNormal, Perlin>>.ScheduleParallel,
			Job<Lattice1D<LatticeTiling, Perlin>>.ScheduleParallel,
			Job<Lattice2D<LatticeNormal, Perlin>>.ScheduleParallel,
			Job<Lattice2D<LatticeTiling, Perlin>>.ScheduleParallel,
			Job<Lattice3D<LatticeNormal, Perlin>>.ScheduleParallel,
			Job<Lattice3D<LatticeTiling, Perlin>>.ScheduleParallel
		},
		{
			Job<Lattice1D<LatticeNormal, Turbulence<Perlin>>>.ScheduleParallel,
			Job<Lattice1D<LatticeTiling, Turbulence<Perlin>>>.ScheduleParallel,
			Job<Lattice2D<LatticeNormal, Turbulence<Perlin>>>.ScheduleParallel,
			Job<Lattice2D<LatticeTiling, Turbulence<Perlin>>>.ScheduleParallel,
			Job<Lattice3D<LatticeNormal, Turbulence<Perlin>>>.ScheduleParallel,
			Job<Lattice3D<LatticeTiling, Turbulence<Perlin>>>.ScheduleParallel
		},
		{
			Job<Lattice1D<LatticeNormal, Value>>.ScheduleParallel,
			Job<Lattice1D<LatticeTiling, Value>>.ScheduleParallel,
			Job<Lattice2D<LatticeNormal, Value>>.ScheduleParallel,
			Job<Lattice2D<LatticeTiling, Value>>.ScheduleParallel,
			Job<Lattice3D<LatticeNormal, Value>>.ScheduleParallel,
			Job<Lattice3D<LatticeTiling, Value>>.ScheduleParallel
		},
		{
			Job<Lattice1D<LatticeNormal, Turbulence<Value>>>.ScheduleParallel,
			Job<Lattice1D<LatticeTiling, Turbulence<Value>>>.ScheduleParallel,
			Job<Lattice2D<LatticeNormal, Turbulence<Value>>>.ScheduleParallel,
			Job<Lattice2D<LatticeTiling, Turbulence<Value>>>.ScheduleParallel,
			Job<Lattice3D<LatticeNormal, Turbulence<Value>>>.ScheduleParallel,
			Job<Lattice3D<LatticeTiling, Turbulence<Value>>>.ScheduleParallel
		},
		{
			Job<Simplex1D<Simplex>>.ScheduleParallel,
			Job<Simplex1D<Simplex>>.ScheduleParallel,
			Job<Simplex2D<Simplex>>.ScheduleParallel,
			Job<Simplex2D<Simplex>>.ScheduleParallel,
			Job<Simplex3D<Simplex>>.ScheduleParallel,
			Job<Simplex3D<Simplex>>.ScheduleParallel
		},
		{
			Job<Simplex1D<Turbulence<Simplex>>>.ScheduleParallel,
			Job<Simplex1D<Turbulence<Simplex>>>.ScheduleParallel,
			Job<Simplex2D<Turbulence<Simplex>>>.ScheduleParallel,
			Job<Simplex2D<Turbulence<Simplex>>>.ScheduleParallel,
			Job<Simplex3D<Turbulence<Simplex>>>.ScheduleParallel,
			Job<Simplex3D<Turbulence<Simplex>>>.ScheduleParallel
		},
		{
			Job<Simplex1D<Value>>.ScheduleParallel,
			Job<Simplex1D<Value>>.ScheduleParallel,
			Job<Simplex2D<Value>>.ScheduleParallel,
			Job<Simplex2D<Value>>.ScheduleParallel,
			Job<Simplex3D<Value>>.ScheduleParallel,
			Job<Simplex3D<Value>>.ScheduleParallel
		},
		{
			Job<Simplex1D<Turbulence<Value>>>.ScheduleParallel,
			Job<Simplex1D<Turbulence<Value>>>.ScheduleParallel,
			Job<Simplex2D<Turbulence<Value>>>.ScheduleParallel,
			Job<Simplex2D<Turbulence<Value>>>.ScheduleParallel,
			Job<Simplex3D<Turbulence<Value>>>.ScheduleParallel,
			Job<Simplex3D<Turbulence<Value>>>.ScheduleParallel
		},
		{
			Job<Voronoi1D<LatticeNormal, Worley, F1>>.ScheduleParallel,
			Job<Voronoi1D<LatticeTiling, Worley, F1>>.ScheduleParallel,
			Job<Voronoi2D<LatticeNormal, Worley, F1>>.ScheduleParallel,
			Job<Voronoi2D<LatticeTiling, Worley, F1>>.ScheduleParallel,
			Job<Voronoi3D<LatticeNormal, Worley, F1>>.ScheduleParallel,
			Job<Voronoi3D<LatticeTiling, Worley, F1>>.ScheduleParallel
		},
		{
			Job<Voronoi1D<LatticeNormal, Worley, F2>>.ScheduleParallel,
			Job<Voronoi1D<LatticeTiling, Worley, F2>>.ScheduleParallel,
			Job<Voronoi2D<LatticeNormal, Worley, F2>>.ScheduleParallel,
			Job<Voronoi2D<LatticeTiling, Worley, F2>>.ScheduleParallel,
			Job<Voronoi3D<LatticeNormal, Worley, F2>>.ScheduleParallel,
			Job<Voronoi3D<LatticeTiling, Worley, F2>>.ScheduleParallel
		},
		{
			Job<Voronoi1D<LatticeNormal, Worley, F2MinusF1>>.ScheduleParallel,
			Job<Voronoi1D<LatticeTiling, Worley, F2MinusF1>>.ScheduleParallel,
			Job<Voronoi2D<LatticeNormal, Worley, F2MinusF1>>.ScheduleParallel,
			Job<Voronoi2D<LatticeTiling, Worley, F2MinusF1>>.ScheduleParallel,
			Job<Voronoi3D<LatticeNormal, Worley, F2MinusF1>>.ScheduleParallel,
			Job<Voronoi3D<LatticeTiling, Worley, F2MinusF1>>.ScheduleParallel
		},
		{
			Job<Voronoi1D<LatticeNormal, Worley, F1>>.ScheduleParallel,
			Job<Voronoi1D<LatticeTiling, Worley, F1>>.ScheduleParallel,
			Job<Voronoi2D<LatticeNormal, Chebyshev, F1>>.ScheduleParallel,
			Job<Voronoi2D<LatticeTiling, Chebyshev, F1>>.ScheduleParallel,
			Job<Voronoi3D<LatticeNormal, Chebyshev, F1>>.ScheduleParallel,
			Job<Voronoi3D<LatticeTiling, Chebyshev, F1>>.ScheduleParallel
		},
		{
			Job<Voronoi1D<LatticeNormal, Worley, F2>>.ScheduleParallel,
			Job<Voronoi1D<LatticeTiling, Worley, F2>>.ScheduleParallel,
			Job<Voronoi2D<LatticeNormal, Chebyshev, F2>>.ScheduleParallel,
			Job<Voronoi2D<LatticeTiling, Chebyshev, F2>>.ScheduleParallel,
			Job<Voronoi3D<LatticeNormal, Chebyshev, F2>>.ScheduleParallel,
			Job<Voronoi3D<LatticeTiling, Chebyshev, F2>>.ScheduleParallel
		},
		{
			Job<Voronoi1D<LatticeNormal, Worley, F2MinusF1>>.ScheduleParallel,
			Job<Voronoi1D<LatticeTiling, Worley, F2MinusF1>>.ScheduleParallel,
			Job<Voronoi2D<LatticeNormal, Chebyshev, F2MinusF1>>.ScheduleParallel,
			Job<Voronoi2D<LatticeTiling, Chebyshev, F2MinusF1>>.ScheduleParallel,
			Job<Voronoi3D<LatticeNormal, Chebyshev, F2MinusF1>>.ScheduleParallel,
			Job<Voronoi3D<LatticeTiling, Chebyshev, F2MinusF1>>.ScheduleParallel
		}
	};

	static int noiseId = Shader.PropertyToID("_Noise");

	[SerializeField]
	Settings noiseSettings = Settings.Default;

	public enum NoiseType {
		Perlin, PerlinTurbulence, Value, ValueTurbulence,
		Simplex, SimplexTurbulence, SimplexValue, SimplexValueTurbulence,
		VoronoiWorleyF1, VoronoiWorleyF2, VoronoiWorleyF2MinusF1,
		VoronoiChebyshevF1, VoronoiChebyshevF2, VoronoiChebyshevF2MinusF1
	}

	[SerializeField]
	NoiseType type;

	[SerializeField, Range(1, 3)]
	int dimensions = 3;

	[SerializeField]
	bool tiling;

	[SerializeField]
	SpaceTRS domain = new SpaceTRS {
		scale = 8f
	};
	
	NativeArray<float4> noise;

	ComputeBuffer noiseBuffer;
	
	protected override void EnableVisualization (
		int dataLength, MaterialPropertyBlock propertyBlock
	) {
		noise = new NativeArray<float4>(dataLength, Allocator.Persistent);
		noiseBuffer = new ComputeBuffer(dataLength * 4, 4);
		propertyBlock.SetBuffer(noiseId, noiseBuffer);
	}

	protected override void DisableVisualization () {
		noise.Dispose();
		noiseBuffer.Release();
		noiseBuffer = null;
	}

	protected override void UpdateVisualization (
		NativeArray<float3x4> positions, int resolution, JobHandle handle
	) {
		noiseJobs[(int)type, 2 * dimensions - (tiling ? 1 : 2)](
			positions, noise, noiseSettings, domain, resolution, handle
		).Complete();
		noiseBuffer.SetData(noise.Reinterpret<float>(4 * 4));
	}
}