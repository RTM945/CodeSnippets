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
    private Image bg;
    
    void Awake()
    {
        bg = GetComponent<Image>();
    }
    
    public void MoveTo(Vector3 target)
    {
        transform.DOLocalMove(target, 0.15f)
            .SetEase(Ease.OutCubic);
    }

    public void SpawnNumber(int number)
    {
        // 底图变色 
        if (number == 2)
        {
            bg.color = new Color(0, 0,0, 255);
            // 咋log来着
            Debug.Log(bg.color);
        }
        else if (number == 4)
        {
            bg.color = new Color(238, 225,201, 255);
            Debug.Log(bg.color);
        }
        // 文字变色
        text.color = new Color(117, 110,82, 255);
        // dotween 动画
        transform.localScale = Vector3.zero;

        transform
            .DOScale(1f, 0.75f)
            .SetEase(Ease.OutBack, 1.5f);
    }
    
    public void PlayMerge()
    {
        transform.DOScale(1.2f, 0.1f)
            .SetEase(Ease.OutBack)
            .OnComplete(() =>
            {
                transform.DOScale(1f, 0.1f);
            });
    }

    public void SetNumber(int number)
    {
        this.number = number;
        text.text = number == 0 ? "" : number.ToString();
    }

    public int GetNumber()
    {
        return number;
    }
    
}
