package acmachine;

import java.util.Queue;
import java.util.Set;
import java.util.Map.Entry;
import java.util.Collections;
import java.util.HashSet;
import java.util.LinkedList;

public class AcMachine {
    private AcNode root = new AcNode('/');

    public void insert(String word) {
        if (word == null || "".equals(word)) {
            return;
        }
        AcNode p = root;
        for (int i = 0; i < word.length(); i++) {
            char data = word.charAt(i);
	    	AcNode next = p.children.get(data);
            if (next == null) {
	    		next = new AcNode(data);
	    		p.children.put(data, next);
	    	}
	    	p = next;
        }
        p.isEndingChar = true;
	    p.length = word.length();
    }

    // 构造fail指针
    public void buildFailPointer() {
        Queue<AcNode> queue = new LinkedList<>();
        root.fail = null;
		queue.add(root);
        while (!queue.isEmpty()) {
            AcNode p = queue.remove();
            for (Entry<Character, AcNode> entry : p.children.entrySet()) {
                AcNode pc = entry.getValue();
                if (p == root) {
					pc.fail = root;
                } else {
                    AcNode q = p.fail;
					while (q != null) {
						AcNode qc = q.children.get(pc.data);
						if (qc != null) {
							pc.fail = qc;
							break;
						}
						q = q.fail;
					}
                    if (q == null) {
						pc.fail = root;
					}
                }
                queue.add(pc);
            }
        }
    }

    public Set<String> match(String content, boolean checkAll) {
        Set<String> result = null;
		AcNode p = root;
        for (int i = 0; i < content.length(); ++i) {
			char data = content.charAt(i);
            while (p.children.get(data) == null && p != root) {
				p = p.fail; 
			}
            p = p.children.get(data);
			if (p == null) {
                p = root;
            }  
			AcNode tmp = p;
            while (tmp != root) {  
				if (tmp.isEndingChar) {
					if (null == result) {
                        result = new HashSet<>();
                    }
					result.add(content.substring(i - tmp.length + 1, i + 1));
					if (!checkAll) {
                        return result;
                    }
				}
				tmp = tmp.fail;
			}
        }

        return null == result ? Collections.emptySet() : result;
    }

    public static void main(String[] args) {
        AcMachine acMachine = new AcMachine();
        String[] str = {"fuck", "wtf", "motherfucker", "good", "lol"};
        for (int i = 0; i < str.length; i++) {
            acMachine.insert(str[i]);
        }
        acMachine.buildFailPointer();
        System.out.println(acMachine.match("fucker", false));
    }
}
