using System.Collections;
using System.Collections.Generic;
using TMPro;
using UnityEngine;
using UnityEngine.UI;

public class GameManager : MonoBehaviour
{

    public Transform cellRoot;
    public Transform tileRoot;
    public GameObject cellPrefab;
    public GameObject tilePrefab;
    public TextMeshProUGUI scoreText;
    public GameObject gameOverPanel;
    
    int[,] grid = new int[4, 4];
    
    Tile[,] tile = new Tile[4,4];
    
    Cell[,] cell = new Cell[4,4];
    
    int score = 0;

    // Start is called before the first frame update
    void Start()
    {
        Debug.Log("2048 Start");

        CreateGrid();
        // 不 Force后面SpawnNumber拿到的position都是000
        Canvas.ForceUpdateCanvases();

        SpawnNumber();
        SpawnNumber();
        UpdateView();
    }

    void CreateGrid()
    {
        for (int i = 0; i < 16; i++)
        {
            // 底图 GridLayout Group 自动布局
            var obj = Instantiate(cellPrefab, cellRoot);
            int x = i % 4;
            int y = i / 4;
            cell[x, y] = obj.GetComponent<Cell>();
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
        tile[x, y] = Instantiate(tilePrefab).GetComponent<Tile>();
        RectTransform rect = tile[x, y].GetComponent<RectTransform>();
        // 动态生成的 tile 设置预定好的坐标
        tile[x, y].transform.position = cell[x, y].transform.position;
        tile[x, y].transform.SetParent(cell[x, y].transform, false);
        rect.localPosition = Vector3.zero;
        rect.localScale = Vector3.one;
        rect.sizeDelta = Vector2.zero;
        
        tile[x, y].SetNumber(number);
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
                        
                        // Tile 移动
                        tile[newX - 1, y] = tile[newX, y];
                        tile[newX, y] = null;
                        tile[newX - 1, y].transform.SetParent(cell[newX - 1, y].transform, false);
                        
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
                        
                        tile[newX + 1, y] = tile[newX, y];
                        tile[newX, y] = null;
                        tile[newX + 1, y].transform.SetParent(cell[newX + 1, y].transform, false);
                        
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
                        
                        tile[x, newY - 1] = tile[x, newY];
                        tile[x, newY] = null;
                        tile[x, newY - 1].transform.SetParent(cell[x, newY - 1].transform, false);
                        
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
                        
                        tile[x, newY + 1] = tile[x, newY];
                        tile[x, newY] = null;
                        tile[x, newY + 1].transform.SetParent(cell[x, newY + 1].transform, false);
                        
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
                    
                    Destroy(tile[x + 1, y].gameObject);
                    tile[x + 1, y] = null;
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
                    
                    // 合并的 tile 销毁
                    Destroy(tile[x - 1, y].gameObject);
                    tile[x - 1, y] = null;
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
                    
                    Destroy(tile[x, y + 1].gameObject);
                    tile[x, y + 1] = null;
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
                    
                    Destroy(tile[x, y - 1].gameObject);
                    tile[x, y - 1] = null;
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
            if (tile[x, y] != null)
            {
                tile[x, y].SetNumber(grid[x, y]);
            }
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
        grid = new int[4, 4];
        for (int x = 0; x < 4; x++)
        {
            for (int y = 0; y < 4; y++)
            {
                Destroy(tile[x, y]?.gameObject);
                tile[x, y] = null;
                Destroy(cell[x, y].gameObject);
                cell[x, y] = null;
            }
        }
        tile = new Tile[4,4];
    
        cell = new Cell[4,4];
        
        CreateGrid();
        
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