/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode() : val(0), next(nullptr) {}
 *     ListNode(int x) : val(x), next(nullptr) {}
 *     ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 */
class Solution {
public:
    ListNode* middleNode(ListNode* head) {
        int size = 0, updateCheck=0;
        ListNode *tempHead = head, *middleNode = head;
        while(tempHead->next != nullptr){
            tempHead = tempHead->next;
            if(size%2) middleNode = middleNode->next;
            size++;
        }
        return size%2 ? middleNode->next : middleNode;
    }
};
