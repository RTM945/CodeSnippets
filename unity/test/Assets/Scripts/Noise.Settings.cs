using System;
using Unity.Burst;
using UnityEngine;

public static partial class Noise {

    [Serializable]
    public struct Settings {

        public int seed;
        
        [Min(1)]
        public int frequency;

        [Range(1, 6)]
        public int octaves;
        
        [Range(2, 4)]
        public int lacunarity;
        
        [Range(0f, 1f)]
        public float persistence;


        public static Settings Default => new Settings {
            frequency = 4,
            octaves = 1,
            lacunarity = 2,
            persistence = 0.5f
        };
    }
	
}