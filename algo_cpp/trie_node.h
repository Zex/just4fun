#pragma once
#include <iostream>

using namespace std;

struct TrieNode {
  vector<TrieNode*> next;
  string word;

  TrieNode() {
    next.resize(26, nullptr);
  }
};

void print_trie(ostream &o, TrieNode *root) {
  if (!root)
    return;

  for (auto n : root->next)
    o << n << ',';
  o << '#' << root->word << endl;

  for (auto n : root->next)
    print_trie(o, n);
}