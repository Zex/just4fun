#include <iostream>
#include <limits.h>
#include "list_node.h"

using namespace std;

void print_list(ListNode* head);

class Solution {
public:
    
    ListNode* merge(ListNode* l, ListNode* r) {
      ListNode *ret = new ListNode(0), *head = ret;

      while (l && r) {
        if (l->val < r->val) {
          head->next = l;
          l = l->next;
        } else {
          head->next = r;
          r = r->next;
        }
        head = head->next;
      }

      while (l) {
          head->next = l;
          l = l->next;
          head = head->next;
      }

      while (r) {
          head->next = r;
          r = r->next;
          head = head->next;
      }

      return ret->next;
    }

    ListNode* sortList(ListNode* head) {
      if (!head || !(head->next))
        return head;

      ListNode *mid = head, *tail = head, *p = head;
      while (tail && tail->next) {
          p = mid;
          mid = mid->next;
          tail = tail->next->next;
      }

      p->next = NULL;
      ListNode* l = sortList(head);
      ListNode* r = sortList(mid);
      return merge(l, r);
    }
};

void print_list(ListNode* head) {
  ListNode* p = head;
  while (p) {
    cout << p->val << ',';
    p = p->next;
  }
  cout << endl;
}

int main() {
  Solution s;
  ListNode* head = new ListNode(4);
  head->next = new ListNode(2);
  head->next->next = new ListNode(7);
  head->next->next->next = new ListNode(1);
  head->next->next->next->next = new ListNode(3);
  head->next->next->next->next->next = new ListNode(5);

  print_list(head);
  head = s.sortList(head);
  print_list(head);
  return 0;
}