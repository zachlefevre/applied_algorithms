"""
Zachary LeFevre
Solution for https://www.hackerrank.com/challenges/binary-search-tree-lowest-common-ancestor/problem
"""

"""
Node is defined as
self.left (the left child of the node)
self.right (the right child of the node)
self.data (the value of the node)
"""
def lca(root, v1, v2):
    if((v1 < root.data and v2 > root.data) or (v1 > root.data and v2 < root.data) or v1 == root.data or v2 == root.data):
        return root
    elif(v1 < root.data):
        return lca(root.left, v1, v2)
    else: return lca(root.right, v1, v2)
