/*
// Definition for a Node.
class Node {
    public int val;
    public Node left;
    public Node right;

    public Node() {}

    public Node(int _val) {
        val = _val;
    }

    public Node(int _val,Node _left,Node _right) {
        val = _val;
        left = _left;
        right = _right;
    }
};
*/
class Solution {
    public Node treeToDoublyList(Node root) {
        
        if (root == null) {
            return null;
        }
        
        ArrayList<Node> list = new ArrayList<Node>();     
        recurse(root, list);
        
        for (int i = 0; i < list.size() - 1; i++) {
            list.get(i).right = list.get(i+1);
            list.get(i+1).left = list.get(i);
        }
        
        Node first = list.get(0);
        Node last = list.get(list.size()-1);
        
        first.left = last;
        last.right = first;
        
        return first;
    }
    
    public void recurse(Node root, ArrayList<Node> list) {
        
        if (root == null) {
            return;
        }
        
        recurse(root.left, list);
        list.add(root);
        recurse(root.right, list);
    }
}