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

		if (hasNode(root, p) && hasNode(root,q)) {

			boolean allInLeft = root.left != null && hasNode(root.left, p) && hasNode(root.left, q);
			boolean allInRight = root.right != null && hasNode(root.right, p) && hasNode(root.right, q);
			
			if (!allInLeft && !allInRight) {
				return root;
			}

			if (allInLeft) {
				return lowestCommonAncestor(root.left, p, q);
			}

			if (allInRight) {
				return lowestCommonAncestor(root.right,p, q);
			}

		}  

		return null;
    }

    public boolean hasNode(TreeNode root, TreeNode node) {

    	if (root == node) {
    		return true;
    	} else if (root.left != null && hasNode(root.left, node)) {
    		return true;
    	}  else if (root.right != null && hasNode(root.right, node)) {
    		return true;
    	} 

    	return false;
    }
}