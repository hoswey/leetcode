/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode(int x) { val = x; }
 * }
 */
class Solution {

    public TreeNode lowestCommonAncestor(TreeNode root, TreeNode p, TreeNode q) {

    	if (root == null) {
    		return null;
    	}        

    	List<TreeNode> pPath = new ArrayList<TreeNode>();
    	List<TreeNode> qPath = new ArrayList<TreeNode>();

    	backtracking(root,pPath,p);
    	backtracking(root,qPath,q);

    	for (int i=0; i < pPath.size(); i++) {
    		if (i >= qPath.size() - 1) {
    			return null;
    		}

    		if (pPath.get(i) != qPath.get(i)) {
    			return pPath.get(i-1);
    		}
    	}

    	return null;
    }

    public boolean backtracking(TreeNode root, List<TreeNode> path, TreeNode node) {

    	path.add(root);
    	if (root.val == node.val) {
    		return true;
    	}

    	if (root.left != null && backtracking(root.left, path, node)) {
    		return true;
	    }

	    if (root.right != null && backtracking(root.right, path, node)) {
    		return true;
	    }

	    path.remove(root);
	    return false;
	}
}