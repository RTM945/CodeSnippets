using System.Collections;
using System.Collections.Generic;
using TMPro;
using UnityEngine;
using UnityEngine.UI;

public class Tile : MonoBehaviour
{
    public TextMeshProUGUI text;
    private int number;

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
