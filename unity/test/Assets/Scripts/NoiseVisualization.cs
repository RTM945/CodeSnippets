﻿using Unity.Collections;
using Unity.Jobs;
using Unity.Mathematics;
using UnityEngine;

using static Noise;

public class NoiseVisualization : Visualization {

    static int noiseId = Shader.PropertyToID("_Noise");
	
    [SerializeField]
    int seed;

    [SerializeField]
    SpaceTRS domain = new SpaceTRS {
        scale = 8f
    };
    
    static ScheduleDelegate[] noiseJobs = {
        Job<Lattice1D>.ScheduleParallel,
        Job<Lattice2D>.ScheduleParallel,
        Job<Lattice3D>.ScheduleParallel
    };
    
    [SerializeField, Range(1, 3)]
    int dimensions = 3;

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
        noiseJobs[dimensions - 1](
            positions, noise, seed, domain, resolution, handle
        ).Complete();
        noiseBuffer.SetData(noise.Reinterpret<float>(4 * 4));
    }
    
}