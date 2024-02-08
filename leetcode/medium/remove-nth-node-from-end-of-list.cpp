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
    void print(ListNode *h){
        while(h){
            std::cout << h->val << ", ";
            h = h->next;
        }
    }
public:
    ListNode* removeNthFromEnd(ListNode* head, int n) {
        ListNode *fast = head, *slow = head;
        int size = 0;
        while(fast && fast->next){
            fast = fast->next->next;
            slow = slow->next;
            size += 2;
        }
        if(fast)    size++;
        int half = size/2+1;
        
        ListNode *temp;
        if(size-half-n >= 0){
            // from middle
            int step = size-half-n;
            while(step--)
                slow = slow->next;
            temp = slow->next;
            slow->next = temp->next;
        }else{
            ListNode *tempHead = head;
            int step = size-n;

            while(step-- > 1)   tempHead = tempHead->next;
            
            temp = tempHead->next;
            if(step>=0) tempHead->next = temp->next;
            else{
                temp = head;
                head = head->next;
            }
        }
        delete temp;
        return head;
    }
};;
