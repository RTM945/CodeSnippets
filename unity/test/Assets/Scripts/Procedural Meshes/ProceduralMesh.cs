using ProceduralMeshes;
using ProceduralMeshes.Generators;
using ProceduralMeshes.Streams;
using UnityEngine;
using UnityEngine.Rendering;

[RequireComponent(typeof(MeshFilter), typeof(MeshRenderer))]
public class ProceduralMesh : MonoBehaviour
{
    Mesh mesh;
    
    [SerializeField, Range(1, 10)]
    int resolution = 1;

    void Awake () {
        mesh = new Mesh {
            name = "Procedural Mesh"
        };
        GetComponent<MeshFilter>().mesh = mesh;
    }
    
    void OnValidate () => enabled = true;

    void Update () {
        GenerateMesh();
        enabled = false;
    }

    void GenerateMesh()
    {
        Mesh.MeshDataArray meshDataArray = Mesh.AllocateWritableMeshData(1);
		Mesh.MeshData meshData = meshDataArray[0];

        MeshJob<SquareGrid, MultiStream>.ScheduleParallel(
            mesh, meshData, resolution, default
        ).Complete();

		Mesh.ApplyAndDisposeWritableMeshData(meshDataArray, mesh);
    }
}