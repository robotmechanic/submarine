﻿using UnityEngine;
using System;

namespace Submarine
{
    [Serializable]
    public static class Constants
    {
        public const string Version = "0.1";
        public const int FrameRate = 30;
        public const int MapLength = 400;

        public const string SubmarinePrefab = "Submarines/Submarine";
        public const string TorpedoPrefab = "Torpedos/Torpedo";
        public const string ExplosionEffectPrefab = "Effects/Explosion";

        public static float Fps { get { return 1f / Time.deltaTime; } }
        public static float FpsRate { get { return Fps / FrameRate; } }
    }
}
