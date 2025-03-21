using Unity.Burst;
using Unity.Collections;
using Unity.Jobs;
using Unity.Mathematics;
using UnityEngine;
using float4x4 = Unity.Mathematics.float4x4;
using quaternion = Unity.Mathematics.quaternion;

using static Unity.Mathematics.math;

public class HashVisualization : Visualization 
{
    
    [BurstCompile(FloatPrecision.Standard, FloatMode.Fast, CompileSynchronously = true)]
    struct HashJob : IJobFor {

        [ReadOnly]
        public NativeArray<float3x4> positions;
		
        [WriteOnly]
        public NativeArray<uint4> hashes;

        public SmallXXHash4 hash;
        
        public float3x4 domainTRS;
        
        float4x3 TransformPositions (float3x4 trs, float4x3 p) => float4x3(
            trs.c0.x * p.c0 + trs.c1.x * p.c1 + trs.c2.x * p.c2 + trs.c3.x,
            trs.c0.y * p.c0 + trs.c1.y * p.c1 + trs.c2.y * p.c2 + trs.c3.y,
            trs.c0.z * p.c0 + trs.c1.z * p.c1 + trs.c2.z * p.c2 + trs.c3.z
        );
        
        public void Execute(int i) {
            float4x3 p = TransformPositions(domainTRS, transpose(positions[i]));

            int4 u = (int4)floor(p.c0);
            int4 v = (int4)floor(p.c1);
            int4 w = (int4)floor(p.c2);

            hashes[i] = hash.Eat(u).Eat(v).Eat(w);
        }
    }
    
    public readonly struct SmallXXHash {

        const uint primeA = 0b10011110001101110111100110110001;
        const uint primeB = 0b10000101111010111100101001110111;
        const uint primeC = 0b11000010101100101010111000111101;
        const uint primeD = 0b00100111110101001110101100101111;
        const uint primeE = 0b00010110010101100110011110110001;
        readonly uint accumulator;

        public SmallXXHash (uint accumulator) {
            this.accumulator = accumulator;
        }
        
        
        public static SmallXXHash Seed (int seed) => (uint)seed + primeE;
        
        public static implicit operator SmallXXHash (uint accumulator) =>
            new SmallXXHash(accumulator);
        
        public static implicit operator uint (SmallXXHash hash) {
            uint avalanche = hash.accumulator;
            avalanche ^= avalanche >> 15;
            avalanche *= primeB;
            avalanche ^= avalanche >> 13;
            avalanche *= primeC;
            avalanche ^= avalanche >> 16;
            return avalanche;
        }
        
        public SmallXXHash Eat (int data) =>
            RotateLeft(accumulator + (uint)data * primeC, 17) * primeD;

        public SmallXXHash Eat (byte data) =>
            RotateLeft(accumulator + data * primeE, 11) * primeA;
        
        static uint RotateLeft (uint data, int steps) =>
            (data << steps) | (data >> 32 - steps);
        
        public static implicit operator SmallXXHash4 (SmallXXHash hash) =>
            new SmallXXHash4(hash.accumulator);
    }

    public readonly struct SmallXXHash4
    {
        const uint primeB = 0b10000101111010111100101001110111;
        const uint primeC = 0b11000010101100101010111000111101;
        const uint primeD = 0b00100111110101001110101100101111;
        const uint primeE = 0b00010110010101100110011110110001;
        
        readonly uint4 accumulator;

        public SmallXXHash4 (uint4 accumulator) {
            this.accumulator = accumulator;
        }

        public static implicit operator SmallXXHash4 (uint4 accumulator) =>
            new SmallXXHash4(accumulator);

        public static SmallXXHash4 Seed (int4 seed) => (uint4)seed + primeE;
	
        static uint4 RotateLeft (uint4 data, int steps) =>
            (data << steps) | (data >> 32 - steps);

        public SmallXXHash4 Eat (int4 data) =>
            RotateLeft(accumulator + (uint4)data * primeC, 17) * primeD;

        public static implicit operator uint4 (SmallXXHash4 hash) 
        {
            uint4 avalanche = hash.accumulator;
            avalanche ^= avalanche >> 15;
            avalanche *= primeB;
            avalanche ^= avalanche >> 13;
            avalanche *= primeC;
            avalanche ^= avalanche >> 16;
            return avalanche;
        }
    }

    private static int hashesId = Shader.PropertyToID("_Hashes");
    
    [SerializeField]
    int seed;
    
    [SerializeField]
    SpaceTRS domain = new SpaceTRS {
        scale = 8f
    };
    
    NativeArray<uint4> hashes;

    ComputeBuffer hashesBuffer;
    
    protected override void EnableVisualization (int dataLength, MaterialPropertyBlock propertyBlock) {
        hashes = new NativeArray<uint4>(dataLength, Allocator.Persistent);
        hashesBuffer = new ComputeBuffer(dataLength * 4, 4);
        propertyBlock.SetBuffer(hashesId, hashesBuffer);
    }

    protected override void DisableVisualization()
    {
        hashes.Dispose();
        hashesBuffer.Release();
        hashesBuffer = null;
    }
    

    // Update is called once per frame
    protected override void UpdateVisualization (
        NativeArray<float3x4> positions, int resolution, JobHandle handle
    ) 
    {
        new HashJob {
            positions = positions,
            hashes = hashes,
            hash = SmallXXHash.Seed(seed),
            domainTRS = domain.Matrix
        }.ScheduleParallel(hashes.Length, resolution, handle).Complete();

        hashesBuffer.SetData(hashes.Reinterpret<uint>(4 * 4));
    }
}
