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
    
    public void MoveTo(Vector3 target)
    {
        transform.DOLocalMove(target, 0.15f)
            .SetEase(Ease.OutCubic);
    }

    public void SpawnNumber(int number)
    {
        text.transform.localScale = Vector3.zero;
        Sequence popUpSequence = DOTween.Sequence();
        popUpSequence.Append(text.transform.DOScale(1.2f, 0.3f).SetEase(Ease.OutBack));
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
