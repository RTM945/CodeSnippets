using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class GameManager : MonoBehaviour
{

    public GameObject tilePrefab;
    public Transform board;

    Tile[] tiles = new Tile[16];
    int[,] grid = new int[4, 4];

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

    List<int> GetEmptyGrid()
    {
        List<int> empty = new List<int>();

        for (int i = 0; i < 16; i++)
        {
            int x = i % 4;
            int y = i / 4;

            if (grid[x, y] == 0)
            {
                empty.Add(i);
            }
        }
        return empty;
    }

    void SpawnNumber()
    {
        List<int> empty = GetEmptyGrid();

        if (empty.Count == 0)
        {
            return;
        }
        
        int index = empty[Random.Range(0, empty.Count)];

        int x = index % 4;
        int y = index / 4;
        
        int number = Random.value < 0.9f ? 2 : 4;

        grid[x, y] = number;

        tiles[index].SetNumber(number);
        
    }

    bool IsGameOver()
    {
        return GetEmptyGrid().Count == 0;
    }

    void CompressLeft()
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

    void CompressRight()
    {
        for (int y = 0; y < 4; y++)
        {
            for (int x = 2; x >= 0; x--)
            {
                if (grid[x, y] != 0)
                {
                    int newX = x;
                    while (newX < 3 && grid[newX + 1, y] == 0)
                    {
                        grid[newX + 1, y] = grid[newX, y];
                        grid[newX, y] = 0;
                        newX++;
                    }
                }
            }
        }
        UpdateView();
    }

    void CompressUp()
    {
        for (int y = 1; y < 4; y++)
        {
            for (int x = 0; x < 4; x++)
            {
                if (grid[x, y] != 0)
                {
                    int newY = y;
                    while (newY > 0 && grid[x, newY - 1] == 0)
                    {
                        grid[x, newY - 1] = grid[x, newY];
                        grid[x, newY] = 0;
                        newY--;
                    }
                }
            }
        }
        UpdateView();
    }

    void CompressDown()
    {
        for (int y = 2; y >= 0; y--)
        {
            for (int x = 0; x < 4; x++)
            {
                if (grid[x, y] != 0)
                {
                    int newY = y;
                    while (newY < 3 && grid[x, newY + 1] == 0)
                    {
                        grid[x, newY + 1] = grid[x, newY];
                        grid[x, newY] = 0;
                        newY++;
                    }
                }
            }
        }
        UpdateView();
    }

    void MergeLeft()
    {
        for (int y = 0; y < 4; y++)
        {
            for (int x = 0; x < 3; x++)
            {
                if (grid[x, y] != 0 && grid[x, y] == grid[x + 1, y])
                {
                    grid[x, y] *= 2;
                    grid[x + 1, y] = 0;
                }
            }
        }
    }

    void MergeRight()
    {
        for (int y = 0; y < 4; y++)
        {
            for (int x = 3; x > 0; x--)
            {
                if (grid[x, y] != 0 && grid[x, y] == grid[x - 1, y])
                {
                    grid[x, y] *= 2;
                    grid[x - 1, y] = 0;
                }
            }
        }
    }

    void MergeUp()
    {
        for (int y = 0; y < 3; y++)
        {
            for (int x = 0; x < 4; x++)
            {
                if (grid[x, y] != 0 && grid[x, y] == grid[x, y + 1])
                {
                    grid[x, y] *= 2;
                    grid[x, y + 1] = 0;
                }
            }
        }
    }

    void MergeDown()
    {
        for (int y = 3; y > 0; y--)
        {
            for (int x = 0; x < 4; x++)
            {
                if (grid[x, y] != 0 && grid[x, y] == grid[x, y - 1])
                {
                    grid[x, y] *= 2;
                    grid[x, y - 1] = 0;
                }
            }
        }
    }

    void MoveLeft()
    {   
        CompressLeft();
        MergeLeft();
        CompressLeft();
        SpawnNumber();
        UpdateView();
    }

    void MoveRight()
    {
        CompressRight();
        MergeRight();
        CompressRight();
        SpawnNumber();
        UpdateView();
    }

    void MoveUp()
    {
        CompressUp();
        MergeUp();
        CompressUp();
        SpawnNumber();
        UpdateView();
    }

    void MoveDown()
    {
        CompressDown();
        MergeDown();
        CompressDown();
        SpawnNumber();
        UpdateView();
    }

    void UpdateView()
    {
        for(int i = 0; i < 16; i++)
        {
            int x = i % 4;
            int y = i / 4;

            tiles[i].SetNumber(grid[x, y]);
        }
    }

    // Update is called once per frame
    void Update()
    {
        if (Input.GetKeyDown(KeyCode.LeftArrow))
        {
            MoveLeft();
        }

        if (Input.GetKeyDown(KeyCode.RightArrow))
        {
            MoveRight();
        }
        
        if (Input.GetKeyDown(KeyCode.UpArrow))
        {
            MoveUp();
        }
        
        if (Input.GetKeyDown(KeyCode.DownArrow))
        {
            MoveDown();
        }
    }
}