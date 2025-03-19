using Unity.Burst;
using Unity.Collections;
using Unity.Jobs;
using Unity.Mathematics;
using UnityEngine;
using float4x4 = Unity.Mathematics.float4x4;
using quaternion = Unity.Mathematics.quaternion;

using static Unity.Mathematics.math;

public class HashVisualization : MonoBehaviour
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
    
    [System.Serializable]
    public struct SpaceTRS {

        public float3 translation, rotation, scale;
        
        public float3x4 Matrix {
            get {
                float4x4 m = float4x4.TRS(
                    translation, quaternion.EulerZXY(math.radians(rotation)), scale
                );
                return math.float3x4(m.c0.xyz, m.c1.xyz, m.c2.xyz, m.c3.xyz);
            }
        }
    }
    
    static int
        hashesId = Shader.PropertyToID("_Hashes"),
        positionsId = Shader.PropertyToID("_Positions"),
        normalsId = Shader.PropertyToID("_Normals"),
        configId = Shader.PropertyToID("_Config");

    [SerializeField]
    Mesh instanceMesh;

    [SerializeField]
    Material material;

    [SerializeField, Range(1, 512)]
    int resolution = 16;
    
    [SerializeField]
    int seed;
    
    [SerializeField, Range(-2f, 2f)]
    float verticalOffset = 1f;
    
    [SerializeField]
    SpaceTRS domain = new SpaceTRS {
        scale = 8f
    };
    
    [SerializeField, Range(-0.5f, 0.5f)]
    float displacement = 0.1f;
    
    [SerializeField]
    Shape shape;
    
    [SerializeField, Range(0.1f, 10f)]
    float instanceScale = 2f;
    
    public enum Shape { Plane, Sphere, Torus }

    static Shapes.ScheduleDelegate[] shapeJobs = {
        Shapes.Job<Shapes.Plane>.ScheduleParallel,
        Shapes.Job<Shapes.Sphere>.ScheduleParallel,
        Shapes.Job<Shapes.Torus>.ScheduleParallel
    };

    NativeArray<uint4> hashes;

    NativeArray<float3x4> positions, normals;

    ComputeBuffer hashesBuffer, positionsBuffer, normalsBuffer;

    MaterialPropertyBlock propertyBlock;
    
    bool isDirty;
    
    void OnEnable () {
        isDirty = true;

        int length = resolution * resolution;
        length = length / 4 + (length & 1);
        hashes = new NativeArray<uint4>(length, Allocator.Persistent);
        positions = new NativeArray<float3x4>(length, Allocator.Persistent);
        normals = new NativeArray<float3x4>(length, Allocator.Persistent);
        hashesBuffer = new ComputeBuffer(length * 4, 4);
        positionsBuffer = new ComputeBuffer(length * 4, 3 * 4);
        normalsBuffer = new ComputeBuffer(length * 4, 3 * 4);

        propertyBlock ??= new MaterialPropertyBlock();
        propertyBlock.SetBuffer(hashesId, hashesBuffer);
        propertyBlock.SetBuffer(positionsId, positionsBuffer);
        propertyBlock.SetBuffer(normalsId, normalsBuffer);
        propertyBlock.SetVector(configId, new Vector4(
            resolution, instanceScale / resolution, displacement
        ));
    }
    
    void OnDisable () {
        hashes.Dispose();
        positions.Dispose();
        normals.Dispose();
        hashesBuffer.Release();
        positionsBuffer.Release();
        normalsBuffer.Release();
        hashesBuffer = null;
        positionsBuffer = null;
        normalsBuffer = null;
    }

    void OnValidate () {
        if (hashesBuffer != null && enabled) {
            OnDisable();
            OnEnable();
        }
    }
    
    // Start is called before the first frame update
    void Start()
    {
        
    }
    
    Bounds bounds;

    // Update is called once per frame
    void Update()
    {
        if (isDirty || transform.hasChanged) {
            isDirty = false;
            transform.hasChanged = false;

            JobHandle handle = shapeJobs[(int)shape](
                positions, normals, resolution, transform.localToWorldMatrix, default
            );

            new HashJob {
                positions = positions,
                hashes = hashes,
                hash = SmallXXHash.Seed(seed),
                domainTRS = domain.Matrix
            }.ScheduleParallel(hashes.Length, resolution, handle).Complete();

            hashesBuffer.SetData(hashes.Reinterpret<uint>(4 * 4));
            positionsBuffer.SetData(positions.Reinterpret<float3>(3 * 4 * 4));
            normalsBuffer.SetData(normals.Reinterpret<float3>(3 * 4 * 4));
            
        }

       
        Graphics.DrawMeshInstancedProcedural(
            instanceMesh, 0, material, bounds, resolution * resolution, propertyBlock
        );
    }
    
    
    public static class Shapes {
        
        public struct Point4 {
            public float4x3 positions, normals;
        }
        
        public interface IShape {
            Point4 GetPoint4 (int i, float resolution, float invResolution);
        }
        
        public delegate JobHandle ScheduleDelegate (
            NativeArray<float3x4> positions, NativeArray<float3x4> normals,
            int resolution, float4x4 trs, JobHandle dependency
        );
        
        public struct Plane : IShape {

            public Point4 GetPoint4 (int i, float resolution, float invResolution) {
                float4x2 uv = IndexTo4UV(i, resolution, invResolution);
                return new Point4 {
                    positions = float4x3(uv.c0 - 0.5f, 0f, uv.c1 - 0.5f),
                    normals = float4x3(0f, 1f, 0f)
                };
            }
        }
        
        public struct Sphere : IShape {

            public Point4 GetPoint4 (int i, float resolution, float invResolution) {
                float4x2 uv = IndexTo4UV(i, resolution, invResolution);

                Point4 p;
                p.positions.c0 = uv.c0 - 0.5f;
                p.positions.c1 = uv.c1 - 0.5f;
                p.positions.c2 = 0.5f - abs(p.positions.c0) - abs(p.positions.c1);
                float4 offset = max(-p.positions.c2, 0f);
                p.positions.c0 += select(-offset, offset, p.positions.c0 < 0f);
                p.positions.c1 += select(-offset, offset, p.positions.c1 < 0f);

                float4 scale = 0.5f * rsqrt(
                    p.positions.c0 * p.positions.c0 +
                    p.positions.c1 * p.positions.c1 +
                    p.positions.c2 * p.positions.c2
                );
                p.positions.c0 *= scale;
                p.positions.c1 *= scale;
                p.positions.c2 *= scale;
                p.normals = p.positions;
                return p;
            }
        }
        
        public struct Torus : IShape {

            public Point4 GetPoint4 (int i, float resolution, float invResolution) {
                float4x2 uv = IndexTo4UV(i, resolution, invResolution);

                float r1 = 0.375f;
                float r2 = 0.125f;
                float4 s = r1 + r2 * cos(2f * PI * uv.c1);

                Point4 p;
                p.positions.c0 = s * sin(2f * PI * uv.c0);
                p.positions.c1 = r2 * sin(2f * PI * uv.c1);
                p.positions.c2 = s * cos(2f * PI * uv.c0);
                p.normals = p.positions;
                p.normals.c0 -= r1 * sin(2f * PI * uv.c0);
                p.normals.c2 -= r1 * cos(2f * PI * uv.c0);
                return p;
            }
        }
        
        public static float4x2 IndexTo4UV (int i, float resolution, float invResolution) {
            float4x2 uv;
            float4 i4 = 4f * i + float4(0f, 1f, 2f, 3f);
            uv.c1 = floor(invResolution * i4 + 0.00001f);
            uv.c0 = invResolution * (i4 - resolution * uv.c1 + 0.5f);
            uv.c1 = invResolution * (uv.c1 + 0.5f);
            return uv;
        }

        [BurstCompile(FloatPrecision.Standard, FloatMode.Fast, CompileSynchronously = true)]
        public struct Job<S> : IJobFor where S : struct, IShape {

            [WriteOnly]
            NativeArray<float3x4> positions, normals;

            public float resolution, invResolution;
            
            public float3x4 positionTRS, normalTRS;
            
            float4x3 TransformVectors (float3x4 trs, float4x3 p, float w = 1f) => float4x3(
                trs.c0.x * p.c0 + trs.c1.x * p.c1 + trs.c2.x * p.c2 + trs.c3.x * w,
                trs.c0.y * p.c0 + trs.c1.y * p.c1 + trs.c2.y * p.c2 + trs.c3.y * w,
                trs.c0.z * p.c0 + trs.c1.z * p.c1 + trs.c2.z * p.c2 + trs.c3.z * w
            );

            public void Execute (int i) {
                Point4 p = default(S).GetPoint4(i, resolution, invResolution);

                positions[i] = transpose(TransformVectors(positionTRS, p.positions));

                float3x4 n = transpose(TransformVectors(normalTRS, p.normals, 0f));
                normals[i] = float3x4(
                    normalize(n.c0), normalize(n.c1), normalize(n.c2), normalize(n.c3)
                );
            }
            
            public static JobHandle ScheduleParallel (
                NativeArray<float3x4> positions, NativeArray<float3x4> normals,int resolution, 
                float4x4 trs, JobHandle dependency
            ) {
                float4x4 tim = transpose(inverse(trs));
                return new Job<S> {
                    positions = positions,
                    normals = normals,
                    resolution = resolution,
                    invResolution = 1f / resolution,
                    positionTRS = float3x4(trs.c0.xyz, trs.c1.xyz, trs.c2.xyz, trs.c3.xyz),
                    normalTRS = float3x4(tim.c0.xyz, tim.c1.xyz, tim.c2.xyz, tim.c3.xyz)
                }.ScheduleParallel(positions.Length, resolution, dependency);
            }
        }
    }
}
