using UnityEngine;

public class Graph : MonoBehaviour {

    [SerializeField]
    Transform pointPrefab;

    [SerializeField, Range(10, 100)]
    int resolution = 10;
    
    [SerializeField]
    FunctionLibrary.FunctionName function;

    Transform[] points;

    void Awake () {
        float step = 2f / resolution;
        var scale = Vector3.one * step;
        //var position = Vector3.zero;
        points = new Transform[resolution * resolution];
        for (int i = 0; i < points.Length; i++) {
            //if (x == resolution) {
            //	x = 0;
            //	z += 1;
            //}
            Transform point = points[i] = Instantiate(pointPrefab);
            //position.x = (x + 0.5f) * step - 1f;
            //position.z = (z + 0.5f) * step - 1f;
            //point.localPosition = position;
            point.localScale = scale;
            point.SetParent(transform, false);
        }
    }

    void Update () {
        FunctionLibrary.Function f = FunctionLibrary.GetFunction(function);
        float time = Time.time;
        float step = 2f / resolution;
        float v = 0.5f * step - 1f;
        for (int i = 0, x = 0, z = 0; i < points.Length; i++, x++) {
            if (x == resolution) {
                x = 0;
                z += 1;
                v = (z + 0.5f) * step - 1f;
            }
            float u = (x + 0.5f) * step - 1f;
            points[i].localPosition = f(u, v, time);
        }
    }
}