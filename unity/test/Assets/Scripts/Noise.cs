using Unity.Burst;
using Unity.Collections;
using Unity.Jobs;
using Unity.Mathematics;

using static Unity.Mathematics.math;

public static partial class Noise {

    public interface INoise {
        float4 GetNoise4 (float4x3 positions, SmallXXHash4 hash, int frequency);
    }

    [BurstCompile(FloatPrecision.Standard, FloatMode.Fast, CompileSynchronously = true)]
    public struct Job<N> : IJobFor where N : struct, INoise
    {
        [ReadOnly]
        public NativeArray<float3x4> positions;

        [WriteOnly]
        public NativeArray<float4> noise;

        public Settings settings;

        public float3x4 domainTRS;

        public void Execute (int i) {
            float4x3 position = domainTRS.TransformVectors(transpose(positions[i]));
            var hash = SmallXXHash4.Seed(settings.seed);
            int frequency = settings.frequency;
            float amplitude = 1f, amplitudeSum = 0f;
            float4 sum = 0f;

            for (int o = 0; o < settings.octaves; o++) {
                sum += amplitude * default(N).GetNoise4(position, hash + o, frequency);
                frequency *= settings.lacunarity;
                amplitude *= settings.persistence;
                amplitudeSum += amplitude;
            }
            noise[i] = sum / amplitudeSum;
        }
        public static JobHandle ScheduleParallel (
            NativeArray<float3x4> positions, NativeArray<float4> noise,
            Settings settings, SpaceTRS domainTRS, int resolution, JobHandle dependency
        ) => new Job<N> {
            positions = positions,
            noise = noise,
            settings = settings,
            domainTRS = domainTRS.Matrix,
        }.ScheduleParallel(positions.Length, resolution, dependency);
    }
	
    public delegate JobHandle ScheduleDelegate (
        NativeArray<float3x4> positions, NativeArray<float4> noise,
        Settings settings, SpaceTRS domainTRS, int resolution, JobHandle dependency
    );
}