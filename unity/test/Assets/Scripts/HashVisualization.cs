using Unity.Burst;
using Unity.Collections;
using Unity.Jobs;
using Unity.Mathematics;
using UnityEngine;

using static Unity.Mathematics.math;

public class HashVisualization : MonoBehaviour
{
    
    [BurstCompile(FloatPrecision.Standard, FloatMode.Fast, CompileSynchronously = true)]
    struct HashJob : IJobFor {

        [WriteOnly]
        public NativeArray<uint> hashes;
        public int resolution;
        
        public SmallXXHash hash;

        public float invResolution;
        
        public void Execute(int i) {
            int v = (int)floor(invResolution * i + 0.00001f);
            int u = i - resolution * v - resolution / 2;
            v -= resolution / 2;
            hashes[i] = hash.Eat(u).Eat(v);
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
    }
    
    static int
        hashesId = Shader.PropertyToID("_Hashes"),
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

    NativeArray<uint> hashes;

    ComputeBuffer hashesBuffer;

    MaterialPropertyBlock propertyBlock;
    
    void OnEnable () {
        int length = resolution * resolution;
        hashes = new NativeArray<uint>(length, Allocator.Persistent);
        hashesBuffer = new ComputeBuffer(length, 4);

        new HashJob {
            hashes = hashes,
            resolution = resolution,
            invResolution = 1f / resolution,
            hash = SmallXXHash.Seed(seed)
        }.ScheduleParallel(hashes.Length, resolution, default).Complete();

        hashesBuffer.SetData(hashes);

        propertyBlock ??= new MaterialPropertyBlock();
        propertyBlock.SetBuffer(hashesId, hashesBuffer);
        propertyBlock.SetVector(configId, new Vector4(
            resolution, 1f / resolution, verticalOffset / resolution
        ));
    }
    
    void OnDisable () {
        hashes.Dispose();
        hashesBuffer.Release();
        hashesBuffer = null;
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

    // Update is called once per frame
    void Update()
    {
        Graphics.DrawMeshInstancedProcedural(
            instanceMesh, 0, material, new Bounds(Vector3.zero, Vector3.one),
            hashes.Length, propertyBlock
        );
    }
}
