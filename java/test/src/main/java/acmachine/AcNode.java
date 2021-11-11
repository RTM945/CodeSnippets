package acmachine;

import java.util.HashMap;
import java.util.Map;

public class AcNode {
    public char data;
	public Map<Character, AcNode> children = new HashMap<>();

    public int length = -1; // 当isEndingChar=true时，记录模式串长度
    public boolean isEndingChar = false;
    public AcNode fail;

    public AcNode(char data) {
        this.data = data;
    }
}
