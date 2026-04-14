using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class GameManager : MonoBehaviour
{
    
    public GameObject tilePrefab;
    public Transform board;
    
    Tile[] tiles = new Tile[16];
    int[,] grid = new int[4,4];
    
    // Start is called before the first frame update
    void Start()
    {
        Debug.Log("2048 Start");

        CreateGrid();
        SpawnNumber();
        SpawnNumber();
    }

    void CreateGrid()
    {
        for (int i = 0; i < 16; i++)
        {
            GameObject obj = Instantiate(tilePrefab, board);
            tiles[i] = obj.GetComponent<Tile>();
        }
    }

    void SpawnNumber()
    {
        int index = Random.Range(0, 16);
        
        int x = index % 4;
        int y = index / 4;
        
        grid[x, y] = 2;
        
        tiles[index].SetNumber(2);
    }

    // Update is called once per frame
    void Update()
    {
        if (Input.GetKeyDown(KeyCode.LeftArrow))
        {
            MoveLeft();
        }
    }

    void MoveLeft()
    {
        for (int y = 0; y < 4; y++)
        {
            for (int x = 1; x < 4; x++)
            {
                if (grid[x, y] != 0)
                {
                    int newX = x;
                    while (newX > 0 && grid[newX - 1, y] == 0)
                    {
                        grid[newX - 1, y] = grid[newX, y];
                        grid[newX, y] = 0;
                        newX--;
                    }
                }
            }
        }

        UpdateView();
    }

    void UpdateView()
    {
        for (int i = 0; i < 16; i++)
        {
            int x = i % 4;
            int y = i / 4;
            
            tiles[i].SetNumber(grid[x, y]);
        }
    }
}
