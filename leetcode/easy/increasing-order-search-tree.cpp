/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode() : val(0), left(nullptr), right(nullptr) {}
 *     TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
 *     TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
 * };
 */
class Solution {
public:
    TreeNode* increasingBST(TreeNode* root) {
        vector<int> arr;
        traverse(root, arr);
        TreeNode *final_node, *cpt;
        final_node = new TreeNode();
        cpt = final_node;
        for(auto x: arr)
        {
            cpt->right = new TreeNode(x);
            cpt = cpt->right;
        }
        return final_node->right;
    }
    void traverse(TreeNode *cpt, vector<int>& arr)
    {
        if(cpt)
        {
            traverse(cpt->left, arr);
            arr.push_back(cpt->val);
            traverse(cpt->right, arr);
        }
    }
};
