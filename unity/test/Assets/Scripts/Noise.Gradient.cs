﻿using Unity.Mathematics;

using static Unity.Mathematics.math;

public static partial class Noise {

    public interface IGradient
    {
        float4 Evaluate (SmallXXHash4 hash, float4 x);
        
        float4 Evaluate (SmallXXHash4 hash, float4 x, float4 y);

        float4 Evaluate (SmallXXHash4 hash, float4 x, float4 y, float4 z);
    }
    
    public struct Value : IGradient {

        public float4 Evaluate (SmallXXHash4 hash, float4 x) => hash.Floats01A * 2f - 1f;

        public float4 Evaluate (SmallXXHash4 hash, float4 x, float4 y) =>
            hash.Floats01A * 2f - 1f;

        public float4 Evaluate (SmallXXHash4 hash, float4 x, float4 y, float4 z) =>
            hash.Floats01A * 2f - 1f;
    }
    
    public struct Perlin : IGradient {

        public float4 Evaluate (SmallXXHash4 hash, float4 x) => (1f + hash.Floats01A) * select(-x, x, ((uint4)hash & 1 << 8) == 0);

        public float4 Evaluate (SmallXXHash4 hash, float4 x, float4 y) {
            float4 gx = hash.Floats01A * 2f - 1f;
            float4 gy = 0.5f - abs(gx);
            gx -= floor(gx + 0.5f);
            return (gx * x + gy * y) * (2f / 0.53528f);
        }

        public float4 Evaluate (SmallXXHash4 hash, float4 x, float4 y, float4 z) {
            float4 gx = hash.Floats01A * 2f - 1f, gy = hash.Floats01D * 2f - 1f;
            float4 gz = 1f - abs(gx) - abs(gy);
            float4 offset = max(-gz, 0f);
            gx += select(-offset, offset, gx < 0f);
            gy += select(-offset, offset, gy < 0f);
            return (gx * x + gy * y + gz * z) * (1f / 0.56290f);
        }
    }
}