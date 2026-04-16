using System.Collections;
using System.Collections.Generic;
using TMPro;
using UnityEngine;
using UnityEngine.UI;

public class GameManager : MonoBehaviour
{

    public GameObject tilePrefab;
    public Transform board;
    public TextMeshProUGUI scoreText;
    public GameObject gameOverPanel;

    Tile[] tiles = new Tile[16];
    int[,] grid = new int[4, 4];
    
    int score = 0;

    // Start is called before the first frame update
    void Start()
    {
        Debug.Log("2048 Start");

        CreateGrid();

        Restart();
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
        if (GetEmptyGrid().Count > 0)
        {
            return false;
        }

        for (int y = 0; y < 4; y++)
        {
            for (int x = 0; x < 3; x++)
            {
                if (grid[x, y] == grid[x + 1, y])
                {
                    return false;
                }
            }
        }

        for (int x = 0; x < 4; x++)
        {
            for (int y = 0; y < 3; y++)
            {
                if (grid[x, y] == grid[x, y + 1])
                {
                    return false;
                }
            }
        }
        return true;
    }

    bool CompressLeft()
    {
        bool compress = false;
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
                        compress = true;
                    }
                }
            }
        }
        return compress;
    }

    bool CompressRight()
    {
        bool compress = false;
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
                        compress = true;
                    }
                }
            }
        }
        return compress;
    }

    bool CompressUp()
    {
        bool compress = false;
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
                        compress = true;
                    }
                }
            }
        }
        return compress;
    }

    bool CompressDown()
    {
        bool compress = false;
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
                        compress = true;
                    }
                }
            }
        }
        return compress;
    }

    bool MergeLeft()
    {
        bool merged = false;
        for (int y = 0; y < 4; y++)
        {
            for (int x = 0; x < 3; x++)
            {
                if (grid[x, y] != 0 && grid[x, y] == grid[x + 1, y])
                {
                    grid[x, y] *= 2;
                    grid[x + 1, y] = 0;
                    score += grid[x, y];
                    merged = true;
                }
            }
        }
        return merged;
    }

    bool MergeRight()
    {
        bool merged = false;
        for (int y = 0; y < 4; y++)
        {
            for (int x = 3; x > 0; x--)
            {
                if (grid[x, y] != 0 && grid[x, y] == grid[x - 1, y])
                {
                    grid[x, y] *= 2;
                    score += grid[x, y];
                    grid[x - 1, y] = 0;
                    merged = true;
                }
            }
        }
        return merged;
    }

    bool MergeUp()
    {
        bool merged = false;
        for (int y = 0; y < 3; y++)
        {
            for (int x = 0; x < 4; x++)
            {
                if (grid[x, y] != 0 && grid[x, y] == grid[x, y + 1])
                {
                    grid[x, y] *= 2;
                    score += grid[x, y];
                    grid[x, y + 1] = 0;
                    merged = true;
                }
            }
        }
        return merged;
    }

    bool MergeDown()
    {
        bool merged = false;
        for (int y = 3; y > 0; y--)
        {
            for (int x = 0; x < 4; x++)
            {
                if (grid[x, y] != 0 && grid[x, y] == grid[x, y - 1])
                {
                    grid[x, y] *= 2;
                    score += grid[x, y];
                    grid[x, y - 1] = 0;
                    merged = true;
                }
            }
        }
        return merged;
    }

    void MoveLeft()
    {   
        bool moved = false;
        moved |= CompressLeft();
        moved |= MergeLeft();
        moved |= CompressLeft();
        if (moved)
        {
            SpawnNumber();
        }
        UpdateView();
    }

    void MoveRight()
    {
        bool moved = false;
        moved |= CompressRight();
        moved |= MergeRight();
        moved |= CompressRight();
        if (moved)
        {
            SpawnNumber();
        }
        UpdateView();
    }

    void MoveUp()
    {
        bool moved = false;
        moved |= CompressUp();
        moved |= MergeUp();
        moved |= CompressUp();
        if (moved)
        {
            SpawnNumber();
        }
        UpdateView();
    }

    void MoveDown()
    {
        bool moved = false;
        moved |= CompressDown();
        moved |= MergeDown();
        moved |= CompressDown();
        if (moved)
        {
            SpawnNumber();
        }
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
        scoreText.text = "Score: " + score;
    }
    
    void ShowGameOver()
    {
        gameOverPanel.SetActive(true);
    }

    public void Restart()
    {
        gameOverPanel.SetActive(false);
        score = 0;
        for (int x = 0; x < 4; x++)
        {
            for (int y = 0; y < 4; y++)
            {
                grid[x, y] = 0;
            }
        }
        SpawnNumber();
        SpawnNumber();

        UpdateView();
    }

    // Update is called once per frame
    void Update()
    {
        UpdateView();
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

        if (IsGameOver())
        {
            ShowGameOver();
        }
        
        if (Input.GetKeyDown(KeyCode.R))
        {
            Restart();
        }
    }
}