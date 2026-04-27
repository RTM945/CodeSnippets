using System;
using System.Collections;
using System.Collections.Generic;
using TMPro;
using UnityEngine;
using UnityEngine.UI;
using DG.Tweening;

public class Tile : MonoBehaviour
{
    public Text text;
    private int number;
    public Image bg;
    public RectTransform rect;

    private readonly Dictionary<int, Color32> _bgColors = new Dictionary<int, Color32>()
    {
        { 2, new Color32(238, 228, 218, 255) },
        { 4, new Color32(237, 224, 200, 255) },
        { 8, new Color32(242, 177, 121, 255) },
        { 16, new Color32(245, 149, 99, 255) },
        { 32, new Color32(246, 124, 95, 255) },
        { 64, new Color32(246, 94, 59, 255) },
        { 128, new Color32(237, 207, 114, 255) },
        { 256, new Color32(237, 204, 97, 255) },
        { 512, new Color32(237, 200, 80, 255) },
        { 1024, new Color32(237, 197, 63, 255) },
        { 2048, new Color32(237, 194, 46, 255) },
        { 2049, new Color32(60, 58, 50, 255) },
    };

    private readonly Dictionary<int, Color32> _textColors = new Dictionary<int, Color32>()
    {
        { 2, new Color32(119, 110, 101, 255) },
        { 4, new Color32(119, 110, 101, 255) },
        { 8, new Color32(249, 246, 242, 255) },
        { 16, new Color32(249, 246, 242, 255) },
        { 32, new Color32(249, 246, 242, 255) },
        { 64, new Color32(249, 246, 242, 255) },
        { 128, new Color32(249, 246, 242, 255) },
        { 256, new Color32(249, 246, 242, 255) },
        { 512, new Color32(249, 246, 242, 255) },
        { 1024, new Color32(249, 246, 242, 255) },
        { 2048, new Color32(249, 246, 242, 255) },
    };

    private void Awake()
    {
        rect = GetComponent<RectTransform>();
    }

    public void SetNumber(int number)
    {
        this.number = number;
        text.text = number == 0 ? "" : number.ToString();
        bg.color = number > 2048 ? _bgColors[2049] : _bgColors[number];
        text.color = number > 2048 ? _textColors[2048] : _textColors[number];
        
        if (number < 100)
        {
            text.fontSize = 60;
        }
        else if (number < 1000)
        {
            text.fontSize = 50;
        }
        else if (number < 10000)
        {
            text.fontSize = 40;
        }
        else
        {
            text.fontSize = 32;
        }
        
        transform.localScale = Vector3.zero;
        
        transform
            .DOScale(1f, 0.5f)
            .SetEase(Ease.OutBack, 2f);
    }

    public int GetNumber()
    {
        return number;
    }
    
}
